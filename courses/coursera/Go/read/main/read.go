package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)
/**
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).
Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
 */

func main() {
	type name struct {
		fname string
		lname string
	}
	var names []name
	var filename string
	fmt.Print("Enter name of file: ")
	fmt.Scan(&filename)
	file, _ := ioutil.ReadFile(filename)
	content := string(file)
	arr := strings.Split(content, "\n")
	for index := 0; index < len(arr); index++ {
		temp := strings.Split(arr[index], " ")
		names = append(names, name { temp[0], temp[1] } )
	}

	for index:= 0; index < len(names); index++ {
		fmt.Println("First Name: ", names[index].fname, "Last Name: ", names[index].lname)
	}
}
