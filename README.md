getlang provides fast natural language detection in Go.

- Offline -- no internet connection required
- Supports 29 languages
- Provides ISO 639 language codes
- Fast

Usage
-----

```go
package main

import (
	"fmt"

	"go.olytico.com/getlang"
)

func main() {
	var (
		shortLang, shortConf = getlang.Detect("short text")
		longLang, longConf   = getlang.Detect("this sentence is a bit longer")
	)

	fmt.Println(shortLang, shortConf) // en 0.15062092062586865
	fmt.Println(longLang, longConf)   // en 0.747458096413626
}
```

Supported Languages
-------------------

| Language       | ISO 639-1 |
| -------------- | --------- |
| Arabic         | ar        |
| Bengali (Bangla) | bn      |
| German         | de        |
| Greek          | el        |
| English        | en        |
| Spanish        | es        |
| French         | fr        |
| Hebrew         | he        |
| Hindi          | hi        |
| Hungarian      | hu        |
| Armenian       | hy        |
| Gujarati       | gu        |
| Italian        | it        |
| Dutch          | nl        |
| Polish         | pl        |
| Portuguese     | pt        |
| Punjabi        | pa        |
| Japanese       | ja        |
| Kannada        | ka        |
| Korean         | ko        |
| Tamil		     | ta        |
| Telugu         | te        |
| Tagalog        | tl        |
| Thai           | th        |
| Russian        | ru        |
| Serbian        | sr        |
| Vietnamese     | vi        |
| Ukrainian      | uk        |
| Chinese        | zh        |

Acknowledgements and Citations
------------------------------
- This is forked from https://github.com/rylans/getlang
- Thanks to https://github.com/abadojack for the trigram generation logic in whatlanggo
- Cavnar, William B., and John M. Trenkle. "N-gram-based text categorization." Ann arbor mi 48113.2 (1994): 161-175.
