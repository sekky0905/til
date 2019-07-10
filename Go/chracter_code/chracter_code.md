# Goの文字コード周りについて

## そもそもの文字コード

- 符号化文字集合 = 文字とビットの組み合わせの規則の集合
- 符号化方式 = > 実際にコンピュータが利用できるデータ列（通常、バイト列）に変換する符号化方式

引用元 : [文字符号化方式 - Wikipedia](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E7%AC%A6%E5%8F%B7%E5%8C%96%E6%96%B9%E5%BC%8F)

参考 
[新人さんに知ってほしい「文字コードのお話」 - Qiita](https://qiita.com/yuji38kwmt/items/b3a7820b4d3b544da4ff)

## Code Pointとrune

- Code Point = 文字コード表中のどこの場所かを示すもの
- rune = Unicode(符号化文字集合)のcode point

参考
[Goのruneを理解するためのUnicode知識 - Qiita](https://qiita.com/seihmd/items/4a878e7fa340d7963fee)

## Goで文字コードを変換する

以下の記事のコードを参考に[書いてみた](https://play.golang.org/p/O0Cb-usczQx)が、Shift Jisへ絵文字を変換すると空文字になる。
絵文字が空文字になる理由は、Shift Jisが使用している[符号化文字集合](https://ja.wikipedia.org/wiki/JIS_X_0208)では、絵文字は表中に定義されていないため。
Shift Jisでも絵文字を表示できたりするのは、表の空き領域に[外字](https://ja.wikipedia.org/wiki/%E5%A4%96%E5%AD%97)である絵文字を定義して使用しているから。

参考
[Go言語で文字コード変換 - Qiita](https://qiita.com/uchiko/items/1810ddacd23fd4d3c934)
