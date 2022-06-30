package utils

import (
	"strings"
	"time"
)

func ExtractAmadeusTime(t string) time.Duration {
	split := strings.Split(t, "PT")[1]
	split = strings.ToLower(split)
	duration, _ := time.ParseDuration(split)

	return duration
}
