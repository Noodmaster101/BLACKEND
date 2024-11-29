package main

import (
	EmployeeController "backend/api/controller/employee"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Employee API Methods
	router.GET("/employee", EmployeeController.GetEmployee)         // ดึงข้อมูลพนักงานทั้งหมด
	router.GET("/employee/:id", EmployeeController.GetEmployeebyID) // ดึงข้อมูลพนักงานตาม ID
	router.POST("/employee", EmployeeController.PostEmployee)       // เพิ่มข้อมูลพนักงาน
	router.PUT("/employee", EmployeeController.PutEmployee)         // แก้ไขข้อมูลพนักงาน
	router.DELETE("/employee", EmployeeController.Deletemployee)    // ลบข้อมูลพนักงาน (แก้ไขเครื่องหมายจุลภาค)

	// เริ่มเซิร์ฟเวอร์
	router.Run() // เริ่มฟังที่ 0.0.0.0:8080 หรือ "localhost:8080"
}
