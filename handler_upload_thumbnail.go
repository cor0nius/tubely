package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bootdotdev/learn-file-storage-s3-golang-starter/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUploadThumbnail(w http.ResponseWriter, r *http.Request) {
	videoIDString := r.PathValue("videoID")
	videoID, err := uuid.Parse(videoIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	fmt.Println("uploading thumbnail for video", videoID, "by user", userID)

	const maxMemory int64 = 10 << 20
	err = r.ParseMultipartForm(maxMemory)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't parse form", err)
		return
	}

	file, header, err := r.FormFile("thumbnail")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't parse form file", err)
		return
	}
	defer file.Close()

	mediaType, _, err := mime.ParseMediaType(header.Header.Get("Content-Type"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't get image type", err)
		return
	}
	if mediaType != "image/jpeg" && mediaType != "image/png" {
		respondWithError(w, http.StatusBadRequest, "Incorrect file type", err)
		return
	}

	video, err := cfg.db.GetVideo(videoID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Video not found", err)
		return
	}
	if video.UserID != userID {
		respondWithError(w, http.StatusUnauthorized, "User is not authorized to update this video", err)
		return
	}

	if video.ThumbnailURL != nil {
		oldThumbnailName := strings.Split(*video.ThumbnailURL, "assets/")[1]
		if err = os.Remove(filepath.Join(cfg.assetsRoot, oldThumbnailName)); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error removing old thumbnail", err)
			return
		}
	}

	thumbnailFileExtension := strings.Split(mediaType, "/")[1]
	fileNameBytes := make([]byte, 32)
	_, err = rand.Read(fileNameBytes)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create new thumbnail name", err)
		return
	}
	thumbnailFileName := fmt.Sprintf("%s.%s", base64.RawURLEncoding.EncodeToString(fileNameBytes), thumbnailFileExtension)
	thumbnailPath := filepath.Join(cfg.assetsRoot, thumbnailFileName)
	thumbnailFile, err := os.Create(thumbnailPath)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create thumbnail file", err)
		return
	}
	defer thumbnailFile.Close()
	if _, err = io.Copy(thumbnailFile, file); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't upload thumbnail", err)
		return
	}
	thumbnailURL := fmt.Sprintf("http://localhost:%s/assets/%s", cfg.port, thumbnailFileName)
	video.ThumbnailURL = &thumbnailURL

	err = cfg.db.UpdateVideo(video)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update video", err)
		return
	}

	respondWithJSON(w, http.StatusOK, video)
}
