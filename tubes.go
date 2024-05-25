package main

import "fmt"

const NMAX int = 2024

type arrUser [NMAX]tabUser

type tabUser struct {
	username, password string
	golongan           int
	panjang            int
}

var jumUser, jumPembeli, M int

type pelanggan struct {
	nama           [NMAX]string
	pembelian      [NMAX]barang
	totalpembelian [NMAX]int
	jumbarang      [NMAX]int
}

type barang struct {
	namaBarang                                               [NMAX]string
	jumlah, hargaJual, hargaBeli                             [NMAX]int
	dataBarang, hargaBeliProduk, hargaJualProduk, hargaTotal int
	namaproduk                                               string
}

type gudang struct {
	totalBarangKeluar, totalBarangMasuk [NMAX]int
}

var T arrUser
var pembeli, cetak pelanggan
var kgudang, produk, temp, bendahara barang

func main() {
	login()
}

func login() {
	var s int
	var exit bool
	var u, p string
	exit = false
	for !exit {
		fmt.Println("======= Akun Login ======")
		fmt.Println("1. Log in")
		fmt.Println("2. Exit")
		fmt.Print("Select menu: ")
		fmt.Scan(&s)
		if s == 1 {
			fmt.Print("Username: ")
			fmt.Scan(&u)
			fmt.Print("Password: ")
			fmt.Scan(&p)
			inputUser(u, p)
		} else if s == 2 {
			exit = true
		} else {
			fmt.Println("Invalid!")
		}
	}
}

func inputUser(username, password string) {
	if username == "Admin" && password == "Admin123" {
		menuAdmin()
	} else {
		if isUser(username, password) {
			menuGolongan(cekGolongan(username, password), username)

		} else {
			fmt.Println("User tidak ditemukan")
		}
	}
}

func menuAdmin() {
	var s int
	var exit bool
	exit = false
	for !exit {
		fmt.Println("======= Menu Admin ======")
		fmt.Println("1. Tambahkan user")
		fmt.Println("2. Cetak username, password dan golongan")
		fmt.Println("3. Hapus user")
		fmt.Println("4. Tips admin")
		fmt.Println("5. Exit")
		fmt.Print("Select menu: ")
		fmt.Scan(&s)
		if s == 1 {
			register()
		} else if s == 2 {
			cetakUser()
		} else if s == 3 {
			hapusUser()
		} else if s == 4 {
			tipsAdmin()
		} else if s == 5 {
			exit = true
		} else {
			fmt.Println("Invalid!")
		}
	}
}

func register() {
	fmt.Print("Jumlah nama yang akan di inputkan: ")
	var jumInput int
	fmt.Scan(&jumInput)
	var i int
	if jumInput > 0 {
		for i < jumInput {
			fmt.Print("Inputkan Username: ")
			fmt.Scan(&T[jumUser].username)
			fmt.Print("Inputkan Password: ")
			fmt.Scan(&T[jumUser].password)
			fmt.Print("Inputkan Golongan: ")
			fmt.Scan(&T[jumUser].golongan)
			if isUsernameExists(i, T[jumUser].username) {

				fmt.Println("======================================================================")
				fmt.Println("       Usename sudah terdaftar. Silahkan daftarkan username lain      ")
				fmt.Println("======================================================================")
				i = i
				jumUser = jumUser
			} else {
				i++
				jumUser++
			}
		}
	}
}

func cetakUser() {
	var i int
	i = 0
	for i < jumUser {
		fmt.Println(T[i].username, T[i].password, T[i].golongan)
		i++
	}
}

func hapusUser() {
	var username string
	fmt.Print("User yang akan dihapuskan: ")
	fmt.Scan(&username)
	for i := 0; i < jumUser; i++ {
		if T[i].username == username {
			for i < jumUser {
				T[i].username = T[i+1].username
				T[i].password = T[i+1].password
				T[i].golongan = T[i+1].golongan
				i++
			}
			jumUser = jumUser - 1
			fmt.Println("User ", username, "telah dihapuskan")
		}
	}
}

