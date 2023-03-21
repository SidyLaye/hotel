package hotel

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Reservation holds data for a single Reservation
type Reservation struct {
	Id_reservation   uuid.UUID `json:"id_reservation" gorm:"primaryKey"`
	Nom              string    `json:"nom"`
	Prenom           string    `json:"prenom"`
	Date_reservation string    `json:"date_reservation,omitempty" gorm:"type:date"`
	Date_entree      string    `json:"date_entree,omitempty" gorm:"type:date"`
	Date_sortie      string    `json:"date_sortie,omitempty" gorm:"type:date"`
	Bar              bool      `json:"bar,omitempty"`
	Petit_dej        bool      `json:"petit_dej,omitempty"`
	Phone            bool      `json:"phone,omitempty"`
	Nuite            int       `json:"nuite,omitempty"`
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
	err = db.Select("id_reservation, nom, prenom").Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

// One returns a single Reservation record from the database
func One_reservation(id_reservation uuid.UUID) (*Reservation, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_reservation.DBUser, conf_reservation.DBPass, conf_reservation.DBHost, conf_reservation.DBPort, conf_reservation.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	r := new(Reservation)
	err = db.First(r, id_reservation).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Delete removes a given record from the database
func Delete_reservation(id_reservation uuid.UUID) error {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_reservation.DBUser, conf_reservation.DBPass, conf_reservation.DBHost, conf_reservation.DBPort, conf_reservation.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	r := new(Reservation)
	err = db.First(r, id_reservation).Error
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

func ParseMysqlDate(mysqlDate string) (time.Time, error) {
	return time.Parse("2006-01-02", mysqlDate)
}

// validate make sure that the record contains valid data
func (r *Reservation) validate_reservation() error {
	entree, _ := ParseMysqlDate(r.Date_entree)
	sortie, _ := ParseMysqlDate(r.Date_sortie)
	reserv, _ := ParseMysqlDate(r.Date_reservation)
	if r.Nom == "" || r.Prenom == "" || entree.After(sortie) || entree.Before(reserv) {
		return ErrRecordInvalid
	}
	return nil
}

func Nightsbeetween(sortie, entree time.Time) int {
	duration := sortie.Sub(entree)
	nights := int(duration.Hours() / 24)
	return nights
}
