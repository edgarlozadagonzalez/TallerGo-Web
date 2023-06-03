package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"

	json "github.com/edgarlozadagonzalez/TallerGo-Web/data"
	"github.com/edgarlozadagonzalez/TallerGo-Web/models"
	"github.com/edgarlozadagonzalez/TallerGo-Web/repositories"
)

type PageData struct {
	Title string
	H1    string
	Icon  string
	Data  interface{}
	Ruta  string
}

func Index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "TallerGO-Web",
	}
	filepath := "views/index.html"
	CargarTemplate(w, filepath, data)
}

func TopMejoresEstudiantes(w http.ResponseWriter, r *http.Request) {

	cursoIdStr := r.FormValue("cursoID")
	topStr := r.FormValue("top")
	opciones := map[int]string{
		0: "Algebra lineal",
		1: "Calculo diferencial",
		2: "POO",
		3: "CTD",
	}
	filepath := "views/topEstudiantesCurso.html"
	data := PageData{
		Title: "TallerGO-Web | Top Mejores Estudiantes",
		H1:    "Top mejores estudiantes por curso",
		Icon: `<svg xmlns="http://www.w3.org/2000/svg" class="bi mt-4 mb-3" style="color: var(--bs-teal);" width="100"
        height="100" fill="currentColor" class="bi bi-trophy-fill" viewBox="0 0 16 16">
            <path d="M2.5.5A.5.5 0 0 1 3 0h10a.5.5 0 0 1 .5.5c0 .538-.012 1.05-.034 1.536a3 3 0 1 1-1.133 5.89c-.79 1.865-1.878 2.777-2.833 3.011v2.173l1.425.356c.194.048.377.135.537.255L13.3 15.1a.5.5 0 0 1-.3.9H3a.5.5 0 0 1-.3-.9l1.838-1.379c.16-.12.343-.207.537-.255L6.5 13.11v-2.173c-.955-.234-2.043-1.146-2.833-3.012a3 3 0 1 1-1.132-5.89A33.076 33.076 0 0 1 2.5.5zm.099 2.54a2 2 0 0 0 .72 3.935c-.333-1.05-.588-2.346-.72-3.935zm10.083 3.935a2 2 0 0 0 .72-3.935c-.133 1.59-.388 2.885-.72 3.935z"/>
          </svg>`,
		Data: nil,
		Ruta: "/curso/topMejoresEstudiantes",
	}
	if cursoIdStr != "" && topStr != "" {
		cursoID, err := strconv.Atoi(cursoIdStr)
		if err != nil {
			http.Error(w, "El numero de codigo no es válido", http.StatusBadRequest)
			return
		}
		top, err := strconv.Atoi(topStr)
		if err != nil {
			http.Error(w, "El numero de top no es válido", http.StatusBadRequest)
			return
		}
		estudiantes := ObtenerEstudiantes()
		mejoresEstudiantesCurso := repositories.MejoresEstudiantesPorCurso(estudiantes, cursoID, top)
		if len(mejoresEstudiantesCurso) == 0 {
			data.H1 = fmt.Sprintf("No hay estudiantes en el curso de %s", opciones[cursoID])
			data.Data = nil
		}
		if len(mejoresEstudiantesCurso) < top {
			data.H1 = fmt.Sprintf("No hay %d estudiantes en el curso de %s", top, opciones[cursoID])
			data.Data = nil
		} else {
			data.H1 = fmt.Sprintf("Top %d de mejores estudiantes del curso %s", top, opciones[cursoID])
			estudiantesConNota := make([]map[string]interface{}, 0)
			for _, estudiante := range mejoresEstudiantesCurso {
				estudianteConNota := map[string]interface{}{
					"Estudiante": estudiante,
					"Nota":       estudiante.BuscarNota(cursoID),
				}
				estudiantesConNota = append(estudiantesConNota, estudianteConNota)
			}
			data.Data = estudiantesConNota
		}
	}
	CargarTemplate(w, filepath, data)
}

