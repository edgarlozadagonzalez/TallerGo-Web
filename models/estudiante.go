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

func (e *Estudiante) CalcularPromedio() float64 {
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

func (e *Estudiante) BuscarNota(cursoId int) float64 {
	for _, curso := range e.Cursos {
		if curso.ID == cursoId {
			return curso.Nota
		}
	}
	return 0.0
}
