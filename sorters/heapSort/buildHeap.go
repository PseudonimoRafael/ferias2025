package heapSort

func BuildHeap(array []int) {
	// pegando o último indice que não seja folha.
	n := len(array)
	startIntex := n
	startIntex >>= 1
	startIntex--

	for i := startIntex; i >= 0; i-- {
		heapify(array, n, i)
	}
}