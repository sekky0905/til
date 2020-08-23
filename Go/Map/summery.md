# Map まとめ

### リテラルでの初期化

```go
package main

import (
	"fmt"
)

type Struct struct {
	m map[string]string
}

func NewStruct() *Struct {
	return &Struct{
		// 何もないと言うリテラルになる
		m: map[string]string{},
	}
}

func main() {
	s := NewStruct()
	s.m["hoge"] = "foo"

	fmt.Println(s.m["hoge"])
}

```

### 結果

```
foo
```