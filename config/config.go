package config // berisi file config.go yang berfungsi untuk mengatur konfigurasi aplikasi seperti database, jwt, dan lain-lain dan juga untuk menginitialisasi konfigurasi

//  import library yang dibutuhkan
import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper" // viper adalah library untuk mengelola konfigurasi aplikasi di golang
)

type App struct { // struct App berfungsi untuk mengatur konfigurasi aplikasi yg berisi port, env, jwt secret key dan issuer
	AppPort string `json:"app_port"` // port aplikasi digunakan untuk menjalankan server
	AppEnv  string `json:"app_env"`  // env aplikasi digunakan untuk menentukan mode aplikasi (development, production, dll)

	JwtSecretKey string `json:"jwt_secret_key"` // secret key untuk jwt digunakan untuk mengenkripsi token jwt
	JwtIssuer    string `json:"jwt_issuer"`     // issuer untuk jwt digunakan untuk menentukan issuer(penerbit) dari token jwt
}

type PsqlDB struct { // struct PsqlDB berfungsi untuk mengatur konfigurasi database postgresql
	Host      string `json:"host"`        // host database digunakan untuk menentukan host dari database postgresql
	Port      string `json:"port"`        // port database digunakan untuk menentukan port dari database postgresql
	User      string `json:"user"`        // user database digunakan untuk menentukan user dari database postgresql
	Password  string `json:"password"`    // password database digunakan untuk menentukan password dari database postgresql
	DBName    string `json:"db_name"`     // db_name database digunakan untuk menentukan nama database postgresql
	DBMaxOpen int    `json:"db_max_open"` // db_max_open database digunakan untuk menentukan jumlah maksimal koneksi database postgresql
	DBMaxIdle int    `json:"db_max_idle"` // db_max_idle database digunakan untuk menentukan jumlah maksimal koneksi idle / standby database postgresql
}

type CloudFlareR2 struct {
	Name      string `json:"name"`       // name cloudflare r2 digunakan untuk menentukan nama dari cloudflare r2
	ApiKey    string `json:"api_key"`    // api_key cloudflare r2 digunakan untuk menentukan api key dari cloudflare r2
	ApiSecret string `json:"api_secret"` // api_secret cloudflare r2 digunakan untuk menentukan api secret dari cloudflare r2
	Token     string `json:"token"`      // token cloudflare r2 digunakan untuk menentukan token dari cloudflare r2
	AccountId string `json:"account_id"` // account_id cloudflare r2 digunakan untuk menentukan account id dari cloudflare r2
	PublicUrl string `json:"public_url"` // public_url cloudflare r2 digunakan untuk menentukan public url dari cloudflare r2
}

type Config struct { // struct Config berfungsi untuk mengatur konfigurasi aplikasi yg berisi konfigurasi aplikasi dan database
	App        App
	Psql       PsqlDB
	CloudFlare CloudFlareR2
}

func Newconfig() *Config { // fungsi Newconfig berfungsi mengembalikan konfigurasi aplikasi yg sudah diatur sebelumnya
	return &Config{
		App: App{
			AppPort:      viper.GetString("APP_PORT"), // viper GetString digunakan untuk mengambil nilai dari file .env / variable dll
			AppEnv:       viper.GetString("APP_PORT"),
			JwtSecretKey: viper.GetString("JWT_SECRET_KEY"),
			JwtIssuer:    viper.GetString("JWT_ISSUER"),
		},
		Psql: PsqlDB{
			Host:      viper.GetString("DATABASE_HOST"),
			Port:      viper.GetString("DATABASE_PORT"),
			User:      viper.GetString("DATABASE_USER"),
			Password:  viper.GetString("DATABASE_PASSWORD"),
			DBName:    viper.GetString("DATABASE_NAME"),
			DBMaxOpen: viper.GetInt("DATABASE_MAX_OPEN_CONNECTION"),
			DBMaxIdle: viper.GetInt("DATABASE_MAX_IDLE_CONNECTION"),
		},
		CloudFlare: CloudFlareR2{
			Name:      viper.GetString("CLOUDFLARE_R2_NAME"),
			ApiKey:    viper.GetString("CLOUDFLARE_R2_API_KEY"),
			ApiSecret: viper.GetString("CLOUDFLARE_R2_API_SECRET"),
			Token:     viper.GetString("CLOUDFLARE_R2_TOKEN"),
			AccountId: viper.GetString("CLOUDFLARE_R2_ACCOUNT_ID"),
			PublicUrl: viper.GetString("CLOUDFLARE_R2_PUBLIC_URL"),
		},
	}
}
