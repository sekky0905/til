# reflect package メモ

## Field は番号でアクセスできる

```go
package main

import (
	"fmt"
	"reflect"
)

type Hoge struct {
	A string
	B string
	C string
	D string
}

func main() {
	h := Hoge{
		A: "a",
		B: "b",
		C: "c",
		D: "d",
	}
	rType := reflect.TypeOf(h)
	for i := 0; i < rType.NumField(); i++ {
		fmt.Printf("%d: %v\n", i, rType.Field(i).Name)
	}
}

```

```
0: A
1: B
2: C
3: D

```
