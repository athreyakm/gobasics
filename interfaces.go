package main

import (
	"fmt"
)

//interface is a type

type writer interface {
	//unlike structs, interfaces does not describe data, it describes behaviour
	//so we define method definitions
	Write([]byte) (int, error) //we defined interface here
}

type incrementer interface {
	increment() int
}

type integer int

func (i integer) read(int) int {
	fmt.Println(i)
	return 1
}

type consoleWriter struct {
}

func (cw consoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func (i integer) Write(data []byte) (int, error) {
	return 100, nil
}

//5:04:00 he explains interface with a db example
//did not understand that

// Naming conventions: try to name your interface by ending with "er"
func interfaceGo() {
	var w writer = consoleWriter{} //create an empty struct object and asssign it to writer
	var s writer = integer(0)      //create an integer object
	a, b := w.Write([]byte("hello Go!"))
	k, _ := s.Write([]byte("lol"))
	fmt.Println(a, b, k)

	// var r reader = integer
}