func TopPeoresEstudiantes(w http.ResponseWriter, r *http.Request) {
	cursoIdStr := r.FormValue("cursoID")
	topStr := r.FormValue("top")
	opciones := map[int]string{
		0: "Algebra lineal",
		1: "Calculo diferencial",
		2: "POO",
		3: "CTD",
	}
	filepath := "views/topEstudiantesCurso.html"
	data := PageData{
		Title: "TallerGO-Web | Top Peores Estudiantes",
		H1:    "Top peores estudiantes por curso",
		Icon: `<svg xmlns="http://www.w3.org/2000/svg" class="bi mt-4 mb-3" style="color: var(--bs-teal);" width="100"
        height="100" fill="currentColor" class="bi bi-trophy-fill" viewBox="0 0 16 16">
            <path d="M2.5.5A.5.5 0 0 1 3 0h10a.5.5 0 0 1 .5.5c0 .538-.012 1.05-.034 1.536a3 3 0 1 1-1.133 5.89c-.79 1.865-1.878 2.777-2.833 3.011v2.173l1.425.356c.194.048.377.135.537.255L13.3 15.1a.5.5 0 0 1-.3.9H3a.5.5 0 0 1-.3-.9l1.838-1.379c.16-.12.343-.207.537-.255L6.5 13.11v-2.173c-.955-.234-2.043-1.146-2.833-3.012a3 3 0 1 1-1.132-5.89A33.076 33.076 0 0 1 2.5.5zm.099 2.54a2 2 0 0 0 .72 3.935c-.333-1.05-.588-2.346-.72-3.935zm10.083 3.935a2 2 0 0 0 .72-3.935c-.133 1.59-.388 2.885-.72 3.935z"/>
          </svg>`,
		Data: nil,
		Ruta: "/curso/topPeoresEstudiantes",
	}
	if cursoIdStr != "" && topStr != "" {
		cursoID, err := strconv.Atoi(cursoIdStr)
		if err != nil {
			http.Error(w, "El numero de codigo no es válido", http.StatusBadRequest)
			return
		}
		top, err := strconv.Atoi(topStr)
		if err != nil {
			http.Error(w, "El numero de top no es válido", http.StatusBadRequest)
			return
		}
		estudiantes := ObtenerEstudiantes()
		peoresEstudiantesCurso := repositories.PeoresEstudiantesPorCurso(estudiantes, cursoID, top)
		if len(peoresEstudiantesCurso) == 0 {
			data.H1 = fmt.Sprintf("No hay estudiantes en el curso de %s", opciones[cursoID])
			data.Data = nil
		}
		if len(peoresEstudiantesCurso) < top {
			data.H1 = fmt.Sprintf("No hay %d estudiantes en el curso de %s", top, opciones[cursoID])
			data.Data = nil
		} else {
			data.H1 = fmt.Sprintf("Top %d de peores estudiantes del curso %s", top, opciones[cursoID])

			estudiantesConNota := make([]map[string]interface{}, 0)
			for _, estudiante := range peoresEstudiantesCurso {
				estudianteConNota := map[string]interface{}{
					"Estudiante": estudiante,
					"Nota":       estudiante.BuscarNota(cursoID),
				}
				estudiantesConNota = append(estudiantesConNota, estudianteConNota)
			}
			data.Data = estudiantesConNota
		}
	}
	CargarTemplate(w, filepath, data)
}

