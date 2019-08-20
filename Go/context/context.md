# Contextまとめ

## context.WithTimeout

### 実装例

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

func main() {
	ctx := context.Background()
	r, err := Call(ctx)
	fmt.Printf("r = %+v, err = %+v", r, err)
}

func Call(ctx context.Context) (int, error) {
	fmt.Println("== Call is called ===")
	c, deadline := context.WithTimeout(ctx, time.Second*1)

	defer deadline()

	return Work(c)
}
func Work(ctx context.Context) (int, error) {
	fmt.Println("== Work is called ===")

	time.Sleep(time.Second * 2)
	fmt.Println("== Do anything heavy process ===")

	select {
	case <-ctx.Done():
		if err := ctx.Err(); err != nil {
			return 0, errors.Wrap(err, "timeout error")
		}
		return 1, nil
	default:
		return 2, nil
	}
}

```

### 結果

```
== Call is called ===
== Work is called ===
== Do anything heavy process ===
r = 0, err = context deadline exceeded
timeout error
```
