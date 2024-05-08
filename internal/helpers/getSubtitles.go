package helpers

import (
	"encoding/xml"
	"errors"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type XML struct {
	XMLName xml.Name `xml:"tt"`
	Body    BodyTag  `xml:"body"`
}

type BodyTag struct {
	XMLName xml.Name `xml:"body"`
	Div     DivTag   `xml:"div"`
}

type DivTag struct {
	XMLName xml.Name `xml:"div"`
	P       []PTag   `xml:"p"`
}

type PTag struct {
	XMLName xml.Name `xml:"p"`
	Text    string   `xml:",chardata"`
	Begin   string   `xml:"begin,attr"`
	End     string   `xml:"end,attr"`
}

type Subtitle struct {
	Text     string `json:"text"`
	Begin    Time   `json:"begin"`
	End      Time   `json:"end"`
	VideoURL string `json:"videoURL"`
}

type Time struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
	Seconds int `json:"seconds"`
}

var pwd, _ = os.Getwd()
var subtitleFolderPath = path.Join(pwd, "subtitles")

func readSubtitlesFolder() ([]os.DirEntry, error) {
	files, err := os.ReadDir(subtitleFolderPath)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func readSubtitlesFile(videoId string) (*os.File, error) {
	files, err := readSubtitlesFolder()
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, errors.New("subtitles folder is empty")
	}

	// Find the subtitle file that starts with the video id
	for _, file := range files {
		if file.Name() == videoId+".tr.ttml" {
			matchingFile := file
			subtitleFilePath := path.Join(subtitleFolderPath, matchingFile.Name())

			subtitleFile, err := os.Open(subtitleFilePath)
			if err != nil {
				return nil, err
			}

			return subtitleFile, nil
		}
	}

	return nil, errors.New("subtitle file not found")
}

func deleteSubtitlesFile(videoId string) error {
	if err := os.Remove(path.Join(subtitleFolderPath, videoId+".tr.ttml")); err != nil {
		return err
	}

	log.Printf("Subtitle file deleted: %s\n", videoId+".tr.ttml")
	return nil
}

func GetSubtitles(videoId string) ([]Subtitle, error) {
	xmlFile, err := readSubtitlesFile(videoId)
	if err != nil {
		return nil, err
	}

	if xmlFile == nil {
		return nil, err
	}

	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)
	xmlParse := XML{}

	if err := xml.Unmarshal(byteValue, &xmlParse); err != nil {
		return nil, err
	}

	subtitles := make([]Subtitle, 0, len(xmlParse.Body.Div.P))

	for _, subtitle := range xmlParse.Body.Div.P {
		const (
			HoursIndex   = 0
			MinutesIndex = 1
			SecondsIndex = 2
		)

		// Parse begin time
		sBegin := strings.Split(subtitle.Begin, ":")
		// Parse end time
		sEnd := strings.Split(subtitle.End, ":")

		// Convert string to int
		// Begin time
		bHours, _ := strconv.Atoi(sBegin[HoursIndex])
		bMinutes, _ := strconv.Atoi(sBegin[MinutesIndex])
		bSeconds, _ := strconv.Atoi(strings.Split(sBegin[SecondsIndex], ".")[0]) // Remove milliseconds

		// End time
		eHours, _ := strconv.Atoi(sEnd[HoursIndex])
		eMinutes, _ := strconv.Atoi(sEnd[MinutesIndex])
		eSeconds, _ := strconv.Atoi(strings.Split(sEnd[SecondsIndex], ".")[0]) // Remove milliseconds

		beginTime := &Time{
			Hours:   bHours,
			Minutes: bMinutes,
			Seconds: bSeconds,
		}

		endTime := &Time{
			Hours:   eHours,
			Minutes: eMinutes,
			Seconds: eSeconds,
		}

		subtitle := Subtitle{
			Text:  subtitle.Text,
			Begin: *beginTime,
			End:   *endTime,
		}

		subtitles = append(subtitles, subtitle)
	}

	// Remove subtitle file
	if err := deleteSubtitlesFile(videoId); err != nil {
		log.Printf("Error removing subtitle file: %s\n", err.Error())
	}

	return subtitles, nil
}
