package main

import (
	"fmt"

	json "github.com/edgarlozadagonzalez/TallerGo-Web/data"
	models "github.com/edgarlozadagonzalez/TallerGo-Web/models"
)

func main() {

	filename := "data/generated.json"
	estudiantes, err := json.LeerJSON(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("MEJOR ESTUDIANTE:")
	mejoresEstudiantes := models.MejoresEstudiantes(estudiantes)
	for _, estudiante := range mejoresEstudiantes {
		fmt.Printf("Nombre: %s, Promedio: %f\n", estudiante.Nombre, estudiante.CalcularPromedio())
	}

	fmt.Println("PEOR ESTUDIANTE:")
	peoresEstudiantes := models.PeoresEstudiantes(estudiantes)
	for _, estudiante := range peoresEstudiantes {
		fmt.Printf("Nombre: %s, Promedio: %f\n", estudiante.Nombre, estudiante.CalcularPromedio())
	}
}
