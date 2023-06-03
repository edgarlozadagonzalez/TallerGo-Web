package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	models "github.com/edgarlozadagonzalez/TallerGo-Web/models"
)

func LeerJSON(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error al leer el archivo: %v", err)
	}

	return data, nil
}

func ParsearEstudiantes(data []byte) ([]models.Estudiante, error) {
	var estudiantes []models.Estudiante
	err := json.Unmarshal(data, &estudiantes)
	if err != nil {
		return nil, fmt.Errorf("Error al parsear el JSON: %v", err)
	}

	return estudiantes, nil
}
