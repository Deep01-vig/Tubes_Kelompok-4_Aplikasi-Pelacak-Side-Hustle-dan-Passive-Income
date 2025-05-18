package main

import "fmt"

const maxPekerjaan = 100

type pekerjaan struct {
	nama        string
	penghasilan int
	pengeluaran int
}

type pengguna struct {
	namaPengguna string
	umur         int
	jumlah       int
	pekerjaan    [maxPekerjaan]pekerjaan
}

var dataPekerjaan [maxPekerjaan]pekerjaan

func main() {
	var user pengguna
	var jumlahPekerjaan int

	fmt.Println("Selamat Datang!")
	fmt.Print("Siapa nama anda? ")
	fmt.Scan(&user.namaPengguna)
	fmt.Print("Berapa usia anda? ")
	fmt.Scan(&user.umur)
	fmt.Print("Berapa banyak pekerjaan anda? ")
	fmt.Scan(&jumlahPekerjaan)

	if jumlahPekerjaan > maxPekerjaan {
		jumlahPekerjaan = maxPekerjaan
	}

	user.jumlah = jumlahPekerjaan

	for i := 0; i < jumlahPekerjaan; i++ {
		fmt.Printf("Nama pekerjaan %d: ", i+1)
		fmt.Scan(&user.pekerjaan[i].nama)
		fmt.Print("Penghasilan: ")
		fmt.Scan(&user.pekerjaan[i].penghasilan)
		fmt.Print("Pengeluaran: ")
		fmt.Scan(&user.pekerjaan[i].pengeluaran)
	}

	fmt.Println("Data sebelum sorting:")
	cetak(user)

	fmt.Println("Data setelah selection sort (penghasilan):")
	selectionSort(&user)
	cetak(user)

	fmt.Println("Data setelah insertion sort (pengeluaran):")
	insertionSort(&user)
	cetak(user)

	hitungLaba(&user)
}

func cetak(user pengguna) {
	for i := 0; i < user.jumlah; i++ {
		fmt.Printf("[%d] %s | Penghasilan: Rp%d | Pengeluaran: Rp%d\n", i+1,
			user.pekerjaan[i].nama, user.pekerjaan[i].penghasilan, user.pekerjaan[i].pengeluaran)
	}
}

func selectionSort(user *pengguna) {
	var minIdx, i, j int
	for i = 0; i < user.jumlah-1; i++ {
		minIdx = i
		for j = i + 1; j < user.jumlah; j++ {
			if user.pekerjaan[j].penghasilan < user.pekerjaan[minIdx].penghasilan {
				minIdx = j
			}
		}
		user.pekerjaan[i], user.pekerjaan[minIdx] = user.pekerjaan[minIdx], user.pekerjaan[i]
	}
}

func insertionSort(user *pengguna) {
	var i, j int
	var temp pekerjaan
	for i = 1; i < user.jumlah; i++ {
		temp = user.pekerjaan[i]
		j = i - 1
		for j >= 0 && user.pekerjaan[j].pengeluaran > temp.pengeluaran {
			user.pekerjaan[j+1] = user.pekerjaan[j]
			j = j - 1
		}
		user.pekerjaan[j+1] = temp
	}
}

func hitungLaba(user *pengguna) {
	var totalPenghasilan, totalPengeluaran, laba int
	var i int

	for i = 0; i < user.jumlah; i++ {
		totalPenghasilan += user.pekerjaan[i].penghasilan
		totalPengeluaran += user.pekerjaan[i].pengeluaran
	}

	laba = totalPenghasilan - totalPengeluaran

	fmt.Println("===== RINGKASAN =====")
	fmt.Printf("Total Penghasilan  : Rp %d\n", totalPenghasilan)
	fmt.Printf("Total Pengeluaran  : Rp %d\n", totalPengeluaran)
	fmt.Printf("Laba Bersih        : Rp %d\n", laba)

	fmt.Println("Masukkan nama pekerjaan untuk dicari: ")
	var namaCari string
	fmt.Scan(&namaCari)
	sequentialSearch(*user, namaCari)

	sortByNama(&user.pekerjaan, user.jumlah)
	fmt.Println("Data disorting berdasarkan nama (untuk binary search).")
	binarySearch(*user, namaCari)

	fmt.Println("Masukkan nama pekerjaan untuk diedit (Jika tidak ada, masukkan -): ")
	fmt.Scan(&namaCari)
	editPekerjaan(user, namaCari)

	fmt.Println("Masukkan nama pekerjaan untuk dihapus (Jika tidak ada, masukkan -): ")
	fmt.Scan(&namaCari)
	deletePekerjaan(user, namaCari)

	fmt.Println("Data setelah edit dan delete:")
	cetak(*user)
}

