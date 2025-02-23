package main

import (
	"fmt"
	"slices"
)

func main()  {
	names := []string{"Paul", "George", "Ringo", "Johnatan"}
	values := []int{100, 80, 60, 50}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "John"))
	fmt.Println(slices.Index(names, "George"))

}
