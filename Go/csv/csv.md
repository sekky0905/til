# Goでcsv

- 基本的には、[Go言語でCSVの読み書き(sjis、euc、utf8対応) - Qiita](https://qiita.com/kesuzuki/items/202cc58db3fd1763c095)に書いてあるような感じでやればok
- ダブルクォーテションをつけたい場合には注意が必要
- 個々の要素にダブルクォーテーションを付与し、 `encoding/csv` を使用した場合、 `""""hoge"""` のような形で出力されてしまう( `"hoge"` としたい)
    - これは `RFC 4180` の仕様によるところらしい
        - 参考: [Strange CSV result for quoted strings in go encoding/csv - Stack Overflow](https://stackoverflow.com/questions/20459038/strange-csv-result-for-quoted-strings-in-go-encoding-csv)
