package main

import "fmt"
import "OrderAlgo/mergeSort"
import "OrderAlgo/quickSort"
import "OrderAlgo/heapSort"
import "OrderAlgo/dataGen"

func main() {
	listaDesordenada := dataGen.Generate() // Propositalmente ordenada de traz pra frente. (ferrar o quickSort)
	var listaMerge []int = mergeSort.MergeSort(listaDesordenada[:])
	
	
	var listaQuick []int = make([]int, 151)
	copy(listaQuick ,listaDesordenada[:])
	quickSort.QuickSort(listaQuick, 0, 151)

	listaHeap := make([]int, 151)
	copy(listaHeap, listaDesordenada[:])
	heapSort.BuildHeap(listaHeap)
	fmt.Printf("\nHeapBuild:\n%v\n", listaHeap)
	heapSort.HeapSort(listaHeap)

	fmt.Printf("Merge Sort\n\n%v\n", listaMerge)
	fmt.Printf("Quick Sort\n\n%v\n", listaQuick)
	fmt.Printf("Heap Sort\n\n%v\n", listaHeap)
}