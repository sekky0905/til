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
