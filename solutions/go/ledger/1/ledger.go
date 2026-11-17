package ledger

import (
    "cmp"
	"errors"
    "slices"
	"strconv"
	"strings"
    "unicode/utf8"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

var meta = map[string][]string{
    "en-US": {"Date", "Description", "Change", "/", ",", "."},
    "nl-NL": {"Datum", "Omschrijving", "Verandering", "-", ".",","},
}

var ccySym = map[string]string{
    "USD": "$",
    "EUR": "€",
}

func FormatLedger(currency, locale string, entries []Entry) (string, error) {
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

    entriesCopy := make([]Entry, len(entries))
    copy(entriesCopy, entries)

	slices.SortFunc(entriesCopy, func(a, b Entry) int {
        if a.Date != b.Date {
            return strings.Compare(a.Date, b.Date)
        }
    
        if a.Description != b.Description {
            return strings.Compare(a.Description, b.Description)
        }
    
        return cmp.Compare(a.Change, b.Change)
    })

	s, parsedOk := buildMetadata(locale)
	if parsedOk != nil {
		return s, parsedOk
    }

    var ret string
	ret += s
	for _, entry := range entriesCopy {
		formattedEntry, ok := formatEntry(entry, locale, currency)
        if ok != nil {
            return "", errors.New("")
        }

		ret += formattedEntry
	}

    return ret, nil
}

func formatEntry(entry Entry, locale, currency string) (string, error) {
    if len(entry.Date) != 10 {
        return "", errors.New("")
    }

    d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]

    if d2 != '-' || d4 != '-' {
        return "", errors.New("")
    }

    de := entry.Description
    if len(de) > 25 {
        de = de[:22] + "..."
    } else {
        de = de + strings.Repeat(" ", 25-len(de))
    }

    d, dateOk := buildDateForLocale(locale, d5, d3, d1)
    if dateOk != nil {
        return "", dateOk
    }

    a, err := buildNumber(currency, locale, entry.Change)
    if err != nil {
        return "", err
    }

    padding := 13 - utf8.RuneCountInString(a)
    if padding < 0 {
        padding = 0
    }

    return d +
        strings.Repeat(" ", 10-len(d)) +
        " | " + de +
        " | " + strings.Repeat(" ", padding) +
        a + "\n", nil
}

func buildMetadata(locale string) (string, error) {
    words, wordsExistErr := getLocaleData(locale)
    if wordsExistErr != nil {
        return "", wordsExistErr
    }
    
    s := words[0] +
			strings.Repeat(" ", 10-len(words[0])) +
			" | " +
			words[1] +
			strings.Repeat(" ", 25-len(words[1])) +
			" | " + words[2] + strings.Repeat(" ", 13-len(words[2])) + "\n"

    return s, nil
}

func buildDateForLocale(locale, day, month, year string) (string, error) {
    words, wordsExistErr := getLocaleData(locale)
    if wordsExistErr != nil {
        return "", wordsExistErr
    }

    var d string
    if locale == "nl-NL" {
    	d = day + words[3] + month + words[3] + year
    } else if locale == "en-US" {
        d = month + words[3] + day + words[3] + year
    } else {
        return "", errors.New("")
    }

    return d, nil
}

func getLocaleData(locale string) ([]string, error) {
    if locale == "" {
        return nil, errors.New("")
    }

    words, wordsExist := meta[locale]
    if !wordsExist {
        return nil, errors.New("")
    }

    return words, nil
}

func buildNumber(currency, locale string, cents int) (string, error) {
    var a string
    var negative bool
    
    if cents < 0 {
        cents = cents * -1
        negative = true
    }

    if symbol, ccyCodeExists := ccySym[currency]; !ccyCodeExists {
        return "", errors.New("")
    } else {
        a += symbol
    }

    if locale == "nl-NL" {
        a += " "
        if negative {
        	a += "-"
        }
    }

    words, wordsExistErr := getLocaleData(locale)
    if wordsExistErr != nil {
        return "", wordsExistErr
    }

    rest := []rune(strconv.Itoa(cents / 100))
    for i := 0; i < len(rest); i++ {
        if i > 0 && (len(rest) - i) % 3 == 0 {
            a += words[4]
        }

        a += string(rest[i])
    }

    a += words[5]
    centsStr := strconv.Itoa(cents % 100)
    if len(centsStr) < 2 {
        a += "0"
    }
    a += centsStr

    if negative && locale == "en-US" {
        a = "(" + a + ")"
    } else {
        a += " "
    }
    
    return a, nil
}
