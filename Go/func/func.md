# 関数系まとめ

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

## Functional Option Pattern 

### 実装してみた

以下を参考に実装してみた
- [Go言語のFunctional Option Pattern - Qiita](https://qiita.com/weloan/items/56f1c7792088b5ede136)
- [Functional Option Pattern](https://blog.web-apps.tech/go-functional-option-pattern/)

[実際の実装](https://play.golang.org/p/irKBjK9F-J7)

```go
package main

import (
	"fmt"
)

type Hoge struct {
	Foo string
	Bar string
}

func NewHoge(options ...func(*Hoge)) *Hoge {
	hoge := &Hoge{}

	for _, option := range options { // 可変長引数で受け取った、Hoge のフィールドを設定する各種関数を range で回し、実行する
		option(hoge) // ポインタなので、各々関数で hoge のフィールドに設定
	}
	return hoge // 設定した hoge を返す
}

func Foo(foo string) func(*Hoge) { // これはあくまでも関数を返すだけで、内部の関数の実行までは行わない
	return func(h *Hoge) { // 関数を返す
		h.Foo = foo
	}
}

func Bar(bar string) func(*Hoge) { // これはあくまでも関数を返すだけで、内部の関数の実行までは行わない
	return func(h *Hoge) { // 関数を返す
		h.Bar = bar
	}
}

func main() {
	h1 := NewHoge(Foo("foo"), Bar("bar"))
	fmt.Printf("h1 = %+v\n", h1)

	h2 := NewHoge(Foo("foo"))
	fmt.Printf("h2 = %+v\n", h2)

	h3 := NewHoge(Bar("bar"))
	fmt.Printf("h3 = %+v\n", h3)
}

```

### 実行結果

```
h1 = &{Foo:foo Bar:bar}
h2 = &{Foo:foo Bar:}
h3 = &{Foo: Bar:bar}
```

### 参考

- [Go言語のFunctional Option Pattern - Qiita](https://qiita.com/weloan/items/56f1c7792088b5ede136)
- [Functional Option Pattern](https://blog.web-apps.tech/go-functional-option-pattern/)

## 引数としての関数にメソッドを割り当てる

### 実装してみた

```go
package main

import (
	"fmt"
)

type Foo struct {
	Field string
}

func (f Foo) Method(arg string) string {
	return f.Field + arg
}

func Func(method func(arg string) string) {
	fmt.Println(method("test"))
}

func main() {
	f := Foo{Field: "x"}
	Func(f.Method)
}
```

### 実行結果

```
xtest
```

