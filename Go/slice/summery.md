# Slice 系まとめ

## Slice の要素の書き換え

indexで配列の要素を指定するようにしないと書き換えられない。

### コード

```go
package main

import (
	"fmt"
)

type Hoge struct {
	Name string
}

func main() {

	s1 := []Hoge{
		{Name: "A"}, {Name: "B"}, {Name: "C"},
	}

	for i, v := range s1 {
		v.Name = fmt.Sprintf("%s_%d", v.Name, i)
	}

	fmt.Printf("s1 = %s\n", s1)

	s2 := []Hoge{
		{Name: "A"}, {Name: "B"}, {Name: "C"},
	}

	for i, v := range s2 {
		s2[i].Name = fmt.Sprintf("%s_%d", v.Name, i)
	}

	fmt.Printf("s2 = %s\n", s2)
}

```

### 結果

```
s1 = [{A} {B} {C}]
s2 = [{A_0} {B_1} {C_2}]
```

## Slice をバリューオブジェクトとして扱いたい場合

以下のように型エイリアスを付与してしまうと、直接要素にアクセスできてしまうので外部から要素を変更できてしまう

```go
type Name string

type Names []Name
```

そこで、型エイリアスではなくて、構造体でラップしてしまい、そのフィールドにpackage privateにしちゃえば書き込みができなくなる

```go
type Names struct {
    values []string
}
```