func tipsAdmin() {
	fmt.Println("Tips untuk admin")
	fmt.Println("1. Pilih 1 ketika anda ingin menginputkan user")
	fmt.Println("2. Pilih 2 ketika anda ingin melihat user yang aktif")
	fmt.Println("3. Pilih 3 ketika anda ingin menonaktifkan user")
	fmt.Println("4. Pilih 5 ketika anda ingin keluar dari menu admin")
	fmt.Println("5. Golongan user terbagi menjadi 3")
	fmt.Println("6. Golongan user 1 adalah sales yang bertugas menginputkan penjualan")
	fmt.Println("7. Golongan user 2 adalah kepala gudang yang bertugas menginputkan barang yang keluar gudang")
	fmt.Println("8. Golongan user 3 adalah bendahara yang bertugas menginputkan data dan harga barang")
}

func isUsernameExists(idx int, username string) bool {
	for i := 0; i < jumUser; i++ {
		if T[i].username == username && i != idx {
			return true
		}
	}
	return false
}

func cekGolongan(username, password string) int {
	for i := 0; i < jumUser; i++ {
		if T[i].username == username && T[i].password == password {
			return T[i].golongan
		}
	}
	return -1
}

func isUser(username, password string) bool {
	var i int
	var ketemu bool
	ketemu = false
	for i < jumUser && !ketemu {
		ketemu = T[i].username == username && T[i].password == password
		i++
	}
	return ketemu
}
func menuGolongan(gol int, username string) {
	if gol == 1 {
		menuSales(username)
	} else if gol == 2 {
		menuGudang(username)
	} else if gol == 3 {
		menuBendahara(username)
	}
}

func menuSales(username string) {
	fmt.Println("Selamat Datang", username)
	var s int
	var exit bool
	exit = false
	for !exit {
		fmt.Println("======= Menu Sales ======")
		fmt.Println("1. Tambahkan Penjualan")
		fmt.Println("2. Cetak Struck Penjualan")
		fmt.Println("3. Exit")
		fmt.Print("Select menu: ")
		fmt.Scan(&s)
		if s == 1 {
			penjualan()
		} else if s == 2 {
			cetakStruck()
		} else if s == 3 {
			exit = true
		} else {
			fmt.Println("Invalid!")
		}
	}
}

func penjualan() {
	var i, j, m, idx, juminput int
	var stop bool
	var namapembeli string
	fmt.Print("Jumlah banyak pembeli yang akan di inputkan: ")
	fmt.Scan(&m)
	for i < m {
		fmt.Print("Inputkan nama pembeli: ")
		fmt.Scan(&namapembeli)
		pembeli.nama[jumPembeli] = namapembeli
		idx = cekIdx(namapembeli)
		fmt.Print("Jumlah banyak barang yang akan dibeli: ")
		fmt.Scan(&juminput)
		stop = false
		j = 0
		if juminput > 0 {
			for !stop {
				fmt.Println("Inputkan nama dan jumlah barang")
				fmt.Scan(&pembeli.pembelian[idx].namaBarang[pembeli.jumbarang[idx]], &pembeli.pembelian[idx].jumlah[pembeli.jumbarang[idx]])
				if j+1 >= juminput {
					gantiAtauTambahBarangPembeli(&stop, idx, &j)
				}
				j++
				pembeli.jumbarang[idx]++
			}
			i++
			jumPembeli = jumPembeli + 1
			M++
		}
	}
}

func gantiAtauTambahBarangPembeli(stop *bool, i int, j *int) {
	var simpan, pilihan int
	fmt.Println("Apakah ada tambahan atau ada yang mau diganti?")
	fmt.Println("1. Ya")
	fmt.Println("2. Tidak")
	fmt.Print("Select: ")
	fmt.Scan(&pilihan)
	if pilihan == 2 {
		*stop = true
	} else if pilihan == 1 {
		fmt.Println("Ganti atau tambah barang?")
		fmt.Println("1. Tambahan")
		fmt.Println("2. Ganti Produk/jumlah")
		fmt.Print("Select: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Println("Ketika menginputkan jumlah tambahan, tambahan tidak boleh kurang dari 1")
			fmt.Print("Jumlah banyak barang yang akan di tambahkan: ")
			fmt.Scan(&simpan)
			M = M + simpan
			*j = *j - 1
		} else if pilihan == 2 {
			gantiArrayBarang(stop, i, j)
		} else {
			fmt.Println("Invalid!")
			gantiAtauTambahBarangPembeli(stop, i, j)
		}
	} else {
		fmt.Println("Invalid!")
		gantiAtauTambahBarangPembeli(stop, i, j)
	}
}

