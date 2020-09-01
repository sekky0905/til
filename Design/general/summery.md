# 設計全般まとめ

## 型でステータスを表す

構造体のフィールドで状態を持って変更するのではなくて、そのステータスのオブジェクトを作成する
こうするとImmutableになる
例えば順番を保証したい場合とかには、次のステータスを返すようなメソッドを作れば良い

こうすると、Mutableになっちゃう


```go
type Hoge struct {
    Status int
   // ...
}
```
こうすると、Immutableになる(~Statusみたいな命名は微妙だけど、例がパッと思いつかなかったので)

```go
type InitStatusHoge struct {
   // ...
}

func NewInitStatusHoge() *InitStatusHoge {
    return &InitStatusHoge{}
}

func (s InitStatusHoge)NextStatus() *SecondStatusHoge {
    return &SecondStatusHoge{}
}
type SecondStatusHoge struct {
       // ...
}
```