package main

import (
	"fmt"

	"github.com/alexander-bautista/testpkg/routes"
	"github.com/google/go-cmp/cmp"
)

func main() {
	routes.Hello()
	fmt.Println(routes.Alert("maria"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
