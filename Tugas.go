package main

import (
    "fmt"
)

func main() {
    // Deklarasi variabel
    var billAmount float64
    var tip float64

    // Input nilai tagihan
    fmt.Print("Masukkan nilai tagihan: ")
    fmt.Scanln(&billAmount)

    // Hitung tip
    if billAmount >= 50 && billAmount <= 300 {
        tip = billAmount * 0.15
    } else {
        tip = billAmount * 0.2
    }

    // Hitung nilai akhir
    var totalAmount float64 = billAmount + tip

    // Tampilkan hasil
    fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", billAmount, tip, totalAmount)
}