package main

import (
	"fmt"
	"main/cars"
)

func calculateSalary(age int, namn string) (salary int, avgift int) {
	salary = 50
	avgift = 12
	if age > 50 && namn == "Stefan" {
		salary = salary + 100
	}
	return salary, avgift
}

// func calculateSalary(age int, namn string ) int{
// 	var result = 50
// 	if age > 50 && namn == "Stefan"{
// 		result = result + 100
// 	}
// 	return result
// }

/*
- Array/listor ung samma från Flip video
- Lite nytt:
- Flera filer/subfolders
- Structs  - "klasser"

stack  - int i = 12;  //0-16G   4 bytes (adress 1000)  int j = 12;
heap - p =  new HockeyPlayer   -    p (adress 10000)

Kan vi inte göra append i array som i python? Och varför?

								// FIZED SIZE COMPILE TRIMNE

Varför behöver vi inte importera slice or list bibliotek(funktioner t.ex append)
som vi behöver för print?
*/

func lab1() {
	var allaTal [4]int
	fmt.Print("Ange ett tal:")
	fmt.Scanln(&allaTal[0])
	fmt.Print("Ange ett tal:")
	fmt.Scanln(&allaTal[1])
	fmt.Print("Ange ett tal:")
	fmt.Scanln(&allaTal[2])
	fmt.Print("Ange ett tal:")
	fmt.Scanln(&allaTal[3])

	largestSoFar := allaTal[0]
	// Vi ska hitta STÖRSTA TALET
	for i := 0; i < 4; i++ {
		if allaTal[i] > largestSoFar {
			largestSoFar = allaTal[i]
		}
	}
	fmt.Println(largestSoFar)

	// for i :=0; i < 4; i++{
	// 	fmt.Print("Ange ett tal:")
	// 	fmt.Scanln(&allaTal[i])
	// 	//var talet int
	// 	// fmt.Scanln(&talet)
	// 	// allaTal[i] = talet
	// }
}

func main() {

	//	lab1()

	allCars := []cars.Car{}

	// c och c2 är objekt skapade ut ritningen Car
	c := cars.NewCar("YUB100", 7000)
	c.Color = "Black"
	c.Manufacturer = "Renault"

	allCars = append(allCars, c)

	c2 := cars.NewCar("ABC113", 1000)
	c2.Manufacturer = "Volvo"
	c2.Color = "Green"

	allCars = append(allCars, c2)
	for _, currentCar := range allCars {
		fmt.Println(currentCar)
		//if cars.IsExpensiveService(currentCar) {
		if currentCar.IsExpensiveService() {
			fmt.Println("Usch det blir dyrt")
		}
	}

	//allCars[0].Color

	// slice LIST = dynamisk
	// slice Javascript a[2:5]

	for _, carManufacturer := range cars.GetCarManufacturersWeSell() {
		fmt.Println(carManufacturer)
		fmt.Scanln()
	}

	r := calculateWhatever(12, 11)
	fmt.Println(r)

	var minaBarn = []string{} // ARRAY - konsekutivt EFTER VARANDRA minne -
	// FIZED SIZE COMPILE TIMNE
	minaBarn[0] = "Fanny"
	minaBarn[1] = "Josefine"
	minaBarn[2] = "Oliver"

	// var players = [...]string{"Stefan", "Mats", "Lisa", "Anna"}

	// for _, vardet := range players {
	// 	fmt.Println(vardet)
	// }

	//Vilken Scenarion använder vi array och inte slice i riktig application
	// en array är  aningen snabbare -
	// om vi har datastrukturer som är FIXED - tex alla månader
	// var swedishCounties = []string{"Västa Götalands län", ""}
	// swedishCounties = append(swedishCounties, "asdjkdasljk")

	//var months = new List<String>{"January", ""}
	// Slice (LIST) - DYNAMISK DATASTRUKTUR - växa och krympa
	//    fortfarande är KONSEKTIVT MINNE
	var players = []string{"Stefan", "Mats", "Lisa", "Anna"}
	var namn string
	for {
		fmt.Print("Ange namn:")
		fmt.Scan(&namn)
		players = append(players, namn)
	}

	//var ages = []int{}; // slice (LIST)

	_, _ = calculateSalary(51, "Stefan")
	for {
		fmt.Println("1. Skapa device")
		fmt.Println("2. Lista alla")
		fmt.Println("3. Ändra device")
		fmt.Println("4. Sök")
		fmt.Println("5. Avsluta")

		fmt.Print("Val:")

		var tal int
		_, e := fmt.Scanln(&tal)
		if e != nil {
			fmt.Printf("Mata in ett tal tack")
		} else {
			switch tal {
			case 1:
				fmt.Println("Skapa device")
			case 2:
				fmt.Println("Lista alla")
			case 3:
				fmt.Println("Ändra device")
			case 4:
				fmt.Println("Sök")
			case 5:
				fmt.Println("Avsluta")
				return
			}
		}

	}

	var name string
	var age int = 18

	fmt.Print("Vad heter du?")
	fmt.Scanln(&name)

	fmt.Print("Vad gammal är du?")
	if age == 15 {

	} else {

	}
	if age >= 65 {
		fmt.Println("Du är pensionär")
	} else if age >= 18 {
		fmt.Println("Du är vuxen")
	} else {
		fmt.Println("Du är barn")
	}
	// and or
	// &&   ||
	if name == "Stefan" && age >= 49 {
		fmt.Println("Du är fortfarande ung")
	}

	// I Go kan funktioner returnera MULTIPLE VALUES
	// ALLA FUNKTIONER I GO RUNTIME RETUNRERAR det den ska PLUS en errorcode
	//FELHANTERING
	_, e := fmt.Scanln(&age)
	if e != nil {
		fmt.Printf("Ngt gick fel")
	} else {
		fmt.Printf("Hej %v ok då du är %v år gammal\n", name, age)
	}

	// Vilken datatyp defaultar var och := ?
	// Så kostnaden av att inte typa med en specifik typ är då kompileringstid?

	// var pi  = 3.1415927
	//pi :=12

	// const country, code = "Sweden", 46

	// pi = pi * 2
	// code = code + 12
}
