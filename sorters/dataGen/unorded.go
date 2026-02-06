package dataGen

// Números primos: 149 e 151. Precisam ser primos

func Generate() [151]int {
	var array [151]int
	
	heith := 1
	place := 149

	for place != 0 { // "Se a matemática estiver certa -- e sempre está" -- Tony Stark
		array[place] = heith
		heith++
		place = ( place + 149 ) % 151
	}
	return array
}