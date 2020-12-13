# Goと文字コード

## そもそもの文字コード周辺の単語の整理

- 文字コード
    - >コンピュータ上で文字（キャラクタ）を利用する目的で各文字に割り当てられるバイト表現。もしくは、バイト表現と文字の対応関係（文字コード体系）のこと
        - 引用元 : [文字コード - Wikipedia](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E3%82%B3%E3%83%BC%E3%83%89)
- 文字集合
    - > 文字を重複なく集めた集合
        - 引用元 : 矢野 啓介 (2018/12/28)『[改訂新版]プログラマのための文字コード技術入門 (WEB+DB PRESS plusシリーズ) 単行本』技術評論社
- 符号化文字集合
    - > 文字集合を定義し、その集合の各文字に対応するビット組み合わせを一意に定めたもの
        - 引用元 : 矢野 啓介 (2018/12/28)『[改訂新版]プログラマのための文字コード技術入門 (WEB+DB PRESS plusシリーズ) 単行本』技術評論社 
- 符号化方式
    - >実際にコンピュータが利用できるデータ列（通常、バイト列）に変換する符号化方式
        - 引用元 : [文字符号化方式 - Wikipedia](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E7%AC%A6%E5%8F%B7%E5%8C%96%E6%96%B9%E5%BC%8F)
    - 複数の符号化文字集合を組み合わせた運用方式
        - 引用元 : 矢野 啓介 (2018/12/28)『[改訂新版]プログラマのための文字コード技術入門 (WEB+DB PRESS plusシリーズ) 単行本』技術評論社 
参考 

- [新人さんに知ってほしい「文字コードのお話」 - Qiita](https://qiita.com/yuji38kwmt/items/b3a7820b4d3b544da4ff)
- [符号化文字集合と文字符号化方式 - 「プログラマのための文字コード技術入門」を読んだ - $shibayu36->blog;](https://blog.shibayu36.org/entry/2015/09/14/102100)

## Goの文字コードの話とrune 

- Unicodeは、符号化文字集合
    - コードポイントは、符号化文字集合内で文字に割り当てるための場所
- UTF-8は、符号化方式
- Goのソースコードや文字列リテラルは、UTF-8
- Goのruneは、Unicode(符号化文字集合)のcode point

参考

- [Strings, bytes, runes and characters in Go - The Go Blog](https://blog.golang.org/strings)
- [Goのruneを理解するためのUnicode知識 - Qiita](https://qiita.com/seihmd/items/4a878e7fa340d7963fee)


## Goで文字コードを変換してみる

[こちらの記事](https://qiita.com/uchiko/items/1810ddacd23fd4d3c934) を参考に以下のようなコードを書いてみた。([playgroundでも書いている](https://play.golang.org/p/hb_LGte3dkV))


```go
package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	a := "abc"
	a2, err := convertFromUTF8ToShiftJis(a)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("a2 = %v, len = %d\n", a2, len(a2))

	b := "あいう"
	b2, err := convertFromUTF8ToShiftJis(b)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("b2 = %v, len = %d\n", b2, len(b2))

	c := "🙇‍♀️"
	c2, err := convertFromUTF8ToShiftJis(c)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("c2 = %v, len = %d\n", c2, len(c2))
}

func convertFromUTF8ToShiftJis(str string) (string, error) {
	iostr := strings.NewReader(str)
	rio := transform.NewReader(iostr, japanese.ShiftJIS.NewEncoder())
	ret, err := ioutil.ReadAll(rio)
	if err != nil {
		return "", err
	}
	return string(ret), err
}
```

`b2, err := convertFromUTF8ToShiftJis(b)` の際には、 `rune not supported by encoding.` のエラーを返す。
これはなぜかというと、

- [Shift_JIS - Wikipedia](https://ja.wikipedia.org/wiki/Shift_JIS) は符号化文字集合として、[JIS X 0208](https://ja.wikipedia.org/wiki/JIS_X_0208)を使用している
    - 参考: [Shift_JIS - Wikipedia](https://ja.wikipedia.org/wiki/Shift_JIS)
- [JIS X 0208](https://ja.wikipedia.org/wiki/JIS_X_0208) では、その文字集合の表の中に絵文字が存在していないため 
    - ただし、[Shift_JIS - Wikipedia](https://ja.wikipedia.org/wiki/Shift_JIS)でも、絵文字を表示できたりする場合がある
        - [JIS X 0208](https://ja.wikipedia.org/wiki/JIS_X_0208) の外字である絵文字を文字集合の表の空き領域内に定義して使用しているから
            - 文字集合に含まれない文字を [外字](https://ja.wikipedia.org/wiki/%E5%A4%96%E5%AD%97) と言う

参考
[Go言語で文字コード変換 - Qiita](https://qiita.com/uchiko/items/1810ddacd23fd4d3c934)

## 参考文献

- 矢野 啓介 (2018/12/28)『[改訂新版]プログラマのための文字コード技術入門 (WEB+DB PRESS plusシリーズ) 単行本』技術評論社

## 参考にさせていただいたサイト

- [新人さんに知ってほしい「文字コードのお話」 - Qiita](https://qiita.com/yuji38kwmt/items/b3a7820b4d3b544da4ff)
- [符号化文字集合と文字符号化方式 - 「プログラマのための文字コード技術入門」を読んだ - $shibayu36->blog;](https://blog.shibayu36.org/entry/2015/09/14/102100)
- [Goのruneを理解するためのUnicode知識 - Qiita](https://qiita.com/seihmd/items/4a878e7fa340d7963fee)
- [golangで日本語（マルチバイト）の文字列を数える - Qiita](https://qiita.com/reiki4040/items/b82bf5056ee747dcf713)
- [Go言語で文字コード変換 - Qiita](https://qiita.com/uchiko/items/1810ddacd23fd4d3c934)
- [文字コード - Wikipedia](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E3%82%B3%E3%83%BC%E3%83%89)
- [文字符号化方式 - Wikipedia](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E7%AC%A6%E5%8F%B7%E5%8C%96%E6%96%B9%E5%BC%8F)
- [Shift_JIS - Wikipedia](https://ja.wikipedia.org/wiki/Shift_JIS)
- [JIS X 0208](https://ja.wikipedia.org/wiki/JIS_X_0208)
- [外字 - Wikipedia](https://ja.wikipedia.org/wiki/%E5%A4%96%E5%AD%97) 