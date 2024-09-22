package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"university-distribution-backend/services"
)

// GetAllJuniorColleges 获取所有专科学校
func GetAllJuniorColleges(c *gin.Context) {
	juniorColleges, err := services.GetAllJuniorColleges()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve junior colleges"})
		return
	}

	c.JSON(http.StatusOK, juniorColleges)
}

// GetJuniorCollegesByProvince 根据省份名称获取专科学校
func GetJuniorCollegesByProvince(c *gin.Context) {
    provinceName := c.Param("provinceName")

    juniorColleges, err := services.GetJuniorCollegesByProvince(provinceName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if len(juniorColleges) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "No junior colleges found for the specified province."})
        return
    }

    c.JSON(http.StatusOK, juniorColleges)
}

// GetJuniorCollegeByName 根据专科学校名称获取专科学校
func GetJuniorCollegeByName(c *gin.Context) {
    juniorCollegeName := c.Param("juniorCollegeName")

    juniorCollege, err := services.GetJuniorCollegeByName(juniorCollegeName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if juniorCollege == nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "No junior college found with the specified name."})
        return
    }

    c.JSON(http.StatusOK, juniorCollege)
}