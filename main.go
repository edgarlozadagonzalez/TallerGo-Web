package main

import (
	"fmt"

	json "github.com/edgarlozadagonzalez/TallerGo-Web/data"
	repositories "github.com/edgarlozadagonzalez/TallerGo-Web/repositories"
)

func main() {

	filename := "data/generated.json"
	data, err := json.LeerJSON(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	estudiantes, err := json.ParsearEstudiantes(data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("MEJOR ESTUDIANTE:")
	mejoresEstudiantes := repositories.MejoresEstudiantes(estudiantes)
	for _, estudiante := range mejoresEstudiantes {
		fmt.Printf("Nombre: %s, Promedio: %f\n", estudiante.Nombre, estudiante.CalcularPromedio())
	}

	fmt.Println("PEOR ESTUDIANTE:")
	peoresEstudiantes := repositories.PeoresEstudiantes(estudiantes)
	for _, estudiante := range peoresEstudiantes {
		fmt.Printf("Nombre: %s, Promedio: %f\n", estudiante.Nombre, estudiante.CalcularPromedio())
	}

	fmt.Println("Top 10 mejores estudiantes por curso:")
	top10MEstudiantesPorCurso := repositories.MejoresEstudiantesPorCurso(estudiantes, 1, 10)
	for _, estudiante := range top10MEstudiantesPorCurso {
		fmt.Printf("- Estudiante: %s %s, Nota: %.2f\n", estudiante.Nombre, estudiante.Apellido, estudiante.BuscarNota(1))
	}

	fmt.Println("Top 10 peores estudiantes por curso:")
	top10PEstudiantesPorCurso := repositories.PeoresEstudiantesPorCurso(estudiantes, 1, 10)
	for _, estudiante := range top10PEstudiantesPorCurso {
		fmt.Printf("- Estudiante: %s %s, Nota: %.2f\n", estudiante.Nombre, estudiante.Apellido, estudiante.BuscarNota(1))
	}
	fmt.Println()

	fmt.Println("Estudiantes masculinos mayores de edad:")
	estudiantesMasculinos := repositories.EstudiantesMasculinos(estudiantes)
	estudiantesMasculinosDeMayorEdad := repositories.EstudiantesMayorEdad(estudiantesMasculinos)
	for _, estudiante := range estudiantesMasculinosDeMayorEdad {
		fmt.Printf("- Estudiante: %s %s, Edad: %d\n", estudiante.Nombre, estudiante.Apellido, estudiante.Edad)
	}
	fmt.Println()
	fmt.Println("Estudiantes femeninos mayores de edad:")
	estudiantesFemeninos := repositories.EstudiantesFemeninos(estudiantes)
	estudiantesFemeninosDeMayorEdad := repositories.EstudiantesMayorEdad(estudiantesFemeninos)
	for _, estudiante := range estudiantesFemeninosDeMayorEdad {
		fmt.Printf("- Estudiante: %s %s, Edad: %d\n", estudiante.Nombre, estudiante.Apellido, estudiante.Edad)
	}

	fmt.Println("Estadisticas por curso: ")
	notas := repositories.ObtenerNotasPorCurso(estudiantes, 1)
	fmt.Println("Promedio: ", repositories.CalcularPromedio(notas))
	fmt.Println("Rango: ", repositories.CalcularRango(notas))
	fmt.Println("Varianza: ", repositories.CalcularVarianza(notas))
	fmt.Println("Desviación estándar: ", repositories.CalcularDesviacionEstandar(notas))

	anio := 2022
	fmt.Println("Reporte estudiantes matriculados en: ", anio)
	estudiantesMatriculados := repositories.ObtenerEstudiantesMatriculadosEnAnio(estudiantes, anio)
	for _, estudiante := range estudiantesMatriculados {
		fmt.Printf("- Estudiante: %s %s, Matriculado: %s\n", estudiante.Nombre, estudiante.Apellido, estudiante.Matriculado)
	}

	edadmin := 20
	edadmax := 29
	fmt.Println("Promedio de estudiantes de: ", edadmin, " a ", edadmax, " años")
	notasRangoEdad := repositories.ObtenerNotasPorRangoEdad(estudiantes, edadmin, edadmax)
	promedioRangoEdad := repositories.CalcularPromedio(notasRangoEdad)
	fmt.Print(promedioRangoEdad)
}