func CalcularEstadisticas(w http.ResponseWriter, r *http.Request) {
	cursoIdStr := r.FormValue("cursoID")
	opciones := map[int]string{
		0: "Algebra lineal",
		1: "Calculo diferencial",
		2: "POO",
		3: "CTD",
	}
	filepath := "views/estadisticas.html"
	data := PageData{
		Title: "TallerGO-Web | Calcular Estadísticas",
		H1:    "Calcular estadísticas de notas por curso",
		Icon: `<svg xmlns="http://www.w3.org/2000/svg" class="bi mt-4 mb-3" style="color: var(--bs-teal);" width="100"
        height="100" fill="currentColor" class="bi bi-trophy-fill" viewBox="0 0 16 16">
            <path d="M2.5.5A.5.5 0 0 1 3 0h10a.5.5 0 0 1 .5.5c0 .538-.012 1.05-.034 1.536a3 3 0 1 1-1.133 5.89c-.79 1.865-1.878 2.777-2.833 3.011v2.173l1.425.356c.194.048.377.135.537.255L13.3 15.1a.5.5 0 0 1-.3.9H3a.5.5 0 0 1-.3-.9l1.838-1.379c.16-.12.343-.207.537-.255L6.5 13.11v-2.173c-.955-.234-2.043-1.146-2.833-3.012a3 3 0 1 1-1.132-5.89A33.076 33.076 0 0 1 2.5.5zm.099 2.54a2 2 0 0 0 .72 3.935c-.333-1.05-.588-2.346-.72-3.935zm10.083 3.935a2 2 0 0 0 .72-3.935c-.133 1.59-.388 2.885-.72 3.935z"/>
          </svg>`,
		Data: nil,
	}
	if cursoIdStr != "" {
		cursoID, err := strconv.Atoi(cursoIdStr)
		if err != nil {
			http.Error(w, "El numero de codigo curso no es válido", http.StatusBadRequest)
			return
		}
		estudiantes := ObtenerEstudiantes()
		notas := repositories.ObtenerNotasPorCurso(estudiantes, cursoID)
		if len(notas) == 0 {
			data.H1 = fmt.Sprintf("No hay registro de notas en el curso %s", opciones[cursoID])
			data.Data = nil
		} else {
			data.H1 = fmt.Sprintf("Estadísticas de notas para el curso %s", opciones[cursoID])
			estadisticas := map[string]float64{
				"Promedio":            repositories.CalcularPromedio(notas),
				"Rango":               repositories.CalcularRango(notas),
				"Varianza":            repositories.CalcularVarianza(notas),
				"Desviacion_estandar": repositories.CalcularDesviacionEstandar(notas),
			}
			data.Data = estadisticas
		}
	}
	CargarTemplate(w, filepath, data)
}

func MejorPromedio(w http.ResponseWriter, r *http.Request) {
	estudiantes := ObtenerEstudiantes()
	data := PageData{
		Title: "TallerGO-Web | Mejor Promedio",
		H1:    "Estudiante con mejor promedio de notas",
		Icon: `<svg xmlns="http://www.w3.org/2000/svg" class="bi mt-4 mb-3" style="color: var(--bs-teal);" width="100"
        height="100" fill="currentColor" class="bi bi-trophy-fill" viewBox="0 0 16 16">
            <path d="M2.5.5A.5.5 0 0 1 3 0h10a.5.5 0 0 1 .5.5c0 .538-.012 1.05-.034 1.536a3 3 0 1 1-1.133 5.89c-.79 1.865-1.878 2.777-2.833 3.011v2.173l1.425.356c.194.048.377.135.537.255L13.3 15.1a.5.5 0 0 1-.3.9H3a.5.5 0 0 1-.3-.9l1.838-1.379c.16-.12.343-.207.537-.255L6.5 13.11v-2.173c-.955-.234-2.043-1.146-2.833-3.012a3 3 0 1 1-1.132-5.89A33.076 33.076 0 0 1 2.5.5zm.099 2.54a2 2 0 0 0 .72 3.935c-.333-1.05-.588-2.346-.72-3.935zm10.083 3.935a2 2 0 0 0 .72-3.935c-.133 1.59-.388 2.885-.72 3.935z"/>
          </svg>`,
		Data: repositories.MejoresEstudiantes(estudiantes),
	}
	filepath := "views/estudiante.html"
	CargarTemplate(w, filepath, data)
}