func cetakStruck() {
	var namapembeli string
	var idx, j int
	fmt.Print("Nama pembeli yang akan di cetak: ")
	fmt.Scan(&namapembeli)
	idx = cekIdx(namapembeli)
	fmt.Println(namapembeli)
	fmt.Println("Nama_barang Jumlah_barang Harga_satuan Total_harga")
	for j < pembeli.jumbarang[idx] {
		pembeli.pembelian[idx].hargaJual[j] = cekHarga(pembeli.pembelian[idx].namaBarang[j])
		fmt.Println(pembeli.pembelian[idx].namaBarang[j], pembeli.pembelian[idx].jumlah[j], pembeli.pembelian[idx].hargaJual[j], pembeli.pembelian[idx].hargaJual[j]*pembeli.pembelian[idx].jumlah[j])
		pembeli.pembelian[idx].hargaTotal = pembeli.pembelian[idx].hargaTotal + pembeli.pembelian[idx].hargaJual[j]*pembeli.pembelian[idx].jumlah[j]
		j++
	}
	fmt.Println("Total: ", pembeli.pembelian[idx].hargaTotal)
	fmt.Println("==================================================")
}

func cekIdx(namapembeli string) int {
	var i int
	for i <= jumPembeli {
		if pembeli.nama[i] == namapembeli {
			return i
		}
		i++
	}
	return -1
}

func cekIdxBarang(namaBarang string, i int) int {
	var j int
	for j <= jumPembeli {
		if pembeli.pembelian[i].namaBarang[j] == namaBarang {
			return j
		}
		j++
	}
	return -1
}

func cekHarga(namaBarang string) int {
	var i int
	for i < produk.dataBarang {
		if namaBarang == produk.namaBarang[i] {
			return produk.hargaJual[i]
		}
		i++
	}
	return -1
}

func gantiArrayBarang(stop *bool, i int, j *int) {
	var pilihan int
	fmt.Println("Hapus atau ganti produk")
	fmt.Println("1. Hapus")
	fmt.Println("2. Ganti")
	fmt.Print("Select: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		hapusArryaBarang(i)
	} else if pilihan == 2 {
		gantiArrayBarangjumlah(i)
	} else {
		fmt.Println("Invalid!")
		gantiArrayBarang(stop, i, j)
	}
	gantiAtauTambahBarangPembeli(stop, i, j)
}

func hapusArryaBarang(i int) {
	var namaBarang string
	var idx, k int
	fmt.Print("Nama barang yang akan di hapus :")
	fmt.Scan(&namaBarang)
	idx = cekIdxBarang(namaBarang, i)
	for k < pembeli.jumbarang[i] {
		pembeli.pembelian[i].namaBarang[idx] = pembeli.pembelian[i].namaBarang[idx+1]
		pembeli.pembelian[i].jumlah[idx] = pembeli.pembelian[i].jumlah[idx+1]
		k++
		idx++
	}
	M = M - 1
	pembeli.jumbarang[i] = pembeli.jumbarang[i] - 1
	fmt.Println("Barang ", namaBarang, "telah di hapuskan")
	fmt.Println("==================================================")
}

func gantiArrayBarangjumlah(i int) {
	var pilihan int
	fmt.Println("Apa yang akan di ganti: ")
	fmt.Println("1. Nama Barang")
	fmt.Println("2. Jumlah Barang")
	fmt.Print("Select: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		gantinamaBarang(i)
	} else if pilihan == 2 {
		gantijumlahbarang(i)
	} else {
		fmt.Println("Invalid!")
		gantiArrayBarangjumlah(i)
	}
}

