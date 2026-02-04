package mergeSort

func MergeSort(array []int) (ordArray []int) {
	/*
		caso um dos arrays seja de tamanho 2,
		chama direto o merge
	*/
	l := len(array)
	m := l >> 1
	delta := l - m

	if l < 2 {
		return array
	}

	var array1 []int
	var array2 []int

	if m < 2 {
		array1 = array[0: m]
	} else {
		array1 = MergeSort(array[0: m])
	}

	if delta < 2{
		array2 = array[m: l]
	} else {
		array2 = MergeSort(array[m: l])
	}

	ordArray = Merge(array1, array2)
	return
}