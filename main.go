package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"

	"golang.org/x/exp/slices"
)

func main() {
	currentTime := time.Now()
	loc, err := time.LoadLocation("Europe/Luxembourg")
	if err != nil {
		panic(err)
	}
	//See https://luxembourg.public.lu/en/living/quality-of-life/jours-feries-legaux.html
	//for Luxembourg Public holidays
	publicHolidays := make([]string, 11)
	publicHolidays[0] = fmt.Sprintf("%d%s", currentTime.Year(), "0101")
	publicHolidays[1] = fmt.Sprintf("%d%s", currentTime.Year(), "0421")
	publicHolidays[2] = fmt.Sprintf("%d%s", currentTime.Year(), "0501")
	publicHolidays[3] = fmt.Sprintf("%d%s", currentTime.Year(), "0509")
	publicHolidays[4] = fmt.Sprintf("%d%s", currentTime.Year(), "0529")
	publicHolidays[5] = fmt.Sprintf("%d%s", currentTime.Year(), "0609")
	publicHolidays[6] = fmt.Sprintf("%d%s", currentTime.Year(), "0623")
	publicHolidays[7] = fmt.Sprintf("%d%s", currentTime.Year(), "0815")
	publicHolidays[8] = fmt.Sprintf("%d%s", currentTime.Year(), "1101")
	publicHolidays[9] = fmt.Sprintf("%d%s", currentTime.Year(), "1225")
	publicHolidays[10] = fmt.Sprintf("%d%s", currentTime.Year(), "1226")
	//See https://stackoverflow.com/a/73880886 for the detail of the trick
	var daysInMonth = 32 - (time.Date(currentTime.Year(), currentTime.Month(), 32, 0, 0, 0, 0, loc)).Day()
	var workingDaysCount = 0
	var workingDaysCountUntilNow = 0
	for d := 1; d <= daysInMonth; d++ {
		var currentDay = time.Date(currentTime.Year(), currentTime.Month(), d, 0, 0, 0, 0, loc)
		var currentDayStr = currentDay.Format("20060102")
		if currentDay.Weekday() != time.Saturday && currentDay.Weekday() != time.Sunday && !slices.Contains(publicHolidays, currentDayStr) {
			workingDaysCount++
			if currentTime.Day() >= d {
				workingDaysCountUntilNow++
			}
		}
	}
	color.Cyan("[+] %s information:\n", currentTime.Format("January"))
	fmt.Printf("%-02d working days so %-03d working hours until now.\n", workingDaysCountUntilNow, workingDaysCountUntilNow*8)
	fmt.Printf("%-02d working days so %-03d working hours in total.\n", workingDaysCount, workingDaysCount*8)
}
