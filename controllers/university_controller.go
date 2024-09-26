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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "大学信息获取失败"})
		return
	}
	c.JSON(http.StatusOK, universities)
}

// 获取某个省份的大学列表
func GetUniversitiesByProvince(c *gin.Context) {
    provinceName := c.Param("provinceName")
    
	// 从 services 获取大学列表
    universities, err := services.GetUniversitiesByProvince(provinceName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if len(universities) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "该省份没有大学"})
        return
    }

    c.JSON(http.StatusOK, universities)
}

// 根据大学名称获取大学信息
func GetUniversityByName(c *gin.Context) {
	universityName := c.Param("universityName")

	// 从 services 获取大学信息
	university, err := services.GetUniversityByName(universityName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if university == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "该大学不存在"})
		return
	}

	// 拼接完整的静态资源 URL
	baseURL := "http://localhost:8080/assets/" // 根据后端静态资源路径
	university.LogoPath = baseURL + university.LogoPath
	university.BackgroundImagePath = baseURL + university.BackgroundImagePath
	university.PlanDiagramPath = baseURL + university.PlanDiagramPath

	// 返回大学信息
	c.JSON(http.StatusOK, university)
}

// 获取每个省的高校数量
func GetUniversityCountByProvince(c *gin.Context) {
	provinceUniversityCounts, err := services.GetUniversityCountByProvince()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取大学数量"})
		return
	}

	c.JSON(http.StatusOK, provinceUniversityCounts)
}

// 获取本科院校数量 TOP10 的省份
func GetTop10Provinces(c *gin.Context) {
	provinces, err := services.GetTop10Provinces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}
	c.JSON(http.StatusOK, provinces)
}