package config

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres" // Correct import for GORM v2 PostgreSQL driver
	"gorm.io/gorm"
)

type Postgres struct { // struct Postgres berfungsi mendefinisikan objek database postgresql yang bertipe gorm.DB dan nanti dikelola oleh GORM
	DB *gorm.DB // ini pointer ke objek gorm.DB yang merupakan objek utama untuk mengelola database postgresql
}

func (cfg Config) ConnectionPostgres() (*Postgres, error) { // fungsi ConnectionPostgres berfungsi untuk membuat koneksi ke database postgresql, setelah dibuka dia mengembalikan objek Postgres dan jika gagal dia mengembalikan error
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Psql.User,
		cfg.Psql.Password,
		cfg.Psql.Host,
		cfg.Psql.Port,
		cfg.Psql.DBName) // dbConnString berfungsi untuk menyimpan string koneksi ke database postgresql yang akan dibangun dengan fmt.Sprintf

	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-1] Failed to connect to database " + cfg.Psql.Host)
		return nil, err
	}

	sqldb, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-2] Failed to get database connection")
		return nil, err
	}

	sqldb.SetMaxOpenConns(cfg.Psql.DBMaxOpen) //  setMaxOpenConns berfungsi untuk mengatur jumlah maksimal koneksi ke database postgresql
	sqldb.SetMaxIdleConns(cfg.Psql.DBMaxIdle) // setMaxIdleConns berfungsi untuk mengatur jumlah maksimal koneksi idle / standby ke database postgresql

	return &Postgres{DB: db}, nil // Correct field name

}
