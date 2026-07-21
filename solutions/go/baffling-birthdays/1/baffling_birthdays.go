package bafflingbirthdays

import (
    "crypto/rand"
    "fmt"
    "math"
    "math/big"
    "time"
)

func SharedBirthday(dates []time.Time) bool {
	bdRef := map[string]struct{}{}
    ret := false

    for _, date := range dates {
        _, month, day := date.Date()

        key := fmt.Sprintf("%s-%d", month, day)
        
        if _, exists := bdRef[key]; exists {
            ret = true
        } else if !exists {
            bdRef[key] = struct{}{}
        }
    }

    return ret
}

func RandomBirthdates(size int) []time.Time {
    randBds := []time.Time{}
    
    for size > 0 {
        month, monthErr := generateRandomMonth()
        if monthErr != nil {
            panic("Error generating random month")
        }

        var date int
        var year int

        if month == 2 {
            for {
                if randDate, dateErr := generateRandomDate(); dateErr != nil {
                    panic("Error generating random date")
                } else if randDate <= 28 {
                    date = randDate
                    break
                }
            }
        } else {
            randDate, dateErr := generateRandomDate()
            
            if dateErr != nil {
                panic("Error generating random date")
            }
            
            date = randDate
        }

        for {
            if randYear, yearErr := generateRandomYear(); yearErr != nil {
                panic("Error generating random year")
            } else if !isLeap(randYear) {
                year = randYear
                break
            }
        }

        randBds = append(randBds, time.Date(year, time.Month(month), date, 0, 0, 0, 0, time.UTC))
        
        size --
    }

    return randBds
}

func generateRandomYear() (int, error) {
    maxYearNum := int64(2300)
    minYearNum := int64(1700)
    yearRange := big.NewInt(maxYearNum - minYearNum)

    randYear, err := rand.Int(rand.Reader, yearRange)

    if err != nil {
        return 0, err
    }

    return int(randYear.Int64()) + int(minYearNum), nil
}

func generateRandomMonth() (int, error) {
    maxMonthOff := int64(12)
    minMonthOff := int64(1)
    monthRange := big.NewInt(maxMonthOff - minMonthOff)

    randMonth, err := rand.Int(rand.Reader, monthRange)

    if err != nil {
        return 0, err
    }

    return int(randMonth.Int64()) + int(minMonthOff), nil
}

func generateRandomDate() (int, error) {
    maxMonthDate := int64(31)
    minMonthDate := int64(1)
    dateRange := big.NewInt(maxMonthDate - minMonthDate)

    randDate, err := rand.Int(rand.Reader, dateRange)

    if err != nil {
        return 0, err
    }

    return int(randDate.Int64()) + int(minMonthDate), nil
}

func isLeap(year int) bool {
    if year % 4 == 0 {
        if year % 100 == 0 {
            return year % 400 == 0;
        }
        
        return true;
    }
    
    return false;
}

func EstimatedProbability(size int) float64 {
	sizeFloat := float64(size)
    prod := float64(1)
    numDaysInYear := float64(365)
	
    for i := float64(0); i < sizeFloat; i += float64(1) {
        prod *= numDaysInYear - i
    }

    return (float64(1) - (prod / math.Pow(numDaysInYear, sizeFloat))) * 100
}
