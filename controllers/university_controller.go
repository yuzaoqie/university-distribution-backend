package controllers

import (
	"net/http"
	"university-distribution-backend/services"
	"github.com/gin-gonic/gin"
)

// 获取所有大学信息
func GetUniversities(c *gin.Context) {
	universities, err := services.GetAllUniversities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve universities"})
		return
	}
	c.JSON(http.StatusOK, universities)
}

// 获取某个省份的大学列表
func GetUniversitiesByProvince(c *gin.Context) {
    provinceName := c.Param("provinceName")
    
    universities, err := services.GetUniversitiesByProvince(provinceName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if len(universities) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No universities found for this province"})
        return
    }

    c.JSON(http.StatusOK, universities)
}

// 根据大学名称获取大学信息
func GetUniversityByName(c *gin.Context) {
	universityName := c.Param("universityName")

	// 从 models 获取大学信息
	university, err := services.GetUniversityByName(universityName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if university == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "University not found"})
		return
	}

	// 拼接完整的静态资源 URL
	baseURL := "http://localhost:8080/assets/" // 根据后端静态资源路径
	university.LogoPath = baseURL + university.LogoPath
	university.BackgroundImagePath = baseURL + university.BackgroundImagePath

	// 返回大学信息
	c.JSON(http.StatusOK, university)
}
