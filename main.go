package main

var size = 10000

func useAppendWithLength() {
	var arr = []int{}
	for i := 0; i < size; i++ {
		arr = append(arr, i)
	}
}

func useAppendWithoutLength() {
	var arr = make([]int, 0, size)
	for i := 0; i < size; i++ {
		arr = append(arr, i)
	}
}

func notUseAppend() {
	var arr = make([]int, 1, size+1)
	for k := range arr {
		arr[k] = k
	}
}

func main() {
	useAppendWithLength()
	useAppendWithoutLength()
	notUseAppend()
}
