# 並行処理

## Timer & Select & error handling

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---start X with limit 1---")
	X(1 * time.Second)

	fmt.Println("---start X with limit 3---")
	X(3 * time.Second)

	fmt.Println("done")
}

func X(limit time.Duration) {
	timer := time.NewTimer(limit)

	select {
	case <-timer.C:
		fmt.Println("timed out")
		return
	case res := <-Y():
		if res.err != nil {
			fmt.Println("error handling")
			return

		}
		fmt.Printf("received : %s\n", res.v)
	}

	fmt.Println("res")
}

type res struct {
	v   string
	err error
}

func Y() <-chan res {
	c := make(chan res)
	go func() {
		time.Sleep(2 * time.Second)
		c <- res{
			v:   "x executed",
			err: nil,
		}
	}()
	return c
}
```

実行結果

```
---start X with limit 1---
timed out
---start X with limit 3---
received : x executed
res
done

```

## sync.Once

実装

```go
package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var hoge *Hoge

type Hoge struct {
	Name       string
	FixedField string
}

func NewHoge() *Hoge {
	once.Do(func() {
		hoge = &Hoge{FixedField: "fixed"}
	})
	return hoge
}

func SetName(name string) {
	h := NewHoge()
	h.Name = name
}

func main() {
	h := NewHoge()
	fmt.Printf("h after NewHoge - %+v\n", h)
	SetName("aaa")
	fmt.Printf("h after SetName - %+v\n", h)
}

```

結果

```
h after NewHoge - &{Name: FixedField:fixed}
h after SetName - &{Name:aaa FixedField:fixed}
```
