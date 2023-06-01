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

	top10MEstudiantesPorCurso := models.Top10MejoresEstudiantesPorCurso(estudiantes)

	fmt.Println("Top 10 mejores estudiantes por curso:")
	for curso, estudiantes := range top10MEstudiantesPorCurso {
		fmt.Printf("Curso: %s\n", curso)
		for _, estudiante := range estudiantes {
			fmt.Printf("- Estudiante: %s %s, Promedio: %.2f\n", estudiante.Nombre, estudiante.Apellido, estudiante.BuscarNota(curso))
		}
		fmt.Println()
	}

	top10PEstudiantesPorCurso := models.Top10PeoresEstudiantesPorCurso(estudiantes)
	fmt.Println("Top 10 peores estudiantes por curso:")
	for curso, estudiantes := range top10PEstudiantesPorCurso {
		fmt.Printf("Curso: %s\n", curso)
		for _, estudiante := range estudiantes {
			fmt.Printf("- Estudiante: %s %s, Promedio: %.2f\n", estudiante.Nombre, estudiante.Apellido, estudiante.BuscarNota(curso))
		}
		fmt.Println()
	}
}
