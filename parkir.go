package main

import "fmt"

func sistemTarifParkir(jam *int, hari *bool, keanggotaan *bool) int {
	var tarif int = 0
	if *jam <= 2 {
		tarif = +5000
	} else {
		tarif = tarif + 5000 + (*jam-2)*2000
	}
	if *keanggotaan == true {
		if *jam <= 5 {
			tarif -= tarif * 50 / 100
		} else {
			tarif -= tarif * 30 / 100
		}
	}
	if *hari == true {
		tarif += 3000
	}
	return tarif
}

func main() {
	var waktu int = 4
	var hari bool = false
	var anggota bool = false
	skenario1 := sistemTarifParkir(&waktu, &hari, &anggota)
	fmt.Println("biaya tarif total selama 4 jam,bukan member dan hari libur :", skenario1)

	var time int = 2
	var day bool = true
	var membership bool = true
	skenario2 := sistemTarifParkir(&time, &day, &membership)
	fmt.Println("biaya tarif total selama 2 jam,member dan hari libur :", skenario2)
}
