package utils

import (
	"fmt"
	"regexp"
)

func videoEmbedURL(videoId string, begin BeginTime, end EndTime) string {
	return fmt.Sprintf("https://www.youtube.com/embed/%s?start=%d&end=%d",
		videoId,
		begin.Hours*3600+begin.Minutes*60+begin.Seconds,
		end.Hours*3600+end.Minutes*60+end.Seconds,
	)
}

func FilterSubtitles(videoId string, subtitles []Subtitle, text string) []Subtitle {
	var filteredSubtitles = make([]Subtitle, 0)

	for _, subtitle := range subtitles {
		if regexp.MustCompile(`(?i)` + text + `\b`).MatchString(subtitle.Text) {
			filteredSubtitles = append(filteredSubtitles, Subtitle{
				Text:  subtitle.Text,
				Begin: subtitle.Begin,
				End:   subtitle.End,
				VideoURL: videoEmbedURL(
					videoId,
					subtitle.Begin,
					subtitle.End,
				),
			})
		}
	}

	return filteredSubtitles
}
