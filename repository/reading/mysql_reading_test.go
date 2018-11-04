package reading

import (
	"database/sql"
	"github.com/VerstraeteBert/WeatherApp/driver"
	"github.com/VerstraeteBert/WeatherApp/models"
	"github.com/VerstraeteBert/WeatherApp/repository"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type ReadingRepoSuite struct {
	suite.Suite
	readingRepo repository.ReadingRepo
	db          *sql.DB
}

func (s *ReadingRepoSuite) SetupSuite() {
	godotenv.Load("../../.env") // TODO Figure out how to set path to root

	connection, err := driver.ConnectSQL(
		os.Getenv("DATABASE_URL"),
	)

	if err != nil {
		s.T().Fatal(err)
	}

	s.db = connection.SQL
	s.readingRepo = NewSQLReadingRepo(connection.SQL)
}

func (s *ReadingRepoSuite) TearDownSuite() {
	// Close the connection after all tests in the suite finish
	s.db.Close()
}

func (s *ReadingRepoSuite) SetupTest() {
	/*
		We delete all entries from the table before each test runs, to ensure a
		consistent state before our tests run. In more complex applications, this
		is sometimes achieved in the form of migrations
	*/
	_, err := s.db.Query("DELETE FROM readings")
	if err != nil {
		s.T().Fatal(err)
	}
}

// This is the actual "test" as seen by Go, which runs the tests defined below
func TestStoreSuite(t *testing.T) {
	s := new(ReadingRepoSuite)
	suite.Run(t, s)
}

func (s *ReadingRepoSuite) TestAddReading() {
	s.readingRepo.AddReading(&models.Reading{
		Timestamp:      "2018-11-03 19:51:09",
		DegreesCelcius: 32.45,
	})

	res, err := s.db.Query(`SELECT COUNT(*) FROM readings WHERE timestamp="2018-11-03 19:51:09" AND degreescelcius=32.45`)
	if err != nil {
		s.T().Fatal(err)
	}

	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}

	assert.Equal(s.T(), 1, count)
}

func (s *ReadingRepoSuite) TestListReading() {
	_, err := s.db.Query(`INSERT INTO readings (timestamp, degreescelcius) VALUES ("2018-11-03 19:51:09", 32.45)`)
	if err != nil {
		s.T().Fatal(err)
	}

	res, err := s.readingRepo.ListReadings()
	if err != nil {
		s.T().Fatal(err)
	}

	assert.Equal(s.T(), 1, len(res))
	assert.Equal(s.T(), "2018-11-03 19:51:09", res[0].Timestamp)
	assert.Equal(s.T(), float32(32.45), res[0].DegreesCelcius)

	_, err = s.db.Query(`INSERT INTO readings (timestamp, degreescelcius) VALUES ("2019-11-03 19:51:09", 33.45)`)
	if err != nil {
		s.T().Fatal(err)
	}

	res, err = s.readingRepo.ListReadings()
	if err != nil {
		s.T().Fatal(err)
	}

	assert.Equal(s.T(), 2, len(res))
}
