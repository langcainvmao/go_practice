#第一行代码
package  main

import (
	"fmt"
	"time"
	"math/rand"
}

func main () {
	rand.seed(time.Now.UnixNano())
	n := rand.Intn(10)
	fmt.Println("Hello World")
	fmt.Println(n)
}
