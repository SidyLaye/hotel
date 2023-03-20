package hotel

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User struct to store user data
type User struct {
	Id       uuid.UUID `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name" gorm:"unique"`
	Password string    `json:"password"`
}

// errors
var (
	ErrRecordInvalid_user = errors.New("User is invalid")
)

var conf_user, _ = Conf()

// All retrieves all Users from the database
func All_users() ([]User, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_user.DBUser, conf_user.DBPass, conf_user.DBHost, conf_user.DBPort, conf_user.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var Users []User
	err = db.Find(&Users).Error
	if err != nil {
		return nil, err
	}
	return Users, nil
}

// One returns a single User record from the database
func One_user(id uuid.UUID) (*User, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_user.DBUser, conf_user.DBPass, conf_user.DBHost, conf_user.DBPort, conf_user.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	u := new(User)
	err = db.First(u, id).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Delete removes a given record from the database
func Delete_user(id uuid.UUID) error {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_user.DBUser, conf_user.DBPass, conf_user.DBHost, conf_user.DBPort, conf_user.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	u := new(User)
	err = db.First(u, id).Error
	if err != nil {
		return err
	}
	return db.Delete(u).Error
}

// Save updates or creates a given record in the database
func (u *User) Save_user() error {
	if err := u.validate_user(); err != nil {
		return err
	}
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_user.DBUser, conf_user.DBPass, conf_user.DBHost, conf_user.DBPort, conf_user.DBName)), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Save(u).Error
}

// validate make sure that the record contains valid data
func (u *User) validate_user() error {
	if u.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}

// Secret key used to sign JWT tokens
var secretKey = []byte("mySecretKey")

// Claims struct used to create a JWT token
type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

// Function that generates a JWT token for a given user ID
func GenerateToken(userId int64) (string, error) {
	// Create the payload (content) of the token
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // expiration in 24 hours
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Encode the token as a string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Function that verifies a JWT token and returns the user ID if it is valid
func VerifyToken(tokenString string) (int64, error) {
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signature method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return 0, err
	}

	// Check the validity of the token and get the payload
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := int64(claims["user_id"].(float64))
		return userId, nil
	} else {
		return 0, fmt.Errorf("Invalid token")
	}
}

// One returns a single User record from the database
func One_user_auth(name string) (*User, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf_user.DBUser, conf_user.DBPass, conf_user.DBHost, conf_user.DBPort, conf_user.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	u := new(User)
	err = db.Where("name = ?", name).First(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}
