package helpers

import (
	"fmt"
	"regexp"
)

func videoEmbedURL(videoId string, begin, end Time) string {
	start := begin.Hours*3600 + begin.Minutes*60 + begin.Seconds
	finish := end.Hours*3600 + end.Minutes*60 + end.Seconds

	return fmt.Sprintf("https://www.youtube.com/embed/%s?start=%d&end=%d", videoId, start, finish)
}

func FilterSubtitles(videoId, text string, subtitles []Subtitle) []Subtitle {
	filteredSubtitles := make([]Subtitle, 0, len(subtitles))

	re := regexp.MustCompile(`(?i)\b` + text + `\b`)

	for _, subtitle := range subtitles {
		if re.MatchString(subtitle.Text) {
			filteredSubtitles = append(filteredSubtitles, Subtitle{
				Text:     subtitle.Text,
				Begin:    subtitle.Begin,
				End:      subtitle.End,
				VideoURL: videoEmbedURL(videoId, subtitle.Begin, subtitle.End),
			})
		}
	}

	return filteredSubtitles
}
