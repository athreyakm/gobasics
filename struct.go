package main

import (
	"fmt"
)

type Doctor struct{		//because its capital, accessable outside?
	number int			//fields are not capital hence these cannot be accessed
	actorName string
	companions []string
}



func struct_(){
	aDoctor := Doctor{
		number: 3,
		actorName: "Paddu",
		companions: []string{
			"lol",
			"str",
		},
	}
	//copy of structure are pass by value, not reference
	//need to use & if u  want to use it as pass by reference

	// anonymous struct:
	//aDoctor := struct{name string}{name: "athreya"}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.number)
	// fmt.Println("struct!")



}
