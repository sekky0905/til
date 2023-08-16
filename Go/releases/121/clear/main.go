package main

import (
	"fmt"
)

func main() {
	fmt.Println("---clear map---")
	ClearMap()

	fmt.Println("---clear slice string---")
	ClearSliceString()

	fmt.Println("---clear slice int---")
	ClearSliceInt()

	fmt.Println("---clear slice bool---")
	ClearSliceBool()

	fmt.Println("---clear slice pointer struct---")
	ClearSlicePointerStruct()

	fmt.Println("---clear slice struct---")
	ClearSliceStruct()
}

func ClearMap() {
	m := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}

	fmt.Println("before clear")
	fmt.Printf("m= %+v\n", m)
	fmt.Printf("len(m) = %v\n", len(m))
	for k, v := range m {
		fmt.Printf("key = %v, value = %v\n", k, v)
	}

	fmt.Println()

	clear(m)
	fmt.Println("after clear")
	fmt.Printf("m= %+v\n", m)
	fmt.Printf("len(m) = %v\n", len(m))
	for k, v := range m {
		fmt.Printf("key = %v, value = %v\n", k, v)
	}
}

func ClearSliceString() {
	s := []string{"foo", "bar", "baz"}

	fmt.Println("before clear")
	fmt.Printf("s= %+v\n", s)
	fmt.Printf("len(s) = %v\n", len(s))
	for i, v := range s {
		fmt.Printf("index = %v, value = %v\n", i, v)
	}

	fmt.Println()
	clear(s)

	fmt.Println("after clear")
	fmt.Printf("s= %+v\n", s)
	// lengthは3なので注意
	fmt.Printf("len(s) = %v\n", len(s))
	for i, v := range s {
		fmt.Printf("index = %v, value = %v\n", i, v)
	}
}

func ClearSliceInt() {
	s := []int{1, 2, 3}

	fmt.Println("before clear")
	fmt.Printf("s= %+v\n", s)
	fmt.Printf("len(s) = %v\n", len(s))
	for i, v := range s {
		fmt.Printf("index = %v, value = %v\n", i, v)
	}

	fmt.Println()
	clear(s)

	fmt.Println("after clear")
	fmt.Printf("s= %+v\n", s)
	// lengthは3なので注意
	fmt.Printf("len(s) = %v\n", len(s))
	for i, v := range s {
		fmt.Printf("index = %v, value = %v\n", i, v)
	}
}

func ClearSliceBool() {
	s := []bool{true, false, true}

	fmt.Println("before clear")
	fmt.Printf("s= %+v\n", s)
	fmt.Printf("len(s) = %v\n", len(s))
	for i, v := range s {
		fmt.Printf("index = %v, value = %v\n", i, v)
	}

	fmt.Println()
	clear(s)

	fmt.Println("after clear")
	fmt.Printf("s= %+v\n", s)
	// lengthは3なので注意
	fmt.Printf("len(s) = %v\n", len(s))
	for i, v := range s {
		fmt.Printf("index = %v, value = %v\n", i, v)
	}
}

func ClearSlicePointerStruct() {
	p := []*struct{ name string }{{"foo"}, {"bar"}, {"baz"}}
	fmt.Println("before clear")
	fmt.Printf("p= %+v\n", p)
	fmt.Printf("len(p) = %v\n", len(p))
	for i, v := range p {
		fmt.Printf("index = %v, value = %v\n", i, v)
	}

	fmt.Println()
	clear(p)

	fmt.Println("after clear")
	fmt.Printf("p= %+v\n", p)
	// lengthは3なので注意
	fmt.Printf("len(p) = %v\n", len(p))
	for i, v := range p {
		fmt.Printf("index = %v, value = %v\n", i, v)
	}
}

func ClearSliceStruct() {
	p := []struct{ name string }{{"foo"}, {"bar"}, {"baz"}}
	fmt.Println("before clear")
	fmt.Printf("p= %+v\n", p)
	fmt.Printf("len(p) = %v\n", len(p))
	for i, v := range p {
		fmt.Printf("index = %v, value = %v\n", i, v)
	}

	fmt.Println()
	clear(p)

	fmt.Println("after clear")
	fmt.Printf("p= %+v\n", p)
	// lengthは3なので注意
	fmt.Printf("len(p) = %v\n", len(p))
	for i, v := range p {
		fmt.Printf("index = %v, value = %v\n", i, v)
	}
}
