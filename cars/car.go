package cars

/*
Struktar används väldigt mycket i Julialang, som är motsvarigheten till vad Go är
gentemot Java/C++, fast mot Python. Jätteflexiblet sätt att skapa egna data typer.


"klass"
*/

// type Address struct {
// 	StreetAddress string
// 	PostalCode    int
// 	City          string
// }

// type Person struct {
// 	Name      string
// 	Age       int
// 	Addressen Address
// }

// OOP encapsulation, constructor VALID STATE
// i videon ->

// Är car(struct) en objekt eller t.ex mileage är objekt?

// KLASS/STRUCT = Riting Blueprint   (ritning på ett hus)
// OBJEKT = ett hus skapat utifrån en ritning

// CONSTRUCTOR - säkerställer att objektet INTE går att skapa i INVALID STATE

type Car struct {
	Manufacturer string
	Milage       int //hur många mil 7000
	Color        string
	Regno        string
}

type Truck struct {
	Manufacturer string
	Milage       int //hur många mil 7000
	Color        string
	Regno        string
	MaxLoad      int
}

// "constructor"
func NewCar(regno string, milage int) Car {
	c := Car{Regno: regno, Milage: milage}
	return c
}

func NewTruck(regno string, milage int) Truck {
	c := Truck{Regno: regno, Milage: milage}
	return c
}

func (car Car) IsExpensiveService() bool {
	if car.Milage >= 7000 {
		return true
	}
	if car.Manufacturer == "BMW" {
		return true
	}
	return false
}
