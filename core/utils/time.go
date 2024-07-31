package utils

import (
	"time"
)

func Time(duration int) time.Time {
	lapse := time.Duration(duration) * time.Second
	now := time.Now()
	tnow := time.Date(now.Year(), now.Month(), now.Day(),now.Hour(), now.Minute(), now.Second(), 0, now.Location())
	t := tnow.Add(lapse)
	return t
} 

func Parse(s string) *time.Time {
	layout := "2006-01-02 15:04:05"
    t, err := time.Parse(layout, s)
    if err != nil {
        return nil
    }
	return &t
}
