package main

import (
	"fmt"

	"github.com/davidhadas/knativesecuritygate/queuegate"
)

func main() {
	str := queuegate.Hello()
	fmt.Println(str)
}
