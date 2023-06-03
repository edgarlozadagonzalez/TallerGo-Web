package repositories

import "math"

//SOLUCIÓN 7 CADA FUNCIÓN RECIBE UNA LISTA DE NOTAS PARA REALIZAR SU RESPECTIVO CÁLCULO.
func CalcularPromedio(notas []float64) float64 {
	var suma float64
	for _, nota := range notas {
		suma += nota
	}
	return suma / float64(len(notas))
}

func CalcularRango(notas []float64) float64 {
	min := notas[0]
	max := notas[0]
	for _, nota := range notas {
		if nota < min {
			min = nota
		}
		if nota > max {
			max = nota
		}
	}
	return max - min
}

func CalcularVarianza(notas []float64) float64 {
	promedio := CalcularPromedio(notas)
	var sumaCuadrados float64
	for _, nota := range notas {
		sumaCuadrados += math.Pow(nota-promedio, 2)
	}
	return sumaCuadrados / float64(len(notas))
}

func CalcularDesviacionEstandar(notas []float64) float64 {
	varianza := CalcularVarianza(notas)
	return math.Sqrt(varianza)
}
