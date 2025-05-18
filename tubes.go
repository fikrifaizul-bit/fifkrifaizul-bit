package main

import "fmt"

const NMAX = 1024

type Topik struct {
	topik string
	isi [NMAX]string
	hashtag [NMAX]string
	konten [NMAX]string
}

type Konten struct {
	topik [NMAX]string
	isi [NMAX]string
	hashtag [NMAX]string
	konten [NMAX]string
	tanggal string
}

type tabKonten struct {
	detailKonten [NMAX]Konten
	n int
}

type tabTopik [NMAX]Topik
type Menu [NMAX]int

func main() {
	var K tabKonten
	var pilihan int
	var isRunning bool = true
	tampilMenu()

	for isRunning {
		fmt.Scan(&pilihan)
		switch pilihan {
			case 1:
				menuBuatKonten(&K)
			case 2:
				tampilData(K)
			case 3:
				fmt.Println("Terima Kasih! Sampai Jumpa Lagi!")
				isRunning = false
			default:
				fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func tampilMenu() {
	fmt.Println("================= SELAMAT DATANG DI =================")
	fmt.Println("====== APLIKASI AI PEMBUAT KONTEN SOSIAL MEDIA ======")
	fmt.Println("1. Buat Konten")
	fmt.Println("2. Daftar Konten")
	fmt.Println("3. Keluar")
	fmt.Println("=====================================================")
	fmt.Print("Masukkan Pilihan: ")
}

func menuBuatKonten(K *tabKonten) {
	var daftarTopik tabTopik
	var caption, hashtag, konten, topik string
	var pilihTopik int

	topik = pilihTopikKonten(&daftarTopik, &pilihTopik)
	konten = rekomendasiIdeKonten(daftarTopik, pilihTopik)
	caption = buatCaptionOtomatis(daftarTopik, pilihTopik)
	hashtag = rekomendasiHashtag(daftarTopik, pilihTopik)

	tambahData(K, topik, konten, caption, hashtag)

	tampilMenu()
}

func pilihTopikKonten(daftarTopik *tabTopik, pilihTopik *int) string {
	fmt.Print("Chat Bot: ")
	fmt.Println("Halo! Saya akan membantu kamu membuat caption otomatis")
	fmt.Print("Chat Bot: ")
	fmt.Println("Topik tentang apa yang ingin kamu buat?")
	isiTopik(daftarTopik)
	cetakTopik(*daftarTopik, 6)

	fmt.Print("User: ")
	fmt.Scan(pilihTopik)
	fmt.Println("\n")
	fmt.Print("Chat Bot: ")
	fmt.Println("Pilihan Bagus, Mari saya bantu untuk mencari ide konten yang cocok untuk kamu!")

	return daftarTopik[*pilihTopik - 1].topik
}

func buatCaptionOtomatis(daftarTopik tabTopik, pilihTopik int) string {
	var pilihOpsi int

	fmt.Print("Chat Bot: ")
	fmt.Printf("Berikut Caption yang saya sediakan untuk Topik %s:\n", daftarTopik[pilihTopik - 1].topik)
	cetakIsiTopik(daftarTopik, 3, pilihTopik - 1)

	fmt.Print("User: ")
	fmt.Scan(&pilihOpsi)
	fmt.Println("\n")
	if pilihOpsi > 3 || pilihOpsi < 0 {
		fmt.Println("Maaf saya belum menambah data.")
		buatCaptionOtomatis(daftarTopik, pilihTopik)
	}
	fmt.Print("Chat Bot: ")
	fmt.Println("Pilihan Bagus, Mari saya bantu untuk membuat hashtag yang cocok untuk kamu!")

	return daftarTopik[pilihTopik - 1].isi[pilihOpsi - 1]
}

func rekomendasiHashtag(daftarTopik tabTopik, pilihTopik int) string {
	var pilihOpsi int

	fmt.Print("Chat Bot: ")
	fmt.Printf("Berikut Hashtag yang saya sediakan untuk Topik %s:\n", daftarTopik[pilihTopik - 1].topik)
	cetakHashtag(daftarTopik, 3, pilihTopik - 1)

	fmt.Print("User: ")
	fmt.Scan(&pilihOpsi)
	fmt.Println("\n")
	fmt.Print("Chat Bot: ")
	fmt.Println("Keren! Kamu sudah membuat konten yang bagus.")

	return daftarTopik[pilihTopik - 1].hashtag[pilihOpsi - 1]
}

func rekomendasiIdeKonten(daftarTopik tabTopik, pilihTopik int) string {
	var pilihOpsi int

	fmt.Print("Chat Bot: ")
	fmt.Printf("Berikut Ide Konten yang saya sediakan untuk Topik %s:\n", daftarTopik[pilihTopik - 1].topik)
	cetakIdeKonten(daftarTopik, 3, pilihTopik - 1)

	fmt.Print("User: ")
	fmt.Scan(&pilihOpsi)
	fmt.Println("\n")
	fmt.Print("Chat Bot: ")
	fmt.Println("Cocok sekali untuk kamu, saya akan bantu merekomendasikan caption yang bagus.")

	return daftarTopik[pilihTopik - 1].konten[pilihOpsi - 1]
}

func tambahData(K *tabKonten, topik, konten, caption, hashtag string) {
	var n int = K.n

	K.detailKonten[n].topik[0] = topik
	K.detailKonten[n].konten[0] = konten
	K.detailKonten[n].isi[0] = caption
	K.detailKonten[n].hashtag[0] = hashtag

	K.n++
}

func tampilData(K  tabKonten) {
	var isRunning bool = true
	var user string
	for isRunning {
		if K.n == 0 {
			fmt.Println("Belum ada data konten yang tersimpan.")
			return
		}
		for i := 0; i < K.n; i++ {
			fmt.Printf("\n%d. ", i + 1)
			fmt.Printf("Topik: %s\n", K.detailKonten[i].topik[0])
			fmt.Printf("Konten: %s\n", K.detailKonten[i].konten[0])
			fmt.Printf("Caption: %s\n", K.detailKonten[i].isi[0])
			fmt.Printf("Hashtag: %s\n", K.detailKonten[i].hashtag[0])
		}
		fmt.Print("\n[Ketik 'exit' untuk kembali ke Menu Utama]: ")
		fmt.Scan(&user)
		if user == "exit" {
			isRunning = false
			tampilMenu()
		}
	}
}

func isiTopik(daftar *tabTopik) {
	daftar[0] = Topik{
		"Makanan", 
		[NMAX 	]string{
			"Yuk makan enak hari ini!", "Makanan enak, hati senang.","Ngemil dulu, kerja nanti!",
		}, [NMAX]string{
			"#makanenak", "#kuliner", "#cemilan",
		}, [NMAX]string{
			"Review makanan lokal favorit", "Cara membuat makanan simpel dan sehat", "Mukbang 5 bungkus mie instan",
		},
	}
	daftar[1] = Topik{
		"Travel",
		[NMAX]string{
			"Petualangan dimulai dari satu langkah.", "Traveling bikin hidup lebih hidup!", "Jalan-jalan yuk!",
		}, [NMAX]string{
			"#jalanjalan", "#travelingindonesia", "#liburan",
		}, [NMAX]string{
			"Vlog perjalanan keluar kota", "Tips liburan hemat", "Tempat hidden gem yang wajib dikunjungi",
		},
	}
	
	daftar[2] = Topik{
		"Motivasi",
		[NMAX]string{
			"Jangan pernah menyerah, kamu bisa!", "Kamu kuat, kamu bisa!", "Jangan lupa untuk selalu berusaha!",
		}, [NMAX]string{
			"#motivasi", "#semangat", "#katabijak",
		}, [NMAX]string{
			"Kisah inspiratif orang sukses", "Langkah kecil untuk semangat besar", "Quotes motivasi untuk hari ini",
		},
	}
	daftar[3] = Topik{
		"Bisnis", [NMAX]string{
			"Bangun brand, bukan cuma jualan.", "Strategi hari ini menentukan hasil besok.", "Jadilah solusi, bukan hanya penjual.",
		}, [NMAX]string{
			"#bisnisonline", "#usahamandiri", "#strategibisnis",
		}, [NMAX]string{
			"Strategi digital marketing", "Cerita kegagalan menjadi pelajaran", "Tips dalam memulai bisnis",
		},
	}
	daftar[4] = Topik{
		"Kesehatan", [NMAX]string{
			"Jaga tubuh, jaga hidup.", "Sehat itu investasi jangka panjang.", "Mulai hidup sehat dari sekarang!",
		}, [NMAX]string{
			"#hidupsehat", "#tipssehat", "#jagakesehatan",
		}, [NMAX]string{
			"Tips hidup sehat", "Cara mengatasi stres", "Olahraga kebugaran untuk pemula",
		},
	}
	daftar[5] = Topik{
		"Edukasi", 
		[NMAX]string{
			"Belajar tidak mengenal usia.", "Ilmu adalah cahaya.", "Pendidikan adalah kunci masa depan.",
		}, [NMAX]string{
			"#belajar", "#pendidikan", "#edukasidigital",
		}, [NMAX]string{
			"Tips belajar efektif", "Penjelasan simpel untuk materi sulit", "Fakta unik dari pelajaran sekolah",
		},
	}
}

func cetakHashtag(daftar tabTopik, N int, pilihan int) {
	for i := 0; i < N; i++ {
		fmt.Println(i + 1, "-", daftar[pilihan].hashtag[i])
	}
}

func cetakIdeKonten(daftar tabTopik, N int, pilihan int) {
	for i := 0; i < N; i++ {
		fmt.Println(i + 1, "-", daftar[pilihan].konten[i])
	}
}

func cetakIsiTopik(daftar tabTopik, N int, pilihan int) {
	for i := 0; i < N; i++ {
		fmt.Println(i + 1, "-", daftar[pilihan].isi[i])
	}
}

func cetakTopik(daftar tabTopik, N int) {
	for i := 0; i < N; i++ {
		fmt.Println(i + 1, "-", daftar[i].topik)
	}
}