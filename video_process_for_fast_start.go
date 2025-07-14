package main

import (
	"fmt"
	"os/exec"
)

func processVideoForFastStart(filepath string) (string, error) {
	outputFilePath := fmt.Sprintf("%s.processing", filepath)
	command := exec.Command("ffmpeg", "-i", filepath, "-c", "copy", "-movflags", "faststart", "-f", "mp4", outputFilePath)
	if err := command.Run(); err != nil {
		return "", err
	}
	return outputFilePath, nil
}
