package main

import (
	"fmt"
	// "reflect"
)

type Animal struct{
	Name string
	Origin string
}



//composition
type Bird struct{
	Animal			//not traditional inheritence. HAS A relation(has animal like charecterstics.) But not an animal
	speed float32
	canFly bool
}

//when to use embedding/composition?


//tag used to describe specific info about a field
//need to import reflect module for this
type Human struct{
	Name string	//`required max :"100"`
	isMale bool  
}

func composition(){
	// b := Bird{}
	// b.Name = "pigon"
	// b.Origin = "asd"
	// b.speed = 123
	// b.canFly = false
	b := Bird{
		Animal : Animal{Name: "emu", Origin:"Australia"},
		speed:48,
		canFly: false,
	}
	
	// c:= Human{
	// 	Name:"Ath",
	// 	isMale:true,
	// }

	fmt.Println(b)
	
	// fmt.Println(c)

	// t := reflect.TypeOf(Human{})		//get the type of the object im working on
	// field, _ := t.FieldByName("Name")	//grab field from this
	// fmt.Println(field.Tag)	//
	// fmt.Println("Composition")
}