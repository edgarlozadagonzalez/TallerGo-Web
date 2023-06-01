package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	models "github.com/edgarlozadagonzalez/TallerGo-Web/models"
)

func LeerJSON(filename string) ([]models.Estudiante, error) {
	// Leer el archivo JSON
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo: %v", err)
	}

	// Crear una variable para almacenar los estudiantes
	var estudiantes []models.Estudiante

	// Parsear el JSON en la estructura
	err = json.Unmarshal(data, &estudiantes)
	if err != nil {
		return nil, fmt.Errorf("error al parsear el JSON: %v", err)
	}

	return estudiantes, nil
}
