package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)


func main() {
	content, err := ioutil.ReadFile("./adjectives.txt")
	if err != nil {
		fmt.Printf("Oj, błąd wywaliło.")
	}
	lines := strings.Split(string(content), "\n")
	content2, err := ioutil.ReadFile("./nouns.txt")
	if err != nil {
		fmt.Printf("Oj, błąd wywaliło.")
	}
	lines2 := strings.Split(string(content2), "\n")
	fmt.Printf("Czy chcesz liczbę w swoim nicku?")
	var input string
	fmt.Scanf("%s", &input)
 if input == "y" {

	 fmt.Printf(lines[rand.Intn(308)] + "_" + lines2[rand.Intn(308)]) + rand.Intn(100)

 } else if input == "n" {
	 fmt.Printf(lines[rand.Intn(308)] + "_" + lines2[rand.Intn(308)])
} }
