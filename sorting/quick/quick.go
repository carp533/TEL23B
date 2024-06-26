package quick

func Sort(a []int) []int {

	quickSort(a, 0, len(a)-1)
	return a
}
func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)
		// Recursively sort elements before partition and after partition
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
