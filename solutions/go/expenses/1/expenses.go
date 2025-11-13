package expenses

import "fmt"
import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	ret := []Record{}

    for _, record := range in {
        if predicate(record) {
            ret = append(ret, record)
        }
    }

    return ret
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	from := p.From
    to := p.To

    return func(r Record) bool {
        return r.Day <= to && r.Day >= from
    }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
        return r.Category == c
    }
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	isInRange := ByDaysPeriod(p)
    sum := 0.0
    
	for _, record := range in {
        if isInRange(record) {
            sum += record.Amount
        }
    }

    return sum
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	isOfCategory := ByCategory(c)
    isInRange := ByDaysPeriod(p)
    
    sum := 0.0
    atLeastOneInCategory := false
    
	for _, record := range in {
        if isOfCategory(record) {
            atLeastOneInCategory = true
        	if isInRange(record) {
            	sum += record.Amount
            }
        }
    }

    if !atLeastOneInCategory {
        return 0, errors.New(fmt.Sprintf("unknown category %s", c))
    }
    
    return sum, nil
}