func gantinamaBarang(i int) {
	var namaBarang string
	var idx int
	fmt.Print("Nama barang yang akan diganti :")
	fmt.Scan(&namaBarang)
	idx = cekIdxBarang(namaBarang, i)
	fmt.Print("Nama barang setelah diganti: ")
	fmt.Scan(&namaBarang)
	pembeli.pembelian[i].namaBarang[idx] = namaBarang
	fmt.Println("Nama barang telah diganti")
	fmt.Println("==================================================")
}

func gantijumlahbarang(i int) {
	var namaBarang string
	var jumlahbarang, idx int
	fmt.Print("Nama barang yang akan diganti :")
	fmt.Scan(&namaBarang)
	idx = cekIdxBarang(namaBarang, i)
	fmt.Print("Masukan jumlah barang setelah diganti: ")
	fmt.Scan(&jumlahbarang)
	pembeli.pembelian[i].jumlah[idx] = jumlahbarang
	fmt.Println("Jumlah barang ", namaBarang, " telah diganti")
	fmt.Println("==================================================")
}

func menuGudang(username string) {
	fmt.Println("Selamat Datang", username)
	var s, tk, tm, bp int
	var exit bool
	exit = false
	for !exit {
		fmt.Println("======= Menu Pengawas Gudang ======")
		fmt.Println("1. Tambahkan Barang Masuk Gudang")
		fmt.Println("2. Tambahkan Barang Keluar Gudang")
		fmt.Println("3. Cetak Stok Gudang")
		fmt.Println("4. Exit")
		fmt.Print("Select menu: ")
		fmt.Scan(&s)
		if s == 1 {
			masukBarang(&bp, &tm)
		} else if s == 2 {
			keluarBarang(&bp, &tk)
		} else if s == 3 {
			cetakStokGudang()
		} else if s == 4 {
			exit = true
		} else {
			fmt.Println("Invalid!")
		}
	}
}

func masukBarang(totalProduk, totalBarangMasuk *int) {
	var banyakproduk, i, j int
	var found bool
	fmt.Println("Masukan banyak produk yang ingin anda masukan ke dalam gudang:")
	fmt.Scan(&banyakproduk)
	i = 0
	j = 0
	for i < banyakproduk {
		fmt.Println("Masukkan nama barang yang ingin anda masukan ke dalam gudang:")
		var namaBarang string
		fmt.Scan(&namaBarang)
		fmt.Println("Masukkan jumlah barang yang ingin anda masukan ke dalam gudang:")
		var jumlahBarang int
		fmt.Scan(&jumlahBarang)
		found = false
		for j < *totalProduk {
			if kgudang.namaBarang[j] == namaBarang {
				kgudang.jumlah[j] += jumlahBarang
				found = true
			}
			j++
		}
		if !found {
			kgudang.namaBarang[*totalProduk] = namaBarang
			kgudang.jumlah[*totalProduk] = jumlahBarang
			*totalProduk++
		}
		*totalBarangMasuk += jumlahBarang
		fmt.Println("Nama:", namaBarang)
		fmt.Println("Jumlah:", jumlahBarang)
		i++
	}
	fmt.Println("Total produk yang masuk ke dalam gudang:", *totalProduk)
	fmt.Println("Total barang masuk ke dalam gudang:", *totalBarangMasuk)
	fmt.Println("==================================================")
}

