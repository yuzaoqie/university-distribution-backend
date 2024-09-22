package services

import (
    "university-distribution-backend/models"
)

// 获取所有专科信息
func GetAllJuniorColleges() ([]models.JuniorCollege, error) {
    return models.GetAllJuniorColleges();
}

// 根据省份名称获取专科列表
func GetJuniorCollegesByProvince(provinceName string) ([]models.JuniorCollege, error) {
    return models.GetJuniorCollegesByProvince(provinceName);
}

// 根据专科名称获取专科信息
func GetJuniorCollegeByName(juniorCollegeName string) (*models.JuniorCollege, error) {
    return models.GetJuniorCollegeByName(juniorCollegeName);
}