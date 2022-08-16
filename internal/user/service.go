package user

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"mendota-backend/config"
)

const (
	ScryptN      = 32768
	ScryptR      = 8
	ScryptP      = 1
	ScryptKeyLen = 32
)

type Service interface {
	Create(user *User) error

	FindById(id uint, result *User) error
	FindByEmail(hashedEmail string, result *User) error
	FindByUserName(name string, result *User) error
	FindProfile(id uint, result *Profile) error

	Update(u *User) error
	Remove(u *User) error
	LoginWithEmail(hashedEmail, password string, result *User) error
	LoginWithName(name, password string, result *User) error
}

type ServiceImpl struct {
	DB *config.DataBase
}

func NewUserService(DB *config.DataBase) Service {
	return &ServiceImpl{DB}
}

func (s *ServiceImpl) Create(user *User) error {
	return nil
}

func (s *ServiceImpl) LoginWithEmail(hashedEmail, password string, result *User) error {
	err := s.FindByEmail(hashedEmail, result)
	if err != nil {
		return err
	}

	return validatePassword(password, result.Password, result.Salt)
}

func (s *ServiceImpl) LoginWithName(name, password string, result *User) error {
	err := s.FindByUserName(name, result)
	if err != nil {
		return err
	}

	return validatePassword(password, result.Password, result.Salt)
}

func (s *ServiceImpl) FindById(id uint, result *User) error {
	return s.DB.Find(&result, id).Error
}

func (s *ServiceImpl) FindByEmail(hashedEmail string, result *User) error {
	return s.DB.Find(&result, &User{HashedEmail: hashedEmail}).Error
}

func (s *ServiceImpl) FindByUserName(name string, result *User) error {
	return s.DB.Find(&result, &User{Name: name}).Error
}

func (s *ServiceImpl) FindProfile(id uint, result *Profile) error {
	return s.DB.Where(&Profile{
		UserID: id,
	}).Find(&result).Error
}

func (s *ServiceImpl) Update(u *User) error {
	return nil
}

func (s *ServiceImpl) Remove(user *User) error {
	return s.DB.Delete(&user).Error
}

func validatePassword(inputPassword, password, salt string) error {
	encodePassword, err := encodePassword(inputPassword, salt)
	if err != nil {
		return err
	}

	if !cmpPassword(encodePassword, []byte(password)) {
		return fmt.Errorf("wrong password")
	}

	return nil
}

func encodePassword(password string, salt string) ([]byte, error) {
	encodePassword, err := scrypt.Key([]byte(password), []byte(salt), ScryptN, ScryptR, ScryptP, ScryptKeyLen)
	return encodePassword, err
}

func cmpPassword(p1 []byte, p2 []byte) bool {
	return bytes.Compare(p1, p2) == 0
}
