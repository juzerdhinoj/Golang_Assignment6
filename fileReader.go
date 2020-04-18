package Question1

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadText(fileName string) {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	fmt.Print(string(dat))
}

func ReadJSON(filename string) {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	var data map[string]interface{}
	if error := json.Unmarshal(dat, &data); err != nil {
		panic(error)
	}

	//Use Struct or read directly as per your need
	//will use struct in next question's solution

	//fmt.Println(data["Operation"])

	// for item := range data {
	// 	fmt.Println(item[1])
	// }
}

func ReadCSV(fileName string) {
	csvFile, err := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	check(err)
	for {
		line, error := reader.Read()
		if error != nil {
			break
		}
		fmt.Println(line)
	}
}

func main() {

	//Question 1
	ReadText("file1.txt")
	//ReadJSON("file2.json")
	//ReadCSV(filename)
}
