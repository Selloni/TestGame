package posgresql

import (
	"math/rand"
)

func generateItem() map[string]int {
	countItem := rand.Intn(4) + 2

	mm := make(map[string]int, countItem)
	//rand.Seed(time.Now().Unix()) // портит
	length := 4
	ranStr := make([]byte, length)

	for ; countItem > 0; countItem-- {
		for i := 0; i < length; i++ {
			ranStr[i] = byte(65 + rand.Intn(25))
		}
		mm[string(ranStr)] = rand.Intn(70) + 10
	}
	return mm
}
