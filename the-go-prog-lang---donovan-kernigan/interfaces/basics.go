package interfaces

import "fmt"

type Driver interface {
	MoveForward(int)
	Turn()
}

type Car struct {
	speed int
}
type Bike struct {
	speed int
}

func (c *Car) MoveForward(speed int) {
	fmt.Println("car is moving with speed: ", c.speed)
}
func (b *Bike) MoveForward(speed int) {
	fmt.Println("bike is moving with speed: ", b.speed)
}

func (c *Car) Turn()  {}
func (c *Bike) Turn() {}

var BikeInstance = Bike{
	15,
}

var CarInstance = Car{
	75,
}

func CompareInterfaces(c, b Driver) {

	c.MoveForward(70)
	b.MoveForward(15)

	fmt.Println(c == b)

}
