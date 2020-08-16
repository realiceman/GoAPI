package user

import (
	"errors"
	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

// User holds data for a single user
type User struct {
	ID bson.ObjectId `json:"id" storm:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

const (
	dbPath = "users.db"
)

// erors
var (
	ErrRecordInvalid = errors.New("Invalid record please correct before proceed.")
)

// retrieve all users from the db
func All([]User, error) ([]User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	users := []User{}
	err = db.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Find a single user record from the db
func One(id bson.ObjectId) (*User, error)  {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func delete(id bson.ObjectId) error  {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return err
	}
	return db.DeleteStruct(u)
}

// Save updates or creates a given record in the db
func (u *User) Save() error  {
	if err := u.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(u)
}

// Make sure the record contains valid data
func (u *User) validate() error {
	if u.Name == ""{
		return ErrRecordInvalid
	}
	return nil
}