// time_operations.go
// Learn the time package for date/time operations

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Current Time ===")
	
	// TODO: Get current time
	now := /* get current time */
	fmt.Printf("Current time: %v\\n", now)
	fmt.Printf("Unix timestamp: %d\\n", /* get unix timestamp */)
	fmt.Printf("Formatted: %s\\n", /* format as RFC3339 */)
	
	fmt.Println("\\n=== Time Creation ===")
	
	// TODO: Create specific time
	birthday := /* create time for Jan 1, 2000, 12:00:00 UTC */
	fmt.Printf("Birthday: %v\\n", birthday)
	
	// TODO: Parse time from string
	timeStr := "2023-12-25 15:30:45"
	layout := "2006-01-02 15:04:05"
	parsed, err := /* parse timeStr using layout */
	if err != nil {
		fmt.Printf("Parse error: %v\\n", err)
	} else {
		fmt.Printf("Parsed: %v\\n", parsed)
	}
	
	fmt.Println("\\n=== Time Formatting ===")
	
	// TODO: Format time in different ways
	fmt.Printf("RFC3339: %s\\n", /* format now as RFC3339 */)
	fmt.Printf("Kitchen: %s\\n", /* format now as Kitchen */)
	fmt.Printf("Custom: %s\\n", /* format now as "Jan 2, 2006 at 3:04pm" */)
	fmt.Printf("Date only: %s\\n", /* format now as "2006-01-02" */)
	fmt.Printf("Time only: %s\\n", /* format now as "15:04:05" */)
	
	fmt.Println("\\n=== Time Arithmetic ===")
	
	// TODO: Add/subtract durations
	future := /* add 2 hours to now */
	past := /* subtract 30 minutes from now */
	
	fmt.Printf("Now: %s\\n", now.Format(time.Kitchen))
	fmt.Printf("Future (+2h): %s\\n", future.Format(time.Kitchen))
	fmt.Printf("Past (-30m): %s\\n", past.Format(time.Kitchen))
	
	// TODO: Calculate duration between times
	diff := /* calculate duration between future and now */
	fmt.Printf("Duration: %v\\n", diff)
	fmt.Printf("Minutes: %.1f\\n", /* get minutes from diff */)
	fmt.Printf("Hours: %.1f\\n", /* get hours from diff */)
	
	fmt.Println("\\n=== Time Comparison ===")
	
	time1 := time.Date(2023, 6, 15, 10, 0, 0, 0, time.UTC)
	time2 := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
	
	// TODO: Compare times
	fmt.Printf("time1.Before(time2): %t\\n", /* check if time1 is before time2 */)
	fmt.Printf("time1.After(time2): %t\\n", /* check if time1 is after time2 */)
	fmt.Printf("time1.Equal(time2): %t\\n", /* check if time1 equals time2 */)
	
	fmt.Println("\\n=== Duration Operations ===")
	
	// TODO: Create durations
	d1 := /* create 2 hours 30 minutes duration */
	d2 := /* create 45 minutes duration */
	
	fmt.Printf("d1: %v\\n", d1)
	fmt.Printf("d2: %v\\n", d2)
	fmt.Printf("d1 + d2: %v\\n", /* add durations */)
	fmt.Printf("d1 - d2: %v\\n", /* subtract durations */)
	
	// TODO: Duration components
	fmt.Printf("d1 in seconds: %.0f\\n", /* get d1 in seconds */)
	fmt.Printf("d1 in minutes: %.1f\\n", /* get d1 in minutes */)
	fmt.Printf("d1 in hours: %.1f\\n", /* get d1 in hours */)
	
	fmt.Println("\\n=== Timers and Tickers ===")
	
	// TODO: Use timer
	fmt.Println("Starting 2-second timer...")
	timer := /* create 2-second timer */
	
	// TODO: Wait for timer
	/* wait for timer to fire */
	fmt.Println("Timer fired!")
	
	// TODO: Use ticker for repeated events
	fmt.Println("Starting ticker (every 500ms, 3 times)...")
	ticker := /* create 500ms ticker */
	
	for i := 0; i < 3; i++ {
		/* wait for ticker */
		fmt.Printf("Tick %d at %s\\n", i+1, time.Now().Format("15:04:05.000"))
	}
	
	/* stop ticker */
	fmt.Println("Ticker stopped")
	
	fmt.Println("\\n=== Time Zones ===")
	
	// TODO: Work with time zones
	utc := time.Now().UTC()
	
	// Load specific time zones
	est, err := /* load "America/New_York" timezone */
	if err != nil {
		fmt.Printf("Timezone error: %v\\n", err)
	} else {
		estTime := /* convert utc to EST timezone */
		fmt.Printf("UTC: %s\\n", utc.Format("15:04:05 MST"))
		fmt.Printf("EST: %s\\n", estTime.Format("15:04:05 MST"))
	}
	
	pst, err := /* load "America/Los_Angeles" timezone */
	if err != nil {
		fmt.Printf("Timezone error: %v\\n", err)
	} else {
		pstTime := /* convert utc to PST timezone */
		fmt.Printf("PST: %s\\n", pstTime.Format("15:04:05 MST"))
	}
}