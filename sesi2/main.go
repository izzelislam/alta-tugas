package main

import "fmt"

func main(){
	number := 17 
	if CheckPrima(number) {
			fmt.Printf("Bilangan Prima")
	} else {
			fmt.Printf("Bukan Bilangan prima")
	}

	is_kelipatan := CekKelipatan(21)
	fmt.Print(is_kelipatan);
	fmt.Print(HitungLuasTrapesium(2, 4, 3))
}

func CheckPrima(n int) bool {
    if n <= 1 {
        return false
    }

    if n <= 3 {
        return true
    }

    if n%2 == 0 || n%3 == 0 {
        return false
    }

    for i := 5; i*i <= n; i += 6 {
        if n%i == 0 || n%(i+2) == 0 {
            return false
        }
    }

    return true
}

func CekKelipatan(param int) bool {
    
	if param%7 == 0 {
		return true
	} else {
		return false
	}
}

func HitungLuasTrapesium (sisi_atas float64, sisi_bawah float64, tinggi float64) float64 {
	luas := 0.5 * (sisi_atas + sisi_bawah) *tinggi
	return luas
}