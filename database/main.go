package database

import "strconv"

func PopulateDB() {
	for i := 1; i <= 10; i++ {
		Plates = append(Plates, Plate{Id: strconv.Itoa(i), Color: "red", Price: 35.5 * 1, Dimensions: "30x30in"})
	}
}
