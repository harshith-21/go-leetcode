package functions

//13. Roman to Integer
//
//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

func RomanToInt(s string) int {
	//priority := []uint8{'M', 'D', 'C', 'L', 'X', 'V', 'I'}
	//value := []int{1000, 500, 100, 50, 10, 5, 1}
	//valueof := make(map[uint8]int)
	//for i, symbol := range priority {
	//	valueof[symbol] = value[i]
	//}
	valueof := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum, current, last := 0, 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		current = valueof[s[i]]
		if last > current {
			sum -= current
		} else {
			sum += current
		}
		last = current
	}
	return sum
}
