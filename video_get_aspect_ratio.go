package main

import (
	"encoding/json"
	"os/exec"
)

type AspectRatio struct {
	Streams []struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"streams"`
}

func getVideoAspectRatio(filePath string) (string, error) {
	command := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)
	buffer, err := command.Output()
	if err != nil {
		return "", err
	}
	aspectRatio := AspectRatio{}
	if err := json.Unmarshal(buffer, &aspectRatio); err != nil {
		return "", err
	}
	if aspectRatio.Streams[0].Width/16 == aspectRatio.Streams[0].Height/9 {
		return "16:9", nil
	} else if aspectRatio.Streams[0].Width/9 == aspectRatio.Streams[0].Height/16 {
		return "9:16", nil
	} else {
		return "other", nil
	}
}
