package main

import(
	"fmt"
)

func maps(){
	// statePop := map[string]int{
	// 	"ath":123,
	// }
	statePop := make(map[string]int)
	//can access this like dictonary
	//delete(statePop, keyname)

	//len(map) gives length of map
	//assignment of map to different variable will be like pass by reference. i.e change new variable will change old variable
	
	//If we try to access some key which doesnot exist it returns 0
	_, is_present := statePop["lol"]
	fmt.Println(is_present)

	//array can be a key but not a slice
}