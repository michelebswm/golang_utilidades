package main

import "fmt"

type Altura float64  // em metros
type Peso float64    // em quilos

type Celsius float64
type Fahrenheit float64

func imc(altura Altura, peso Peso) float64{
	return  float64(peso) / float64(altura * altura)
}

func celsiusToFahrenheit(celsius Celsius) Fahrenheit{
	return Fahrenheit(float64(celsius) * 9./5. + 32)
}

func main() {
	alturaMi := Altura(1.62)
	pesoMi := Peso(85.00)

	//res := alturaMi + pesoMi  // invalid operation: alturaMi + pesoMi (mismatched types Altura and Peso)compilerMismatchedTypes

	res := imc(alturaMi, pesoMi)
	fmt.Printf("Cálculo do IMC: %.2f\n", res)

	
	grausCelsius := Celsius(21.4)
	grausFahrenheit := celsiusToFahrenheit(grausCelsius)

	fmt.Printf("%f graus Celsius correspondem a %f graus Fahrenheit.\n", grausCelsius, grausFahrenheit)


}
