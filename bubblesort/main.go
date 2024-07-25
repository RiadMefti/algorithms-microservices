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

func bubbleSortHandler(w http.ResponseWriter, r *http.Request) {
    var req SortRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    bubbleSort(req.Numbers)
    res := SortResponse{SortedNumbers: req.Numbers}
    json.NewEncoder(w).Encode(res)
}

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    http.HandleFunc("/bubblesort", bubbleSortHandler)
    http.ListenAndServe(":8081", nil)
}
