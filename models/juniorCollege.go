package models

import (
	"university-distribution-backend/config"
)

// JuniorCollege 表示专科学校的模型
type JuniorCollege struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	ProvinceName        string `json:"province_name"`
	EstablishedYear     int    `json:"established_year"`
	Type                string `json:"type"`
	Address             string `json:"address"`
	Website             string `json:"website"`
	Motto               string `json:"motto"`
	Description         string `json:"description"`
	History             string `json:"history"`
	PublicPrivate       string `json:"public_private"`
	LogoPath            string `json:"logo_path"`
	BackgroundImagePath string `json:"background_image_path"`
	DisciplineCategory  string `json:"discipline_category"`
	GraduatePoints      string `json:"graduate_points"`
	FacultyStrength     string `json:"faculty_strength"`
	ResearchStrength    string `json:"research_strength"`
	NotableAlumni       string `json:"notable_alumni"`
	ContactPhone        string `json:"contact_phone"`
	ContactEmail        string `json:"contact_email"`
	CampusActivities    string `json:"campus_activities"`
}

// GetAllJuniorColleges 获取所有专科学校
func GetAllJuniorColleges() ([]JuniorCollege, error) {
	query := `
		SELECT 
			id, 
			name, 
			province_name, 
			established_year, 
			type, 
			address, 
			website, 
			motto, 
			description, 
			history, 
			public_private, 
			logo_path, 
			background_image_path, 
			discipline_category, 
			graduate_points, 
			faculty_strength, 
			research_strength, 
			notable_alumni, 
			contact_phone, 
			contact_email, 
			campus_activities 
		FROM junior_colleges`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	juniorColleges := []JuniorCollege{}
	for rows.Next() {
		var jc JuniorCollege
		if err := rows.Scan(&jc.ID, &jc.Name, &jc.ProvinceName, &jc.EstablishedYear, &jc.Type, &jc.Address, &jc.Website, &jc.Motto, &jc.Description, &jc.History, &jc.PublicPrivate, &jc.LogoPath, &jc.BackgroundImagePath, &jc.DisciplineCategory, &jc.GraduatePoints, &jc.FacultyStrength, &jc.ResearchStrength, &jc.NotableAlumni, &jc.ContactPhone, &jc.ContactEmail, &jc.CampusActivities); err != nil {
			return nil, err
		}
		juniorColleges = append(juniorColleges, jc)
	}

	return juniorColleges, nil
}

// 根据省份名称获取专科学校
// GetJuniorCollegesByProvince 根据省份名称获取专科学校
func GetJuniorCollegesByProvince(provinceName string) ([]JuniorCollege, error) {
	query := `
		SELECT 
			id, 
			name, 
			province_name, 
			established_year, 
			type, 
			address, 
			website, 
			motto, 
			description, 
			history, 
			public_private, 
			logo_path, 
			background_image_path, 
			discipline_category, 
			graduate_points, 
			faculty_strength, 
			research_strength, 
			notable_alumni, 
			contact_phone, 
			contact_email, 
			campus_activities 
        FROM junior_colleges 
    	WHERE province_name = $1`

	rows, err := config.DB.Query(query, provinceName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	juniorColleges := []JuniorCollege{}
	for rows.Next() {
		var jc JuniorCollege
		if err := rows.Scan(&jc.ID, &jc.Name, &jc.ProvinceName, &jc.EstablishedYear, &jc.Type, &jc.Address, &jc.Website, &jc.Motto, &jc.Description, &jc.History, &jc.PublicPrivate, &jc.LogoPath, &jc.BackgroundImagePath, &jc.DisciplineCategory, &jc.GraduatePoints, &jc.FacultyStrength, &jc.ResearchStrength, &jc.NotableAlumni, &jc.ContactPhone, &jc.ContactEmail, &jc.CampusActivities); err != nil {
			return nil, err
		}
		juniorColleges = append(juniorColleges, jc)
	}

	return juniorColleges, nil
}

// 根据专科名称获取专科信息
func GetJuniorCollegeByName(juniorCollegeName string) (*JuniorCollege, error) {
	query := `
		SELECT 
			id, 
			name, 
			province_name, 
			established_year, 
			type, 
			address, 
			website, 
			motto, 
			description, 
			history, 
			public_private, 
			logo_path, 
			background_image_path, 
			discipline_category, 
			graduate_points, 
			faculty_strength, 
			research_strength, 
			notable_alumni, 
			contact_phone, 
			contact_email, 
			campus_activities 
		FROM junior_colleges 
		WHERE name = $1`

	row := config.DB.QueryRow(query, juniorCollegeName)

	// TODO 有必要吗: 将数据库查询结果映射到JuniorCollege结构体
	var jc JuniorCollege
	if err := row.Scan(&jc.ID, &jc.Name, &jc.ProvinceName, &jc.EstablishedYear, &jc.Type, &jc.Address, &jc.Website, &jc.Motto, &jc.Description, &jc.History, &jc.PublicPrivate, &jc.LogoPath, &jc.BackgroundImagePath, &jc.DisciplineCategory, &jc.GraduatePoints, &jc.FacultyStrength, &jc.ResearchStrength, &jc.NotableAlumni, &jc.ContactPhone, &jc.ContactEmail, &jc.CampusActivities); err != nil {
		return nil, err
	}

	return &jc, nil
}