func keluarBarang(totalProduk, totalBarangKeluar *int) {
	var banyakproduk, jumlahKeluar, i int
	var namaBarang string
	fmt.Println("Masukkan banyak produk yang akan keluar dari gudang:")
	fmt.Scan(&banyakproduk)
	i = 0
	for i < banyakproduk {
		fmt.Println("Masukkan nama barang yang akan dikeluarkan dari gudang:")
		fmt.Scan(&namaBarang)
		fmt.Println("Masukkan jumlah barang yang ingin anda keluarkan dari gudang:")
		fmt.Scan(&jumlahKeluar)
		var jumlahKeluarTotal, j int
		var found bool
		jumlahKeluarTotal = jumlahKeluar
		j = 0
		found = false
		for j < *totalProduk {
			if kgudang.namaBarang[j] == namaBarang {
				found = true
				if kgudang.jumlah[j] >= jumlahKeluarTotal {
					kgudang.jumlah[j] -= jumlahKeluarTotal
					*totalBarangKeluar += jumlahKeluarTotal
					fmt.Printf("Berhasil mengeluarkan %d %s dari gudang\n", jumlahKeluarTotal, namaBarang)
					jumlahKeluarTotal = 0
					j = *totalProduk
				} else {
					jumlahKeluarTotal -= kgudang.jumlah[j]
					*totalBarangKeluar += kgudang.jumlah[j]
					fmt.Printf("Mengeluarkan %d %s (semua stok yang tersedia) dari gudang\n", kgudang.jumlah[j], namaBarang)
					kgudang.jumlah[j] = 0
				}
			}
			j++
		}
		if !found {
			fmt.Printf("Barang %s tidak ditemukan di gudang\n", namaBarang)
		} else if jumlahKeluarTotal > 0 {
			fmt.Printf("Tidak cukup stok untuk %s. Sisa yang tidak bisa dikeluarkan: %d\n", namaBarang, jumlahKeluarTotal)
		}
		i++
	}
	fmt.Println("Total barang yang keluar dari gudang:", *totalBarangKeluar)
	fmt.Println("==================================================")
}

func cetakStokGudang() {
	fmt.Println("======= Stok Barang di Gudang =======")
	var i int
	i = 0
	for i < len(kgudang.namaBarang) {
		if kgudang.namaBarang[i] != "" {
			fmt.Printf("Nama Barang: %s, berjumlah: %d\n", kgudang.namaBarang[i], kgudang.jumlah[i])
		}
		i++
	}
	fmt.Println("=====================================")
}

func menuBendahara(username string) {

	fmt.Println("Selamat Datang", username)
	var s int
	var exit bool
	exit = false
	for !exit {
		fmt.Println("========== Menu Bendahara ===========")
		fmt.Println("1. input barang")
		fmt.Println("2. Total Pendapatan")
		fmt.Println("3. Cetak Pembelian")
		fmt.Println("4. Exit")
		fmt.Print("Select menu: ")
		fmt.Scan(&s)
		if s == 1 {
			inputbarang()
		} else if s == 2 {
			cetaktotalpendapatan()
		} else if s == 3 {
			cetakPembelian()
		} else if s == 4 {
			exit = true
		} else {
			fmt.Println("Invalid!")
		}
	}
}

func inputbarang() {

	var s int
	var exit bool
	exit = false
	for !exit {
		fmt.Println("======= Menu Input Barang ======")
		fmt.Println("1. Tambahkan barang")
		fmt.Println("2. Rubah harga/nama barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Cetak Daftar barang")
		fmt.Println("5. Exit")
		fmt.Print("Select menu: ")
		fmt.Scan(&s)
		if s == 1 {
			tambahBarang()
		} else if s == 2 {
			rubahHargaBarang()
		} else if s == 3 {
			HapusBarang()
		} else if s == 4 {
			cetakDaftarBarang()
		} else if s == 5 {
			exit = true
		} else {
			fmt.Println("Invalid!")
		}
	}
}

func tambahBarang() {
	var i, banyakproduk int
	fmt.Println("Masukkan berapa banyak produk yang akan di inputkan:")
	fmt.Scan(&banyakproduk)
	i = 0
	if banyakproduk > 0 {
		for i < banyakproduk {
			fmt.Print("Masukkan nama barang: ")
			fmt.Scan(&produk.namaBarang[produk.dataBarang])
			fmt.Print("Masukkan harga beli barang satuan: ")
			fmt.Scan(&produk.hargaBeli[produk.dataBarang])
			fmt.Print("Masukkan harga jual barang satuan: ")
			fmt.Scan(&produk.hargaJual[produk.dataBarang])
			produk.dataBarang++
			i++
		}
	}
}

func rubahHargaBarang() {
	var pilihan int
	fmt.Println("Apa yang akan diganti: ")
	fmt.Println("1. Nama Barang")
	fmt.Println("2. Jumlah Barang")
	fmt.Print("Select: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		gantinamaBarangBendahara()
	} else if pilihan == 2 {
		gantiHargaBarangBendahara()
	} else {
		fmt.Println("Invalid!")
		rubahHargaBarang()
	}
}

