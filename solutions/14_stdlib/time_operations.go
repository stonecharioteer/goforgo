// time_operations.go - SOLUTION
// Learn the time package for date/time operations

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Current Time ===")
	
	// Get current time
	now := time.Now()
	fmt.Printf("Current time: %v\n", now)
	fmt.Printf("Unix timestamp: %d\n", now.Unix())
	fmt.Printf("Formatted: %s\n", now.Format(time.RFC3339))
	
	fmt.Println("\n=== Time Creation ===")
	
	// Create specific time
	birthday := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)
	fmt.Printf("Birthday: %v\n", birthday)
	
	// Parse time from string
	timeStr := "2023-12-25 15:30:45"
	layout := "2006-01-02 15:04:05"
	parsed, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
	} else {
		fmt.Printf("Parsed: %v\n", parsed)
	}
	
	fmt.Println("\n=== Time Formatting ===")
	
	// Format time in different ways
	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("Kitchen: %s\n", now.Format(time.Kitchen))
	fmt.Printf("Custom: %s\n", now.Format("Jan 2, 2006 at 3:04pm"))
	fmt.Printf("Date only: %s\n", now.Format("2006-01-02"))
	fmt.Printf("Time only: %s\n", now.Format("15:04:05"))
	
	fmt.Println("\n=== Time Arithmetic ===")
	
	// Add/subtract durations
	future := now.Add(2 * time.Hour)
	past := now.Add(-30 * time.Minute)
	
	fmt.Printf("Now: %s\n", now.Format(time.Kitchen))
	fmt.Printf("Future (+2h): %s\n", future.Format(time.Kitchen))
	fmt.Printf("Past (-30m): %s\n", past.Format(time.Kitchen))
	
	// Calculate duration between times
	diff := future.Sub(now)
	fmt.Printf("Duration: %v\n", diff)
	fmt.Printf("Minutes: %.1f\n", diff.Minutes())
	fmt.Printf("Hours: %.1f\n", diff.Hours())
	
	fmt.Println("\n=== Time Comparison ===")
	
	time1 := time.Date(2023, 6, 15, 10, 0, 0, 0, time.UTC)
	time2 := time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)
	
	// Compare times
	fmt.Printf("time1.Before(time2): %t\n", time1.Before(time2))
	fmt.Printf("time1.After(time2): %t\n", time1.After(time2))
	fmt.Printf("time1.Equal(time2): %t\n", time1.Equal(time2))
	
	fmt.Println("\n=== Duration Operations ===")
	
	// Create durations
	d1 := 2*time.Hour + 30*time.Minute
	d2 := 45 * time.Minute
	
	fmt.Printf("d1: %v\n", d1)
	fmt.Printf("d2: %v\n", d2)
	fmt.Printf("d1 + d2: %v\n", d1+d2)
	fmt.Printf("d1 - d2: %v\n", d1-d2)
	
	// Duration components
	fmt.Printf("d1 in seconds: %.0f\n", d1.Seconds())
	fmt.Printf("d1 in minutes: %.1f\n", d1.Minutes())
	fmt.Printf("d1 in hours: %.1f\n", d1.Hours())
	
	fmt.Println("\n=== Timers and Tickers ===")
	
	// Use timer
	fmt.Println("Starting 2-second timer...")
	timer := time.NewTimer(2 * time.Second)
	
	// Wait for timer
	<-timer.C
	fmt.Println("Timer fired!")
	
	// Use ticker for repeated events
	fmt.Println("Starting ticker (every 500ms, 3 times)...")
	ticker := time.NewTicker(500 * time.Millisecond)
	
	for i := 0; i < 3; i++ {
		<-ticker.C
		fmt.Printf("Tick %d at %s\n", i+1, time.Now().Format("15:04:05.000"))
	}
	
	ticker.Stop()
	fmt.Println("Ticker stopped")
	
	fmt.Println("\n=== Time Zones ===")
	
	// Work with time zones
	utc := time.Now().UTC()
	
	// Load specific time zones
	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Printf("Timezone error: %v\n", err)
	} else {
		estTime := utc.In(est)
		fmt.Printf("UTC: %s\n", utc.Format("15:04:05 MST"))
		fmt.Printf("EST: %s\n", estTime.Format("15:04:05 MST"))
	}
	
	pst, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Printf("Timezone error: %v\n", err)
	} else {
		pstTime := utc.In(pst)
		fmt.Printf("PST: %s\n", pstTime.Format("15:04:05 MST"))
	}
}