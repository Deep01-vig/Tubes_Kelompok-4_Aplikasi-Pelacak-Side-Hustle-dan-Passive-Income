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
	pekerjaan    tabPekerjaan
}

type tabPekerjaan [maxPekerjaan]pekerjaan

func bacaData(user *pengguna) {
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
}

func cetakData(user pengguna) {
	for i := 0; i < user.jumlah; i++ {
		fmt.Printf("[%d] %s | Penghasilan: Rp%d | Pengeluaran: Rp%d\n", i+1, user.pekerjaan[i].nama, user.pekerjaan[i].penghasilan, user.pekerjaan[i].pengeluaran)
	}
}

func totalPenghasilan(user *pengguna) int {
	var total, i int
	for i = 0; i < user.jumlah; i++ {
		total += user.pekerjaan[i].penghasilan
	}
	return total
}

func totalPengeluaran(user *pengguna) int {
	var total, i int
	for i = 0; i < user.jumlah; i++ {
		total += user.pekerjaan[i].pengeluaran
	}
	return total
}

func hitungLaba(user *pengguna) int {
	var penghasilan, pengeluaran, laba int
	
	penghasilan = totalPenghasilan(user)
	pengeluaran = totalPengeluaran(user)
	laba = penghasilan - pengeluaran

	fmt.Println("===== RINGKASAN =====")
	fmt.Printf("Total Penghasilan  : Rp %d\n", penghasilan)
	fmt.Printf("Total Pengeluaran  : Rp %d\n", pengeluaran)
	fmt.Printf("Laba Bersih        : Rp %d\n", laba)

	return laba
}

func saranOptimasi(totalPenghasilan int, totalPengeluaran int) {
	var selisih int
	var persen float64
	
	fmt.Println("\n===== SARAN OPTIMASI =====")

	if totalPenghasilan == 0 && totalPengeluaran == 0 {
		fmt.Println("Data penghasilan dan pengeluaran kosong. Tidak dapat memberikan saran.")
		return
	}

	if totalPenghasilan < totalPengeluaran {
		selisih = totalPengeluaran - totalPenghasilan
		persen = float64(selisih) / float64(totalPengeluaran) * 100
		fmt.Println("Pengeluaran anda lebih besar dari penghasilan.")
		fmt.Printf("Disarankan untuk mengurangi pengeluaran minimal sebesar %.2f%%.\n", persen)
	} else if totalPenghasilan > totalPengeluaran {
		selisih = totalPenghasilan - totalPengeluaran
		persen = float64(selisih) / float64(totalPenghasilan) * 100
		fmt.Println("Keuangan anda bagus! Penghasilan melebihi pengeluaran.")
		fmt.Printf("Anda berhasil menyimpan %.2f%% dari penghasilan Anda.\n", persen)
	} else {
		fmt.Println("Penghasilan dan pengeluaran anda sama.")
		fmt.Println("Disarankan untuk meningkatkan penghasilan atau mengurangi pengeluaran agar memiliki laba bersih.")
	}
}

func cariGajiTertinggi(user pengguna) {
	var maxIdx, i int
	
	if user.jumlah == 0 {
		fmt.Println("Tidak ada data pekerjaan.")
		return
	}

	maxIdx = 0
	for i = 1; i < user.jumlah; i++ {
		if user.pekerjaan[i].penghasilan > user.pekerjaan[maxIdx].penghasilan {
			maxIdx = i
		}
	}

	fmt.Println("\n===== PEKERJAAN DENGAN GAJI TERTINGGI =====")
	fmt.Printf("Nama Pekerjaan : %s\n", user.pekerjaan[maxIdx].nama)
	fmt.Printf("Penghasilan    : Rp%d\n", user.pekerjaan[maxIdx].penghasilan)
	fmt.Printf("Pengeluaran    : Rp%d\n", user.pekerjaan[maxIdx].pengeluaran)
}

