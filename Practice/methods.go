package main

import "fmt"

type reckt struct{
	width, height int
}

func (r *reckt) area() int {
	return r.width * r.height
}

func (r reckt) perim() int {
	return 2*r.width + 2*r.height
}

func main(){
	r := reckt{width: 10, height:5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim", r.perim())

	rp := &r
	fmt.Println("area ", rp.area())
	fmt.Println("area ", rp.perim())
}
