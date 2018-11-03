package reading

import (
	"WeatherApp/models"
	"WeatherApp/repository"
	"context"
	"database/sql"
)

func NewSQLReadingRepo(Conn *sql.DB) repository.ReadingRepo {
	return &mysqlReadingRepo {
		Conn: Conn,
	}
}

type mysqlReadingRepo struct {
	Conn *sql.DB
}

func (m *mysqlReadingRepo) ListReadings(ctx context.Context) ([]*models.Reading, error) {
	rows, err := m.Conn.QueryContext(ctx, "select * from readings")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	payload := make([]*models.Reading, 0)
	for rows.Next() {
		data := new(models.Reading)
		err := rows.Scan(
			&data.ID,
			&data.Timestamp,
			&data.DegreesCelcius,
		)
		if err != nil {
			panic(err)
		}
		payload = append(payload, data)
	}

	return payload, nil
}

