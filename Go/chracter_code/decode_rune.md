# utf8.DecodeRune の注意点

> DecodeRune unpacks the first UTF-8 encoding in p and returns the rune and its width in bytes.

引用元: [utf8 - The Go Programming Language](https://golang.org/pkg/unicode/utf8/#DecodeRune)

とあるように、普通にやると引数で与えた []byte の最初の 最初のUTF-8エンコード(1文字)を対象にする。

なので、以下のように実装すると

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("アイウエオ")
	r, n := utf8.DecodeRune(b)
	fmt.Printf("r = %v, n = %d\n", r, n)

	fmt.Printf("b[:n] の文字列 = %s", b[:n])
}

```

[Go Playground でのコード](https://play.golang.org/p/KH6gN1QKxMm)


実行結果は以下のようになる。

```
r = 12450, n = 3
b[:n] の文字列 = ア
```

## 参考

[utf8 - The Go Programming Language](https://golang.org/pkg/unicode/utf8/#DecodeRune)
