package hotel

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Client holds data for a single client
type Client struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
	Tel    string `json:"tel"`
}

// errors
var (
	ErrRecordInvalid = errors.New("record is invalid")
)

// All retrieves all clients from the database
func All() ([]Client, error) {
	db, err := gorm.Open(mysql.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var clients []Client
	err = db.Find(&clients).Error
	if err != nil {
		return nil, err
	}
	return clients, nil
}

// One returns a single client record from the database
func One(id uint) (*Client, error) {
	db, err := gorm.Open(mysql.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	c := new(Client)
	err = db.First(c, id).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Delete removes a given record from the database
func Delete(id uint) error {
	db, err := gorm.Open(mysql.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}
	c := new(Client)
	err = db.First(c, id).Error
	if err != nil {
		return err
	}
	return db.Delete(c).Error
}

// Save updates or creates a given record in the database
func (c *Client) Save() error {
	if err := c.validate(); err != nil {
		return err
	}
	db, err := gorm.Open(mysql.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Save(c).Error
}

// validate make sure that the record contains valid data
func (c *Client) validate() error {
	if c.Nom == "" {
		return ErrRecordInvalid
	}
	return nil
}
