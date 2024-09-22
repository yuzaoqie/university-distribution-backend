package routes

import (
	"university-distribution-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// 获取所有大学信息
	router.GET("/api/universities", controllers.GetUniversities)
	// 根据省份名称获取大学列表 /api/universitiesByProvinceMap/:provinceName
	router.GET("/api/universitiesByProvinceMap/:provinceName", controllers.GetUniversitiesByProvince)
	// 根据大学名称获取大学信息
	router.GET("/api/universities/:universityName", controllers.GetUniversityByName)

	// 获取所有专科学校信息
	router.GET("/api/juniorColleges", controllers.GetAllJuniorColleges)
	// 根据省份名称获取专科学校列表
	router.GET("/api/juniorCollegesByProvinceMap/:provinceName", controllers.GetJuniorCollegesByProvince)
	// 根据专科学校名称获取专科学校信息
	router.GET("/api/juniorColleges/:juniorCollegeName", controllers.GetJuniorCollegeByName)
}
