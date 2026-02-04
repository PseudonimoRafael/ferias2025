package heapSort

func HeapSort(array []int) {
	for n:= len(array) - 1; n > 0; n-- {
		// troca
		array[0], array[n] = array[n], array[0]
		// reorganiza o heap
		heapify(array, n, 0)
	} 
}