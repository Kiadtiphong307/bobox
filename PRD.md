
# Project Requirement Document (PRD)

## Project Name
BoBox: ระบบขายไอดีเกมและร้านเติมเกมพร้อมระบบส่วนลดและจัดการโดยแอดมิน

---

## Objective
สร้างระบบ e-commerce สำหรับขายไอดีเกม/เติมเกมแบบออนไลน์ที่มีฟีเจอร์การจัดการส่วนลด, การจำกัดการใช้คูปอง, ระบบผู้ใช้ และระบบแอดมิน โดยใช้ Nuxt.js 3 (frontend), 
Go + Fiber (backend),
MySQL (database), 
Docker (runtime), 
Nginx (reverse proxy)

---

## 1. Features

### 1.1 ระบบผู้ใช้ (User System)
- สมัครสมาชิก / เข้าสู่ระบบ (JWT Authentication)
- ดูและแก้ไขโปรไฟล์ส่วนตัว
- จัดการรายการคำสั่งซื้อย้อนหลังของตนเอง

### 1.2 การขายสินค้า (Game Top-Up / Game ID Selling)
- รายการสินค้า: ไอดีเกม, เติมเกม, บัตรเติมเงิน
- สินค้าแต่ละประเภทอาจมีรายละเอียดแตกต่างกัน (เช่น ระบุเกม, server, วิธีติดต่อ)

### 1.3 ระบบส่วนลด (Discount System)
- สร้างคูปองส่วนลดจากฝั่งแอดมิน
- จำกัดจำนวนการใช้คูปองแต่ละรายการ / จำกัดผู้ใช้ / จำกัดวันหมดอายุ
- ตรวจสอบการใช้คูปองจาก backend ก่อนชำระเงิน

### 1.4 ระบบแอดมิน (Admin Dashboard)
- จัดการสินค้า (เพิ่ม, แก้ไข, ลบ)
- จัดการคูปองส่วนลด
- ดูคำสั่งซื้อทั้งหมด
- จัดการผู้ใช้ (ดูข้อมูล, ลบบัญชี)

### 1.5 คำสั่งซื้อ (Order System)
- ผู้ใช้สามารถเลือกสินค้าและทำรายการสั่งซื้อ
- ระบบสามารถบันทึกสถานะ (รอดำเนินการ / สำเร็จ / ยกเลิก)

---

## 2. Tech Stack

- **Frontend:** Nuxt.js 3
- **Styling:** Tailwind CSS
- **Backend:** Golang + Fiber
- **Database:** MySQL
- **Runtime:** Docker
- **Proxy:** Nginx

---

## 3. Workflow Plan

### Phase 1: Auth & User System
- Register / Login (JWT)
- Profile view/edit

### Phase 2: Admin Backend
- Admin Role Authentication Middleware
- CRUD สินค้า + คูปองส่วนลด

### Phase 3: Public Frontend
- แสดงสินค้าทั้งหมด
- ค้นหาสินค้า
- กรอกคูปองส่วนลดก่อน checkout

### Phase 4: Order & Checkout
- สร้างคำสั่งซื้อ
- ตรวจสอบการใช้คูปองแบบ realtime
- แสดงประวัติคำสั่งซื้อ

---

## 4. Folder Structure Proposal

### Frontend (Nuxt.js 3)
```
frontend/
├── assets/
├── components/
│   ├── base/
│   ├── forms/
│   └── layout/
├── composables/
├── layouts/
├── middleware/
├── pages/
│   ├── index.vue
│   ├── login.vue
│   ├── register.vue
│   ├── admin/
│   ├── profile.vue
│   ├── product/
│   ├── checkout.vue
├── plugins/
├── utils/
├── app.vue
└── nuxt.config.ts
```

### Backend (Go + Fiber)
```
backend/
├── controller/
│   ├── auth_controller.go
│   ├── admin_controller.go
│   ├── product_controller.go
│   ├── discount_controller.go
│   └── order_controller.go
├── middleware/
│   └── jwt.go
├── models/
│   ├── user.go
│   ├── product.go
│   ├── discount.go
│   ├── order.go
│   └── migration.go
├── validation/
│   ├── user_validation.go
│   ├── product_validation.go
│   └── discount_validation.go
├── service/
│   └── business_logic.go
├── routes/
│   └── routes.go
├── database/
│   └── init.go
├── utils/
│   └── helper.go
├── .env
└── main.go
```

---

## 5. Database Schema Overview

### users
- id (PK)
- username
- fullname
- email
- password
- role ["user", "admin"]

### products
- id (PK)
- name
- category ["topup", "game_id"]
- description
- price
- stock
- create_at

### discounts
- id (PK)
- code
- percentage
- usage_limit
- expires_at

### orders
- id (PK)
- user_id (FK)
- product_id (FK)
- discount_id (FK / nullable)
- status ["pending", "completed", "canceled"]
- created_at

---

## 6. API Overview

### Public
- POST /api/auth/register
- POST /api/auth/login
- GET /api/products

### User Authenticated
- GET /api/user/profile
- PUT /api/user/profile
- POST /api/orders
- GET /api/orders

### Admin Only
- GET /api/admin/users
- POST /api/admin/products
- PUT /api/admin/products/:id
- DELETE /api/admin/products/:id
- POST /api/admin/discounts
- PUT /api/admin/discounts/:id
- DELETE /api/admin/discounts/:id

---

## 7. หมายเหตุพิเศษจาก Feedback Boblog

- แยก model / controller / validation / logic ชัดเจน (Clean Architecture)
- แยก Business Logic (service layer)
- แยก Middleware สำหรับ Role (admin/user)
- ใช้ helper function และ Reusable components
- ตรวจสอบ API response ด้วยมาตรฐาน JSON response
- ใช้ Docker + .env สำหรับ config ทั้ง backend/frontend



## Command 

```
## Frontend
<!-- isntall nuxt 3 & Tailwind -->
npm create nuxt@latest 
npm install tailwindcss @tailwindcss/vite
npm install @nuxtjs/axios

//npm i & npm run build
```

```
## Backend
<!-- Install Fiber(Go) & Library  -->
//go mod init backend
go get github.com/gofiber/fiber/v2
go get -u github.com/gofiber/contrib/jwt
go get -u github.com/golang-jwt/jwt/v5
go get github.com/go-sql-driver/mysql
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/joho/godotenv
go get golang.org/x/crypto/bcrypt
go install github.com/air-verse/air@latest
go get github.com/go-playground/validator/v10
//go mod tidy 
```

## 
frontend = localhost/
backeend = localhost/api
