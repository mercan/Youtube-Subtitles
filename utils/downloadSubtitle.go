package utils

import (
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

func DownloadSubtitles(videoId string) error {
	cmd := exec.Command("yt-dlp",
		skipDownload,
		writeAutoSub,
		subLanguage,
		subFormatExt,
		outputFlag,
		outputPath,
		"--",
		videoId,
	)

	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	log.Printf("Subtitles Downloaded Successfully VideoID: %s\n", videoId)
	return nil
}
