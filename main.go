package main

import (
	"net/http"

	"github.com/edgarlozadagonzalez/TallerGo-Web/controllers"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("https://tallergo-web-production.up.railway.app/", controllers.Index)
	http.HandleFunc("https://tallergo-web-production.up.railway.app/curso/topMejoresEstudiantes", controllers.TopMejoresEstudiantes)
	http.HandleFunc("https://tallergo-web-production.up.railway.app/curso/topPeoresEstudiantes", controllers.TopPeoresEstudiantes)
	http.HandleFunc("https://tallergo-web-production.up.railway.app/curso/calcularEstadisticas", controllers.CalcularEstadisticas)

	http.HandleFunc("https://tallergo-web-production.up.railway.app/estudiante/mejorPromedio", controllers.MejorPromedio)
	http.HandleFunc("https://tallergo-web-production.up.railway.app/estudiante/peorPromedio", controllers.PeorPromedio)
	http.HandleFunc("https://tallergo-web-production.up.railway.app/estudiante/masculinoMasLongevo", controllers.MasculinoMasLongevo)
	http.HandleFunc("https://tallergo-web-production.up.railway.app/estudiante/femeninaMasLongeva", controllers.FemeninaMasLongeva)

	http.HandleFunc("https://tallergo-web-production.up.railway.app/reporteMatriculas", controllers.ReporteMatriculas)
	http.HandleFunc("https://tallergo-web-production.up.railway.app/promedioRangoEdad", controllers.PromedioRangoEdad)
	http.ListenAndServe(":8080", nil)
}