func gantiHargaBarangBendahara() {
	var pilihan int
	fmt.Println("Apa yang akan diganti: ")
	fmt.Println("1. Harga jual")
	fmt.Println("2. Harga beli")
	fmt.Print("Select: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		gantihargaJual()
	} else if pilihan == 2 {
		gantihargaBeli()
	} else {
		fmt.Println("Invalid!")
		rubahHargaBarang()
	}
}

func gantinamaBarangBendahara() {
	var namaBarang string
	var idx int
	fmt.Print("Nama barang yang akan diganti :")
	fmt.Scan(&namaBarang)
	idx = cekIdxBarangBendahara(namaBarang)
	fmt.Print("Nama barang setelah diganti: ")
	fmt.Scan(&namaBarang)
	produk.namaBarang[idx] = namaBarang
	fmt.Println("Nama barang telah diganti")
	fmt.Println("==================================================")
}

func gantihargaBeli() {
	var namaBarang string
	var idx, hargaBeli int
	fmt.Print("Nama barang yang akan diganti :")
	fmt.Scan(&namaBarang)
	idx = cekIdxBarangBendahara(namaBarang)
	fmt.Print("Harga beli barang setelah diganti: ")
	fmt.Scan(&hargaBeli)
	produk.hargaJual[idx] = hargaBeli
	fmt.Println("Harga beli dari produk ", namaBarang, "telah diganti menjadi ", hargaBeli)
	fmt.Println("==================================================")
}

func gantihargaJual() {
	var namaBarang string
	var idx, hargaJual int
	fmt.Print("Nama barang yang akan diganti :")
	fmt.Scan(&namaBarang)
	idx = cekIdxBarangBendahara(namaBarang)
	fmt.Print("Harga barang setelah diganti: ")
	fmt.Scan(&hargaJual)
	produk.hargaJual[idx] = hargaJual
	fmt.Println("Harga jual dari produk ", namaBarang, "telah diganti menjadi ", hargaJual)
	fmt.Println("==================================================")
}

func cekIdxBarangBendahara(namaBarang string) int {
	var i int
	for i < produk.dataBarang {
		if produk.namaBarang[i] == namaBarang {
			return i
		}
		i++
	}
	return -1
}

func HapusBarang() {
	var namaBarang string
	var i int
	fmt.Print("Nama barang yang akan dihapus :")
	fmt.Scan(&namaBarang)
	i = cekIdxBarangBendahara(namaBarang)
	for i < produk.dataBarang {
		produk.namaBarang[i] = produk.namaBarang[i+1]
		produk.hargaBeli[i] = produk.hargaJual[i+1]
		produk.hargaJual[i] = produk.hargaJual[i+1]
		i++
	}
	produk.dataBarang = produk.dataBarang - 1
	fmt.Println("Barang", namaBarang, "telah dihapuskan")
	fmt.Println("==================================================")
}

func cetakDaftarBarang() {
	var s int
	var exit bool
	exit = false
	for !exit {
		fmt.Println("======= Menu cetak daftar barang ======")
		fmt.Println("1. Cetak daftar sesuai harga jual")
		fmt.Println("2. Cetak daftar sesuai harga beli")
		fmt.Println("3. Cetak biasa")
		fmt.Println("4. Exit")
		fmt.Print("Select menu: ")
		fmt.Scan(&s)
		if s == 1 {
			cetakBendaharaHargaJual()
		} else if s == 2 {
			cetakBendaharahargaBeli()
		} else if s == 3 {
			cetakProduk()
		} else if s == 4 {
			exit = true
		} else {
			fmt.Println("Invalid!")
		}
	}
}

func cetakBendaharaHargaJual() {
	var s int
	var exit bool
	exit = false
	for !exit {
		fmt.Println("======= Cetak daftar barang seuai harga jual ======")
		fmt.Println("1. Harga jual terurut membesar")
		fmt.Println("2. Harga jual terurut mengecil")
		fmt.Println("3. Exit")
		fmt.Println("masuk")
		fmt.Print("Select menu: ")
		fmt.Scan(&s)
		if s == 1 {
			ascHargaJual()
			cetakProduk()
		} else if s == 2 {
			desHargaJual()
			cetakProduk()
		} else if s == 3 {
			exit = true
		} else {
			fmt.Println("Invalid!")
		}
	}
}

