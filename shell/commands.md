# よく使うコマンドまとめ

## Draft の prefix が存在する branch を一括削除したい

```bash
git branch | grep 'draft.*' | xargs git branch -D
```
