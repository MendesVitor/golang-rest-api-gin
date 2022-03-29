package routes

import (
	"api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/students", controllers.ShowStudents)
	r.GET("/students/:id", controllers.GetStudentById)
	r.GET("/:name", controllers.Hello)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdatedStuden)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)
	r.Run()
}
