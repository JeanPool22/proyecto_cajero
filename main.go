package main

import (
	"fmt"

	"github.com/JeanPool22/proyecto_cajero/cajero"
)

func main() {
	var dineroSolicitado int

	fmt.Println("---------- BIENVENIDO A SU CAJERO ----------")
	fmt.Print(`¡Importante!
	* Solo ingresa valores que sean conformados por la
	  combinación de billetes de $10.000, $20.000 y $50.000
	* El monto minimo de retiro son $20.000 pesos`)

	fmt.Print("\n\nPor favor ingrese el dinero que desea retirar: ")
	fmt.Scanln(&dineroSolicitado)

	cajero.Dinero()
	resultado := cajero.Retirar(dineroSolicitado)
	fmt.Printf("\nEl nuevo saldo de su cuenta es: %d\n", resultado)
}
