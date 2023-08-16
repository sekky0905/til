package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("---min and max behavior for basic pattern---")
	MinAndMaxBehaviorForBasicPattern()
	fmt.Println("---min and max behavior for extreme pattern---")
	MinAndMaxBehaviorForExtremePattern()
}

func MinAndMaxBehaviorForBasicPattern() {
	// m == x
	fmt.Printf("min(1) = %v\n", min(1))
	// x と y のうち小さい方
	fmt.Printf("min(1, 2) = %v\n", min(1, 2))
	// x と y のうち大きい方、ただし最低10(両方とも10未満)
	fmt.Printf("min(1, 2, 10) = %v\n", min(1, 2, 10))
	// x と y のうち大きい方、ただし最低10（片方だけ10未満)
	fmt.Printf("min(1, 20, 10) = %v\n", min(1, 20, 10))
	// x と y のうち大きい方、ただし最低10（両方だけ10より大きい)
	fmt.Printf("min(1, 2, 10) = %v\n", min(15, 20, 10))
	// c == 10.0 (浮動小数点の種類)
	fmt.Printf("max(1, 2.0, 10) = %v\n", max(1, 2.0, 10))
	// f の型は float32
	f := max(0, float32(1))
	fmt.Printf("max(0, float32(1)) = %v, type = %T \n", f, f)

	//var s []string = []string{"foo", "bar", "baz"}
	//_ = min(s...)              // 無効: スライスの引数は許可されていません

	t := max("", "foo", "bar") // t == "foo" (文字列の種類)
	fmt.Printf("max(\"\", \"foo\", \"bar\") = %v, type = %T \n", t, t)
}

func MinAndMaxBehaviorForExtremePattern() {
	//var s []string = []string{"foo", "bar", "baz"}
	//_ = min(s...)              // 無効: スライスの引数は許可されていません

	t := max("", "foo", "bar") // t == "foo" (文字列の種類)
	fmt.Printf("max(\"\", \"foo\", \"bar\") = %v, type = %T \n", t, t)

	// 負のゼロは（非負の）ゼロよりも小さい
	fmt.Printf("min(-0.0, 0.0): %v\n", math.Min(-0.0, 0.0))
	fmt.Printf("max(-0.0, 0.0): %v\n", math.Max(-0.0, 0.0))

	// 負の無限大は他の任意の数よりも小さい
	fmt.Printf("min(-Inf, 1.0): %v\n", math.Min(math.Inf(-1), 1.0))
	fmt.Printf("max(-Inf, 1.0): %v\n", math.Max(math.Inf(-1), 1.0))

	// 正の無限大は他の任意の数よりも大きい
	fmt.Printf("min(+Inf, 1.0): %v\n", math.Min(math.Inf(1), 1.0))
	fmt.Printf("max(+Inf, 1.0): %v\n", math.Max(math.Inf(1), 1.0))

	// 引数のいずれかが NaN の場合、結果は NaN
	fmt.Printf("min(NaN, 1.0): %v\n", math.Min(math.NaN(), 1.0))
	fmt.Printf("max(NaN, 1.0): %v\n", math.Max(math.NaN(), 1.0))
}
