package hotel

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Info_hotel holds data for a single Info_hotel
type Info_hotel struct {
	Id_info_hotel   uuid.UUID `json:"id_info_hotel" gorm:"primaryKey"`
	Nom_hotel       string    `json:"nom"`
	Date_debut      time.Time `json:"date_debut" gorm:"type:date"`
	Nombre_chambres int       `json:"nombre_chambres"`
	Nombre_niveaux  int       `json:"nombre_niveaux"`
	Tel_hotel       string    `json:"tel_hotel"`
}

// errors
var (
	ErrRecordInvalid_info_hotel = errors.New("record is invalid")
)

var conf_info_hotel, _ = Conf()

// All retrieves all infos_hotel from the database
func All_info_hotel() ([]Info_hotel, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_info_hotel.DBUser, conf_info_hotel.DBPass, conf_info_hotel.DBHost, conf_info_hotel.DBPort, conf_info_hotel.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var infos_hotel []Info_hotel
	err = db.Find(&infos_hotel).Error
	if err != nil {
		return nil, err
	}
	return infos_hotel, nil
}

// One returns a single Info_hotel record from the database
func One_info_hotel(id uuid.UUID) (*Info_hotel, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_info_hotel.DBUser, conf_info_hotel.DBPass, conf_info_hotel.DBHost, conf_info_hotel.DBPort, conf_info_hotel.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	i := new(Info_hotel)
	err = db.First(i, id).Error
	if err != nil {
		return nil, err
	}
	return i, nil
}

// Delete removes a given record from the database
func Delete_info_hotel(id uuid.UUID) error {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_info_hotel.DBUser, conf_info_hotel.DBPass, conf_info_hotel.DBHost, conf_info_hotel.DBPort, conf_info_hotel.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	i := new(Info_hotel)
	err = db.First(i, id).Error
	if err != nil {
		return err
	}
	return db.Delete(i).Error
}

// Save updates or creates a given record in the database
func (i *Info_hotel) Save_info_hotel() error {
	if err := i.validate_info_hotel(); err != nil {
		return err
	}
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_info_hotel.DBUser, conf_info_hotel.DBPass, conf_info_hotel.DBHost, conf_info_hotel.DBPort, conf_info_hotel.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Save(i).Error
}

// validate make sure that the record contains valid data
func (i *Info_hotel) validate_info_hotel() error {
	if i.Nom_hotel == "" {
		return ErrRecordInvalid
	}
	return nil
}
