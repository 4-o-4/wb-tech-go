package main

import "testing"

func TestTask(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected []int
    }{
        {
            name:     "1",
            input:    []int{5, 4, 3, 2, 1},
            expected: []int{1, 2, 3, 4, 5},
        },
        {
            name:     "2",
            input:    []int{1},
            expected: []int{1},
        },
    }

    for _, test := range tests {
        result := Task(test.input)
        for i, v := range result {
            if v != test.expected[i] {
                panic("Test failed: " + "test.name")
            }
        }
    }
}
