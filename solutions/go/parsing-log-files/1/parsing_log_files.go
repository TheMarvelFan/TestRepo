package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)

    if err != nil {
        panic(err)
    }

    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<[-*~=]*>`)

    if err != nil {
        panic(err)
    }

    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	pattern := `"[^"]*password[^"]*"`
    re, err := regexp.Compile(`(?i)` + pattern)

    if err != nil {
        panic(err)
    }
    
    count := 0
    for _, line := range lines {
        if re.MatchString(line) {
            count ++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end-of-line\d+`)

    if err != nil {
        panic(err)
    }

    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	pattern := `User\s+(\S+)`
    re, err := regexp.Compile(pattern)

    if err != nil {
        panic(err)
    }
    
    result := make([]string, len(lines))
    
    for i, line := range lines {
        matches := re.FindStringSubmatch(line)
        if matches != nil {
            username := matches[1]
            result[i] = "[USR] " + username + " " + line
        } else {
            result[i] = line
        }
    }
    
    return result
}
