package main

import(
	"fmt"
	"math"
)

const ss string = "constant"

func main(){
	fmt.Print(ss)

	const n = 500000000

	const d = 3e20 /n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
