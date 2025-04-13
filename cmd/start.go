package cmd

import (
	"movie-catalog/internal/app"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{ // startCmd gunanya untuk mendefinisikan perintah baru dari cobra, cobra adalah package untuk menghandle perintah atau argument dari aplikasi go
	Use:   "start", // nama perintah yang akan digunakan untuk menjalankan perintah pada CLI
	Short: "start", // deskripsi singkat tentang perintah yang akan digunakan
	Long:  `start`, // deskripsi panjang tentang perintah yang akan digunakan
	Run: func(cmd *cobra.Command, args []string) { // Run berfungsi untuk menjalankan perintah
		// Call the function to run the server
		app.RunServer()

	},
}

func init() { // init di panggil saat aplikasi di jalankan, ini adalah fungsi yang pertama kali di panggil
	rootCmd.AddCommand(startCmd)
}
