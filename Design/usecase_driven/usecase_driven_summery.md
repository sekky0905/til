## ユースケース駆動設計関係まとめ

### 資料

-  [なぜDDD初心者はググり出してすぐに心がくじけてしまうのか - little hands' lab](https://little-hands.hatenablog.com/entry/2017/09/24/005903)
-  [明日から使えるDDDのためのユースケース駆動開発（ICONIXプロセス） - Qiita](https://qiita.com/hirodragon/items/e2330edc1d1a329d17f5)
-  [DDD時代に考えたいICONIXプロセス/ICONIX in DDD - Speaker Deck](https://speakerdeck.com/hirodragon112/iconix-in-ddd)
-  [ドメイン駆動設計の正しい歩き方](https://www.slideshare.net/slideshow/embed_code/key/tn4xarCWbpAabm?startSlide=24)

### 考えまとめ

- ICONIX は、用件定義~設計をレビューを挟みながら、イテラブルに回していくイメージ
    - 動的な部分と静的な部分を交互に繰り返しながら
- 上流工程でイテラブルにレビューを挟みながら作業を実施していくので、実装後にやっぱり違いましたみたいな手戻りが発生することは少なさそう
- 意味のある実装にも繋がったドキュメントが残るので、他の人にも理解しやすくなる
    - 実装がどうかということの理解だけではなくて、実際の利用者のアクションとそれに対応するシステム側の反応がわかるので