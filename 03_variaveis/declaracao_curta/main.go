package main

import "fmt"

func main(){
	x:= 45
	x = 50 // Reatribuição com mesmo tipo
	y:= 10
	z:= 3.14
	w:= 2.71
	v:= 1.41
	a:= "Estudo"
	b:= "Golang"
	c:= "today"
	fmt.Println("x =", x)
	fmt.Printf("x = %T\n", x)
	fmt.Println("y =", y)
	fmt.Printf("z = %f | w = %f | v = %f \n", z, w, v)
	fmt.Printf("z = %T | w = %T | v = %T \n", z, w, v)
	fmt.Println(a, b, c)
}
