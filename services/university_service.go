package services

import (
    "university-distribution-backend/models"
)

// 获取所有大学信息
func GetAllUniversities() ([]models.University, error) {
    return models.GetAllUniversities();
}

// 根据省份名称获取大学列表
func GetUniversitiesByProvince(provinceName string) ([]models.University, error) {
    return models.GetUniversitiesByProvince(provinceName);
}

// 根据大学名称获取大学信息
func GetUniversityByName(universityName string) (*models.University, error) {
    return models.GetUniversityByName(universityName);
}
