package main

import(
	"fmt"
)

func loops(){
	//easy...read when u need it lol
	k := make([]int, 10, 100)
	// s := []int{1,2,3,4}
	for i,j := range(k){
		fmt.Println(i,j)
	}

	d := map[int]int{
		1:2,
		3:4,
		5:6,
	}
	for _, j := range(d){
		fmt.Println(j)
	}
}