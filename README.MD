下面是你的后端项目的 README 模板，基于你的项目结构：

---

# 中国大学分布图 - 后端部分

## 项目简介
本项目后端基于 Go 语言和 Gin 框架开发，主要负责处理中国大学分布数据的 API 请求。前端与后端通过 RESTful API 进行交互，后端从数据库中获取大学相关信息，并返回给前端用于展示。

## 功能特点
- 提供各省高校分布信息的 RESTful API。
- 提供高校详细信息的 API 接口。
- 使用 Go 语言结合 Gin 框架进行高效的 HTTP 请求处理。
- 与 PostgreSQL 数据库交互，存储高校数据。

## 项目结构
```bash
university-distribution-backend/
├── assets/                # 静态资源文件夹
├── config/                # 配置文件
│   └── config.go          # 数据库连接等配置
├── controllers/           # 控制器，处理 HTTP 请求和响应
│   └── university_controller.go   # 处理大学相关的请求
├── models/                # 数据模型和数据库交互逻辑
│   └── university.go      # 大学模型
├── routes/                # 路由定义
│   └── routes.go          # 注册 API 路由
├── services/              # 业务逻辑
│   └── university_service.go      # 大学相关的业务逻辑
└── main.go                # 项目入口，初始化服务器和数据库连接
```

## 技术栈
- **编程语言**：Go
- **Web 框架**：Gin
- **数据库**：PostgreSQL
- **API 格式**：RESTful

## 安装与运行

### 前提条件
- Go 1.23.1 版本或更高
- PostgreSQL 数据库

### 安装步骤

1. 克隆项目代码：
   ```bash
   git clone https://github.com/yourname/university-distribution-backend.git
   cd university-distribution-backend
   ```

2. 配置环境变量：
   在项目的根目录下创建 `.env` 文件，配置数据库连接等必要的信息：

   ```bash
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=yourusername
   DB_PASSWORD=yourpassword
   DB_NAME=university_db
   ```

3. 安装依赖：
   ```bash
   go mod tidy
   ```

4. 启动服务器：
   ```bash
   go run main.go
   ```

5. 打开浏览器访问 `http://localhost:8080`，或者根据你定义的 API 路径进行请求。

## API 接口说明
后端提供了以下 API 路径用于获取高校信息：

- **获取省份高校信息**：
  - `GET /api/provinces/:provinceName`：获取指定省份的高校分布信息。
  
- **获取高校详细信息**：
  - `GET /api/universities/:universityId`：获取指定高校的详细信息。

## 数据库设计
数据库包含两个主要表：
- **provinces**：存储中国各省份的信息。
- **universities**：存储高校信息，包含高校名称、类别、所属省份、是否为 985/211/双一流 等字段。
