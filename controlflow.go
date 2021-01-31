package main

import(
	"fmt"
)

func ifstatement(){
	stateP := map[string]int{
		"athreya": 1,
		"adithya":0,
		"prithvi":3,
	}
	if _, ok := stateP["athreya"]; ok{
		fmt.Println("inside loop")
	}
	
	a := 20
	b := 20

	if a < b{
		fmt.Println("a < b")
	}else if a > b{
		fmt.Println("a > b")
	} else {
		fmt.Println("a = b")
	}
	// fmt.Println("if statements")
}

func switchstatement(){
	// cases := "asdeqwe"
	// switch cases{
	// case "asd", "asdeqwe", "qwsadd":
	// 	fmt.Println(cases)
	// case "qww":
	// 	fmt.Println(cases)
	// case "qwe":
	// 	fmt.Println(cases)
	// default:
	// 	fmt.Println(cases)
	// }
	// cases := "asdeqwe"
	switch cases := "1"+"23"; cases{
	case "asd", "asdeqwe", "qwsadd":
		fmt.Println(cases)   //by  default it is break. If  u want to cosider next case even if this case is successful then add
		//fallthrough			//to undo break. execution doesnot stop and continues to next case even if this case was true
	case "qww":
		fmt.Println(cases)
	case "qwe":
		fmt.Println(cases)
	default:
		fmt.Println(cases)
	}
}