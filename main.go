package main

import (
	"net/http"

	"github.com/edgarlozadagonzalez/TallerGo-Web/controllers"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/curso/topMejoresEstudiantes", controllers.TopMejoresEstudiantes)
	http.HandleFunc("/curso/topPeoresEstudiantes", controllers.TopPeoresEstudiantes)
	http.HandleFunc("/curso/calcularEstadisticas", controllers.CalcularEstadisticas)

	http.HandleFunc("/estudiante/mejorPromedio", controllers.MejorPromedio)
	http.HandleFunc("/estudiante/peorPromedio", controllers.PeorPromedio)
	http.HandleFunc("/estudiante/masculinoMasLongevo", controllers.MasculinoMasLongevo)
	http.HandleFunc("/estudiante/femeninaMasLongeva", controllers.FemeninaMasLongeva)

	http.HandleFunc("/reporteMatriculas", controllers.ReporteMatriculas)
	http.HandleFunc("/promedioRangoEdad", controllers.PromedioRangoEdad)
	http.ListenAndServe(":8080", nil)
}
