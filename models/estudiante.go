package models

type Estudiante struct {
	Index       int     `json:"index"`
	Nombre      string  `json:"nombre"`
	Apellido    string  `json:"apellido"`
	Edad        int     `json:"edad"`
	Gender      string  `json:"gender"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Address     string  `json:"address"`
	About       string  `json:"about"`
	Matriculado string  `json:"matriculado"`
	Cursos      []Curso `json:"cursos"`
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
