package hotel

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Fourniture holds data for a single Fourniture
type Fourniture struct {
	Id_fourniture   uuid.UUID `json:"id_fourniture" gorm:"primaryKey"`
	Nom_fourniture  string    `json:"nom_fourniture"`
	Prix_fourniture uint      `json:"prix_fourniture"`
}

// errors
var (
	ErrRecordInvalid_fourniture = errors.New("record is invalid")
)

var conf_fourniture, _ = Conf()

// All retrieves all fournitures from the database
func All_fournitures() ([]Fourniture, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_fourniture.DBUser, conf_fourniture.DBPass, conf_fourniture.DBHost, conf_fourniture.DBPort, conf_fourniture.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var fournitures []Fourniture
	err = db.Find(&fournitures).Error
	if err != nil {
		return nil, err
	}
	return fournitures, nil
}

// One returns a single Fourniture record from the database
func One_fourniture(id_fourniture uuid.UUID) (*Fourniture, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_fourniture.DBUser, conf_fourniture.DBPass, conf_fourniture.DBHost, conf_fourniture.DBPort, conf_fourniture.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	f := new(Fourniture)
	err = db.Where("id_fourniture = ?", id_fourniture).First(f).Error
	if err != nil {
		return nil, err
	}
	return f, nil
}

// Delete removes a given record from the database
func Delete_fourniture(id_fourniture uuid.UUID) error {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_fourniture.DBUser, conf_fourniture.DBPass, conf_fourniture.DBHost, conf_fourniture.DBPort, conf_fourniture.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	f := new(Fourniture)
	err = db.First(f, id_fourniture).Error
	if err != nil {
		return err
	}
	return db.Delete(f).Error
}

// Save updates or creates a given record in the database
func (f *Fourniture) Save_fourniture() error {
	if err := f.validate_fourniture(); err != nil {
		return err
	}
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_fourniture.DBUser, conf_fourniture.DBPass, conf_fourniture.DBHost, conf_fourniture.DBPort, conf_fourniture.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Save(f).Error
}

// validate make sure that the record contains valid data
func (f *Fourniture) validate_fourniture() error {
	if f.Nom_fourniture == "" {
		return ErrRecordInvalid
	}
	return nil
}
