package main

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Day struct {
	Unix         int64
	Weekday      string
	Day          int
	IsToday      bool
	IsOutOfMonth bool
}

type MonthButton struct {
	Month int
	Year  int
}

type Calendar struct {
	Weekdays []string
	Month    string
	Previous MonthButton
	Next     MonthButton
	Year     int
	Days     []Day
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	mux := http.NewServeMux()

	tmpl := NewTemplateRenderer()

	mux.Handle("GET /days/{dayId}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dayId := r.PathValue("dayId")
		if dayId == "" {
			http.Error(w, "Missing day ID", http.StatusBadRequest)
			return
		}
		unix, err := strconv.ParseInt(dayId, 10, 64)
		if err != nil {
			http.Error(w, "Invalid day ID", http.StatusBadRequest)
			return
		}
		day := generateDay(unix)
		if isHtmxRequest(r) {
			err = tmpl.RenderSnippet(w, "views/day.html", "body", day)
		} else {
			err = tmpl.RenderPage(w, "views/day.html", day)
		}
		if err != nil {
			http.Error(w, "Could not execute template", http.StatusInternalServerError)
			logger.Error("Could not execute template", "error", err.Error())
		}
	}))

	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		monthQuery := query.Get("month")
		yearQuery := query.Get("year")

		if (monthQuery != "" && yearQuery == "") || (monthQuery == "" && yearQuery != "") {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		var month int
		var year int

		if monthQuery != "" && yearQuery != "" {
			var err error
			month, err = strconv.Atoi(monthQuery)
			if err != nil {
				http.Error(w, "Invalid month", http.StatusBadRequest)
			}
			year, err = strconv.Atoi(yearQuery)
			if err != nil {
				http.Error(w, "Invalid year", http.StatusBadRequest)
			}
			if month < 1 || month > 12 {
				http.Error(w, "Invalid month", http.StatusBadRequest)
			}
		} else {
			now := time.Now()
			month = int(now.Month())
			year = now.Year()
		}

		cal := generateCalendar(month, year)

		var err error
		if isHtmxRequest(r) {
			err = tmpl.RenderSnippet(w, "views/home.html", "body", cal)
		} else {
			err = tmpl.RenderPage(w, "views/home.html", cal)
		}
		if err != nil {
			http.Error(w, "Could not execute template", http.StatusInternalServerError)
			logger.Error("Could not execute template", "error", err.Error())
		}
	}))

	mux.Handle("GET /assets/{path...}", Assets())

	var handler http.Handler
	handler = mux

	server := &http.Server{
		Handler: handler,
		Addr:    ":4000",
	}

	logger.Info("App Listening", "addr", server.Addr)
	panic(server.ListenAndServe())
}

func isHtmxRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") != ""
}
