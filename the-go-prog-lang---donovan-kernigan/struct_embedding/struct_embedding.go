package structembedding

import (
	"fmt"
	"time"
)

type Model struct {
	Name string
	Year int
}

func (m *Model) Rename(newName string) {
	m.Name = newName
}

type Coordinates struct {
	X float64
	Y float64
}

type Pilot struct {
	Firstname   string
	Lastname    string
	DOB         time.Time
	Nationality string
}

func (p *Pilot) ChangePilot(firstname string, lastname string) {
	p.Firstname = firstname
	p.Lastname = lastname
}

type Plane struct {
	Model
	Coordinates
	Pilot
}

var RedDragon = Plane{
	Model: Model{
		Name: "Red Dragon",
		Year: 1994,
	},
	Coordinates: Coordinates{
		X: 25.4,
		Y: 66.2,
	},
	Pilot: Pilot{
		Firstname:   "Ivan",
		Lastname:    "Petrov",
		DOB:         time.Date(1973, 10, 12, 0, 0, 0, 0, time.UTC),
		Nationality: "Japanese",
	},
}

func (p Plane) GPS() {
	fmt.Printf("plane with name '%s' right now at %.2f, %.2f\n", p.Model.Name, p.Coordinates.X, p.Coordinates.Y)
}

func Embedding() {

	fmt.Printf("%#v\n", RedDragon)

	RedDragon.GPS()

	RedDragon.X = 44.12
	RedDragon.Y = 152.15

	RedDragon.GPS()

	RedDragon.Rename("Blue Wave")

	RedDragon.GPS()

	RedDragon.ChangePilot("Petr", "Chekhov")

	fmt.Printf("%#v\n", RedDragon)
}
