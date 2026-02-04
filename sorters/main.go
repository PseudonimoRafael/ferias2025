package main

import "fmt"
import "OrderAlgo/mergeSort"
import "OrderAlgo/quickSort"
import "OrderAlgo/heapSort"

func main() {
	listaDesordenada := []int{15,14,15,12,11,10,9,8,7,6,5,4,3,2,1} // Propositalmente ordenada de traz pra frente. (ferrar o quickSort)
	var listaMerge []int = mergeSort.MergeSort(listaDesordenada)
	
	
	var listaQuick []int = make([]int, len(listaDesordenada))
	copy(listaQuick ,listaDesordenada)
	quickSort.QuickSort(listaQuick, 0, 15)

	listaHeap := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	heapSort.BuildHeap(listaHeap)
	fmt.Printf("\nHeapBuild:\n%v\n", listaHeap)
	heapSort.HeapSort(listaHeap)

	fmt.Printf("Merge Sort\n\n%v\n", listaMerge)
	fmt.Printf("Quick Sort\n\n%v\n", listaQuick)
	fmt.Printf("Heap Sort\n\n%v\n", listaHeap)
}