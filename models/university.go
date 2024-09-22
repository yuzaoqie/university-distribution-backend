package models

import (
    "university-distribution-backend/config"
)

type University struct {
    ID                   int    `json:"id"`                     // 自增的大学ID
    Name                 string `json:"name"`                   // 大学名称
    ProvinceID           int    `json:"province_id"`            // 省份ID，外键
    EstablishedYear      int    `json:"established_year"`       // 成立年份
    Type                 string `json:"type"`                   // 类型，例如 '本科' 或 '专科'
    Address              string `json:"address"`                // 地址
    Website              string `json:"website"`                // 官方网站
    StrengthDescription  string `json:"strength_description"`   // 学校实力的具体描述
    Motto                string `json:"motto"`                  // 校训
    Description          string `json:"description"`            // 大学描述
    History              string `json:"history"`                // 学校历史
    EnrollmentWebsite     string `json:"enrollment_website"`     // 招生办网站
    PublicPrivate        string `json:"public_private"`         // 公办/民办
    LogoPath             string `json:"logo_path"`              // 校徽（logo）图片的文件路径
    BackgroundImagePath  string `json:"background_image_path"`  // 学校背景（bg）图片的文件路径
    DisciplineCategory    string `json:"discipline_category"`     // 学科类别
    GraduatePoints       string `json:"graduate_points"`        // 研究生点、博士生点的详细信息
    FacultyStrength      string `json:"faculty_strength"`       // 师资力量信息
    ResearchStrength     string `json:"research_strength"`      // 科研实力
    NotableAlumni       string `json:"notable_alumni"`        // 知名校友
    ContactPhone         string `json:"contact_phone"`          // 联系电话
    ContactEmail         string `json:"contact_email"`          // 联系邮箱
    EnrollmentAddress    string `json:"enrollment_address"`     // 招生办地址
    CampusActivities     string `json:"campus_activities"`      // 校园活动与文化
}

// 获取所有大学
func GetAllUniversities() ([]University, error) {
    rows, err := config.DB.Query("SELECT id, name, province_id, established_year, type, address, website, strength_description, motto, description, history, enrollment_website, public_private, logo_path, background_image_path, discipline_category, graduate_points, faculty_strength, research_strength, notable_alumni, contact_phone, contact_email, enrollment_address, campus_activities FROM universities")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    universities := []University{}
    for rows.Next() {
        var u University
        if err := rows.Scan(&u.ID, &u.Name, &u.ProvinceID, &u.EstablishedYear, &u.Type, &u.Address, &u.Website, &u.StrengthDescription, &u.Motto, &u.Description, &u.History, &u.EnrollmentWebsite, &u.PublicPrivate, &u.LogoPath, &u.BackgroundImagePath, &u.DisciplineCategory, &u.GraduatePoints, &u.FacultyStrength, &u.ResearchStrength, &u.NotableAlumni, &u.ContactPhone, &u.ContactEmail, &u.EnrollmentAddress, &u.CampusActivities); err != nil {
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
            u.id, 
            u.name, 
            u.province_id, 
            u.established_year, 
            u.type, 
            u.address, 
            u.website, 
            u.strength_description, 
            u.motto, 
            u.description, 
            u.history, 
            u.enrollment_website, 
            u.public_private, 
            u.logo_path, 
            u.background_image_path, 
            u.discipline_category, 
            u.graduate_points, 
            u.faculty_strength, 
            u.research_strength, 
            u.notable_alumni, 
            u.contact_phone, 
            u.contact_email, 
            u.enrollment_address, 
            u.campus_activities
        FROM universities u
        JOIN provinces p ON u.province_id = p.id
        WHERE p.name = $1;
    `

    rows, err := config.DB.Query(query, provinceName)  // 根据省份名称查询
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    universities := []University{}
    for rows.Next() {
        var u University
        if err := rows.Scan(&u.ID, &u.Name, &u.ProvinceID, &u.EstablishedYear, &u.Type, &u.Address, &u.Website, &u.StrengthDescription, &u.Motto, &u.Description, &u.History, &u.EnrollmentWebsite, &u.PublicPrivate, &u.LogoPath, &u.BackgroundImagePath, &u.DisciplineCategory, &u.GraduatePoints, &u.FacultyStrength, &u.ResearchStrength, &u.NotableAlumni, &u.ContactPhone, &u.ContactEmail, &u.EnrollmentAddress, &u.CampusActivities); err != nil {
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
            province_id, 
            established_year, 
            type, 
            address, 
            website, 
            strength_description, 
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
            enrollment_address, 
            campus_activities
        FROM universities 
        WHERE name = $1;
    `

    var university University
    err := config.DB.QueryRow(query, universityName).Scan(
        &university.ID,
        &university.Name,
        &university.ProvinceID,
        &university.EstablishedYear,
        &university.Type,
        &university.Address,
        &university.Website,
        &university.StrengthDescription,
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
        &university.EnrollmentAddress,
        &university.CampusActivities,
    )

    if err != nil {
        return nil, err
    }

    return &university, nil
}