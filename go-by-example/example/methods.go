package example

import "fmt"

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (r rectangle) perim() int {
	return 2*r.width + 2*r.height
}

func Methods() {
	r := rectangle{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
