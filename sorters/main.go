package main

import "fmt"
import "OrderAlgo/mergeSort"
import "OrderAlgo/quickSort"

func main() {
	listaDesordenada := []int{15,14,15,12,11,10,9,8,7,6,5,4,3,2,1} // Propositalmente ordenada de traz pra frente. (ferrar o quickSort)
	var listaMerge []int = mergeSort.MergeSort(listaDesordenada)
	var listaQuick []int = make([]int, len(listaDesordenada))
	copy(listaQuick ,listaDesordenada)
	quickSort.QuickSort(listaQuick, 0, 15)

	fmt.Printf("Merge Sort\n\nAntes: %v\n Depois: %v\n", listaDesordenada, listaMerge)
	fmt.Printf("Quick Sort\n\nAntes: %v\n Depois: %v\n", listaDesordenada, listaQuick)
}