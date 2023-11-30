package main

import (
	"fmt"
)

type Mobil struct {
	Ban       [4]string
	PintuKanan string
	PintuKiri  string
}

func main() {
	mobil := Mobil{
		Ban:       [4]string{"karet", "karet", "karet", "karet"},
		PintuKanan: "Knock",
		PintuKiri:  "beep",
	}

	fmt.Println("Bunyi pintu kanan ketika diketuk:", mobil.PintuKanan)
	fmt.Println("Bunyi pintu kiri ketika diketuk:", mobil.bunyiPintuKiri())

	BanBaru := [4]string{"kayu", "besi", "karet", "besi"}
	mobil.gantiBan(BanBaru)

	fmt.Println("Bunyi pintu kanan ketika dibuka:", mobil.PintuKanan)
	fmt.Println("Bunyi pintu kiri ketika dibuka:", mobil.bunyiPintuKiri())
}

func (m *Mobil) gantiBan(BanBaru [4]string) {
	if len(BanBaru) == 4 {
		m.Ban = BanBaru
		fmt.Println("Ban mobil telah diganti.")
	} else {
		fmt.Println("Jumlah Ban tidak sesuai. Harap ganti dengan 4 Ban.")
	}
}

func (m Mobil) bunyiPintuKiri() string {
	if m.PintuKanan == "Knock" {
		return "beep"
	}
	return "Knock"
}
