# JobSeek API

JobSeek API adalah backend untuk aplikasi pencari kerja yang dibangun dengan Golang dan framework Gin. API ini menyediakan fitur lengkap untuk manajemen pekerjaan, user (perusahaan & freelancer), komunikasi real-time melalui WebSocket, serta fitur tambahan seperti review dan penyimpanan pekerjaan/freelancer favorit.

## 🔥 Fitur Utama

- **Autentikasi & Manajemen Pengguna**
  - Registrasi & Login dengan JWT
  - Middleware Role-based (Admin, Freelancer, Perusahaan)
  - Manajemen Profil User (Update Profil, Upload Foto Profil via Cloudinary)

- **Manajemen Pekerjaan (Job Management)**
  - CRUD Job (Hanya Perusahaan yang bisa membuat job)
  - Pencarian & Filter Job (Kategori, Lokasi, Level Pengalaman)
  - Menyimpan pekerjaan favorit (Saved Jobs)

- **Manajemen Freelancer**
  - Melihat daftar freelancer
  - Menyimpan freelancer favorit (Saved Freelancer)
  - Memberikan review kepada freelancer

- **Manajemen Perusahaan**
  - Melihat daftar perusahaan
  - Memberikan review kepada perusahaan

- **Chat & Notifikasi Real-time**
  - Chat antara freelancer & perusahaan menggunakan WebSocket
  - Notifikasi real-time untuk pesan baru dan update penting

- **Review & Rating**
  - Freelancer dapat diberi review oleh perusahaan
  - Perusahaan dapat diberi review oleh freelancer

## 🚀 Teknologi yang Digunakan

- **Backend**: Golang (Gin Framework)
- **Database**: MySQL
- **ORM**: GORM
- **Autentikasi**: JWT
- **Real-time Communication**: Gorilla WebSocket
- **File Storage**: Cloudinary (untuk upload gambar profil)

## 📌 Instalasi & Menjalankan Project

1. **Clone Repository**

   ```bash
   git clone https://github.com/habbazettt/jobseek-api.git
   cd jobseek-api
   ```

2. **Buat file `.env`** berdasarkan template `.env.example`

3. **Jalankan Aplikasi**

   ```bash
   go mod tidy
   go run main.go
   ```

## 📡 Dokumentasi API (Swagger)

API ini menggunakan Swagger untuk dokumentasi endpoint. Setelah server berjalan, buka:

   ```
   http://localhost:8080/swagger/index.html
   ```

## ✅ Testing API

Untuk menguji endpoint, gunakan **Postman** atau **cURL**.

### 1️⃣ Testing WebSocket Chat

Gunakan WebSocket client seperti **Postman WebSocket** atau **wscat**:

```bash
wscat -c ws://localhost:8080/api/v1/chat/ws
```

## 🤝 Kontribusi

Jika ingin berkontribusi, silakan buat pull request atau buka issue baru di [repository ini](https://github.com/habbazettt/jobseek-api.git).

## 📄 Lisensi

MIT License - Silakan gunakan dan kontribusi! 🎉
