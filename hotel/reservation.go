package hotel

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Reservation holds data for a single Reservation
type Reservation struct {
	Id_reservation   uuid.UUID `json:"id" gorm:"primaryKey"`
	Nom              string    `json:"nom"`
	Prenom           string    `json:"prenom"`
	Date_reservation string    `json:"date_reservation" gorm:"type:date"`
	Date_entree      string    `json:"date_entree" gorm:"type:date"`
	Date_sortie      string    `json:"date_sortie" gorm:"type:date"`
	Bar              bool      `json:"bar"`
	Petit_dej        bool      `json:"petit_dej"`
	Phone            bool      `json:"phone"`
	Nuite            int       `json:"nuite"`
}

// errors
var (
	ErrRecordInvalid_reservation = errors.New("reservation is invalid")
)

var conf_reservation, _ = Conf()

// All retrieves all reservations from the database
func All_reservations() ([]Reservation, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_reservation.DBUser, conf_reservation.DBPass, conf_reservation.DBHost, conf_reservation.DBPort, conf_reservation.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var reservations []Reservation
	err = db.Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

// One returns a single Reservation record from the database
func One_reservation(id uuid.UUID) (*Reservation, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_reservation.DBUser, conf_reservation.DBPass, conf_reservation.DBHost, conf_reservation.DBPort, conf_reservation.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	r := new(Reservation)
	err = db.First(r, id).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Delete removes a given record from the database
func Delete_reservation(id uuid.UUID) error {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_reservation.DBUser, conf_reservation.DBPass, conf_reservation.DBHost, conf_reservation.DBPort, conf_reservation.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	r := new(Reservation)
	err = db.First(r, id).Error
	if err != nil {
		return err
	}
	return db.Delete(r).Error
}

// Save updates or creates a given record in the database
func (r *Reservation) Save_reservation() error {
	if err := r.validate_reservation(); err != nil {
		return err
	}
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_reservation.DBUser, conf_reservation.DBPass, conf_reservation.DBHost, conf_reservation.DBPort, conf_reservation.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Save(r).Error
}

// validate make sure that the record contains valid data
func (r *Reservation) validate_reservation() error {
	if r.Nom == "" {
		return ErrRecordInvalid
	}
	return nil
}
