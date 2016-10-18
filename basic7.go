package main

import "fmt"

func out(sun int) (x, y int){
	x = sun * 7 / 2
	y = sun - x
	return
}

func main(){
	fmt.Print(out(10))
}
