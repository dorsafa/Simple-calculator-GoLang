package main
import (
	"fmt"

)
func main() {
	fmt.Println("dodo")
	var power int
	power = 9000
	fmt.Printf("It's over %d\n", power)
	powers := 50
	fmt.Println("powers :", powers)
	x := getPower()
	fmt.Println("powers:", x)

	//// Initialiser structure
	fmt.Println("person :",person{"Dorsaf", 29})
	//// & pointeur sur la structure
	fmt.Println("person poiteur :", &person{"youyou", 1})
	//// Acceder champs de la struvture person
	s := person{"fatoum", 2}
	fmt.Println("age fattouma :", s.age)
	////  pointeurs de structure
	sp := &s
	fmt.Println(sp.age)




}

func getPower() int {
	return 5000
}
///////////////	Structure
type person struct {
	name string
	age  int
}