func PeorPromedio(w http.ResponseWriter, r *http.Request) {
	estudiantes := ObtenerEstudiantes()
	data := PageData{
		Title: "TallerGO-Web | Peor Promedio",
		H1:    "Estudiante con peor promedio de notas",
		Icon: `<svg xmlns="http://www.w3.org/2000/svg" class="bi mt-4 mb-3" style="color: var(--bs-teal);" width="100"
		height="100" fill="currentColor" fill="currentColor" class="bi bi-emoji-frown-fill" viewBox="0 0 16 16">
		<path
			d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zM7 6.5C7 7.328 6.552 8 6 8s-1-.672-1-1.5S5.448 5 6 5s1 .672 1 1.5zm-2.715 5.933a.5.5 0 0 1-.183-.683A4.498 4.498 0 0 1 8 9.5a4.5 4.5 0 0 1 3.898 2.25.5.5 0 0 1-.866.5A3.498 3.498 0 0 0 8 10.5a3.498 3.498 0 0 0-3.032 1.75.5.5 0 0 1-.683.183zM10 8c-.552 0-1-.672-1-1.5S9.448 5 10 5s1 .672 1 1.5S10.552 8 10 8z" />
	</svg>`,
		Data: repositories.PeoresEstudiantes(estudiantes),
	}
	filepath := "views/estudiante.html"
	CargarTemplate(w, filepath, data)
}

func MasculinoMasLongevo(w http.ResponseWriter, r *http.Request) {
	estudiantes := ObtenerEstudiantes()
	estudiantes = repositories.EstudiantesMasculinos(estudiantes)
	data := PageData{
		Title: "TallerGO-Web | Masculino Más Longevo",
		H1:    "Estudiante masculino mas longevo",
		Icon: `<svg xmlns="http://www.w3.org/2000/svg" class="bi mt-4 mb-3" style="color: var(--bs-teal);" width="100"
		height="100" fill="currentColor" class="bi bi-gender-male" viewBox="0 0 16 16">
		  <path fill-rule="evenodd" d="M9.5 2a.5.5 0 0 1 0-1h5a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-1 0V2.707L9.871 6.836a5 5 0 1 1-.707-.707L13.293 2H9.5zM6 6a4 4 0 1 0 0 8 4 4 0 0 0 0-8z"/>
		</svg>`,
		Data: repositories.EstudiantesMayorEdad(estudiantes),
	}
	filepath := "views/estudiante.html"
	CargarTemplate(w, filepath, data)
}

func FemeninaMasLongeva(w http.ResponseWriter, r *http.Request) {
	estudiantes := ObtenerEstudiantes()
	estudiantes = repositories.EstudiantesFemeninos(estudiantes)
	data := PageData{
		Title: "TallerGO-Web | Femenina Más Longeva",
		H1:    "Estudiante femenina mas longeva",
		Icon: `<svg xmlns="http://www.w3.org/2000/svg" class="bi mt-4 mb-3" style="color: var(--bs-teal);" width="100"
		height="100" fill="currentColor" class="bi bi-gender-female" viewBox="0 0 16 16">
		  <path fill-rule="evenodd" d="M8 1a4 4 0 1 0 0 8 4 4 0 0 0 0-8zM3 5a5 5 0 1 1 5.5 4.975V12h2a.5.5 0 0 1 0 1h-2v2.5a.5.5 0 0 1-1 0V13h-2a.5.5 0 0 1 0-1h2V9.975A5 5 0 0 1 3 5z"/>
		</svg>`,
		Data: repositories.EstudiantesMayorEdad(estudiantes),
	}
	filepath := "views/estudiante.html"
	CargarTemplate(w, filepath, data)
}

