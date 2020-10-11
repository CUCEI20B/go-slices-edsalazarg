package main

import "fmt"

func main()  {
	var numeros int
	var numero int
	var s []int
	fmt.Print("Numeros: ")
	fmt.Scan(&numeros)

	for i := 0; i < numeros; i++{
		fmt.Scan(&numero)
		s = append(s,numero)
	}
	var resultado = 0
	for _, v := range s{
		resultado += v
	}
	fmt.Println(resultado)
	
}