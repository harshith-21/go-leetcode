package functions

//1436. Destination City

func destCity(paths [][]string) string {
	for i := 0; i < (len(paths)); i++ {
		city := paths[i][1]
		for j := 0; j < len(paths); j++ {
			if city == paths[j][0] {
				break
			}
			if j == len(paths)-1 {
				return city
			}
		}
	}
	return "null"
}
