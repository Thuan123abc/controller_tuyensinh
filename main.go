package controller_tuyensinh

import (
	"controller_tuyensinh/controller"
	"controller_tuyensinh/database"
	"controller_tuyensinh/view"
)

func main() {
	db := database.NewDB()
	eRepo := database.NewStudentRepo(db)
	eController := controller.NewStudentController(eRepo)
	eView := view.NewStudentView(eController)

	eView.CreateStudent()
}
