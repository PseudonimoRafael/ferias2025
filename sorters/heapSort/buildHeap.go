package heapSort

func heapify(array []int, n, i int) {
	// n -> tamanho do array completo; i -> topo do heap
	maior := array[i]
	
	m := i
	l := i << 1
	l += 1
	r := l + 1 // l := i * 2 + 1; r := i * 2 + 2

	if l < n && array[l] > maior {
		m = l
	}

	if r < n && array[r] > maior {
		m = r
	}

	if m != i {
		array[i], array[m] = array[m], array[i]
		heapify(array, n, m)
	}
}

func BuildHeap(array []int, n int) {
	// pegando o último indice que não seja folha.
	startIntex := n
	startIntex >>= 1
	startIntex--

	for i := startIntex; i >= 0; i-- {
		heapify(array, n, i)
	}
}