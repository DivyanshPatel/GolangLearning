package main

import "fmt"

func main(){
	greet("Tony" , "Stark")
}

func greet(fname , lname string){
	fmt.Println(fname,lname)
}