func cetakBendaharahargaBeli() {
	var s int
	var exit bool
	exit = false
	for !exit {
		fmt.Println("======= Cetak daftar barang seuai harga jual ======")
		fmt.Println("1. Harga jual terurut membesar")
		fmt.Println("2. Harga jual terurut mengecil")
		fmt.Println("3. Exit")
		fmt.Println("masuk")
		fmt.Print("Select menu: ")
		fmt.Scan(&s)
		if s == 1 {
			ascHargaBeli()
			cetakProduk()
		} else if s == 2 {
			desHargaBeli()
			cetakProduk()
		} else if s == 3 {
			exit = true
		} else {
			fmt.Println("Invalid!")
		}
	}
}

func cetakProduk() {
	var i int
	fmt.Println("No. Nama_Produk Harga_Beli Harga_Jual")
	for i < produk.dataBarang {
		fmt.Println(i+1, produk.namaBarang[i], produk.hargaBeli[i], produk.hargaJual[i])
		i++
	}
	fmt.Println("==================================================")
}

func ascHargaBeli() {
	var i, j, idx int
	i = 1
	for i <= produk.dataBarang {
		idx = i - 1
		j = i
		for j < produk.dataBarang {
			if produk.hargaBeli[idx] > produk.hargaBeli[j] {
				idx = j
			}
			j = j + 1
		}
		temp.namaproduk = produk.namaBarang[idx]
		temp.hargaBeliProduk = produk.hargaBeli[idx]
		temp.hargaJualProduk = produk.hargaJual[idx]
		produk.namaBarang[idx] = produk.namaBarang[i-1]
		produk.hargaBeli[idx] = produk.hargaBeli[i-1]
		produk.hargaJual[idx] = produk.hargaJual[i-1]
		produk.namaBarang[i-1] = temp.namaproduk
		produk.hargaBeli[i-1] = temp.hargaBeliProduk
		produk.hargaJual[i-1] = temp.hargaJualProduk
		i = i + 1
	}
}

func ascHargaJual() {
	var i, j, idx int
	i = 1
	for i <= produk.dataBarang {
		idx = i - 1
		j = i
		for j < produk.dataBarang {
			if produk.hargaJual[idx] > produk.hargaJual[j] {
				idx = j
			}
			j = j + 1
		}
		temp.namaproduk = produk.namaBarang[idx]
		temp.hargaBeliProduk = produk.hargaBeli[idx]
		temp.hargaJualProduk = produk.hargaJual[idx]
		produk.namaBarang[idx] = produk.namaBarang[i-1]
		produk.hargaBeli[idx] = produk.hargaBeli[i-1]
		produk.hargaJual[idx] = produk.hargaJual[i-1]
		produk.namaBarang[i-1] = temp.namaproduk
		produk.hargaBeli[i-1] = temp.hargaBeliProduk
		produk.hargaJual[i-1] = temp.hargaJualProduk
		i = i + 1
	}
}

func desHargaBeli() {
	var pass, i int
	pass = 1
	for pass < produk.dataBarang {
		i = pass
		temp.hargaJualProduk = produk.hargaJual[pass]
		temp.hargaBeliProduk = produk.hargaBeli[pass]
		temp.namaproduk = produk.namaBarang[pass]
		for i > 0 && temp.hargaBeliProduk > produk.hargaBeli[i-1] {
			produk.hargaJual[i] = produk.hargaJual[i-1]
			produk.hargaBeli[i] = produk.hargaBeli[i-1]
			produk.namaBarang[i] = produk.namaBarang[i-1]
			i = i - 1
		}
		produk.hargaJual[i] = temp.hargaJualProduk
		produk.hargaBeli[i] = temp.hargaBeliProduk
		produk.namaBarang[i] = temp.namaproduk
		pass++
	}
}

