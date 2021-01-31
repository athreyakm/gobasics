package main

import(
	"fmt"
)
func arrays(){
	// grades:= [...]int{1,2,3}		//declare an array that can hold all the elements that are passed
	
	//to specify value to an array we do like this
	/*
	var students [2]string
	students[0] = "lololol"
	students[1] = "askdkl"
	fmt.Println(students)		//prints [lololol]
	fmt.Println(string(students[1]))   //prints lololol again but not as an array
	fmt.Println(string(students[1][0]))		//
	fmt.Println(len(students))		//length of array student
	fmt.Println(len(students[0]))		//length of 0th element of string array 
	*/


	//var twoDarray [3][3]int = [3][3]int{[3]int{1,2,3}, [3]int{4,5,6}, [3]int{7,8,9}}
	//var twoDarray [3][3]int = [3][3]int{{1,2,3}, {4,5,6}, {7,8,9}}
	
	
	/*
	var twoDarray [3][3]int
	twoDarray[0] = [3]int{1,2,3}
	twoDarray[1] = [3]int{1,2,3}
	twoDarray[2] = [3]int{1,2,3}
	fmt.Println(twoDarray)
	*/
	

	/*
	a:= [3]int{1,2,3}
	b := a	//everything gets copied. Not just pointing
	fmt.Println(a)
	fmt.Println(b)
	*/

	//b := &a //now a and b are pointing to same thing

	//arrays need to know their size during compile time
	//to solve this problem we have slice
	/*

	a := []int{1,2,3,4,5,6}	//this is slice. Not array. So by default its reference
	
	//SLICEING OPERATIONS WORK BOTH FOR ARRAY AND SLICE
	//i.e a[m:n] way of slicing can be done both for array and slice
	

	fmt.Printf("a is %v\n", a)
	b := a
	fmt.Printf("b is %v\n", b)
	b[2] = 10
	fmt.Printf("After changing b, a is %v, b is %v\n", a, b)
	d := a[:3]
	e := a[2:5]    //elements between 2 and 5, including 2?
	fmt.Printf("d is %v\na is %v\ne is %v\n", d, a, e)
	d[0] = 100		//all elements till 3rd index
	fmt.Printf("After changing d, slice a is  %v, b is %v\n",a,b)
	fmt.Printf("Length: %v\n",len(d))
	fmt.Printf("Capacity: %v....IMP length is not same as capacity\n",cap(d))
	
	*/

	/*

	MAKE
	takes 3 argument

	everytime we append an element and if exisiting ds cant fit it in, new ds is created and everyting is copied. it might be a costly operation. Inorder to make it fast we can use make. especially when we might have some idea that we will use approximately 100 size
	*/

			//type, length, capacity
	// wq := make([]int, 3, 100)	//slice can be of length 3, but underlying array can have 100 elements
								//this because, unlike arrays slices dont need to have fixed size. We can add or remove elements
	wq := []int{}
	//append -> to add new element to the slice
	fmt.Printf("WQ: %v, Capacity: %v  Length: %v\n",wq, cap(wq), len(wq))
	wq = append(wq, 1)
	fmt.Printf("WQ: %v, Capacity: %v, Length: %v\n",wq, cap(wq), len(wq))
	wq = append(wq, 1,2,3,4,5)
	fmt.Printf("WQ: %v, Capacity: %v, Length: %v\n",wq, cap(wq), len(wq))
	
	// wq = append(wq, []int{1,2,3,4,5})		//error because u cant add array
	wq = append(wq, []int{1,2,3,4,5}...)		//spread operator.
	
	fmt.Printf("WQ: %v, Capacity: %v, Length: %v\n",wq, cap(wq), len(wq))
	
	//Stack operations:
	//to remove last element, we can do wq[:len(wq)-1]
	//to remove first element, w can do wq[1:]
	//to remove somewhere in between we can do 
	//k := append(wq[:2], wq[3:]...) //remove 2nd index		//spread operator is a must
				//be careful when u do this. wq will get changed since everything is reference  

	// fmt.Println(wq)


}