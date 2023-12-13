package repository

import "bookings/internals/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res *models.Reservation) error
}
