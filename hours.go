package main

import (
	"time"
)

type Hour struct {
	Label string
	Unix  int64
}

type DayData struct {
	Unix     int64
	Month    string
	MonthInt int
	Year     int
	Weekday  string
	Day      int
	Hours    []Hour
}

func generateDay(unix int64) *DayData {
	t := time.Unix(unix, 0)
	hours := make([]Hour, 24)

	for i := 0; i < len(hours); i++ {
		time := t.Add(time.Hour * time.Duration(i))
		hours[i] = Hour{
			Label: time.Format("3:04 PM"),
			Unix:  time.Unix(),
		}
	}

	return &DayData{
		Month:    t.Month().String(),
		MonthInt: int(t.Month()),
		Year:     t.Year(),
		Weekday:  t.Weekday().String(),
		Day:      t.Day(),
		Unix:     t.Unix(),
		Hours:    hours,
	}
}
