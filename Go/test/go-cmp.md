# go-cmp メモ

- テスト等で `expected` と `got` を比較するには、[google/go-cmp](https://github.com/google/go-cmp)を使用すると便利
- 基礎的なところは、以下の記事が参考になる
    - [go-cmp を実際に使ってみた - Qiita](https://qiita.com/iszk/items/e799ec4d6f1f5eece706)
    - [構造体などをテストするのに便利なgoogle/go-cmpの使い方 - Qiita](https://qiita.com/hgsgtk/items/bd78bada902c91745fa5)

## [Comparer function](https://godoc.org/github.com/google/go-cmp/cmp#Comparer)を使用して slice のテストを簡単にする


### [Comparer function](https://godoc.org/github.com/google/go-cmp/cmp#Comparer)とは
>  Comparer returns an Option that determines whether two values are equal to each other.

[意訳] Comparer は、2つの値がお互いに等しいかどうかを決定する Option を返す。

引用元: [Comparer function](https://godoc.org/github.com/google/go-cmp/cmp#Comparer)

Comparer 2つの値がお互いに等しいかどうかを決定する関数を渡してやれば良い

### テストに用いる

以下のような実装があるとする

```go
package main

import "time"

type User struct {
	ID uint32
	Name string
	Age uint8
	CreatedAt time.Time
}

func GetUsers()[]User {
	// ここで仮にDBの操作が存在するものとする
	// そして、CreatedAt は DB 側で NOW() で自動で追加されているものとする


	return []User {
		{
			ID:        1,
			Name:      "hoge",
			Age:       20,
			CreatedAt:time.Now(), // 擬似的にここで入力
		},
		{
			ID:        1,
			Name:      "foo",
			Age:       20,
			CreatedAt:time.Now(),
		},
		{
			ID:        1,
			Name:      "bar",
			Age:       20,
			CreatedAt:time.Now(),
		},
	}
}
```

これをテストする際に以下のようにすると当然だがエラーになる。

```go
package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetUsers(t *testing.T) {
	tests := []struct {
		name string
		want []User
	}{
		{
			name: "3usersを取得すること",
			want: []User{
				{
					ID:   1,
					Name: "hoge",
					Age:  20,
				},
				{
					ID:   1,
					Name: "foo",
					Age:  20,
				},
				{
					ID:   1,
					Name: "bar",
					Age:  20,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUsers(); !reflect.DeepEqual(got, tt.want) {
			    t.Errorf("GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

```

エラー

```go
--- FAIL: TestGetUsers (0.00s)
--- FAIL: TestGetUsers/3usersを取得すること (0.00s)
    main_test.go:37: GetUsers() = [{1 hoge 20 2019-08-28 16:40:19.772311 +0900 JST m=+0.000646976} {1 foo 20 2019-08-28 16:40:19.772312 +0900 JST m=+0.000647073} {1 bar 20 2019-08-28 16:40:19.772312 +0900 JST m=+0.000647166}], want [{1 hoge 20 0001-01-01 00:00:00 +0000 UTC} {1 foo 20 0001-01-01 00:00:00 +0000 UTC} {1 bar 20 0001-01-01 00:00:00 +0000 UTC}]
```

このような感じで1つ１つを range で回しながら比較することもできる

```go
package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetUsers(t *testing.T) {
	tests := []struct {
		name string
		want []User
	}{
		{
			name: "3usersを取得すること",
			want: []User{
				{
					ID:   1,
					Name: "hoge",
					Age:  20,
				},
				{
					ID:   1,
					Name: "foo",
					Age:  20,
				},
				{
					ID:   1,
					Name: "bar",
					Age:  20,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetUsers()
			opt := cmpopts.IgnoreFields(User{}, "CreatedAt")
			for i, v := range got {
				if diff := cmp.Diff(v, tt.want[i], opt); diff != "" {
					t.Errorf("diff (-got +expected)\n%s", diff)
				}
			}
		})
	}
}

```

 [Comparer function](https://godoc.org/github.com/google/go-cmp/cmp#Comparer)を使用するとこんな感じでも書くことができる。


```go
package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetUsers(t *testing.T) {
	tests := []struct {
		name string
		want []User
	}{
		{
			name: "3usersを取得すること",
			want: []User{
				{
					ID:   1,
					Name: "hoge",
					Age:  20,
				},
				{
					ID:   1,
					Name: "foo",
					Age:  20,
				},
				{
					ID:   1,
					Name: "bar",
					Age:  20,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetUsers()
			opt := cmp.Comparer(func(x, y []User) bool {
				nx, ny := len(x), len(y)
				if nx != ny {
					return false
				}

				for i := 0; i < nx; i++ {
					if x[i].ID!= y[i].ID || x[i].Name != y[i].Name || x[i].Age != y[i].Age  {
						return false
					}
				}
				return true
			})

			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("diff (-got +expected)\n%s", diff)
			}

		})
	}
}

```
