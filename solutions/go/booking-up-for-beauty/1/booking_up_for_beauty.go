package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    timeFormat := "1/2/2006 15:04:05"
	t, err := time.Parse(timeFormat, date)

    if err != nil {
        panic(err)
    }
    
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    timeFormat := "January 2, 2006 15:04:05"
	appt, err := time.Parse(timeFormat, date)

    if err != nil {
        panic(err)
    }

    return appt.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    timeFormat := "Monday, January 2, 2006 15:04:05"
    hour, err := time.Parse(timeFormat, date)

    if err != nil {
        panic(err)
    }
	
    return hour.Hour() >= 12 && hour.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    timeFormat := "1/2/2006 15:04:05"
    appt, err := time.Parse(timeFormat, date)

    if err != nil {
        panic(err)
    }
    
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%02d.", appt.Weekday(), appt.Month(), appt.Day(), appt.Year(), appt.Hour(), appt.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
