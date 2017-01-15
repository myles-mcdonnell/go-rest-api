package repository

import (
	"github.com/myles-mcdonnell/go-rest-api/models"
)

type Scanner interface {
	Scan(dest ...interface{}) error
}

func GetAllClients() ([]*models.Client, error) {

	rows, err := GetClientsStmt.Queryx()

	if err != nil {
		return nil, err
	}

	results := make([]*models.Client, 0)

	for rows.Next() {

		client := new(models.Client)
		err := rows.StructScan(client)

		if err != nil {
			return nil, err
		}

		results = append(results, client)
	}

	rows.Close()

	return results, nil
}
