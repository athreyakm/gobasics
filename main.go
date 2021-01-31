package main

import (
	"fmt"
	// "github.com/athreyakm/"
	"strconv" //to convert into to string
)

//declaring variable at package level(?)
//cant declare as a := 20
//need to  use full way of declaration

var i float32 = 27 //gives better control on type.U can declare integer and call it a float

//other way of declaring multiple variables
var (
	actorName string = "athreya"
	doctor    string = "paddu"
	cost      int    = 25
) //package level variables are not compulsory to be used?

//package level declaration and is lowercase, then any file in the same package can access the variable
//package level declaration and is uppercase, then its globally visible

func variable() {
	//inside this also  uppercase variable will export and make globally visible or vvisible for entire package? -> try this out

	//every time u initialise a value it will be 0
	// variables declared here are restricted to this particular scope
	//variables should always be used. U declare and dont use it, it will throw errors. It will be a compile time errors
	var k int
	//if a variable is declared both at package level and inner level, inner most level takes presedence

	// var i float32 = 27		//gives better control on type.U can declare integer and call it a float

	k = 25
	j := 28.

	var convertType float32
	convertType = float32(k)
	fmt.Println(convertType) //should convert explicitily because sometimes when u convert float to integer, u lose some decimal point so gives error if u try to do below
	//convertType = k   //error if u are losing some info. i.e decimal values

	fmt.Println(strconv.Itoa(k))

	fmt.Printf("%v -%T\n %v - %T\n %v - %T\n", i, i, j, j, k, k) //%v -> value     %T -> type
}

func primitives() {
	// var n bool = true
	// k := 1==2
	// j := 3==3
	// fmt.Println(n, k, j)
	// n_ := 42

	//cant add 2 number of different types
	//cant divide multiple types?

	//has or and (^ -> xor), (&^ -> andnot)?

	// a << 3   -> leftshit by 3 points
	//a >> 3    ->       ''   ''

	//there is a type called byte as well?

	//there is also a type called complex       var n complex64 = 1+2i		//real(n) givs real side, imag(n) givees imaginary part

	//var n uint16 = 42
	//fmt.Printf("%v %T\n", n, n)

	// s := "Hello this is athreya "	//strings in go are aliases for bytes.       //strings are immutable

	// fmt.Printf("%v %T", s[2], s[2])		//printz integer, not charecter
	// fmt.Printf("%v %T", string(s[2]), s[2])

	// y := "string to test concantination"
	// z := s+y

	// can convert string to collection of bytes
	// b := []byte(s)		//usually helpful when we want to  send as response in webserver or in some other cases
	// fmt.Println(b)
	//needs to be very careful with bytes

	// r:='a'		//runes
	//same as
	var r rune
	fmt.Printf("%v %T", r, r)
}

// const a = iota		//iota is a counter that can be used when we create enumaritive constant

const (
	k = iota
	j = iota
	l = iota
)

const (
	_ = iota //i dont care about this value. So i dont need a variable name for that
	a
	b
	c
)

//below example is very interesting
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

// 1:43:59 has an excellent example on efficiently using iota property

func constants() {
	// const myConst int =42		//const has to be signable at runtime(?) 1:29:40 in video
	//constants should never be assigned at runtime

	//create a constant at package level and then redeclare inside a function, inner declaraition wins
	// fmt.Printf("%v %T", myConst,myConst)

	// fmt.Printf("%v, %T", a, a)

	// fmt.Printf("%v\n", k)
	// fmt.Printf("%v\n", j)
	// fmt.Printf("%v\n", l)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	// myConst++ //error
}

func main() {
	fmt.Println("HI")
	// ifstatement()
	// switchstatement()
	// loops()
	// deferGo()
	// panicGo()
	// pointers()
	// a,b := handlingErrors(1,0)
	interfaceGo()
	// interface2Go()
	// fmt.Println(a, b)
}
