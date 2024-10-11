package main

import "fmt"

// Fungsi untuk menghitung luas persegi panjang
func hitungLuas(panjang float64, lebar float64) float64 {
	return panjang * lebar
}

// Fungsi untuk menghitung keliling persegi panjang
func hitungKeliling(panjang float64, lebar float64) float64 {
	return 2 * (panjang + lebar)
}

func main() {
	// Mendefinisikan panjang dan lebar
	var panjang, lebar float64

	// Input user
	fmt.Print("Masukkan panjang: ")
	fmt.Scan(&panjang)

	fmt.Print("Masukkan lebar: ")
	fmt.Scan(&lebar)

	// Memanggil fungsi hitungLuas dan menampilkan hasilnya
	luas := hitungLuas(panjang, lebar)
	fmt.Printf("Luas persegi panjang: %.2f\n", luas)

	// Memanggil fungsi hitungKeliling dan menampilkan hasilnya
	keliling := hitungKeliling(panjang, lebar)
	fmt.Printf("Keliling persegi panjang: %.2f\n", keliling)
}