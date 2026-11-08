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

func main(){
    info := getlang.FromString("Wszyscy ludzie rodzą się wolni i równi w swojej godności i prawach")
    fmt.Println(info.LanguageCode(), info.Confidence())
}
```

Acknowledgements and Citations
------------------------------
* Thanks to [abadojack](https://github.com/abadojack) for the trigram generation logic in whatlanggo
* Cavnar, William B., and John M. Trenkle. "N-gram-based text categorization." Ann arbor mi 48113.2 (1994): 161-175.
