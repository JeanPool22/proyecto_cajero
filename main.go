package main

import (
	"fmt"

	"github.com/JeanPool22/proyecto_cajero/cajero"
)

func main() {
	var dineroSolicitado int

	fmt.Println("---------- BIENVENIDO A SU CAJERO ----------")
	fmt.Println(`¡Importante!
	* Solo ingresa valores que sean conformados por la
	combinación de billetes de $20.000 y $50.000
	* El monto minimo de retiro son $20.000 pesos`)

	fmt.Print("Por favor ingrese el dinero que desea retirar: ")
	fmt.Scanln(&dineroSolicitado)

	cajero.Dinero()
	resultado := cajero.Retirar(dineroSolicitado)
	fmt.Printf("El nuevo saldo de su cuenta es: %d\n", resultado)
}
