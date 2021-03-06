package reading

import (
	"database/sql"
	"github.com/VerstraeteBert/WeatherApp/models"
	"github.com/VerstraeteBert/WeatherApp/repository"
)

func NewSQLReadingRepo(Conn *sql.DB) repository.ReadingRepo {
	return &mysqlReadingRepo{
		Conn: Conn,
	}
}

type mysqlReadingRepo struct {
	Conn *sql.DB
}

func (m *mysqlReadingRepo) AddReading(reading *models.Reading) (int64, error) {
	pstmt, err := m.Conn.Prepare("INSERT INTO readings (timestamp, degreescelsius) VALUES (?, ?)")
	if err != nil {
		return -1, err
	}
	defer pstmt.Close()

	res, err := pstmt.Exec(reading.Timestamp, reading.DegreesCelsius)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlReadingRepo) ListReadings() ([]*models.Reading, error) {
	rows, err := m.Conn.Query("SELECT * FROM readings")
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
			&data.DegreesCelsius,
		)
		if err != nil {
			panic(err)
		}
		payload = append(payload, data)
	}

	return payload, nil
}
