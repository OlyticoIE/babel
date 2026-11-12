// Package babel provides fast natural language detection.
//
// babel compares the text to a characteristic profile of each supported
// language and returns the language that best matches the text.
package babel

import (
	"cmp"
	"math"
	"slices"
	"sort"
	"unicode"
)

// Detect the script and language, returns the script and language codes and
// confidence as [0, 1.0] inclusive.
func Detect[T string | []byte](text T) (Script, Language, float64) {
	var t string
	switch tt := any(text).(type) {
	case string:
		t = tt
	case []byte:
		t = string(tt)
	}

	script, conf := detectScript(t)
	if script == ScriptUnknown {
		return ScriptUnknown, LanguageUnknown, 1
	}

	scriptLangs := languages[script]
	// Scripts with only one language: don't need to do anything more here.
	if len(scriptLangs) == 0 {
		return script, oneLanguageScripts[script], conf
	}

	matches := map[Language]int{LanguageUnknown: 1}
	trigs := sortedTrigs(t)
	for k, v := range scriptLangs {
		matchWith(k, trigs, v, matches)
	}

	smx := softMax(matches)
	lang := maxKey(matches)
	return script, lang, smx[lang]
}

func softMax(mapping map[Language]int) map[Language]float64 {
	const (
		rescale     = 0.5
		expOverflow = 7.09e+02
	)
	var (
		softMaxMap = make(map[Language]float64)
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

func maxKey(mapping map[Language]int) Language {
	var (
		max int
		key Language
	)
	for k, v := range mapping {
		if v > max {
			max, key = v, k
		}
	}
	return key
}

func matchWith(lang Language, trigs []trigram, langProfile []string, matches map[Language]int) {
	prof := make(map[string]int)
	for _, x := range langProfile {
		prof[x] = 1
	}

	const undeterminedRate = 41
	var undeterminedCount int
	for _, trig := range trigs {
		if _, exists := prof[trig.trigram]; exists {
			matches[lang] += trig.count
		} else {
			undeterminedCount++
			if (undeterminedCount % undeterminedRate) == 0 {
				matches[LanguageUnknown]++
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

func detectScript(t string) (Script, float64) {
	match := make(map[Script]int)
	for _, r := range t {
		if unicode.IsSymbol(r) || unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsDigit(r) {
			continue
		}
		for _, s := range scripts {
			if unicode.Is(s.tbl, r) {
				match[s.s]++
				break
			}
		}
	}
	switch len(match) {
	case 0:
		return ScriptUnknown, 1
	case 1:
		for k := range match {
			return k, 1
		}
	}

	type m struct {
		s Script
		n int
	}
	sorted := make([]m, 0, len(match))
	var total int
	for s, n := range match {
		sorted = append(sorted, m{s, n})
		total += n
	}
	slices.SortFunc(sorted, func(a, b m) int { return cmp.Compare(b.n, a.n) })

	if sorted[0].n == sorted[1].n {
		return ScriptUnknown, 0
	}
	return sorted[0].s, float64(sorted[0].n)/float64(total) + 0.15
}
