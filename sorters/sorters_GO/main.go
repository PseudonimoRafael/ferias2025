package main

import "fmt"
import "mergeSRT/mergeSort"

func main() {
	listaDesordenada := []int{15,14,15,12,11,10,9,8,7,6,5,4,3,2,1}
	var listaOrdenada []int = mergeSRT.MergeSort(listaDesordenada)

	fmt.Printf("Antes: %v\n, Depois: %v\n", listaDesordenada, listaOrdenada)
}