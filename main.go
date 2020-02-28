package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//alex := person{"Alex", "Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)

	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"

	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	//jim.print()

	//&jim esta atribuindo o ENDEREÇO DE MEMORIA de "jim"
	//para jimPointer, assim podemos dar update no nome
	//de Jim, mudando direitamente pelo endereço de memoria
	//de "jim"
	//jimPointer := &jim
	//jimPointer.updateName("Jimmy")

	//Atalho: como na func updateName tem o receiver *person
	//e "jim" também é do tipo person, Go consegue interpretar
	//que ambos são do mesmo tipo e que você quer o ENDEREÇO
	//de jim, e não o seu valor, ja que ambos "jim" e "pointerToPerson"
	//são do tipo... person
	jim.updateName("Jimmy")
	jim.print()

}

//o *person informa q queremos um endereço de memoria
// chamado "pointerToPerson" do Type person
func (pointerToPerson *person) updateName(newFirstName string) {
	//Enquanto *pointerToPerson quer pegar o VALOR
	//existente dentro do endereço de memoria "pointerToPerson"
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
