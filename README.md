# Go E-Commerce API

**Backend ini adalah API sederhana untuk aplikasi e-commerce**, dibangun menggunakan **Golang** dengan framework **Gin** dan **GORM** untuk ORM (Object Relational Mapping).  
Proyek ini memiliki fitur **CRUD Produk** dan **autentikasi JWT** untuk mengamankan endpoint tertentu.

---

## 📌 Fitur Utama

### 1. **Autentikasi JWT**
- Login dan register user.
- Token JWT digunakan untuk mengakses endpoint yang dilindungi.

### 2. **Manajemen Produk**
- Membuat produk baru.
- Mengambil daftar produk (dengan kategori dan gambar).
- Mengambil detail produk berdasarkan ID.
- Memperbarui data produk.
- Menghapus produk.

### 3. **Manajemen Kategori Produk**
- Membuat, membaca, memperbarui, dan menghapus kategori.
- Mengatur status aktif/non-aktif kategori.

### 4. **Upload Gambar Produk**
- Mendukung penyimpanan path gambar untuk produk.
- Upload file gambar ke folder `uploads/products`.

---

## 🛠 Teknologi yang Digunakan
- **Bahasa Pemrograman:** Go (Golang)
- **Framework:** Gin
- **ORM:** GORM
- **Database:** MySQL
- **Autentikasi:** JWT (JSON Web Token)
- **File Upload:** Multipart Form Data

---

## 📂 Struktur Endpoint

### **Autentikasi**
| Method | Endpoint     | Deskripsi |
|--------|-------------|-----------|
| POST   | `/register` | Registrasi pengguna baru |
| POST   | `/login`    | Login dan mendapatkan JWT token |

### **Produk**
| Method | Endpoint             | Deskripsi |
|--------|----------------------|-----------|
| GET    | `/products`          | Mendapatkan semua produk |
| GET    | `/products/:id`      | Mendapatkan detail produk |
| POST   | `/products`          | Menambahkan produk baru *(butuh JWT)* |
| PUT    | `/products/:id`      | Memperbarui produk *(butuh JWT)* |
| DELETE | `/products/:id`      | Menghapus produk *(butuh JWT)* |

### **Kategori Produk**
| Method | Endpoint                     | Deskripsi |
|--------|------------------------------|-----------|
| GET    | `/categories`                | Mendapatkan semua kategori |
| POST   | `/categories`                | Menambahkan kategori |
| PUT    | `/categories/:id`            | Memperbarui kategori |
| DELETE | `/categories/:id`            | Menghapus kategori |
| PATCH  | `/categories/:id/status`     | Mengubah status aktif kategori |

---

## 🚀 Cara Menjalankan Proyek

1. **Clone repository**
   ```bash
   git clone https://github.com/username/go-ecommerce-api.git
   cd go-ecommerce-api
   go run main.go
