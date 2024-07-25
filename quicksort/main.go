package main

import (
	"encoding/json"
	"net/http"
)

type SortRequest struct {
	Numbers []int `json:"numbers"`
}

type SortResponse struct {
	SortedNumbers []int `json:"sortedNumbers"`
}

func quickSortHandler(w http.ResponseWriter, r *http.Request) {
	var req SortRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	quickSort(req.Numbers)
	res := SortResponse{SortedNumbers: req.Numbers}
	json.NewEncoder(w).Encode(res)
}

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	pivot := len(arr) / 2

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

func main() {
	http.HandleFunc("/quicksort", quickSortHandler)
	http.ListenAndServe(":8080", nil)
}
