package main

import "fmt"

func main() {

	for {
		fmt.Println("1. Skapa device")
		fmt.Println("2. Lista alla")
		fmt.Println("3. Ändsra device")
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
