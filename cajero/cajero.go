package cajero

import (
	"fmt"
	"sort"
)

// Declaración de mapa que almacena los billetes y la cantidad que tiene
var billetesDisponibles = map[int]int{
	50000: 25,
	20000: 20,
	10000: 20,
}

// Declaración de la porción que servirá para alamcenar el mapa ordenado
var keys = make([]int, 0, len(billetesDisponibles))

// Declaración del dinero que tiene disponible en el "banco"
var dineroDisponible int = 1_200_000

// Declaración del monto minimo que se puede retirar
var montoMinimoRetiro int = 20_000

// Función que devuelve el dinero disponible para el retiro
func Dinero() {
	fmt.Printf("\nDinero disponible: $%d\n", dineroDisponible)
}

// Función que ordena el mapa por su clave de mayor a menor (50.000 - 20.000 - 10.000)
func ordenarMapa() {
	for k := range billetesDisponibles {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
}

// Función que realiza el conteo de los billetes que recibirá el usuario
func calcularBilletesADevolver(dineroRecibido int) {
	for _, billete := range keys {
		div := int(dineroRecibido / billete)
		fmt.Println(div, "billetes de", billete)
		if int(dineroRecibido%billete) < billete {
			dineroRecibido = dineroRecibido - (billete * div)
			continue
		}
	}
}

// Función que controla el flujo del dinero que se solicita para ser retirado
func Retirar(dineroRecibido int) int {
	nuevoSaldo := 0
	totalDineroRecibido := dineroRecibido

	if dineroDisponible > 0 && dineroRecibido <= dineroDisponible && !(dineroRecibido <= 0) && dineroRecibido >= montoMinimoRetiro {
		fmt.Println("Es posible realizar el retiro de dinero...")
		fmt.Println()
		fmt.Println("Para su retiro recibió la siguiente cantidad de billetes:")

		ordenarMapa()
		calcularBilletesADevolver(dineroRecibido)
		nuevoSaldo = dineroDisponible - totalDineroRecibido

	} else if dineroRecibido <= 0 {
		fmt.Println("Ingrese un valor mayor a 0")
	} else if dineroRecibido < montoMinimoRetiro {
		fmt.Printf("Ingrese un monto mayor a $%d\n", montoMinimoRetiro)
	} else {
		fmt.Println("No tiene fondos suficientes")
	}

	return nuevoSaldo
}
