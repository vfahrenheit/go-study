package main

import "fmt"

type Vertex struct {
	X, Y int
}

func swap(x, y string) (string, string) {
	return y, x
}


type Vertex1 struct {
	Lat, Long float64
}

var m = map[string]Vertex1{
	"Bell Labs": Vertex1{
		40.68433, -74.39967,
	},
	"Google": Vertex1{
		37.42202, -122.08408,
	},
}
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	var x, y, z int = 1, 2, 3
	c, python, java := true, false, "no!"

	fmt.Println(x, y, z, c, python, java)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)

	fmt.Println(m)
}