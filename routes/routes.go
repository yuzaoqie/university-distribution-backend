package routes

import (
    "github.com/gin-gonic/gin"
    "university-distribution-backend/controllers"
)

func SetupRoutes(router *gin.Engine) {
    // 获取所有大学信息
    router.GET("/api/universities", controllers.GetUniversities);
    // TODO 待修改，根据省份名称获取大学列表 /api/provinceMap/:provinceName
    // router.GET("/api/universities/:provinceName", controllers.GetUniversitiesByProvince);
     // 根据大学名称获取大学信息
     router.GET("/api/universities/:universityName", controllers.GetUniversityByName)
}