func desHargaJual() {
	var pass, i int
	pass = 1
	for pass < produk.dataBarang {
		i = pass
		temp.hargaJualProduk = produk.hargaJual[pass]
		temp.hargaBeliProduk = produk.hargaBeli[pass]
		temp.namaproduk = produk.namaBarang[pass]
		for i > 0 && temp.hargaJualProduk > produk.hargaJual[i-1] {
			produk.hargaJual[i] = produk.hargaJual[i-1]
			produk.hargaBeli[i] = produk.hargaBeli[i-1]
			produk.namaBarang[i] = produk.namaBarang[i-1]
			i = i - 1
		}
		produk.hargaJual[i] = temp.hargaJualProduk
		produk.hargaBeli[i] = temp.hargaBeliProduk
		produk.namaBarang[i] = temp.namaproduk
		pass++
	}
}

func cetaktotalpendapatan() {
	var i int
	for i < jumPembeli {
		bendahara.hargaTotal = bendahara.hargaTotal + pembeli.pembelian[i].hargaTotal
		i++
	}
	fmt.Println("Total pendapatan perusahaan:", bendahara.hargaTotal)
	fmt.Println("Dengan rincian: ")
	i = 0
	for i < jumPembeli {
		fmt.Println(pembeli.nama[i], pembeli.pembelian[i].hargaTotal)
		i++
	}
	fmt.Println("==================================================")
}

func cetakPembelian() {
	var s int
	var exit bool
	exit = false
	for !exit {
		fmt.Println("======= Menu cetak pembelian======")
		fmt.Println("1. Cetak pembelian nama terurut")
		fmt.Println("2. Cetak pembelian berdasarkan nama")
		fmt.Println("3. Exit")
		fmt.Println("masuk")
		fmt.Print("Select menu: ")
		fmt.Scan(&s)
		if s == 1 {
			cetakPembelianurut()
		} else if s == 2 {
			cetakPembeliankhusus()
		} else if s == 3 {
			exit = true
		} else {
			fmt.Println("Invalid!")
		}
	}
}

func cetakPembelianurut() {
	var i int
	urutnama()
	fmt.Println("Berikut adalah daftar nama pembeli dan jumlah pembelian")
	fmt.Println("No. Nama Jumlah_Pembelian")
	for i < jumPembeli {
		fmt.Println(i+1, pembeli.nama[i], pembeli.pembelian[i].hargaTotal)
		i++
	}
	fmt.Println("==================================================")
}

func cetakPembeliankhusus() {
	var idx int
	var X string
	urutnama()
	fmt.Print("Tuliskan nama yang akan di cetak: ")
	fmt.Scan(&X)
	idx = binary(X)
	fmt.Println("Nama Jumlah_Pembelian")
	fmt.Println(cetak.nama[idx], cetak.pembelian[idx].hargaTotal)
	fmt.Println("==================================================")
}

func binary(X string) int {
	var bawah, atas, tengah, ketemu int
	bawah = 0
	atas = jumPembeli
	ketemu = -1
	for bawah <= atas && ketemu == -1 {
		tengah = (bawah + atas) / 2
		if X > cetak.nama[tengah] {
			bawah = tengah + 1
		} else if X < cetak.nama[tengah] {
			atas = tengah - 1
		} else {
			ketemu = tengah
		}
	}
	return ketemu
}

func urutnama() {
	var i, j, idx int
	i = 1
	isiArray()
	for i <= jumPembeli {
		idx = i - 1
		j = i
		for j < jumPembeli {
			if cetak.nama[idx] > cetak.nama[j] {
				idx = j
			}
			j = j + 1
		}
		temp.namaproduk = cetak.nama[idx]
		temp.hargaBeliProduk = cetak.pembelian[idx].hargaTotal
		cetak.nama[idx] = cetak.nama[i-1]
		cetak.pembelian[idx].hargaTotal = cetak.pembelian[i-1].hargaTotal
		cetak.nama[i-1] = temp.namaproduk
		cetak.pembelian[i-1].hargaTotal = temp.hargaBeliProduk
		i = i + 1
	}
}

func isiArray() {
	var i int
	for i < jumPembeli {
		cetak.nama[i] = pembeli.nama[i]
		cetak.pembelian[i].hargaTotal = pembeli.pembelian[i].hargaTotal
		i++
	}
}