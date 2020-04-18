package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type calculations struct {
	Op    string `json:"Operation"`
	Units units  `json:"Units"`
}

type units struct {
	Num1 int `json:"one"`
	Num2 int `json:"two"`
}

func checker(e error) {
	if e != nil {
		panic(e)
	}
}

var output string = ""

var cals []calculations

func ReadCals(filename string) {
	dat, err := ioutil.ReadFile(filename)
	checker(err)

	if error := json.Unmarshal(dat, &cals); err != nil {
		panic(error)
	}

	for _, cal := range cals {
		num1 := cal.Units.Num1
		num2 := cal.Units.Num2

		switch cal.Op {
		case "Add":
			defer fmt.Println("Addition: ", Add(num1, num2))
		case "Subtract":
			defer fmt.Println("Subtraction: ", Subtract(num1, num2))
		case "Multiply":
			defer fmt.Println("Multiplication: ", Multiply(num1, num2))
		case "Divide":
			defer fmt.Println("Division: ", Divide(num1, num2))
		}
	}

}

func Add(num1, num2 int) int {
	return num1 + num2
}
func Subtract(num1, num2 int) int {
	return num1 - num2
}
func Multiply(num1, num2 int) int {
	return num1 * num2
}
func Divide(num1, num2 int) int {
	return num1 / num2
}

func main() {
	//Question 2
	ReadCals("file2.json")
}
