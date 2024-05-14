package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYoutubeURLValidation(t *testing.T) {
	type testCase struct {
		url      string
		expected bool
	}

	testCases := []testCase{
		{"https://www.youtube.com/watch?v=dQw4w9WgXcQ", true}, // Valid URL (Standard URL)
		{"http://youtu.be/dQw4w9WgXcQ", true},                 // Valid URL (Shortened URL)
		{"https://www.youtube.com/watch?v=", false},           // Invalid URL (Video ID missing)
		{"https://www.youtube.com/watch?v=invalidID", false},  // Invalid URL (Invalid Video ID)
		{"https://www.youtu.be/watch?v=dQw4w9WgXcQ", false},   // Invalid URL (Incorrect domain)
		{"https://www.google.com", false},                     // Invalid URL (Google URL)
		{"", true},                                            // Empty URL
		{"random text", false},                                // Invalid URL Format
	}

	for _, tc := range testCases {
		t.Run(tc.url, func(t *testing.T) {
			v := validator.New()
			if err := v.RegisterValidation("youtubeURL", youtubeURLValidation); err != nil {
				t.Fatalf("failed to register custom validation: %v", err)
			}

			var req struct {
				Url string `validate:"youtubeURL"`
			}

			req.Url = tc.url
			err := v.Struct(req)

			if tc.expected {
				assert.NoError(t, err, "unexpected error for URL: "+tc.url)
			} else {
				assert.Error(t, err, "expected error for URL: "+tc.url)
			}
		})
	}
}
