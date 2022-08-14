package controller

import "controller_tuyensinh/database"

type StudentController struct {
	student *database.StudentRepo
}

func (s StudentController) toStudentModel(student Student) database.StudentModel {
	return database.StudentModel{
		Name:       student.Name,
		DoB:        student.DoB,
		Sex:        student.Sex,
		Phone:      student.Phone,
		University: student.University,

		GradelLevel: database.Grade(Grade(student.GradelLevel)),

		GPA:            student.GPA,
		BestRewardName: student.BestRewardName,

		EnglishScore:   student.EnglishScore,
		EntryTestScore: student.EntryTestScore,
	}
}

func NewStudentController(student *database.StudentRepo) *StudentController {
	return &StudentController{
		student: student,
	}
}
func (s StudentController) CreateStudent(student Student) error {
	studentModel := toStudentModel(student)
	err := s.student.CreateStudentRepo(studentModel).Error()
	if err != nil {
		return err
	}
	return nil
}
