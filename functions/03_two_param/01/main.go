package main

import "fmt"

func main(){
	greet("Steve" , "Rogers")
}

func greet(fname string, lname string){
	fmt.Println(fname,lname)
}

