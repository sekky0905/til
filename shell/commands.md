# よく使うコマンドまとめ

## Draft の prefix が存在する branch を一括削除したい

```bash
git branch | grep 'draft.*' | xargs git branch -D
```

## grep + 無視

```bash
grep -rn 'target' . | grep -v 'ignore target'
```
