package services

import (
    "university-distribution-backend/models"
)

func FetchUniversities() ([]models.University, error) {
    return models.GetAllUniversities();
}

// 根据省份名称获取大学列表
func GetUniversitiesByProvince(provinceName string) ([]models.University, error) {
    return models.GetUniversitiesByProvince(provinceName);
}