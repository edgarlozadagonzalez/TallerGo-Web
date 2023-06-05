package main

import (
	"log"
	"net/http"
	"os"

	"github.com/edgarlozadagonzalez/TallerGo-Web/controllers"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/curso/topMejoresEstudiantes", controllers.TopMejoresEstudiantes)
	http.HandleFunc("/curso/topPeoresEstudiantes", controllers.TopPeoresEstudiantes)
	http.HandleFunc("/curso/calcularEstadisticas", controllers.CalcularEstadisticas)
	http.HandleFunc("/curso/agregarCurso", controllers.AgregarCurso)

	http.HandleFunc("/estudiante/mejorPromedio", controllers.MejorPromedio)
	http.HandleFunc("/estudiante/peorPromedio", controllers.PeorPromedio)
	http.HandleFunc("/estudiante/masculinoMasLongevo", controllers.MasculinoMasLongevo)
	http.HandleFunc("/estudiante/femeninaMasLongeva", controllers.FemeninaMasLongeva)
	http.HandleFunc("/estudiante/agregarEstudiante", controllers.AgregarEstudiante)

	http.HandleFunc("/reporteMatriculas", controllers.ReporteMatriculas)
	http.HandleFunc("/promedioRangoEdad", controllers.PromedioRangoEdad)
	log.Println("listening on", port)
	http.ListenAndServe(":"+port, nil)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
