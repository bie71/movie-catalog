# 🎬 Movie Catalog

Sistem katalog film sederhana dengan dukungan migrasi database menggunakan [`golang-migrate`](https://github.com/golang-migrate/migrate).

---

## 📁 Struktur Folder

```
.
├── database
│   └── migrations
│       ├── 1_init.up.sql
│       ├── 1_init.down.sql
│       └── ...
└── ...
```

---

## 🧩 Requirement

- [migrate CLI](https://github.com/golang-migrate/migrate/releases)
- PostgreSQL (atau sesuaikan DBMS kamu)

---

## ⚙️ Perintah Migrasi

### 🔼 Menjalankan Migrasi (Up)

```bash
migrate -path database/migrations -database "postgres://username:password@localhost:5432/dbname?sslmode=disable" up
```

### 🔽 Rollback Migrasi (Down)

```bash
migrate -path database/migrations -database "postgres://username:password@localhost:5432/dbname?sslmode=disable" down
```

### ⬆️ Menjalankan Hanya 1 Step Migrasi

```bash
migrate -path database/migrations -database "postgres://username:password@localhost:5432/dbname?sslmode=disable" step 1
```

### ⬇️ Rollback 1 Step

```bash
migrate -path database/migrations -database "postgres://username:password@localhost:5432/dbname?sslmode=disable" step -1
```

---

## 🧼 Reset Database (Danger Zone)

```bash
migrate -path database/migrations -database "postgres://username:password@localhost:5432/dbname?sslmode=disable" drop -f
```

> ⚠️ Hati-hati! Ini akan menghapus semua data dan struktur tabel.

---

## 📝 Membuat File Migrasi Baru

```bash
migrate create -ext sql -dir database/migrations -seq nama_migrasi
```

Contoh:

```bash
migrate create -ext sql -dir database/migrations -seq create_movies_table
```

Akan menghasilkan:

```
database/migrations
├── 000001_create_movies_table.up.sql
└── 000001_create_movies_table.down.sql
```

---

## 📦 License

MIT — enjoy and use responsibly 🎬
