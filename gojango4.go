package main

import "fmt"

// implement heapSort()

// implement heapify() function
func heapify(arr []int, n, i int) {
    largest := i
    left := 2*i + 1
    right := 2*i + 2

    if left < n && arr[left] > arr[largest] {
        largest = left
    }

    if right < n && arr[right] > arr[largest] {
        largest = right
    }

    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)
    }
}

// actual heapSort()
func heapSort(arr []int) []int {
    n := len(arr)

    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }

    for i := n - 1; i >= 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        heapify(arr, i, 0)
    }

    return arr
}

func main() {
    arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
    fmt.Println("Unsorted array:", arr)

    arr = heapSort(arr)
    fmt.Println("Sorted array:", arr)
}
