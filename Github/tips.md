## リポジトリ内のドキュメント内で他のファイルを参照する
リポジトリ内のドキュメント内で他のファイルを参照するときには以下のような感じで相対パスでリンクを貼ることができる

```
[リンク](path/to/file)
```

## issue上で、ファイルの diff を綺麗に見せたいとき

- diffをとるときには、以下のようにする
    - [Git で変更を patch ファイルにする / patch コマンドで適用する - Qiita](https://qiita.com/sea_mountain/items/7d9c812e68a26bd1a292)
- `diff.patch` の中身をコピーして、issue等に ` ```patch` という形でシンタックスハイライトで書くと、綺麗に diff が表示される
