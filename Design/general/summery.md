# 設計全般まとめ

## 型でステータスを表す

- 構造体のフィールドで状態を持って変更するのではなくて、そのステータスのオブジェクトを作成する
- こうするとImmutableになる
- 例えば順番を保証したい場合とかには、次のステータスを返すようなメソッドを作れば良い

```go
type Hoge struct {
    Status int
   // ...
}
```

```go

func New()InitStatusHoge *InitStatusHoge {
    return &InitStatusHoge{}
}

type InitStatusHoge struct {
   // ...
}

func (s InitStatusHoge)NextStatus()SecondStatusHoge *SecondStatusHoge {
    return &SecondStatusHoge{}
}


// SecondStatus は、名前的に微妙だけど...
type SecondStatusHoge struct {
       // ...
}
```