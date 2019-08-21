# Contextまとめ

## context.WithTimeout

### 実装例

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func Work(ctx context.Context) {
	fmt.Println("work is called")
	fmt.Println("start do any heavy process")
	time.Sleep(3 * time.Second)
	fmt.Println("end do any heavy process")
}

func Call() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	go Work(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("timeout:", ctx.Err())
	}
}

func main() {
	Call()
}

```

### 結果

```
work is called
start do any heavy process
timeout: context deadline exceeded
```

### 注意点

呼び出す関数を goroutine で呼び出さないと、その関数がブロックしてしまい、関数が完了するまで select にたどり着かない
そのため、全部完了してから select で Done されるという意味のない状態になる

#### 実装例

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func Work(ctx context.Context) {
	fmt.Println("work is called")
	fmt.Println("start do any heavy process")
	time.Sleep(3 * time.Second)
	fmt.Println("end do any heavy process")
}

func Call() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	Work(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("timeout:", ctx.Err())
	}
}

func main() {
	Call()
}
```

#### 結果

```
work is called
start do any heavy process
end do any heavy process
timeout: context deadline exceeded
```
