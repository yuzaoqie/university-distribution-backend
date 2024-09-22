package controllers

import (
    "net/http"
    "university-distribution-backend/models"
    "github.com/gin-gonic/gin"
    "university-distribution-backend/services"
)

// 获取所有大学信息
func GetUniversities(c *gin.Context) {
    universities, err := models.GetAllUniversities()
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

    c.JSON(http.StatusOK, universities)
}

// 根据大学名称获取大学信息
func GetUniversityByName(c *gin.Context) {
    universityName := c.Param("universityName")

    university, err := models.GetUniversityByName(universityName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if university == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "University not found"})
        return
    }

    c.JSON(http.StatusOK, university)
}