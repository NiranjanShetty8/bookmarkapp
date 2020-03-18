package main

import (
	"fmt"

	"github.com/NiranjanShetty8/bookmarkapp/repository"
)

func main() {
	fmt.Println("Hello")
	test := repository.NewGormRepository()
	fmt.Println(test)

}
