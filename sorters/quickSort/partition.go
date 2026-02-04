package quickSort
/*
Quem é lomuto?
*/
func partition(array []int, in, fin int) int {
	comp := array[in]
	j := in + 1
	for i := in + 1; i < fin; i++ {
		if array[i] < comp {
			// tem uma forma mais simples de fazer essa troca, Não me diga o que já sei
			tmp := array[j]
			array[j] = array[i]
			array[i] = tmp
			j++
		}
	}
	j-- // necessário
	array[in] = array[j]
	array[j] = comp
	return j
}