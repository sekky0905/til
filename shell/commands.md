# よく使うコマンドまとめ

## Draft の prefix が存在する branch を一括削除したい

```bash
git branch | grep 'draft.*' | xargs git branch -D
```

## grep + 無視

```bash
grep -rn 'target' . | grep -v 'ignore target'
```

## 連続してコマンド実行系

- [Linuxコマンドを連続して使うには - Qiita](https://qiita.com/egawa_kun/items/714394609eef6be8e0bf)
