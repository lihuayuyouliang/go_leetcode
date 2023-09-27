package mysort

func BubblingSort(arr []int) ([]int, int) {
	n := 0
	arrLen := len(arr)
	for i := arrLen - 1; i >= 0; i-- {
		flag := false
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				n++
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
	return arr, n
}

func QuickSort(arr []int) []int {
	quick(arr, 0, len(arr)-1)
	return arr
}
func quick(arr []int, l, r int) {
	if l >= r {
		return
	}
	i := l
	j := r
	for i < j {
		for i < j && arr[j] >= arr[l] {
			j--
		}
		for i < j && arr[i] <= arr[l] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[l], arr[i] = arr[i], arr[l]
	quick(arr, l, i-1)
	quick(arr, i+1, r)
}
