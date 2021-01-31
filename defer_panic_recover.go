package main

import(
	"fmt"
	// "log"
)

func deferGo(){
	fmt.Println("start")
	//open something and close immediately. Defer makes sure it gets closed at the end of funcc call
	defer fmt.Println("middle, called with defer keyword")		//executed last. Just before return statement
										//if there are multiple defer keywords then they get executed in last in first out order 
										//helpful because sometimes we open some session and hav to do many operations before closing it. later might forget to close it. So use defer keeyword and close it as soon as u open it. 
	fmt.Println("end")
	// fmt.Println("defeer func")


	a := "start"
	defer fmt.Println(a)
	a = "end"		//doesnt print end because argument is taken at the time it was called. So a will be replaced by "start"



}


//we dont have excptions in go
//for ex opening a non existing file is caught as exception in other languages. But in go its not illegal. It returns tuple which has (value, error) So no exception is needed to handle it. 

/*
	WATCH 3:55:00 again
*/


//read more on panic and recover
func panicGo(){
	a:=1
	b := 0
	// defer func(){
	// 	if err := recover(); err != nil{
	// 		log.Println("error:", err)
	// 	}
	// }
	if b == 0{
		panic("divide by 0")	//panic happen after defer statment are executed
								//defer gets called even after panic. We can use this to run some function defer func(){//do something in case of panic. U call recover inside that function...read more about recover}
	}
	ans := a/b
	fmt.Println(ans)
}