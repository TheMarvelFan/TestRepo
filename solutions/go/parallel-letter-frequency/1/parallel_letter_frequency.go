package parallelletterfrequency

import (
	"sync"
	"unicode"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	runeTextArr := []rune(text)
	for _, r := range runeTextArr {
		if unicode.IsLetter(r) {
			frequencies[unicode.ToLower(r)]++	
		}
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	retFreqMap := FreqMap{}
	wg := new(sync.WaitGroup)
	writeMutex := sync.Mutex{}
	for _, text := range texts {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			freqMap := Frequency(text)
			writeMutex.Lock()
			for r, count := range freqMap {
				retFreqMap[r] += count
			}
			writeMutex.Unlock()
		}(text)
	}
	wg.Wait()
	return retFreqMap
}
