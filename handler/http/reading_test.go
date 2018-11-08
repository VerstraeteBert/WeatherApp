package http

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VerstraeteBert/WeatherApp/models"
	"github.com/VerstraeteBert/WeatherApp/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestReadingHandler_ListReadings(t *testing.T) {
	tcs := []struct {
		name             string
		returnedFromMock []*models.Reading
		expected         []models.Reading
		status           int
		err              bool
	}{
		{"erroneous request", nil, nil, http.StatusInternalServerError, true},
		{
			"successful request",
			[]*models.Reading{
				{
					ID:             1,
					Timestamp:      "1997-06-02 06:10:12",
					DegreesCelsius: 24.00,
				},
			},
			[]models.Reading{
				{
					ID:             1,
					Timestamp:      "1997-06-02 06:10:12",
					DegreesCelsius: 24.00,
				},
			},
			http.StatusOK,
			false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mocks.ReadingRepo{}
			req, err := http.NewRequest("GET", "localhost:1337/readings", nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			rec := httptest.NewRecorder()

			rh := ReadingHandler{mockRepo}
			if tc.err {
				mockRepo.On("ListReadings").Return(
					nil,
					errors.New("error"),
				)
			} else {
				mockRepo.On("ListReadings").Return(
					tc.returnedFromMock,
					nil,
				)
			}

			rh.ListReadings(rec, req)

			res := rec.Result()
			defer res.Body.Close()
			if res.StatusCode != tc.status {
				t.Errorf("Expected status %v, got %v", tc.status, res.Status)
			}

			if res.StatusCode == http.StatusOK {
				b, err := ioutil.ReadAll(res.Body)
				resModels := make([]models.Reading, 0)
				err = json.Unmarshal(b, &resModels)
				if err != nil {
					t.Errorf("Failed to unmarshal response json, %v", err)
				}

				assert.Equal(
					t,
					tc.expected,
					resModels,
				)
			}
		})
	}
}
