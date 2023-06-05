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

func AgregarEstudianteJSON(filename string, estudiantes []models.Estudiante) error {
	data, err := json.MarshalIndent(estudiantes, "", "  ")
	if err != nil {
		return fmt.Errorf("error al codificar los datos: %v", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error al escribir en el archivo: %v", err)
	}
	return nil
}
