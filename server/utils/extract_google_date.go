package utils

import (
	"log"
	"time"
)

func ExtractGoogleDate(t int, timezone string) time.Time {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		log.Print("[WARNING]: Could not parse timezone.\nFalling back on locale time...")
		return time.Unix(int64(t), 0)
	}
	return time.Unix(int64(t), 0).In(loc)
}
