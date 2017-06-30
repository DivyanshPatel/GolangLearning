package main

import "fmt"

func main() {

	greet("Alex")
	greet("Steve")

}

func greet(name string){
	fmt.Println(name)
}

