package hotel

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Etat string

const (
	Libre   Etat = "libre"
	Occupe  Etat = "occupe"
	Reserve Etat = "reserve"
)

type Classe string

const (
	Economie Classe = "Economie"
	Standing Classe = "Standing"
	Affaire  Classe = "Affaire"
)

// Chambre holds data for a single Chambre
type Chambre struct {
	Id_chambre     uuid.UUID `json:"id_chambre" gorm:"primaryKey"`
	Id_categorie   uuid.UUID `json:"id_categorie"`
	Id_reservation uuid.UUID `json:"id_reservation"`
	Num_chambre    string    `json:"num_chambre"`
	Etat           Etat      `json:"etat"`
	Classe         Classe    `json:"classe"`
}

// errors
var (
	ErrRecordInvalid_chambre = errors.New("record is invalid")
)

var conf_chambre, _ = Conf()

// All retrieves all chambres from the database
func All_chambres() ([]Chambre, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_chambre.DBUser, conf_chambre.DBPass, conf_chambre.DBHost, conf_chambre.DBPort, conf_chambre.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var chambres []Chambre
	err = db.Find(&chambres).Error
	if err != nil {
		return nil, err
	}
	return chambres, nil
}

// One returns a single Chambre record from the database
func One_chambre(id uuid.UUID) (*Chambre, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_chambre.DBUser, conf_chambre.DBPass, conf_chambre.DBHost, conf_chambre.DBPort, conf_chambre.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	c := new(Chambre)
	err = db.First(c, id).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Delete removes a given record from the database
func Delete_chambre(id uuid.UUID) error {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_chambre.DBUser, conf_chambre.DBPass, conf_chambre.DBHost, conf_chambre.DBPort, conf_chambre.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	c := new(Chambre)
	err = db.First(c, id).Error
	if err != nil {
		return err
	}
	return db.Delete(c).Error
}

// Save updates or creates a given record in the database
func (c *Chambre) Save_chambre() error {
	if err := c.validate_chambre(); err != nil {
		return err
	}
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_chambre.DBUser, conf_chambre.DBPass, conf_chambre.DBHost, conf_chambre.DBPort, conf_chambre.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Save(c).Error
}

// validate make sure that the record contains valid data
func (c *Chambre) validate_chambre() error {
	if c.Num_chambre == "" {
		return ErrRecordInvalid
	}
	return nil
}
