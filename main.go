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
		mapName := `pemain-` + strconv.Itoa(i) + ``
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
		for i := 1; i <= len(result); i++ {
			fmt.Println(`pemain-`+strconv.Itoa(i), " : ", result[`pemain-`+strconv.Itoa(i)])
		}

		fmt.Println("========= SETELAH EVALUASI =========")
		exist := 0
		for i := 1; i <= len(result); i++ {
			arrayBaru := evaluasiAngka6(result[`pemain-`+strconv.Itoa(i)], make([]int, 0), `pemain-`+strconv.Itoa(i))
			arrayBaru2 := evaluasiAngka1(arrayBaru, make([]int, 0), `pemain-`+strconv.Itoa(i))

			if exist == i {
				arrayBaru2 = append(arrayBaru2, 1)
			}

			if i == 1 {
				for _, val := range result["pemain-"+strconv.Itoa(len(result))] {
					if val == 1 {
						arrayBaru2 = append(arrayBaru2, 1)
					}
				}
			}

			poin := jumlahDadu - len(arrayBaru2)
			response := `pemain-` + strconv.Itoa(i) + ` mendapat ` + strconv.Itoa(poin) + " poin "
			fmt.Println(response, arrayBaru2)

			for _, val := range result["pemain-"+strconv.Itoa(i)] {
				if val == 1 {
					exist = i + 1
				}
			}
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

func evaluasiAngka1(data []int, data2 []int, index string) []int {
	for i := 0; i < len(data); i++ {
		if data[i] != 1 && data[i] != 0 {
			data2 = append(data2, data[i])
		}
	}
	return data2
}
