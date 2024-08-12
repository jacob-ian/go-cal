package main

import "time"

func generateCalendar(month int, year int) *Calendar {
	tStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	if tStart.Weekday() != time.Sunday {
		tStart = tStart.AddDate(0, 0, int(time.Sunday)-int(tStart.Weekday()))
	}

	tEnd := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
	if tEnd.Weekday() != time.Saturday {
		tEnd = tEnd.AddDate(0, 0, int(time.Saturday)-int(tEnd.Weekday()))
	}

	now := time.Now()
	days := make([]Day, 5*7)
	cur := tStart
	for i := 0; i < len(days); i++ {
		days[i] = Day{
			Weekday:      cur.Weekday().String(),
			Day:          cur.Day(),
			IsOutOfMonth: cur.Month() != time.Month(month),
			IsToday:      now.Month() == cur.Month() && now.Year() == cur.Year() && now.Day() == cur.Day(),
			Unix:         cur.Unix(),
		}
		cur = cur.AddDate(0, 0, 1)
	}

	var previous MonthButton
	if month == 1 {
		previous = MonthButton{
			Year:  year - 1,
			Month: 12,
		}
	} else {
		previous = MonthButton{
			Year:  year,
			Month: month - 1,
		}
	}

	var next MonthButton
	if month == 12 {
		next = MonthButton{
			Year:  year + 1,
			Month: 1,
		}
	} else {
		next = MonthButton{
			Year:  year,
			Month: month + 1,
		}
	}

	return &Calendar{
		Weekdays: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		Month:    time.Month(month).String(),
		Previous: previous,
		Next:     next,
		Year:     year,
		Days:     days,
	}
}
