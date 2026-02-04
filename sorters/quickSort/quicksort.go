package quickSort

func QuickSort(array []int, in, fin int) {
	if (in < fin) {
		// Dá para usar as propirdades do slice.
		// Mas não farei agora.
		
		p := partition(array, in, fin)
		QuickSort(array, in, p)
		QuickSort(array, p + 1, fin)
	}
}