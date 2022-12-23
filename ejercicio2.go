package main

import "fmt"

func main() {

	calificaciones := []float64{}
	var cantidad int
	var nota float64
	fmt.Println("Cuantas notas desea agregar?")
	fmt.Scanln(&cantidad)
	for i := 0; i < cantidad; i++ {
		for {
			fmt.Printf("Nota %v: ", i+1)
			fmt.Scanln(&nota)
			if nota > 0 {
				break
			} else {
				fmt.Println("La nota no puede ser menor a 0")
				fmt.Println("Ingrese nuevamente la nota porfavor")
			}
		}
		calificaciones = append(calificaciones, nota)
	}

	result := Promedio(calificaciones...)
	fmt.Println("El promedio total es ", result)
}

func Promedio(calificaciones ...float64) float64 {
	var promedio float64
	for _, v := range calificaciones {
		promedio += v
	}
	return promedio / float64(len(calificaciones))
}
