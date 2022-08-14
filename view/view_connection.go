package view

import (
	"controller_tuyensinh/controller"
	"fmt"
)

type StudentView struct {
	studentController *controller.StudentController
}

func NewStudentView(studentController *controller.StudentController) *StudentView {
	return &StudentView{
		studentController: studentController,
	}
}

func (v *StudentView) CreateStudent() {
	var e controller.Student
	v.InputStudent(&e)
	err := v.studentController.CreateStudent(e)
	if err != nil {
		fmt.Println("CreateEmployee failed")
		return
	}
	fmt.Println("CreateEmployee ok")
}
