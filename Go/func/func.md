# 　関数系まとめ

## 可変長引数

コード

```go
package main

import (
	"fmt"
)

func main() {
	s1 := []string{"a", "b"}
	A(s1...)
	s2 := "c"
	A(s2)
	A("d", "e")
	A()
}

func A(x ...string) {
	fmt.Println(len(x))
}

```

実行結果

```go
2
1
2
0
```


## 構造体の中に埋め込んだ関数に別の構造体のメソッドをブチ込んでみる

[実際の実装](https://play.golang.org/p/01Pfy4ZZFMV)

### 実装

```go
package main

import (
	"fmt"
)

type Hoge struct {
	Foo string
}

func (h *Hoge) Bar(name string) {
	if name == h.Foo {
		fmt.Println("同じ")
	} else {
		fmt.Println("違う")
	}
}

func main() {
	h := &Hoge{
		Foo: "hoge",
	}

	fmt.Print(`h.Bar("hoge") => `)
	h.Bar("hoge")

	x := &X{
		Bar: h.Bar,
	}

	fmt.Print(`x.Bar("foo") => `)
	x.Bar("foo")

	fmt.Print(`x.Bar("hoge") => `)
	x.Bar("hoge")
}

type X struct {
	Bar func(name string)
}
```

### 結果

```
h.Bar("hoge") => 同じ
x.Bar("foo") => 違う
x.Bar("hoge") => 同じ
```
