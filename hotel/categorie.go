package hotel

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Categorie holds data for a single Categorie
type Categorie struct {
	Id_categorie  uuid.UUID `json:"id" gorm:"primaryKey"`
	Nom           string    `json:"nom"`
	Tarif_normal  uint      `json:"tarif_normal"`
	Tarif_special uint      `json:"tarif_special"`
}

// errors
var (
	ErrRecordInvalid_categorie = errors.New("record is invalid")
)

var conf_categorie, _ = Conf()

// All retrieves all categories from the database
func All_categories() ([]Categorie, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_categorie.DBUser, conf_categorie.DBPass, conf_categorie.DBHost, conf_categorie.DBPort, conf_categorie.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var categories []Categorie
	err = db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// One returns a single Categorie record from the database
func One_categorie(id uuid.UUID) (*Categorie, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_categorie.DBUser, conf_categorie.DBPass, conf_categorie.DBHost, conf_categorie.DBPort, conf_categorie.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	c := new(Categorie)
	err = db.First(c, id).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Delete removes a given record from the database
func Delete_categorie(id uuid.UUID) error {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_categorie.DBUser, conf_categorie.DBPass, conf_categorie.DBHost, conf_categorie.DBPort, conf_categorie.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	c := new(Categorie)
	err = db.First(c, id).Error
	if err != nil {
		return err
	}
	return db.Delete(c).Error
}

// Save updates or creates a given record in the database
func (c *Categorie) Save_categorie() error {
	if err := c.validate_categorie(); err != nil {
		return err
	}
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_categorie.DBUser, conf_categorie.DBPass, conf_categorie.DBHost, conf_categorie.DBPort, conf_categorie.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Save(c).Error
}

// validate make sure that the record contains valid data
func (c *Categorie) validate_categorie() error {
	if c.Nom == "" {
		return ErrRecordInvalid
	}
	return nil
}
