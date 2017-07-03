package main

import "fmt"

func main(){
	s:="Tony"
	changeMe(s)
	fmt.Println(s)


}

func changeMe(z string){
	z="Stark"
}
