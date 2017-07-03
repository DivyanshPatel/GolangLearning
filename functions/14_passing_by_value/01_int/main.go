package main

import (
	"fmt"
)

func main(){
	x:=44
	changeMe(x)
	fmt.Println(x)

}

func changeMe(z int){
	fmt.Println(z)
	z=24
}
