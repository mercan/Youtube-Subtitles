package utils

import (
	"bytes"
	"log"
	"os/exec"
)

const (
	skipDownload = "--skip-download"
	writeAutoSub = "--write-auto-sub"
	subLanguage  = "--sub-lang=tr"
	outputFlag   = "-o"
	outputPath   = "subtitles/%(id)s.%(ext)s"
	subFormatExt = "--sub-format=ttml"
)

func buildFlags(videoId string) []string {
	return []string{skipDownload, writeAutoSub, subLanguage, subFormatExt, outputFlag, outputPath, "--", videoId}
}

func DownloadSubtitles(videoId string) error {
	flags := buildFlags(videoId)
	cmd := exec.Command("yt-dlp", flags...)

	var out, err bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &err

	if err := cmd.Start(); err != nil {
		log.Printf("Error downloading subtitles for video ID %s: %v\n", videoId, err)
		return err
	}

	if err := cmd.Wait(); err != nil {
		log.Printf("Error downloading subtitles for video ID %s: %v\n", videoId, err)
		return err
	}

	log.Printf("Subtitles downloaded for video ID %s: %s\n", videoId, out.String())
	return nil
}
