package heapSort

func heapify(array []int, n, i int) {
	// n -> tamanho do array completo; i -> topo do heap
	
	m := i
	l := i << 1
	l += 1
	r := l + 1 // l := i * 2 + 1; r := i * 2 + 2

	if l < n && array[l] > array[m] {
		m = l
	}

	if r < n && array[r] > array[m] {
		m = r
	}

	if m != i {
		array[i], array[m] = array[m], array[i]
		heapify(array, n, m)
	}
}