func sortByNama(data *tabPekerjaan, n int) {
	var i, j int
	for i = 0; i < n-1; i++ {
		for j = 0; j < n-i-1; j++ {
			if data[j].nama > data[j+1].nama {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func selectionSort(user *pengguna) {
	var minIdx, i, j int
	for i = 0; i < user.jumlah-1; i++ {
		minIdx = i
		for j = i + 1; j < user.jumlah; j++ {
			if user.pekerjaan[j].penghasilan > user.pekerjaan[minIdx].penghasilan {
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
		for j >= 0 && user.pekerjaan[j].pengeluaran < temp.pengeluaran {
			user.pekerjaan[j+1] = user.pekerjaan[j]
			j = j - 1
		}
		user.pekerjaan[j+1] = temp
	}
}

func sequentialSearch(user pengguna, nama string) {
	var i int
	var ditemukan bool
	
	ditemukan = false
	fmt.Println("===== SEQUENTIAL SEARCH =====")
	for i = 0; i < user.jumlah; i++ {
		if user.pekerjaan[i].nama == nama {
			fmt.Printf("Nama pekerjaan: %s\nPenghasilan: Rp%d\nPengeluaran: Rp%d\n",
				user.pekerjaan[i].nama, user.pekerjaan[i].penghasilan, user.pekerjaan[i].pengeluaran)
			ditemukan = true
			break
		}
	}
	if !ditemukan {
		fmt.Println("Pekerjaan tidak ditemukan.")
	}
}

func binarySearch(user pengguna, keyword string) {
	var kiri, kanan, tengah int
	var ditemukan bool
	
	kiri = 0
	kanan = user.jumlah - 1
	ditemukan = false

	fmt.Println("===== BINARY SEARCH =====")
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if user.pekerjaan[tengah].nama == keyword {
			fmt.Printf("Nama pekerjaan: %s\nPenghasilan: Rp%d\nPengeluaran: Rp%d\n",
				user.pekerjaan[tengah].nama, user.pekerjaan[tengah].penghasilan, user.pekerjaan[tengah].pengeluaran)
			ditemukan = true
			break
		} else if user.pekerjaan[tengah].nama < keyword {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
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
			fmt.Printf("Data Lama - Penghasilan: Rp%d, Pengeluaran: Rp%d\n", user.pekerjaan[i].penghasilan, user.pekerjaan[i].pengeluaran)
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
	var idx, i int

	idx = -1
	for i = 0; i < user.jumlah; i++ {
		if user.pekerjaan[i].nama == namaCari {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("Pekerjaan tidak ditemukan. Tidak ada yang dihapus.")
		return
	}
	for i = idx; i < user.jumlah-1; i++ {
		user.pekerjaan[i] = user.pekerjaan[i+1]
	}
	user.jumlah--
	fmt.Println("Data berhasil dihapus.")
}

func main() {
	var user pengguna
	var pilihan int
	var namaCari string

	bacaData(&user)
	fmt.Println()
	fmt.Println("Data Keuangan Anda: ")
	cetakData(user)
	cariGajiTertinggi(user)

	for {
		fmt.Println("\n===== MENU UTAMA =====")
		fmt.Println("1. Hitung Laba dan Saran Optimasi")
		fmt.Println("2. Sortir Berdasarkan Nama")
		fmt.Println("3. Sortir Berdasarkan Penghasilan")
		fmt.Println("4. Sortir Berdasarkan Pengeluaran")
		fmt.Println("5. Pencarian Pekerjaan")
		fmt.Println("6. Edit Pekerjaan")
		fmt.Println("7. Hapus Pekerjaan")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println()
			hitungLaba(&user)
			fmt.Println()
			penghasilan := totalPenghasilan(&user)
			pengeluaran := totalPengeluaran(&user)
			saranOptimasi(penghasilan, pengeluaran)			
		case 2:
			sortByNama(&user.pekerjaan, user.jumlah)
			fmt.Println("Data berhasil disortir berdasarkan nama.")
			cetakData(user)
		case 3:
			selectionSort(&user)
			fmt.Println("Data berhasil disortir berdasarkan penghasilan.")
			cetakData(user)
		case 4:
			insertionSort(&user)
			fmt.Println("Data berhasil disortir berdasarkan pengeluaran.")
			cetakData(user)
		case 5:
			fmt.Print("Masukkan nama pekerjaan yang ingin dicari: ")
			fmt.Scan(&namaCari)
			sequentialSearch(user, namaCari)
			fmt.Println()
			sortByNama(&user.pekerjaan, user.jumlah)
			binarySearch(user, namaCari)
		case 6:
			fmt.Print("Masukkan nama pekerjaan yang ingin diedit: ")
			fmt.Scan(&namaCari)
			editPekerjaan(&user, namaCari)
			cetakData(user)
		case 7:
			fmt.Print("Masukkan nama pekerjaan yang ingin dihapus: ")
			fmt.Scan(&namaCari)
			deletePekerjaan(&user, namaCari)
			cetakData(user)
		case 0:
			fmt.Println("Program selesai. Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih antara 0-6.")
		}
	}
}