func sequentialSearch(user pengguna, nama string) {
	var i int
	var ditemukan = false
	fmt.Println("===== SEQUENTIAL SEARCH =====")
	for i = 0; i < user.jumlah; i++ {
		if user.pekerjaan[i].nama == nama {
			fmt.Printf("Nama pekerjaan: %s\nPenghasilan: Rp%d\nPengeluaran: Rp%d\n",
				user.pekerjaan[i].nama, user.pekerjaan[i].penghasilan, user.pekerjaan[i].pengeluaran)
			ditemukan = true
			i = user.jumlah
		}
	}
	if !ditemukan {
		fmt.Println("Pekerjaan tidak ditemukan.")
	}
}

func sortByNama(data *[maxPekerjaan]pekerjaan, n int) {
	var i, j int
	for i = 0; i < n-1; i++ {
		for j = 0; j < n-i-1; j++ {
			if data[j].nama > data[j+1].nama {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func binarySearch(user pengguna, keyword string) {
	var kiri, kanan, tengah int
	var ditemukan = false
	kiri = 0
	kanan = user.jumlah - 1

	fmt.Println("===== BINARY SEARCH =====")
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if user.pekerjaan[tengah].nama == keyword {
			fmt.Printf("Nama pekerjaan: %s\nPenghasilan: Rp%d\nPengeluaran: Rp%d\n",
				user.pekerjaan[tengah].nama, user.pekerjaan[tengah].penghasilan, user.pekerjaan[tengah].pengeluaran)
			ditemukan = true
			kiri = kanan + 1 // akhiri loop
		} else {
			if user.pekerjaan[tengah].nama < keyword {
				kiri = tengah + 1
			} else {
				kanan = tengah - 1
			}
		}
	}
	if !ditemukan {
		fmt.Println("Pekerjaan tidak ditemukan (binary search).")
	}
}

func editPekerjaan(user *pengguna, namaCari string) {
	var i int
	for i = 0; i < user.jumlah; i++ {
		if user.pekerjaan[i].nama == namaCari {
			fmt.Println("===== EDIT DATA =====")
			fmt.Printf("Data Lama - Penghasilan: Rp%d, Pengeluaran: Rp%d\n",
				user.pekerjaan[i].penghasilan, user.pekerjaan[i].pengeluaran)
			fmt.Print("Masukkan penghasilan baru: ")
			fmt.Scan(&user.pekerjaan[i].penghasilan)
			fmt.Print("Masukkan pengeluaran baru: ")
			fmt.Scan(&user.pekerjaan[i].pengeluaran)
			fmt.Println("Data berhasil diperbarui.")
			return
		}
	}
	fmt.Println("Pekerjaan tidak ditemukan. Tidak ada yang diubah.")
}

func deletePekerjaan(user *pengguna, namaCari string) {
	var i int
	var found = false
	for i = 0; i < user.jumlah; i++ {
		if user.pekerjaan[i].nama == namaCari && !found {
			found = true
		}
		if found && i < user.jumlah-1 {
			user.pekerjaan[i] = user.pekerjaan[i+1]
		}
	}
	if found {
		user.jumlah--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Pekerjaan tidak ditemukan. Tidak ada yang dihapus.")
	}
}