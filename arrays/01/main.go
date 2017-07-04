package main

import "fmt"



func main(){

	var x [58]string
	fmt.Println(x)
	fmt.Println(x[42])

	for i:=65 ; i <=100 ; i++ {
		x[i-65] = string(i)


	}

	fmt.Println(x)
	fmt.Println(x[42])



}
