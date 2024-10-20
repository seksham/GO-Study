package main


import "testing"

func TestFindLastOccurrence(t *testing.T) {
    tests := []struct {
        name     string
        arr      []int
        target   int
        expected int
    }{
        {"Empty slice", []int{}, 5, -1},
        {"Single element, found", []int{5}, 5, 0},
        {"Single element, not found", []int{3}, 5, -1},
        {"Multiple elements, found at end", []int{1, 2, 3, 4, 5, 5, 5}, 5, 6},
        {"Multiple elements, found in middle", []int{1, 2, 3, 3, 3, 4, 5}, 3, 4},
        {"Multiple elements, found at beginning", []int{1, 1, 1, 2, 3, 4, 5}, 1, 2},
        {"Multiple elements, not found", []int{1, 2, 3, 4, 6, 7, 8}, 5, -1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := FindLastOccurrence(tt.arr, tt.target)
            if result != tt.expected {
                t.Errorf("FindLastOccurrence(%v, %d) = %d; want %d", tt.arr, tt.target, result, tt.expected)
            }
        })
    }
}