// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/VerstraeteBert/WeatherApp/models"

// ReadingRepo is an autogenerated mock type for the ReadingRepo type
type ReadingRepo struct {
	mock.Mock
}

// AddReading provides a mock function with given fields: reading
func (_m *ReadingRepo) AddReading(reading *models.Reading) (int64, error) {
	ret := _m.Called(reading)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*models.Reading) int64); ok {
		r0 = rf(reading)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Reading) error); ok {
		r1 = rf(reading)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReadings provides a mock function with given fields:
func (_m *ReadingRepo) ListReadings() ([]*models.Reading, error) {
	ret := _m.Called()

	var r0 []*models.Reading
	if rf, ok := ret.Get(0).(func() []*models.Reading); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Reading)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}