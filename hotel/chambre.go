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

// Chambre holds data for a single Chambre
type Chambre struct {
	Num_chambre    int       `json:"num_chambre" gorm:"primaryKey"`
	Nom_categorie  Nom       `json:"nom_categorie,omitempty" gorm:"type:enum('Economique','Standing','Affaire')"`
	Id_reservation uuid.UUID `json:"-"`
	Etat           Etat      `json:"etat,omitempty" gorm:"type:enum('libre','occupe','reserve')"`
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
	err = db.Select("num_chambre").Find(&chambres).Error
	if err != nil {
		return nil, err
	}
	return chambres, nil
}

// One returns a single Chambre record from the database
func One_chambre(num_chambre int) (*Chambre, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_chambre.DBUser, conf_chambre.DBPass, conf_chambre.DBHost, conf_chambre.DBPort, conf_chambre.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	c := new(Chambre)
	err = db.Select("num_chambre, nom_categorie, etat").First(c, num_chambre).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Delete removes a given record from the database
func Delete_chambre(num_chambre int) error {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_chambre.DBUser, conf_chambre.DBPass, conf_chambre.DBHost, conf_chambre.DBPort, conf_chambre.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	c := new(Chambre)
	err = db.First(c, num_chambre).Error
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
	//if c.Num_chambre == 0 {
	//	return ErrRecordInvalid
	//}
	return nil
}
