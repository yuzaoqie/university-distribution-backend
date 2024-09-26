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

// 获取每个省的高校数量
func GetUniversityCountByProvince() ([]models.ProvinceUniversityCount, error) {
    return models.GetUniversityCountByProvince();
}

// 获取本科院校数量 TOP10 的省份
func GetTop10Provinces() ([]models.ProvinceUniversityCount, error) {
    return models.GetTop10Provinces();
}