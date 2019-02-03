package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	car1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "green",
		},
		fourWheel: false,
	}

	car2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	car3 := struct {
		vehicle
		hasgps bool
	}{
		vehicle: vehicle{
			doors: 2,
			color: "green"},
		hasgps: true,
	}
	fmt.Println(car2)
	fmt.Println(car1)
	fmt.Println(car3)

	fmt.Println("Car #1 is", car1.color)
	fmt.Println("Is car #2 a luxury car?", car2.luxury)
	fmt.Println("Does car #3 have GPS?", car3.hasgps)

}
