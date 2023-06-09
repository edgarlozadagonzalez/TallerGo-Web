package repositories

import (
	"sort"
	"time"

	models "github.com/edgarlozadagonzalez/TallerGo-Web/models"
)

// SOLUCIÓN 1 Y 2 FUNCIONES NECESARIAS PARA OBTENER MEJORES Y PEORES PROMEDIOS DE ESTUDIANTES
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

// SOLUCIÓN 3 Y 4 FUNCIONES NECESARIAS PARA OBTENER EL TOP N MEJORES Y PEORES ESTUDIANTES DE CADA CURSO

func ObtenerEstudiantesPorCurso(estudiantes []models.Estudiante, idCurso int) []models.Estudiante {
	var estudiantesCurso []models.Estudiante
	for _, estudiante := range estudiantes {
		for _, curso := range estudiante.Cursos {
			if curso.ID == idCurso {
				estudiantesCurso = append(estudiantesCurso, estudiante)
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

// SOLUCIÓN 5 Y 6 FUNCIONES NECESARIAS PARA OBTENER LOS ESTUDIANTES MASCULINOS, FEMENINOS DE MAYOR EDAD
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

// SOLUCIÓN 7 FUNCIONES NECESARIAS PARA REALIZAR CÁLCULOS ESTADÍSTICOS A LAS NOTAS DE CADA CURSO

func ObtenerNotasPorCurso(estudiantes []models.Estudiante, idCurso int) []float64 {
	estudiantesCurso := ObtenerEstudiantesPorCurso(estudiantes, idCurso)
	notas := []float64{}
	for _, estudiante := range estudiantesCurso {
		for _, curso := range estudiante.Cursos {
			if curso.ID == idCurso {
				notas = append(notas, curso.Nota)
			}
		}
	}
	return notas
}

// SOLUCIÓN 8 FUNCIONES NECESARIAS PARA REALIZAR EL REPORTE DE LOS ESTUDIANTES QUE SE MATRICULARON EN UN AÑO X
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

// SOLUCIÓN 9 FUNCIONES NECESARIAS PARA OBTENER EL PROMEDIO DE NOTA DE LOS ESTUDIANTES EN UN RANGO DE EDAD
func ObtenerEstudiantesPorRangoEdad(estudiantes []models.Estudiante, edadMin int, edadMax int) []models.Estudiante {
	var estudiantesRango []models.Estudiante

	for _, estudiante := range estudiantes {
		if estudiante.Edad >= edadMin && estudiante.Edad <= edadMax {
			estudiantesRango = append(estudiantesRango, estudiante)
		}
	}
	return estudiantesRango
}

func ObtenerNotas(estudiantes []models.Estudiante) []float64 {
	notas := []float64{}
	for _, estudiante := range estudiantes {
		for _, curso := range estudiante.Cursos {
			notas = append(notas, curso.Nota)
		}
	}
	return notas
}

func ObtenerNotasPorRangoEdad(estudiantes []models.Estudiante, edadMin int, edadMax int) []float64 {
	estudiantesRangoEdad := ObtenerEstudiantesPorRangoEdad(estudiantes, edadMin, edadMax)
	return ObtenerNotas(estudiantesRangoEdad)
}

func AgregarEstudiante(estudiantes []models.Estudiante, estudiante models.Estudiante) []models.Estudiante {
	estudiantes = append(estudiantes, estudiante)
	return estudiantes
}

func BuscarEstudiante(estudiantes []models.Estudiante, index int) *models.Estudiante {
	for _, estudiante := range estudiantes {
		if estudiante.Index == index {
			return &estudiante
		}
	}
	return nil
}

func ActualizarEstudiante(estudiantes []models.Estudiante, estudiante models.Estudiante) []models.Estudiante {
	for i, est := range estudiantes {
		if est.Index == estudiante.Index {
			estudiantes[i] = estudiante
		}
	}
	return estudiantes
}

func CursoExistente(cursoAgregar models.Curso, estudiante models.Estudiante) bool {
	for _, curso := range estudiante.Cursos {
		if curso.ID == cursoAgregar.ID {
			return true
		}
	}
	return false
}

func AgregarCurso(curso models.Curso, estudiante models.Estudiante) models.Estudiante {
	estudiante.Cursos = append(estudiante.Cursos, curso)
	return estudiante
}
