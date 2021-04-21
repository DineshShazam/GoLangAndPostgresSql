package techList

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// custom dataType with slice of string props
type tech []string

// func that produces techlogies list
func newStack() tech {

	techList := tech{}

	frontEnd := []string{"React-js", "react-material-ui", "redux", "React-Hooks"}
	backend := []string{"Go-Lang", "Node-js", "Express-js", "Next-js"}
	database := []string{"postgresSql", "Redis", "MongoDB", "DynomoDB"}
	Cloud := []string{"AWS", "Docker", "Linux", "ECS & Fargate"}

	for _, front := range frontEnd {
		techList = append(techList, front)
	}
	for _, back := range backend {
		techList = append(techList, back)
	}

	for _, databs := range database {
		techList = append(techList, databs)
	}
	for _, cld := range Cloud {
		techList = append(techList, cld)
	}

	return techList
}

// receiver func that prints value
func (t tech) print() {
	for _, techPrint := range t {
		fmt.Println(techPrint)
	}
}

// receiver func that converts to string
//// converts custom type to slice and then to string
func (ss tech) toString() string {
	return strings.Join([]string(ss), ",")
}

// receiver func that writes file
func (techL tech) writeFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(techL.toString()), 0644)
}

// receiver func that reads file
func readFileTech(fileName string) tech {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File Not Found")
	}
	// passing the byte and converting to string
	s := strings.Split(string(bs), ",") // you will get []strig as output
	return tech(s)
}
