# nil まとめ

## nil の型

- [絶対ハマる、不思議なnil - Qiita](https://qiita.com/umisama/items/e215d49138e949d7f805#nil%E3%81%AF%E3%82%AD%E3%83%A3%E3%82%B9%E3%83%88%E5%87%BA%E6%9D%A5%E3%82%8B)

以下で nil を cast できる

```go
var hoge = (*Hoge)(nil)
```

## Exists

```go
package main

import (
	"fmt"
)

type Hoge struct {
}

func (h *Hoge) Exists() bool {
	return h != nil
}

func (h *Hoge) Foo1() {
	fmt.Printf("h.Exists = %t @Foo1\n", h.Exists())
}

func (h Hoge) Foo2() {
	fmt.Printf("h.Exists = %t @Foo2\n", h.Exists())
}

func main() {
	var h *Hoge
	call1(h)

}

func call1(h *Hoge) {
	h.Foo1()
	h.Foo2()
}

```

### 結果

```
h.Exists = false @Foo1
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x492f90]

goroutine 1 [running]:
main.call1(0x0)
	/tmp/sandbox813117633/prog.go:30 +0x30
main.main()
	/tmp/sandbox813117633/prog.go:24 +0x2a
```

## Exists2

```go
package main

import (
	"fmt"
)

type X interface {
	Exists() bool
}

type xImp struct {
}

func (x *xImp) Exists() bool {
	return x != nil
}

func NewX() X {
	return &xImp{}
}

func NewXNil() X {
	return nil
}

func main() {
	x := NewX()
	Hoge(x)

	var x2 *xImp
	fmt.Printf("x2 = %+v\n", x2)
	Hoge(x2)
}

func Hoge(x X) {
	fmt.Println(x.Exists())
}

```

### 結果

```
true
x2 = <nil>
false
```