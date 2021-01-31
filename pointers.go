package main

import("fmt")


type mystruct struct{
	name int
}
func pointers(){
	var a int = 42
	var b *int = &a

	var ms *mystruct		//by default <nil> will be initialised
//(*ms).name = something		//way to initialise

	ms = &mystruct{
		name : 12,
	}

	//ms  = new(mystruct)		//initialises an object with default 0

	//compiler is smart
	//it reads ms.foo as (*ms).foo to make our life better 

	fmt.Println(ms)		//prints &{12} ms is holding address of an object having 42 in it. 
						//reason for this is most of the time wee dont care about the address. So its not displaying address but its simply saying that it is an address space having 12 in it

	//if u really want address of structure, then u can get it through %p access specifier
	// fmt.Printf("%p", ms)	//prints address and not &{42}
	//no pointer arithmatic with go pointers
	//use unsafe package if u want to use pointer arithmatic 



	//array is not samee as pointer. Slice is projection of an array. So slice contains pointer to the first element the slice is pointing to on the underlying array

	//maps also same as slice. a = map...b := a...means b points to a

	fmt.Println(a, *b)
}