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
	// 获取每个省的高校数量
	router.GET("/api/universityCount", controllers.GetUniversityCountByProvince)
	// 获取本科院校数量 TOP10 的省份
	router.GET("/api/top10Provinces", controllers.GetTop10Provinces)
}
