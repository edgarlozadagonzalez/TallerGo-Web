package models

import "sort"

type Estudiante struct {
	Index       int
	Nombre      string
	Apellido    string
	Edad        int
	Gender      string
	Email       string
	Phone       string
	Address     string
	About       string
	Matriculado string
	Cursos      []Curso
}

func (e Estudiante) CalcularPromedio() float64 {
	totalCursos := len(e.Cursos)
	if totalCursos == 0 {
		return 0.0
	}

	sumaNotas := 0.0
	for _, curso := range e.Cursos {
		sumaNotas += curso.Nota
	}

	promedio := sumaNotas / float64(totalCursos)
	return promedio
}
func (e Estudiante) BuscarNota(cursoN string) float64 {
	for _, curso := range e.Cursos {
		if curso.Nombre == cursoN {
			return curso.Nota
		}
	}
	return 0.0
}

func MejoresEstudiantes(estudiantes []Estudiante) []Estudiante {
	mejor_promedio := 0.0
	var mejores_estudiantes []Estudiante

	for _, estudiante := range estudiantes {
		promedio := estudiante.CalcularPromedio()
		if promedio > mejor_promedio {
			mejores_estudiantes = []Estudiante{estudiante}
			mejor_promedio = promedio
		} else if promedio == mejor_promedio {
			mejores_estudiantes = append(mejores_estudiantes, estudiante)
		}
	}
	return mejores_estudiantes
}

func PeoresEstudiantes(estudiantes []Estudiante) []Estudiante {
	peor_promedio := 5.0
	var peores_estudiantes []Estudiante

	for _, estudiante := range estudiantes {
		promedio := estudiante.CalcularPromedio()
		if promedio < peor_promedio {
			peores_estudiantes = []Estudiante{estudiante}
			peor_promedio = promedio
		} else if promedio == peor_promedio {
			peores_estudiantes = append(peores_estudiantes, estudiante)
		}
	}
	return peores_estudiantes
}

func Top10MejoresEstudiantesPorCurso(estudiantes []Estudiante) map[string][]Estudiante {
	top10 := make(map[string][]Estudiante)

	for _, estudiante := range estudiantes {
		for _, curso := range estudiante.Cursos {
			estudiantesCurso, existe := top10[curso.Nombre]
			if !existe {
				estudiantesCurso = make([]Estudiante, 0, 10)
			}
			if len(estudiantesCurso) < 10 {
				estudiantesCurso = append(estudiantesCurso, estudiante)
			} else {
				// Encontrar el estudiante con la nota más baja en el top 10
				indiceNotaBaja := 0
				for i := 1; i < 10; i++ {
					if estudiantesCurso[i].BuscarNota(curso.Nombre) < estudiantesCurso[indiceNotaBaja].BuscarNota(curso.Nombre) {
						indiceNotaBaja = i
					}
				}

				// Reemplazar el estudiante con la nota más baja si el estudiante actual tiene una nota más alta
				if estudiante.BuscarNota(curso.Nombre) > estudiantesCurso[indiceNotaBaja].BuscarNota(curso.Nombre) {
					estudiantesCurso[indiceNotaBaja] = estudiante
				}
			}
			// Ordenar los estudiantes por nota de mayor a menor
			sort.Slice(estudiantesCurso, func(i, j int) bool {
				return estudiantesCurso[i].BuscarNota(curso.Nombre) > estudiantesCurso[j].BuscarNota(curso.Nombre)
			})
			top10[curso.Nombre] = estudiantesCurso
		}

	}
	return top10
}

func Top10PeoresEstudiantesPorCurso(estudiantes []Estudiante) map[string][]Estudiante {
	top10 := make(map[string][]Estudiante)

	for _, estudiante := range estudiantes {
		for _, curso := range estudiante.Cursos {
			estudiantesCurso, existe := top10[curso.Nombre]
			if !existe {
				estudiantesCurso = make([]Estudiante, 0, 10)
			}
			if len(estudiantesCurso) < 10 {
				estudiantesCurso = append(estudiantesCurso, estudiante)
			} else {
				// Encontrar el estudiante con la nota más alta en el top 10
				indiceNotaAlta := 0
				for i := 1; i < 10; i++ {
					if estudiantesCurso[i].BuscarNota(curso.Nombre) > estudiantesCurso[indiceNotaAlta].BuscarNota(curso.Nombre) {
						indiceNotaAlta = i
					}
				}

				// Reemplazar el estudiante con la nota más alta si el estudiante actual tiene una nota más baja
				if estudiante.BuscarNota(curso.Nombre) < estudiantesCurso[indiceNotaAlta].BuscarNota(curso.Nombre) {
					estudiantesCurso[indiceNotaAlta] = estudiante
				}
			}
			// Ordenar los estudiantes por nota de menor a mayor
			sort.Slice(estudiantesCurso, func(i, j int) bool {
				return estudiantesCurso[i].BuscarNota(curso.Nombre) < estudiantesCurso[j].BuscarNota(curso.Nombre)
			})
			top10[curso.Nombre] = estudiantesCurso
		}

	}
	return top10
}
