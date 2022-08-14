package database

import (
	"fmt"
	"gorm.io/gorm"
)

type StudentRepo struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *StudentRepo {
	db.AutoMigrate(&db)
	return &StudentRepo{
		db: db,
	}
}
func (s StudentRepo) CreateStudentRepo(model StudentModel) error {
	err := s.db.Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}
func (s StudentRepo) DeleteStudentRepo(id int64) error {
	err := s.db.Delete("id=?", id).Error
	if err != nil {
		fmt.Println("Delete khong thanh cong")
		return err
	}
	fmt.Println("Delete thanh cong")
	return nil
}

func (s StudentRepo) UpdateStudentRepo(model StudentModel) error {
	err := s.db.Model(&model).Omit("id").Updates(model).Error
	if err != nil {
		fmt.Println("update khong thanh cong")
		return err
	}
	fmt.Println("update thanh cong")
	return nil
}

func (s StudentRepo) GetStudentRepo(id int64) (*StudentModel, error) {
	var student StudentModel
	err := s.db.Where("id=", id).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (s StudentRepo) GetListStudentRepo() ([]StudentModel, error) {
	var list []StudentModel
	err := s.db.Find(list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
