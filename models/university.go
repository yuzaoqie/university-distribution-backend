package models

import (
	"university-distribution-backend/config"
)

type University struct {
	ID                  int    `json:"id"`                    // 自增的大学ID
	Name                string `json:"name"`                  // 大学名称
	ProvinceName        string `json:"province_name"`         // 省份名称
	EstablishedYear     int    `json:"established_year"`      // 成立年份
	Type                string `json:"type"`                  // 类型，例如 985、211、双一流等
	Address             string `json:"address"`               // 地址
	Website             string `json:"website"`               // 官方网站
	Motto               string `json:"motto"`                 // 校训
	Description         string `json:"description"`           // 大学描述
	History             string `json:"history"`               // 学校历史
	EnrollmentWebsite   string `json:"enrollment_website"`    // 招生办网站
	PublicPrivate       string `json:"public_private"`        // 公办/民办
	LogoPath            string `json:"logo_path"`             // 校徽（logo）图片的文件路径
	BackgroundImagePath string `json:"background_image_path"` // 学校背景（bg）图片的文件路径
	DisciplineCategory  string `json:"discipline_category"`   // 学科类别
	GraduatePoints      string `json:"graduate_points"`       // 研究生点、博士生点的详细信息
	FacultyStrength     string `json:"faculty_strength"`      // 师资力量信息
	ResearchStrength    string `json:"research_strength"`     // 科研实力
	NotableAlumni       string `json:"notable_alumni"`        // 知名校友
	ContactPhone        string `json:"contact_phone"`         // 联系电话
	ContactEmail        string `json:"contact_email"`         // 联系邮箱
	CampusActivities    string `json:"campus_activities"`     // 校园活动与文化
	PlanDiagramPath     string `json:"plan_diagram_path"`     // 校园平面图的文件路径
	Level               string `json:"level"`                 // 学校层次，例如 本科、专科
}

// ProvinceUniversityCount 存储省份名称和高校数量
type ProvinceUniversityCount struct {
	ProvinceName    string `json:"province_name"`
	UniversityCount int    `json:"university_count"`
}

// 获取所有大学
func GetAllUniversities() ([]University, error) {
	rows, err := config.DB.Query("SELECT id, name, province_name, established_year, type, address, website, motto, description, history, enrollment_website, public_private, logo_path, background_image_path, discipline_category, graduate_points, faculty_strength, research_strength, notable_alumni, contact_phone, contact_email, campus_activities, plan_diagram_path, level FROM universities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	universities := []University{}
	for rows.Next() {
		var u University
		if err := rows.Scan(&u.ID, &u.Name, &u.ProvinceName, &u.EstablishedYear, &u.Type, &u.Address, &u.Website, &u.Motto, &u.Description, &u.History, &u.EnrollmentWebsite, &u.PublicPrivate, &u.LogoPath, &u.BackgroundImagePath, &u.DisciplineCategory, &u.GraduatePoints, &u.FacultyStrength, &u.ResearchStrength, &u.NotableAlumni, &u.ContactPhone, &u.ContactEmail, &u.CampusActivities, &u.PlanDiagramPath, &u.Level); err != nil {
			return nil, err
		}
		universities = append(universities, u)
	}

	return universities, nil
}

// 根据省份名称获取大学
func GetUniversitiesByProvince(provinceName string) ([]University, error) {
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
            enrollment_website, 
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
            campus_activities,
			plan_diagram_path,
			level
        FROM universities
        WHERE province_name = $1;
    `

	rows, err := config.DB.Query(query, provinceName) // 根据省份名称查询
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	universities := []University{}
	for rows.Next() {
		var u University
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.ProvinceName,
			&u.EstablishedYear,
			&u.Type,
			&u.Address,
			&u.Website,
			&u.Motto,
			&u.Description,
			&u.History,
			&u.EnrollmentWebsite,
			&u.PublicPrivate,
			&u.LogoPath,
			&u.BackgroundImagePath,
			&u.DisciplineCategory,
			&u.GraduatePoints,
			&u.FacultyStrength,
			&u.ResearchStrength,
			&u.NotableAlumni,
			&u.ContactPhone,
			&u.ContactEmail,
			&u.CampusActivities,
			&u.PlanDiagramPath,
			&u.Level); err != nil {
			return nil, err
		}
		universities = append(universities, u)
	}

	return universities, nil
}

// 根据大学名称获取大学信息
func GetUniversityByName(universityName string) (*University, error) {
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
            enrollment_website, 
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
            campus_activities,
			plan_diagram_path,
			level
        FROM universities 
        WHERE name = $1;
    `

	var university University
	err := config.DB.QueryRow(query, universityName).Scan(
		&university.ID,
		&university.Name,
		&university.ProvinceName,
		&university.EstablishedYear,
		&university.Type,
		&university.Address,
		&university.Website,
		&university.Motto,
		&university.Description,
		&university.History,
		&university.EnrollmentWebsite,
		&university.PublicPrivate,
		&university.LogoPath,
		&university.BackgroundImagePath,
		&university.DisciplineCategory,
		&university.GraduatePoints,
		&university.FacultyStrength,
		&university.ResearchStrength,
		&university.NotableAlumni,
		&university.ContactPhone,
		&university.ContactEmail,
		&university.CampusActivities,
		&university.PlanDiagramPath,
		&university.Level,
	)

	if err != nil {
		return nil, err
	}

	return &university, nil
}

// 获取每个省的高校数量
func GetUniversityCountByProvince() ([]ProvinceUniversityCount, error) {
	query := `
        SELECT province_name, COUNT(*) as university_count
        FROM universities
        GROUP BY province_name;
    `

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	provinceUniversityCounts := []ProvinceUniversityCount{}
	for rows.Next() {
		var puc ProvinceUniversityCount
		if err := rows.Scan(&puc.ProvinceName, &puc.UniversityCount); err != nil {
			return nil, err
		}
		provinceUniversityCounts = append(provinceUniversityCounts, puc)
	}

	return provinceUniversityCounts, nil
}

// 获取本科院校数量 TOP10 的省份
func GetTop10Provinces() ([]ProvinceUniversityCount, error) {
	query := `
		SELECT province_name, COUNT(*) as count
		FROM universities
		WHERE level = '本科' -- 只计算本科院校
		GROUP BY province_name
		ORDER BY count DESC
		LIMIT 10;
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var provinces []ProvinceUniversityCount
	for rows.Next() {
		var province ProvinceUniversityCount
		if err := rows.Scan(&province.ProvinceName, &province.UniversityCount); err != nil {
			return nil, err
		}
		provinces = append(provinces, province)
	}

	return provinces, nil
}
