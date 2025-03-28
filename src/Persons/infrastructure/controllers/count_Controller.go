package controllers

import "fmt"

func (mysql *MySQL) CountGender() ([]int, error) {

	var countH, countF int
	var counts []int

	query1 := "SELECT COUNT(id_person) FROM person WHERE gender = 'Masculino'"
	query2 := "SELECT COUNT(id_person) FROM person WHERE gender = 'Femenino'"

	// Query para la parte del conteo de la cantidad de hombres
	rows1 := mysql.conn.FetchRows(query1)

	defer rows1.Close()

	for rows1.Next() {

		if err := rows1.Scan(&countH); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err.Error())
		}

		counts = append(counts, countH)

	}

	// Query para la parte del conteo de la cantidad de mujeres
	rows2 := mysql.conn.FetchRows(query2)

	defer rows2.Close()

	for rows2.Next() {

		if err := rows2.Scan(&countF); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err.Error())
		}

		counts = append(counts, countF)

	}

	return counts, nil

}
