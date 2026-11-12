babel provides fast natural language detection in Go.

Usage
-----
```go
package main

import (
	"fmt"

	"go.olytico.com/babel"
)

func main() {
	var (
		script1, lang1, conf1 = babel.Detect("short text")
		script2, lang2, conf2 = babel.Detect("this sentence is a bit longer")
	)

	fmt.Println(script1, lang1, conf1) // Latn en 0.15062092062586865
	fmt.Println(script2, lang2, conf2) // Latn en 0.747458096413626
}
```

Supported Languages
-------------------

| Language         | ISO 639-1 |
| ---------------- | --------- |
| Arabic           | ar        |
| Armenian         | hy        |
| Bengali (Bangla) | bn        |
| Chinese          | zh        |
| Dutch            | nl        |
| English          | en        |
| French           | fr        |
| German           | de        |
| Greek            | el        |
| Gujarati         | gu        |
| Hebrew           | he        |
| Hindi            | hi        |
| Hungarian        | hu        |
| Italian          | it        |
| Japanese         | ja        |
| Kannada          | ka        |
| Korean           | ko        |
| Polish           | pl        |
| Portuguese       | pt        |
| Punjabi          | pa        |
| Russian          | ru        |
| Serbian          | sr        |
| Spanish          | es        |
| Tagalog          | tl        |
| Tamil		       | ta        |
| Telugu           | te        |
| Thai             | th        |
| Ukrainian        | uk        |
| Vietnamese       | vi        |

Acknowledgements and Citations
------------------------------
- This is forked from https://github.com/rylans/getlang
- Thanks to https://github.com/abadojack for the trigram generation logic in whatlanggo
- Cavnar, William B., and John M. Trenkle. "N-gram-based text categorization." Ann arbor mi 48113.2 (1994): 161-175.
