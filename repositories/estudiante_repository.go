package repositories

import (
	"sort"
	"time"

	models "github.com/edgarlozadagonzalez/TallerGo-Web/models"
)

// SOLUCION 1 Y 2 FUNCIONES NECESARIAS PARA OBTENER MEJORES Y PEORES PROMEDIOS DE ESTUDIANTES
func MejoresEstudiantes(estudiantes []models.Estudiante) []models.Estudiante {
	mejor_promedio := 0.0
	var mejores_estudiantes []models.Estudiante

	for _, estudiante := range estudiantes {
		promedio := estudiante.CalcularPromedio()
		if promedio > mejor_promedio {
			mejores_estudiantes = []models.Estudiante{estudiante}
			mejor_promedio = promedio
		} else if promedio == mejor_promedio {
			mejores_estudiantes = append(mejores_estudiantes, estudiante)
		}
	}
	return mejores_estudiantes
}

func PeoresEstudiantes(estudiantes []models.Estudiante) []models.Estudiante {
	peor_promedio := 5.0
	var peores_estudiantes []models.Estudiante

	for _, estudiante := range estudiantes {
		promedio := estudiante.CalcularPromedio()
		if promedio < peor_promedio {
			peores_estudiantes = []models.Estudiante{estudiante}
			peor_promedio = promedio
		} else if promedio == peor_promedio {
			peores_estudiantes = append(peores_estudiantes, estudiante)
		}
	}
	return peores_estudiantes
}

// SOLUCION 3 Y 4 FUNCIONES NECESARIAS PARA OBTENER EL TOP N MEJORES Y PEORES ESTUDIANTES DE CADA CURSO

func ObtenerEstudiantesPorCurso(estudiantes []models.Estudiante, idCurso int) []models.Estudiante {
	var estudiantesCurso []models.Estudiante
	for _, estudiante := range estudiantes {
		for _, curso := range estudiante.Cursos {
			if curso.ID == idCurso {
				estudiantesCurso = append(estudiantesCurso, estudiante)
				break
			}
		}
	}
	return estudiantesCurso
}

func MejoresEstudiantesPorCurso(estudiantes []models.Estudiante, idCurso int, n int) []models.Estudiante {
	estudiantesCurso := ObtenerEstudiantesPorCurso(estudiantes, idCurso)

	sort.Slice(estudiantesCurso, func(i, j int) bool {
		return estudiantesCurso[i].BuscarNota(idCurso) > estudiantesCurso[j].BuscarNota(idCurso)
	})
	if n > len(estudiantesCurso) {
		n = len(estudiantesCurso)
	}
	return estudiantesCurso[:n]
}

func PeoresEstudiantesPorCurso(estudiantes []models.Estudiante, idCurso int, n int) []models.Estudiante {
	estudiantesCurso := ObtenerEstudiantesPorCurso(estudiantes, idCurso)

	sort.Slice(estudiantesCurso, func(i, j int) bool {
		return estudiantesCurso[i].BuscarNota(idCurso) < estudiantesCurso[j].BuscarNota(idCurso)
	})
	if n > len(estudiantesCurso) {
		n = len(estudiantesCurso)
	}
	return estudiantesCurso[:n]
}

// SOLUCION 5 Y 6 FUNCIONES NECESARIAS PARA OBTENER LOS ESTUDIANTES MASCULINOS, FEMENINOS DE MAYOR EDAD

func EstudiantesFemeninos(estudiantes []models.Estudiante) []models.Estudiante {
	estudiantesFemeninos := []models.Estudiante{}

	for _, estudiante := range estudiantes {
		if estudiante.Gender == "female" {
			estudiantesFemeninos = append(estudiantesFemeninos, estudiante)
		}
	}

	return estudiantesFemeninos
}

func EstudiantesMasculinos(estudiantes []models.Estudiante) []models.Estudiante {
	estudiantesMasculinos := []models.Estudiante{}

	for _, estudiante := range estudiantes {
		if estudiante.Gender == "male" {
			estudiantesMasculinos = append(estudiantesMasculinos, estudiante)
		}
	}

	return estudiantesMasculinos
}

func EstudiantesMayorEdad(estudiantes []models.Estudiante) []models.Estudiante {
	estudiantesMayorEdad := []models.Estudiante{}
	mayorEdad := -1

	for _, estudiante := range estudiantes {
		if estudiante.Edad > mayorEdad {
			mayorEdad = estudiante.Edad
			estudiantesMayorEdad = []models.Estudiante{estudiante}
		} else if estudiante.Edad == mayorEdad {
			estudiantesMayorEdad = append(estudiantesMayorEdad, estudiante)
		}
	}

	return estudiantesMayorEdad
}

// SOLUCION 7 FUNCIONES NECESARIAS PARA REALIZAR CALCULOS ESTADISTICOS A LAS NOTAS DE CADA CURSO

func ObtenerNotas(estudiantes []models.Estudiante) []float64 {
	notas := []float64{}
	for _, estudiante := range estudiantes {
		for _, curso := range estudiante.Cursos {
			notas = append(notas, curso.Nota)
		}
	}
	return notas
}

func ObtenerNotasPorCurso(estudiantes []models.Estudiante, idCurso int) []float64 {
	estudiantesCurso := ObtenerEstudiantesPorCurso(estudiantes, idCurso)
	return ObtenerNotas(estudiantesCurso)
}

// SOLUCION 8 FUNCIONES NECESARIAS PARA REALIZAR EL REPORTE DE LOS ESTUDIANTES QUE SE MATRICULARON EN UN AÑO X
func ObtenerEstudiantesMatriculadosEnAnio(estudiantes []models.Estudiante, anio int) []models.Estudiante {
	var estudiantesMatriculados []models.Estudiante

	for _, estudiante := range estudiantes {
		matriculado, err := time.Parse("2006-01-02T15:04:05 -07:00", estudiante.Matriculado)
		if err != nil {
			continue
		}

		if matriculado.Year() == anio {
			estudiantesMatriculados = append(estudiantesMatriculados, estudiante)
		}
	}
	return estudiantesMatriculados
}

// SOLUCION 9 FUNCIONES NECESARIAS PARA OBTENER EL PROMEDIO DE NOTA DE LOS ESTUDIANTES EN UN RANGO DE EDAD
func ObtenerEstudiantesPorRangoEdad(estudiantes []models.Estudiante, edadMin int, edadMax int) []models.Estudiante {
	var estudiantesRango []models.Estudiante

	for _, estudiante := range estudiantes {
		if estudiante.Edad >= edadMin && estudiante.Edad <= edadMax {
			estudiantesRango = append(estudiantesRango, estudiante)
		}
	}
	return estudiantesRango
}

func ObtenerNotasPorRangoEdad(estudiantes []models.Estudiante, edadMin int, edadMax int) []float64 {
	estudiantesRangoEdad := ObtenerEstudiantesPorRangoEdad(estudiantes, edadMin, edadMax)
	return ObtenerNotas(estudiantesRangoEdad)
}
