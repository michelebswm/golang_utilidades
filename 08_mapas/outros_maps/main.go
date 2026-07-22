package main

import "fmt"

func main(){
	notas := map[string][]float64{}

	notas["michele"] = []float64{10, 9.85, 9.89, 10, 10}
	notas["ana"] = []float64{8, 9.85, 9.89, 10}

	fmt.Println(notas)
	fmt.Println("Notas michele:", notas["michele"])

	cursos := map[string][]string{
		"michele": []string{"go", "python", "java"},
		"ana": []string{"viagem", "motorhome", "lazer"},
	}
	cursos["lara"] = []string{"Arquitetura", "Design"}
	fmt.Println(cursos)


}