package cars

//folder är bättre än classer , att förstå som nybörjare

// protected, private, public
// i GO finns inte access modifiers
// Jo  private, public finns i GO fast inte som KEYWORD
// liten bokstav = private
// stor bokstav = public
// acess modifiers är inte KLASS -scopade
// MODUL/PACKAGE-scope
func GetCarManufacturersWeSell() []string {
	allCars := []string{"Volvo", "Renault", "Saab"}
	return allCars
}
