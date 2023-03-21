package hotel

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Nomss string

const (
	Petit_dej Nomss = "petit_dej"
	Phone     Nomss = "phone"
	Bar       Nomss = "bar"
)

// Service_supp holds data for a single Service_supp
type Service_supp struct {
	Nom            Nomss     `json:"nom" gorm:"primaryKey;type:enum('petit_dej','phone','bar')"`
	Id_reservation uuid.UUID `json:"-"`
	Tarif          uint      `json:"tarif"`
}

// errors
var (
	ErrRecordInvalid = errors.New("record is invalid")
)

var conf_service_supp, _ = Conf()

// All retrieves all services_supp from the database
func All_service_supp() ([]Service_supp, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_service_supp.DBUser, conf_service_supp.DBPass, conf_service_supp.DBHost, conf_service_supp.DBPort, conf_service_supp.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var services_supp []Service_supp
	err = db.Find(&services_supp).Error
	if err != nil {
		return nil, err
	}
	return services_supp, nil
}

// One returns a single Service_supp record from the database
func One_service_supp(id_service_supp Nomss) (*Service_supp, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_service_supp.DBUser, conf_service_supp.DBPass, conf_service_supp.DBHost, conf_service_supp.DBPort, conf_service_supp.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	s := new(Service_supp)
	err = db.First(s, id_service_supp).Error
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Delete removes a given record from the database
func Delete_service_supp(id_service_supp Nomss) error {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_service_supp.DBUser, conf_service_supp.DBPass, conf_service_supp.DBHost, conf_service_supp.DBPort, conf_service_supp.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	s := new(Service_supp)
	err = db.First(s, id_service_supp).Error
	if err != nil {
		return err
	}
	return db.Delete(s).Error
}

// Save updates or creates a given record in the database
func (s *Service_supp) Save_service_supp() error {
	if err := s.validate_service_supp(); err != nil {
		return err
	}
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_service_supp.DBUser, conf_service_supp.DBPass, conf_service_supp.DBHost, conf_service_supp.DBPort, conf_service_supp.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Save(s).Error
}

// validate make sure that the record contains valid data
func (s *Service_supp) validate_service_supp() error {
	if s.Nom == "" {
		return ErrRecordInvalid
	}
	return nil
}
