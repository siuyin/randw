package main

import (
	"fmt"

	"github.com/siuyin/randw"
)

func main() {
	fmt.Println(randw.Word())
	fmt.Printf("rand: %v\n", randw.Select(5))
}
