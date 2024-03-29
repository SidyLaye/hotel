package hotel

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Facture holds data for a single Facture
type Facture struct {
	Id_facture      uuid.UUID    `json:"id_facture" gorm:"primaryKey"`
	Id_reservation  uuid.UUID    `json:"id_reservation" gorm:"index"`
	Reservation     Reservation  `json:"reservation" gorm:"foreignKey:Id_reservation"`
	Id_service_supp Nomss        `json:"id_service_supp" gorm:"index"`
	Service_supp    Service_supp `json:"service_supp" gorm:"foreignKey:nom"`
	Tarif_chambre   uint         `json:"tarif_chambre"`
	Tarif_bar       uint         `json:"tarif_bar"`
	Tarif_petit_dej uint         `json:"tarif_petit_dej"`
	Tarif_phone     uint         `json:"tarif_phone"`
	Tarif_special   uint         `json:"tarif_special"`
	Total           uint         `json:"total"`
}

// errors
var (
	ErrRecordInvalid_facture = errors.New("record is invalid")
)

var conf_facture, _ = Conf()

// All retrieves all factures from the database
func All_factures() ([]Facture, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_facture.DBUser, conf_facture.DBPass, conf_facture.DBHost, conf_facture.DBPort, conf_facture.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var factures []Facture
	err = db.Preload("Reservation").Find(&factures).Error
	if err != nil {
		return nil, err
	}
	return factures, nil
}

// One returns a single Facture record from the database
func One_facture(id uuid.UUID) (*Facture, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_facture.DBUser, conf_facture.DBPass, conf_facture.DBHost, conf_facture.DBPort, conf_facture.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	c := new(Facture)
	err = db.Preload("reservation").Where("id_reservation = ?", id).First(&c).Error
	if err != nil {
		return nil, err
	}
	err = db.Preload("Service_supp").Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Delete removes a given record from the database
func Delete_facture(id uuid.UUID) error {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_facture.DBUser, conf_facture.DBPass, conf_facture.DBHost, conf_facture.DBPort, conf_facture.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	c := new(Facture)
	err = db.Preload("reservation").Where("id_reservation = ?", id).First(&c).Error
	if err != nil {
		return err
	}
	return db.Delete(c).Error
}

// Save updates or creates a given record in the database
func (c *Facture) Save_facture() error {
	if err := c.validate_facture(); err != nil {
		return err
	}
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_facture.DBUser, conf_facture.DBPass, conf_facture.DBHost, conf_facture.DBPort, conf_facture.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	var r Reservation
	var s Service_supp
	_ = db.First(r, c.Id_reservation)
	_ = db.First(s, c.Id_service_supp)
	_ = db.Preload("Reservation").Where("id_reservation = ? ", r.Id_reservation)
	_ = db.Preload("Service_supp").Where(" id_service_supp = ?", s.Nom)
	return db.Save(c).Error
}

// validate make sure that the record contains valid data
func (f *Facture) validate_facture() error {
	if f.Reservation.Nom == "" {
		return ErrRecordInvalid
	}
	return nil
}
