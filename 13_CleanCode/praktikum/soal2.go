package main

import "fmt"

type Kendaraan struct {
    totalRoda       int
    kecepatanPerJam int
}

type Mobil struct {
    Kendaraan
}

func (m *Mobil) Berjalan() {
    m.tambahKecepatan(10)
}

func (m *Mobil) tambahKecepatan(kecepatanBaru int) {
    m.kecepatanPerJam += kecepatanBaru
}

func main() {
    mobilCepat := Mobil{}
    mobilCepat.Berjalan()
    mobilCepat.Berjalan()
    mobilCepat.Berjalan()

    mobilLamban := Mobil{}
    mobilLamban.Berjalan()

    fmt.Printf("Mobil cepat: %d km/h\n", mobilCepat.kecepatanPerJam)
    fmt.Printf("Mobil lamban: %d km/h\n", mobilLamban.kecepatanPerJam)
}
