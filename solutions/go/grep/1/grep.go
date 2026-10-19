package grep

import (
    "maps"
    "slices"
    "strconv"
    "strings"
)

var fileContents = map[string]string{
    "iliad.txt": `Achilles sing, O Goddess! Peleus' son;
His wrath pernicious, who ten thousand woes
Caused to Achaia's host, sent many a soul
Illustrious into Ades premature,
And Heroes gave (so stood the will of Jove)
To dogs and to all ravening fowls a prey,
When fierce dispute had separated once
The noble Chief Achilles from the son
Of Atreus, Agamemnon, King of men.`,
    "midsummer-night.txt": `I do entreat your grace to pardon me.
I know not by what power I am made bold,
Nor how it may concern my modesty,
In such a presence here to plead my thoughts;
But I beseech your grace that I may know
The worst that may befall me in this case,
If I refuse to wed Demetrius.`,
    "paradise-lost.txt": `Of Mans First Disobedience, and the Fruit
Of that Forbidden Tree, whose mortal tast
Brought Death into the World, and all our woe,
With loss of Eden, till one greater Man
Restore us, and regain the blissful Seat,
Sing Heav'nly Muse, that on the secret top
Of Oreb, or of Sinai, didst inspire
That Shepherd, who first taught the chosen Seed`,
}

const colon = ":"

func Search(pattern string, flags, files []string) []string {
	if len(files) == 0 || len(pattern) == 0 {
        return nil
    }
    
    ret := map[string][]string{}
	
    for _, file := range files {
        checkPatternPresence(flags, &ret, len(files), pattern, file, fileContents[file])
    }

    return processResults(&ret)
}

func checkPatternPresence(flags []string, ret *map[string][]string, filesInScope int, pattern, file, content string) {
    lines := strings.Split(content, "\n")
    matchFullLine := slices.Contains(flags, "-x")
    matchIgnoreCase := slices.Contains(flags, "-i")
    invertMatchCriteria := slices.Contains(flags, "-v")
    filenamesOnly := slices.Contains(flags, "-l")
    includeLineNumbers := slices.Contains(flags, "-n")

    for lineNum, line := range lines {
        if matchFullLine {
        	processMatchFullLine(ret, line, pattern, file, filesInScope, lineNum, matchIgnoreCase, invertMatchCriteria, filenamesOnly, includeLineNumbers)
        } else {
            if matchIgnoreCase {
                processMatchIgnoreCase(ret, line, pattern, file, filesInScope, lineNum, invertMatchCriteria, filenamesOnly, includeLineNumbers, false)
            } else {
                processMatchRegularCase(ret, line, pattern, file, filesInScope, lineNum, invertMatchCriteria, filenamesOnly, includeLineNumbers)
            }
        }
	}
}

func processMatchFullLine(ret *map[string][]string, line, pattern, file string, filesInScope, lineNum int, matchIgnoreCase, invertMatchCriteria, filenamesOnly, includeLineNumbers bool) {
    if matchIgnoreCase {
        processMatchIgnoreCase(ret, line, pattern, file, filesInScope, lineNum, invertMatchCriteria, filenamesOnly, includeLineNumbers, true)
    } else if invertMatchCriteria {
            processInvertedMatchCriteria(ret, line, pattern, file, filesInScope, lineNum, filenamesOnly, includeLineNumbers, true)
    } else if line == pattern {
        appendFilenameOrLine(ret, line, file, lineNum, filesInScope, includeLineNumbers, filenamesOnly)
    }
}

func processMatchIgnoreCase(ret *map[string][]string, line, pattern, file string, filesInScope, lineNum int, invertMatchCriteria, filenamesOnly, includeLineNumbers, matchFullLine bool) {
    lCaseLine := strings.ToLower(line)
    lCasePat := strings.ToLower(pattern)

    if matchFullLine && ((!invertMatchCriteria && lCaseLine == lCasePat) || (invertMatchCriteria && lCaseLine != lCasePat)) {
        appendFilenameOrLine(ret, line, file, lineNum, filesInScope, includeLineNumbers, filenamesOnly)       
    } else if !matchFullLine && ((!invertMatchCriteria && strings.Contains(lCaseLine, lCasePat)) || (invertMatchCriteria && !strings.Contains(lCaseLine, lCasePat))) {
    	appendFilenameOrLine(ret, line, file, lineNum, filesInScope, includeLineNumbers, filenamesOnly)
    }
}

func processMatchRegularCase(ret *map[string][]string, line, pattern, file string, filesInScope , lineNum int, invertMatchCriteria, filenamesOnly, includeLineNumbers bool) {
    if invertMatchCriteria {
        processInvertedMatchCriteria(ret, line, pattern, file, filesInScope, lineNum, filenamesOnly, includeLineNumbers, false)
    } else if strings.Contains(line, pattern) {
        appendFilenameOrLine(ret, line, file, lineNum, filesInScope, includeLineNumbers, filenamesOnly)
    }
}

func processInvertedMatchCriteria(ret *map[string][]string, line, pattern, file string, filesInScope, lineNum int, filenamesOnly, includeLineNumbers, matchFullLine bool) {
    if (matchFullLine && line != pattern) || !(matchFullLine || strings.Contains(line, pattern)) {
        appendFilenameOrLine(ret, line, file, lineNum, filesInScope, includeLineNumbers, filenamesOnly)
    }
}

func appendFilenameOrLine(ret *map[string][]string, line, file string, lineNum, filesInScope int, includeLineNumbers, filenamesOnly bool) {
	if _, fileAlreadyExists := (*ret)[file]; filenamesOnly && !fileAlreadyExists {
        (*ret)[file] = []string{file}
        return
    }
    
    if filenamesOnly {
        return
    }
    
	pref := ""
    
    if filesInScope > 1 {
        pref = file + colon
    }

    if includeLineNumbers {
        (*ret)[file] = append((*ret)[file], pref + strconv.Itoa(lineNum + 1) + colon + strings.TrimSpace(line))
    } else {
        (*ret)[file] = append((*ret)[file], pref + strings.TrimSpace(line))
    }
}

func processResults(results *map[string][]string) []string {
	if len(*results) == 0 {
        return nil
    }

    sortedFiles := slices.Sorted(maps.Keys(*results))
    
	ret := []string{}

    for _, file := range sortedFiles {
        ret = append(ret, (*results)[file]...)
    }

    return ret
}
