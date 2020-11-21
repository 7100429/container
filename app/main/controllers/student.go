package controllers

import (
	"fmt"
	"gitee.com/chensyf/container/app/main/models"
	"gitee.com/chensyf/container/app/main/models/request"
	"gitee.com/chensyf/container/app/main/services"
	"gitee.com/chensyf/container/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type studentController struct {}

var StudentController = &studentController{}


func (s *studentController) Router(g *gin.RouterGroup){
	g.GET("/student", s.FindStudentByName)
	g.GET("/students", s.FindAll)
	g.POST("/student", s.CreateStudent)
}

// @Summary 根据姓名查询学生
// @Description 根据姓名查询学生
// @Tags 学生
// @Accept  json
// @Produce json
// @Param name query string true "姓名"
// @Success 200 {object} models.Student
// @Router /student [get]
func (s *studentController) FindStudentByName(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "name is required",
		})
		return
	}
	var student models.Student

	fmt.Println(c.Request.URL.String())

	ok := utils.GetResponse(c.Request.URL.String(), &student)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"student": student,
		})
		return
	}

	student, err := services.StudentService.FindStudentByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	_ = utils.SetResponse(c.Request.URL.String(), student)

	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
	return
}

// @Summary 查询所有学生
// @Description 查询所有学生
// @Tags 学生
// @Accept  json
// @Produce json
// @Success 200 {array} models.Student
// @Router /students [get]
func (s *studentController) FindAll(c *gin.Context) {
	students := services.StudentService.FindAll()
	c.JSON(http.StatusOK, gin.H{
		"students": students,
	})
}

// @Summary 新增学生信息
// @Description 新增学生信息
// @Tags 学生
// @Accept  json
// @Produce json
// @Param   studentRequest body request.Student true "学生信息"
// @Success 200 {string} msg "创建成功"
// @Router /student [post]
func (s *studentController) CreateStudent(c *gin.Context) {
	var studentRequest request.Student
	if err := c.Bind(&studentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	err := services.StudentService.CreateStudent(&studentRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "创建成功！",
	})
}