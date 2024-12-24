package main

func Task(arr []int, target int) int {
    return binarySearchHelper(arr, target, 0, len(arr)-1)
}

func binarySearchHelper(arr []int, target, start, end int) int {
    if start > end {
        return -1
    }
    mid := start + (end-start)/2
    if arr[mid] > target {
        return binarySearchHelper(arr, target, start, mid-1)
    } else if arr[mid] < target {
        return binarySearchHelper(arr, target, mid+1, end)
    }
    return mid
}
