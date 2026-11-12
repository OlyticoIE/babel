// Package babel provides fast natural language detection.
//
// babel compares the text to a characteristic profile of each supported
// language and returns the language that best matches the text.
package babel

import (
	"math"
	"sort"
	"strings"
	"unicode"
)

var scripts = map[string][]*unicode.RangeTable{
	"ar": {unicode.Arabic},
	"bn": {unicode.Bengali},
	"el": {unicode.Greek},
	"gu": {unicode.Gujarati},
	"he": {unicode.Hebrew},
	"hy": {unicode.Armenian},
	"ja": {unicode.Hiragana, unicode.Katakana},
	"kn": {unicode.Kannada},
	"ko": {unicode.Hangul},
	"pa": {unicode.Gurmukhi},
	"ta": {unicode.Tamil},
	"te": {unicode.Telugu},
	"th": {unicode.Thai},
	"zh": {unicode.Han},
}

// Detetect the language, returns the ISO 639-1 code and confidence as [0, 1.0]
// inclusive.
//
// [golang.org/x/text/language.Parse] can be used to get the language name from
// the ISO code.
func Detect[T string | []byte](text T) (string, float64) {
	var t string
	switch tt := any(text).(type) {
	case string:
		t = tt
	case []byte:
		t = string(tt)
	}

	matches := map[string]int{"und": 1}
	trigs := sortedTrigs(t)
	for k, v := range langs {
		matchWith(k, trigs, v, matches)
	}
	for k, v := range scripts {
		matchScript(k, t, matches, v...)
	}

	smx := softMax(matches)
	lang := maxKey(matches)
	langCode, _, _ := strings.Cut(lang, "-")
	return langCode, smx[lang]
}

func softMax(mapping map[string]int) map[string]float64 {
	const (
		rescale     = 0.5
		expOverflow = 7.09e+02
	)
	var (
		softMaxMap = make(map[string]float64)
		denom      float64
		overflowed = false
	)
	for _, v := range mapping {
		denom += math.Exp(float64(v) * rescale)
		if v > expOverflow {
			overflowed = true
		}
	}
	for k := range mapping {
		if !overflowed {
			softMaxMap[k] = math.Exp(rescale*float64(mapping[k])) / denom
		} else {
			softMaxMap[k] = 1.0
		}
	}
	return softMaxMap
}

func maxKey(mapping map[string]int) string {
	var (
		max int
		key string
	)
	for k, v := range mapping {
		if v > max {
			max, key = v, k
		}
	}
	return key
}

func matchScript(langName, text string, matches map[string]int, ranges ...*unicode.RangeTable) {
	for _, r := range text {
		if unicode.In(r, ranges...) {
			matches[langName] += 2
		}
	}
}

func matchWith(langName string, trigs []trigram, langProfile []string, matches map[string]int) {
	prof := make(map[string]int)
	for _, x := range langProfile {
		prof[x] = 1
	}

	const undeterminedRate = 41
	var undeterminedCount int
	for _, trig := range trigs {
		if _, exists := prof[trig.trigram]; exists {
			matches[langName] += trig.count
		} else {
			undeterminedCount++
			if (undeterminedCount % undeterminedRate) == 0 {
				matches["und"]++
			}
		}
	}
}

func countedTrigrams(text string) map[string]int {
	var (
		trigrams = map[string]int{}
		txt      = make([]rune, 0, len(text)+1)
	)
	for _, r := range text {
		txt = append(txt, unicode.ToLower(toTrigramChar(r)))
	}
	txt = append(txt, ' ')

	var (
		r1 = ' '
		r2 = txt[0]
		r3 rune
	)
	for i := 1; i < len(txt); i++ {
		r3 = txt[i]
		if !(r2 == ' ' && (r1 == ' ' || r3 == ' ')) {
			trigram := []rune{r1, r2, r3}
			if trigrams[string(trigram)] == 0 {
				trigrams[string(trigram)] = 1
			} else {
				trigrams[string(trigram)]++
			}
		}
		r1, r2 = r2, r3
	}

	return trigrams
}

type trigram struct {
	trigram string
	count   int
}

func sortedTrigs(s string) []trigram {
	var (
		counterMap = countedTrigrams(s)
		trigrams   = make([]trigram, len(counterMap))
		i          int
	)
	for tg, count := range counterMap {
		trigrams[i] = trigram{tg, count}
		i++
	}
	sort.SliceStable(trigrams, func(i, j int) bool {
		if trigrams[i].count == trigrams[j].count {
			return trigrams[i].trigram < trigrams[j].trigram
		}
		return trigrams[i].count > trigrams[j].count
	})
	return trigrams
}

func toTrigramChar(ch rune) rune {
	if unicode.IsPunct(ch) || unicode.IsSpace(ch) {
		return ' '
	}
	return ch
}
