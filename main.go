package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	game(4, 3)
}

func game(jumlahDadu, jumlahPemain int) {
	result := make(map[string][]int, jumlahPemain)
	for i := 1; i <= jumlahPemain; i++ {
		mapName := `hasil-lemparan-pemain-` + strconv.Itoa(i) + ``
		result[mapName] = make([]int, jumlahDadu)
	}
	counter := 0
	for 1 > 0 {
		listPemenang := make([]interface{}, 0)
		counter++
		// lempar dadu
		for i, v := range result {
			result[i] = make([]int, 0)
			if counter > 1 {
				arrayBaru := evaluasiAngka6(v, make([]int, 0), i)
				lemparDadu(len(arrayBaru), i, result)
			} else {
				lemparDadu(jumlahDadu, i, result)
			}
		}
		fmt.Println("========= HASIL LEMPAR DADU =========")
		for i, _ := range result {
			fmt.Println(i, " : ", result[i])
		}

		fmt.Println("========= SETELAH EVALUASI =========")
		for i, v := range result {
			arrayBaru := evaluasiAngka6(v, make([]int, 0), i)
			poin := jumlahDadu - len(arrayBaru)
			response := i + ` mendapat ` + strconv.Itoa(poin) + " poin "
			fmt.Println(response, arrayBaru)
		}

		for i, _ := range result {
			if len(result[i]) == 0 {
				listPemenang = append(listPemenang, i)
			}
		}

		if jumlahPemain-len(listPemenang) == 1 {
			fmt.Println("Game berakhir pemenang nya ", listPemenang)
			break
		}

		time.Sleep(1 * time.Second)
	}

}

func lemparDadu(jumlahDadu int, indexName string, mapPemain map[string][]int) {
	max := 6
	min := 1
	for i := 0; i < jumlahDadu; i++ {
		resultRoll := rand.Intn(max) + min
		mapPemain[indexName] = append(mapPemain[indexName], resultRoll)
	}
}

func evaluasiAngka6(data []int, data2 []int, index string) []int {
	for i := 0; i < len(data); i++ {
		if data[i] != 6 && data[i] != 0 {
			data2 = append(data2, data[i])
		}
	}
	return data2
}
