package main

import (
	"sync"
)

var wg sync.WaitGroup
var days = []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

func processDays(ch chan<- string) {
	for _, day := range days {
		wg.Add(1)
		go determineWorkingDay(day, ch)
	}
}

func determineWorkingDay(day string, ch chan<- string) {
	defer wg.Done()

	isWeekend := (day == "Sabtu" || day == "Minggu")

	if !isWeekend {

		isWorkingDay := len(day) > 0

		if isWorkingDay {
			ch <- day
		}
	}
}
