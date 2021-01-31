package main

import(
	"fmt"
)

func functions(msg string) int{		//just like variabls, uppercase functions will be accessable outside package
	fmt.Println(msg)
	return len(msg)
}


func variableParameter(msg string, value ...int){
	for _, v := range(value){
		fmt.Println(v)
	}
}


func interesting() *int{
	result := 0
	return &result		//seg fault in other languages but not in go 
}

func syntaticsugarReturn() (result int){		//not a good practice. Because if func is long, u need to come all the way to top to 										see what is being returned
	result++
	return			//dont have to return with var name since we specified it along with function return type
}


func handlingErrors(a int, b int) (int, error){
	if( b == 0){
		return 0.0, fmt.Errorf("cannot divide by 0")		//what exactly is  the use of errorf?
	}
	return a/b, nil
}


// video 4:45:00 -> talks about async



//methods in go....-> these are same as functions but little different for better usage

//we can create a method for any datatype. For ex in case of structure,
type greeter struct{
	name string
	age int
}

func (g greeter) greet(){			//g* greeter if u want to change value of g
	fmt.Println(g.name, g.age)
}

// lets say we have a struct object called g. now wee can call g.greet()