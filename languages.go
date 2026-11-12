package babel

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"
	"unicode"
	"unicode/utf8"
)

type (
	// Script is a ISO 15924 script code.
	//
	// https://en.wikipedia.org/wiki/ISO_15924
	Script string
	// Language is a ISO 639-1 language code.
	//
	// https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes
	Language string
)

// Name of this script, in English.
func (s Script) Name() string {
	if n, ok := scriptNames[s]; ok {
		return n
	}
	return string(s)
}

// Name of this language, in English.
func (s Language) Name() string {
	if n, ok := languageNames[s]; ok {
		return n
	}
	return string(s)
}

// Script constants.
const (
	ScriptArabic     = Script("Arab")
	ScriptArmenian   = Script("Armn")
	ScriptBengali    = Script("Beng")
	ScriptCyrillic   = Script("Cyrl")
	ScriptDevanagari = Script("Deva")
	ScriptEthiopic   = Script("Ethi")
	ScriptGeorgian   = Script("Geor")
	ScriptGreek      = Script("Grek")
	ScriptGujarati   = Script("Gujr")
	ScriptGurmukhi   = Script("Guru")
	ScriptHan        = Script("Hani")
	ScriptHangul     = Script("Hang")
	ScriptHebrew     = Script("Hebr")
	ScriptHiragana   = Script("Hira")
	ScriptKannada    = Script("Knda")
	ScriptKatakana   = Script("Kana")
	ScriptKhmer      = Script("Khmr")
	ScriptLatin      = Script("Latn")
	ScriptMalayalam  = Script("Mlym")
	ScriptMyanmar    = Script("Mymr")
	ScriptOriya      = Script("Orya")
	ScriptSinhala    = Script("Sinh")
	ScriptTamil      = Script("Taml")
	ScriptTelugu     = Script("Telu")
	ScriptThai       = Script("Thai")
	ScriptUnknown    = Script("")
)

const (
	LanguageArabic     = Language("ar")
	LanguageArmenian   = Language("hy")
	LanguageChinese    = Language("zh")
	LanguageDutch      = Language("nl")
	LanguageEnglish    = Language("en")
	LanguageFrench     = Language("fr")
	LanguageGerman     = Language("de")
	LanguageGreek      = Language("el")
	LanguageGujarati   = Language("gu")
	LanguageHebrew     = Language("he")
	LanguageHungarian  = Language("hu")
	LanguageItalian    = Language("it")
	LanguageJapanese   = Language("ja")
	LanguageKannada    = Language("kn")
	LanguageKorean     = Language("ko")
	LanguagePolish     = Language("pl")
	LanguagePortuguese = Language("pt")
	LanguagePunjabi    = Language("pa")
	LanguageRussian    = Language("ru")
	LanguageSerbian    = Language("sr")
	LanguageSpanish    = Language("es")
	LanguageTagalog    = Language("tl")
	LanguageTamil      = Language("ta")
	LanguageTelugu     = Language("te")
	LanguageThai       = Language("th")
	LanguageUkrainian  = Language("uk")
	LanguageUnknown    = Language("")
	LanguageVietnamese = Language("vi")
)

//go:embed languages/*
var languageFiles embed.FS

var languages = func() map[Script]map[Language][]string {
	ls, err := languageFiles.ReadDir("languages")
	if err != nil {
		panic(err)
	}

	langs := make(map[Script]map[Language][]string)
	for _, f := range ls {
		scriptCode, langCode, ok := strings.Cut(f.Name(), "-")
		if !ok {
			panic(fmt.Sprintf("not as Script-languageCode: %q", "languages/"+f.Name()))
		}
		d, err := fs.ReadFile(languageFiles, "languages/"+f.Name())
		if err != nil {
			panic(err)
		}
		terms := strings.Fields(string(d))
		for i := range terms {
			if utf8.RuneCountInString(terms[i]) != 3 {
				panic(fmt.Sprintf("%q: not three characters: %q", f.Name(), terms[i]))
			}
			terms[i] = strings.ReplaceAll(terms[i], "-", " ")
		}
		if langs[Script(scriptCode)] == nil {
			langs[Script(scriptCode)] = make(map[Language][]string)
		}
		langs[Script(scriptCode)][Language(langCode)] = terms
	}
	return langs
}()

