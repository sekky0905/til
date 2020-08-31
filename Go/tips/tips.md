# Go の Tips

## iota を使用する際には、1 を初期値にした方が良さそう


## switch 文の fallthrough

- Go に限った話ではないが、次のケースの処理をさせたい時に `fallthrough` という技がある
    - ただ分かりにくいし、バグの温床になる可能性が高いので、知っておくだけであまり使わない方がいいかも

### コード

- [実際のコード](https://play.golang.org/p/dYXdytUCIXq)

```go
package main

import (
	"fmt"
)

func main() {
	x1 := judge("A")
	fmt.Println(x1)

	x2 := judge("B")
	fmt.Println(x2)

	x3 := judge("D")
	fmt.Println(x3)
}

func judge(lang string) string {
	switch lang {
	case "A":
		fallthrough
	case "B":
		return "b"
	default:
		return "c"
	}
}
```

### 実行結果

```
b
b
c
```

参考: [switch文 - Wikipedia](https://ja.wikipedia.org/wiki/Switch%E6%96%87#%E3%83%95%E3%82%A9%E3%83%BC%E3%83%AB%E3%82%B9%E3%83%AB%E3%83%BC)

## ディレクトリ構成のスタンダード 

- [Goにはディレクトリ構成のスタンダードがあるらしい。 - Qiita](https://qiita.com/sueken/items/87093e5941bfbc09bea8)

## Enum の関数

- [4 iota enum examples · YourBasic Go](https://yourbasic.org/golang/iota/)

## buffer

```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer

	buf.WriteString("a")
	buf.WriteString("b")

	fmt.Println(buf.String())
}
```

## 空の構造体はメモリを消費しない

### コード

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(struct{}{}))
	fmt.Println(unsafe.Sizeof(true))
}

```

### 実行結果

```
0
1
```

## iota

[golang の iota の使い所 - pospomeのプログラミング日記](https://www.pospome.work/entry/2017/08/20/153604)

## Constructor内でのバリデーション

以下のようなパターンはあまり好ましくない。
生成時のルールが追加されたときに、ガードのロジックがむき出しになると、この関数の保守性が下がっていくから。

```go
type Name string

func NewName(name string) (Name, error) {

	s := string(n)
	if utf8.RuneCountInString(s) > 10 {
		return "", errors.New("invalid lenmgth")
	}

	return Name(name), nil
}

```

以下のように仕様オブジェクトに切り出しちゃうのが良い

```go
package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type Name string

func NewName(name string) (Name, error) {
	n := nameValidator(name)

	if err := n.validateLength(); err != nil {
		return "", err
	}

	return Name(n.string()), nil
}

type nameValidator string

func (n nameValidator) validateLength() error {
	s := string(n)
	if utf8.RuneCountInString(s) > 10 {
		return errors.New("invalid lenmgth")
	}
	return nil
}

func (n nameValidator) string() string {
	return string(n)
}
```

## package名と構造体の関係

- package名についている名詞は構造体には付与しない
    - 例えば、hoge packageには、hogeRaderではなくReaderで良い
        - [標準ライブラリでもこんな感じになっている](https://golang.org/pkg/encoding/csv/#example_Reader)