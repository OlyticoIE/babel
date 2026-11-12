package babel

import (
	"embed"
	"io/fs"
	"strings"
)

//go:embed languages/*
var languageFiles embed.FS

var langs = func() map[string][]string {
	ls, err := languageFiles.ReadDir("languages")
	if err != nil {
		panic(err)
	}

	langs := make(map[string][]string)
	for _, f := range ls {
		d, err := fs.ReadFile(languageFiles, "languages/"+f.Name())
		if err != nil {
			panic(err)
		}
		terms := strings.Fields(string(d))
		for i := range terms {
			terms[i] = strings.ReplaceAll(terms[i], "-", " ")
		}
		langs[f.Name()] = terms
	}
	return langs
}()
