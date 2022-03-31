package main

import (
	"api-gin/controllers"
	"api-gin/database"
	"api-gin/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

var ID int

func CreateStudentMock() {
	student := models.Student{Name: "Test", Cpf: "12345678901", Rg: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestVerifyHelloStatusCode(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:name", controllers.Hello)
	req, _ := http.NewRequest("GET", "/vitor", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "should be equal")
	resMock := `{"message":"vitor praise the sun"}`
	resBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, resMock, string(resBody))
}

func TestGetAllStudents(t *testing.T) {
	database.ConnDB()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.GET("/students", controllers.ShowStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetByCpf(t *testing.T) {
	database.ConnDB()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)
	req, _ := http.NewRequest("GET", "/students/cpf/12345678901", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetStudentById(t *testing.T) {
	database.ConnDB()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.GET("/students/:id", controllers.GetStudentById)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var studentMock models.Student
	json.Unmarshal(res.Body.Bytes(), &studentMock)
	assert.Equal(t, "Test", studentMock.Name)
	assert.Equal(t, "12345678901", studentMock.Cpf)
	assert.Equal(t, "123456789", studentMock.Rg)

}

func TestDeleteStudent(t *testing.T) {
	database.ConnDB()
	CreateStudentMock()
	r := SetupTestRoutes()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

}

func TestUpdateStudent(t *testing.T) {
	database.ConnDB()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.PATCH("/students/:id", controllers.UpdatedStuden)
	path := "/students/" + strconv.Itoa(ID)
	student := models.Student{Name: "Test", Cpf: "12345678999", Rg: "123456799"}
	dataJson, _ := json.Marshal(student)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(dataJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var updatedStudentMock models.Student
	json.Unmarshal(res.Body.Bytes(), &updatedStudentMock)
	assert.Equal(t, "12345678999", updatedStudentMock.Cpf)
	assert.Equal(t, "123456799", updatedStudentMock.Rg)
	assert.Equal(t, "Test", updatedStudentMock.Name)

}
