package mergeSort

func Merge(lista1, lista2 []int ) (mergedList []int) {
	lenNew := len(lista1) + len(lista2)
	var len1, len2 int = len(lista1), len(lista2)
	var i, j, k int = 0, 0, 0

	mergedList = make([]int, lenNew)

	for i < len1 && j < len2 {
		if lista1[i] <= lista2[j] {
			mergedList[k] = lista1[i]
			i++
		} else {
			mergedList[k] = lista2[j]
			j++
		}
		k++
	}

	for i < len1 {
		mergedList[k] = lista1[i]
		i++
		k++
	}
	for j < len2 {
		mergedList[k] = lista2[j]
		j++
		k++
	}
	return
}