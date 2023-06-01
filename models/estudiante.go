package models

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
