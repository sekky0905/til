# Goでcsv

- 基本的には、[Go言語でCSVの読み書き(sjis、euc、utf8対応) - Qiita](https://qiita.com/kesuzuki/items/202cc58db3fd1763c095)に書いてあるような感じでやればok

## ダブルクォーテーション (`"`) について

### RFC側

- 各要素が `"` で囲まれているかもしれないし、囲まれていないかもしれない
> Each field may or may not be enclosed in double quotes
- 各要素を `"` で囲む場合には、要素中の `"` は、 `"` を重ねてエスケープする
> If double-quotes are used to enclose fields, then a double-quote appearing inside a field must be escaped by preceding it with another double quote

引用元: [RFC 4180 - Common Format and MIME Type for Comma-Separated Values (CSV) Files](https://tools.ietf.org/html/rfc4180)

### encoding/csv

- 以下のような振る舞い
    - 要素内に `"` が存在する場合
        - その要素を `"` で囲むパターンで返してくる(エスケープも行われる)
    - 要素内に `"` が存在しない場合
        - その要素を `"` で囲まないパターンで返してくる
    - 要素内に `"` がないのに各要素を `"` で囲むパターンができない

#### 実際のコード

```go
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	out := os.Stdout

	records := [][]string{
		[]string{`a1`, `b1`},
		[]string{`"a2`, `b2`},
		[]string{`a3`, `"b3`},
		[]string{`a4"`, `b4`},
		[]string{`a5`, `b5"`},
		[]string{`"a6"`, `b6`},
		[]string{`"a7"`, `"b7`},
		[]string{`"a8"`, `"b8"`},
	}

	w := csv.NewWriter(out)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			fmt.Printf("error has occurred. %+v", err)
		}
		w.Flush()
	}
}
```

#### 実行結果

```
a1,b1
"""a2",b2
a3,"""b3"
"a4""",b4
a5,"b5"""
"""a6""",b6
"""a7""","""b7"
"""a8""","""b8"""
```

## CSV読み込む処理が入ったメソッドのテストがしたい

- データストリームを変換する
strings.NewReader > csv.NewReader 

```go
 s := `ここにcsvの文字列を書く`
 r := csv.NewReader(strings.NewReader(s))
```