var oneLanguageScripts = map[Script]Language{
	ScriptArabic:   LanguageArabic,
	ScriptArmenian: LanguageArmenian,
	ScriptGreek:    LanguageGreek,
	ScriptGujarati: LanguageGujarati,
	ScriptGurmukhi: LanguagePunjabi,
	ScriptHan:      LanguageChinese,
	ScriptHangul:   LanguageKorean,
	ScriptHebrew:   LanguageHebrew,
	ScriptHiragana: LanguageJapanese,
	ScriptKannada:  LanguageKannada,
	ScriptKatakana: LanguageJapanese,
	ScriptTamil:    LanguageTamil,
	ScriptTelugu:   LanguageTelugu,
	ScriptThai:     LanguageThai,
}

var scripts = []struct {
	s   Script
	tbl *unicode.RangeTable
}{
	{ScriptLatin, unicode.Latin},
	{ScriptCyrillic, unicode.Cyrillic},
	{ScriptArabic, unicode.Arabic},
	{ScriptDevanagari, unicode.Devanagari},
	{ScriptHebrew, unicode.Hebrew},
	{ScriptBengali, unicode.Bengali},
	{ScriptHan, unicode.Han},
	{ScriptKatakana, unicode.Katakana},
	{ScriptHiragana, unicode.Hiragana},
	{ScriptGeorgian, unicode.Georgian},
	{ScriptHangul, unicode.Hangul},
	{ScriptGreek, unicode.Greek},
	{ScriptKannada, unicode.Kannada},
	{ScriptTamil, unicode.Tamil},
	{ScriptThai, unicode.Thai},
	{ScriptGujarati, unicode.Gujarati},
	{ScriptGurmukhi, unicode.Gurmukhi},
	{ScriptTelugu, unicode.Telugu},
	{ScriptMalayalam, unicode.Malayalam},
	{ScriptOriya, unicode.Oriya},
	{ScriptMyanmar, unicode.Myanmar},
	{ScriptSinhala, unicode.Sinhala},
	{ScriptKhmer, unicode.Khmer},
	{ScriptEthiopic, unicode.Ethiopic},
	{ScriptArmenian, unicode.Armenian},
}

var scriptNames = map[Script]string{
	ScriptArabic:     "Arabic",
	ScriptArmenian:   "Armenian",
	ScriptBengali:    "Bengali",
	ScriptCyrillic:   "Cyrillic",
	ScriptDevanagari: "Devanagari",
	ScriptEthiopic:   "Ethiopic",
	ScriptGeorgian:   "Georgian",
	ScriptGreek:      "Greek",
	ScriptGujarati:   "Gujarati",
	ScriptGurmukhi:   "Gurmukhi",
	ScriptHan:        "Han",
	ScriptHangul:     "Hangul",
	ScriptHebrew:     "Hebrew",
	ScriptHiragana:   "Hiragana",
	ScriptKannada:    "Kannada",
	ScriptKatakana:   "Katakana",
	ScriptKhmer:      "Khmer",
	ScriptLatin:      "Latin",
	ScriptMalayalam:  "Malayalam",
	ScriptMyanmar:    "Myanmar",
	ScriptOriya:      "Oriya",
	ScriptSinhala:    "Sinhala",
	ScriptTamil:      "Tamil",
	ScriptTelugu:     "Telugu",
	ScriptThai:       "Thai",
	ScriptUnknown:    "(Unknown Script)",
}

var languageNames = map[Language]string{
	LanguageArabic:     "Arabic",
	LanguageArmenian:   "Armenian",
	LanguageChinese:    "Chinese",
	LanguageDutch:      "Dutch",
	LanguageEnglish:    "English",
	LanguageFrench:     "French",
	LanguageGerman:     "German",
	LanguageGreek:      "Greek",
	LanguageGujarati:   "Gujarati",
	LanguageHebrew:     "Hebrew",
	LanguageHungarian:  "Hungarian",
	LanguageItalian:    "Italian",
	LanguageJapanese:   "Japanese",
	LanguageKannada:    "Kannada",
	LanguageKorean:     "Korean",
	LanguagePolish:     "Polish",
	LanguagePortuguese: "Portuguese",
	LanguagePunjabi:    "Punjabi",
	LanguageRussian:    "Russian",
	LanguageSerbian:    "Serbian",
	LanguageSpanish:    "Spanish",
	LanguageTagalog:    "Tagalog",
	LanguageTamil:      "Tamil",
	LanguageTelugu:     "Telugu",
	LanguageThai:       "Thai",
	LanguageUkrainian:  "Ukrainian",
	LanguageUnknown:    "(Unknown Language)",
	LanguageVietnamese: "Vietnamese",
}
