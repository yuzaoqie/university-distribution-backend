CREATE TABLE universities (
    id SERIAL PRIMARY KEY,  -- 主键，自增的大学ID
    name VARCHAR(100) NOT NULL,  -- 大学名称
    province_name VARCHAR(100),  -- 省份名称
    established_year INT,  -- 成立年份
    type VARCHAR(50) NOT NULL,  -- 类型，例如985、211、双一流等信息
    address VARCHAR(255),  -- 地址，可选
    website VARCHAR(255),  -- 官方网站，可选
    motto TEXT,  -- 校训
    description TEXT,  -- 大学描述
    history TEXT,  -- 学校历史
    enrollment_website VARCHAR(255),  -- 招生办网站
    public_private VARCHAR(20),  -- 公办/民办
    logo_path VARCHAR(255),  -- 校徽（logo）图片的文件路径
    background_image_path VARCHAR(255),  -- 学校背景（bg）图片的文件路径
    discipline_category VARCHAR(50),  -- 学科类别：理工类、文科类、综合类等
    graduate_points TEXT,  -- 研究生点、博士生点的详细信息
    faculty_strength TEXT,  -- 师资力量信息
    research_strength TEXT,  -- 科研实力：实验室、研究中心、重大科研成果
    notable_alumni TEXT,  -- 知名校友
    contact_phone VARCHAR(20),  -- 联系电话
    contact_email VARCHAR(100),  -- 联系邮箱
    campus_activities TEXT,  -- 校园活动与文化
    plan_diagram_path VARCHAR(255) -- 校园平面图的文件路径
);
