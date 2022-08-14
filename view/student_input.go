package view

import (
	"bufio"
	"controller_tuyensinh/controller"
	"fmt"
	"os"
	"strings"
)

func (v *StudentView) InputStudent(s *controller.Student) {
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Nhap ten sinh vien\n")
	name, _ := consoleReader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	s.Name = name

	fmt.Println("Nhap ngay thang nam sinh cua sinh vien\n")
	doB, _ := consoleReader.ReadString('\n')
	doB = strings.Replace(name, "\n", "", -1)
	s.DoB = doB

	fmt.Println("Nhap gioi tinh cua sinh vien\n")
	sex, _ := consoleReader.ReadString('\n')
	sex = strings.Replace(sex, "\n", "", -1)
	s.Sex = sex

	fmt.Println("Nhap sdt cua nhan vien\n")
	phone, _ := consoleReader.ReadString('\n')
	phone = strings.Replace(phone, "\n", "", -1)
	s.Phone = phone

}
