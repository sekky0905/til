# JSON 系 Tips


- `json:",omitempty"` のように `omitempty"` を付与すると、付与されたフィールドが空だった場合に Marshal の際に出力されないようにすることができる

参考: [golangのenconding/jsonのタグについて - Qiita](https://qiita.com/pside/items/0fc7f85746211371b23e#marshal%E6%99%82%E3%81%AB%E5%87%BA%E5%8A%9B%E3%81%97%E3%81%AA%E3%81%84)


```go
package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
}

func (v A) MarshalJSON() ([]byte, error) {
	return []byte{}, nil
}

type B struct {
}

func (v B) MarshalJSON() ([]byte, error) {
	return []byte("{}"), nil
}

type C struct {
}

func (v C) MarshalJSON() ([]byte, error) {
	return nil, nil
}

type D struct {
}

func main() {
	b1, err := json.Marshal(A{})
	fmt.Printf("b1 - %+v, err  - %+v\n", string(b1), err)

	b2, err := json.Marshal(B{})
	fmt.Printf("b2 - %+v, err  - %+v\n", string(b2), err)

	b3, err := json.Marshal(C{})
	fmt.Printf("b3 - %+v, err  - %+v\n", string(b3), err)

	b4, err := json.Marshal(D{})
	fmt.Printf("b4 - %+v, err  - %+v\n", string(b4), err)

}
```

```
b1 - , err  - json: error calling MarshalJSON for type main.A: unexpected end of JSON input
b2 - {}, err  - <nil>
b3 - , err  - json: error calling MarshalJSON for type main.C: unexpected end of JSON input
b4 - {}, err  - <nil>
```
