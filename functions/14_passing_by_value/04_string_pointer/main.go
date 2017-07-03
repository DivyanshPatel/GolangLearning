package main

import "fmt"

func main(){
	s:="Tony"
	fmt.Println(&s)
	changeMe(&s)
	fmt.Println(s)


}

func changeMe(z *string){
	*z="Stark"
}
