# time についてのメモ

> These are predefined layouts for use in Time.Format and time.Parse. The reference time used in the layouts is the specific time:
>  Mon Jan 2 15:04:05 MST 2006

【意訳】
Time.Format と time.Parse で使用するための事前に定義されたレイアウトが存在する。レイアウトで使用される参照時間は、具体的な時間。

> To define your own format, write down what the reference time would look like formatted your way

【意訳】自分自身のフォーマットを定義するには、参照時間がどのようにフォーマットされるかを書く。

引用元: [time - The Go Programming Language](https://golang.org/pkg/time/#pkg-constants)


具体的な使用方法は、[time - The Go Programming Language](https://golang.org/pkg/time/#pkg-constants)を参照。

フォーマットを守っていれば以下のような感じにすることもできる。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2019, 9, 5, 0, 0, 0, 0, time.Local)
	fmt.Println(t.Format("blog-20060102.md"))
}
```

実行結果

```
blog-20190905.md
```

ちなみに以下のようにすると、フォーマットが不正( `20050102` )なので、うまくフォーマットできない

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2019, 9, 5, 0, 0, 0, 0, time.Local)
	fmt.Println(t.Format("blog-20050102.md"))
}
```

実行結果

```
blog-50000905.md
```
