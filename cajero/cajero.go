package cajero

import "fmt"

var billetesDisponibles = map[int]int{
	50000: 25,
	20000: 20,
	10000: 20,
}

var dineroDisponible int = 1_200_000

func Billetes() {
	fmt.Println(billetesDisponibles)
}

func Dinero() {
	fmt.Printf("Dinero disponible: %d\n", dineroDisponible)
}

func Retirar(dineroRecibido int) int {
	nuevoSaldo := 0
	totalDineroRecibido := dineroRecibido
	if dineroDisponible > 0 && dineroRecibido < dineroDisponible && !(dineroRecibido <= 0) {
		fmt.Println("Es posible realizar el retiro de dinero...")
		fmt.Println("Para su retiro recibio la siguiente cantidad de billetes:")
		for billete := range billetesDisponibles {
			div := int(dineroRecibido / billete)
			fmt.Println(div, "billetes de", billete)
			if int(dineroRecibido%billete) < billete {
				dineroRecibido = dineroRecibido - (billete * div)
				continue
			}
		}
		nuevoSaldo = dineroDisponible - totalDineroRecibido
	} else if dineroRecibido <= 0 {
		fmt.Println("Ingrese un valor mayor a 0")
	} else {
		fmt.Println("No tiene fondos suficientes")
	}

	return nuevoSaldo
}
