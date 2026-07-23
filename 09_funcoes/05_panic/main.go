package main

import (
	"fmt"
	"strconv"
)

func validaEntrada() float64 {
	defer func(){
		txt := recover()
		fmt.Println(txt)
	}()

	var input string
	fmt.Println("Informe um float: ")
	fmt.Scanf("%s", &input)

	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic("Entrada inválida")
	} 

	return  num

}

func main() {

	num := validaEntrada()
	fmt.Println(num)
}