func ReporteMatriculas(w http.ResponseWriter, r *http.Request) {
	yearStr := r.FormValue("year")
	filepath := "views/reporteMatriculas.html"

	data := PageData{
		Title: "TallerGO-Web | Reporte Matriculas",
		H1:    "Reporte de matriculas por año",
		Icon: `<svg xmlns="http://www.w3.org/2000/svg" class="bi mt-4 mb-3" style="color: var(--bs-teal);" width="100"
			height="100" fill="currentColor" class="bi bi-table" viewBox="0 0 16 16">
			  <path d="M0 2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2zm15 2h-4v3h4V4zm0 4h-4v3h4V8zm0 4h-4v3h3a1 1 0 0 0 1-1v-2zm-5 3v-3H6v3h4zm-5 0v-3H1v2a1 1 0 0 0 1 1h3zm-4-4h4V8H1v3zm0-4h4V4H1v3zm5-3v3h4V4H6zm4 4H6v3h4V8z"/>
			</svg>`,
		Data: nil,
	}

	if yearStr != "" {
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			http.Error(w, "El año no es válido", http.StatusBadRequest)
			return
		}

		estudiantes := ObtenerEstudiantes()
		reporte := repositories.ObtenerEstudiantesMatriculadosEnAnio(estudiantes, year)

		data.H1 = fmt.Sprintf("Reporte de matriculas para el año %d", year)
		data.Data = reporte
		if len(reporte) == 0 {
			data.H1 = fmt.Sprintf("No hay estudiantes matriculados para el año %d ", year)
		}
	}
	CargarTemplate(w, filepath, data)
}

func PromedioRangoEdad(w http.ResponseWriter, r *http.Request) {
	edad1Str := r.FormValue("edad1")
	edad2Str := r.FormValue("edad2")
	filepath := "views/promedioRangoEdad.html"

	data := PageData{
		Title: "TallerGO-Web | Promedio Rango Edad",
		H1:    "Promedio de notas de los estudiantes dentro de un rango de edad",
		Icon: `<svg xmlns="http://www.w3.org/2000/svg" class="bi mt-4 mb-3" style="color: var(--bs-teal);" width="100"
			height="100" fill="currentColor" class="bi bi-calendar3-range-fill" viewBox="0 0 16 16">
			  <path fill-rule="evenodd" d="M2 0a2 2 0 0 0-2 2h16a2 2 0 0 0-2-2H2zM0 8V3h16v2h-6a1 1 0 1 0 0 2h6v7a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-4h6a1 1 0 1 0 0-2H0z"/>
			</svg>`,
		Data: nil,
	}

	if edad1Str != "" && edad2Str != "" {
		edad1, err := strconv.Atoi(edad1Str)
		if err != nil {
			http.Error(w, "Edad1 no es un número válido", http.StatusBadRequest)
			return
		}

		edad2, err := strconv.Atoi(edad2Str)
		if err != nil {
			http.Error(w, "Edad2 no es un número válido", http.StatusBadRequest)
			return
		}

		if edad1 > edad2 {
			aux := edad1
			edad1 = edad2
			edad2 = aux
		}

		estudiantes := ObtenerEstudiantes()
		notas := repositories.ObtenerNotasPorRangoEdad(estudiantes, edad1, edad2)
		data.Data = repositories.CalcularPromedio(notas)
		data.H1 = fmt.Sprintf("Promedio de notas de los estudiantes entre %d y %d años:", edad1, edad2)
		if len(notas) == 0 {
			data.H1 = fmt.Sprintf("No hay estudiantes entre %d y %d años", edad1, edad2)
			data.Data = nil
		}
	}
	CargarTemplate(w, filepath, data)
}

func CargarTemplate(w http.ResponseWriter, filePath string, data interface{}) {
	templates := template.Must(template.ParseFiles("templates/head.html", "templates/header.html", "templates/footer.html", filePath))

	err := templates.ExecuteTemplate(w, filepath.Base(filePath), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ObtenerEstudiantes() []models.Estudiante {
	filename := "data/generated.json"
	data, err := json.LeerJSON(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	listestudiantes, err := json.ParsearEstudiantes(data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	return listestudiantes
}
