# cmp package の tips

## cmp.Comparer 関数

### slice の比較

- [slice や struct を走査し](https://github.com/google/go-cmp/blob/5a6f75716e1203a923a78c9efb94089d857df0f6/cmp/compare.go#L404) `Comparer` に渡した比較関数の[型が一致していればそれを適用する](https://github.com/google/go-cmp/blob/5a6f75716e1203a923a78c9efb94089d857df0f6/cmp/options.go#L359)と言う形なので、比較関数の中でわざわざ for 等で回さなくてもOK

実装例

```go
package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type Person struct {
	ID   int
	Name string
}

func comparePerson(x, y *Person) bool {
	if x == nil && y == nil {
		return true
	}

	return x != nil && y != nil &&
		x.ID == y.ID &&
		x.Name == y.Name
}

func main() {
	s1 := []*Person{&Person{ID: 1, Name: "a"}, &Person{ID: 2, Name: "b"}, &Person{ID: 3, Name: "v"}}
	s2 := []*Person{&Person{ID: 1, Name: "a"}, &Person{ID: 2, Name: "b"}, &Person{ID: 3, Name: "v"}}
	s3 := []*Person{&Person{ID: 1, Name: "d"}, &Person{ID: 2, Name: "b"}, &Person{ID: 3, Name: "v"}}

	opt := cmp.Comparer(comparePerson)
	if diff := cmp.Diff(s1, s2, opt); diff != "" {
		fmt.Printf("diff (s1 +s2)\n%s", diff)
	}

	if diff := cmp.Diff(s1, s3, opt); diff != "" {
		fmt.Printf("diff (s1 +s3)\n%s", diff)
	}

}
```

実行結果

```
diff (s1 +s3)
  []*main.Person{
- 	&{ID: 1, Name: "a"},
+ 	&{ID: 1, Name: "d"},
  	&{ID: 2, Name: "b"},
  	&{ID: 3, Name: "v"},
  }

```

[実際のコード](https://play.golang.org/p/Bs0160wQq-g)
