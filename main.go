package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/uber-go/zap"
)

var (
	day    = flag.Int("d", 1, "day")
	month  = flag.Int("m", int(time.Now().Month()), "month")
	year   = flag.Int("y", time.Now().Year(), "year")
	tz     = flag.String("tz", "Europe/Berlin", "time zone")
	layout = flag.String("l", time.RFC3339, "layout")
	log    = zap.New(zap.NewTextEncoder(zap.TextTimeFormat(*layout)))
)

func main() {
	flag.Parse()
	loc, err := time.LoadLocation(*tz)
	if err != nil {
		log.Fatal("failed to load location", zap.Error(err))
	}
	t := time.Date(*year, time.Month(*month), *day, 0, 0, 0, 0, loc)
	fmt.Println(t.Format(*layout))
}
