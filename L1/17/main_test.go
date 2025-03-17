package main

import "testing"

func TestTask(t *testing.T) {
    tests := []struct {
        arr      []int
        target   int
        expected int
    }{
        {
            arr:      []int{1, 2, 3, 4, 5},
            target:   3,
            expected: 2,
        },
        {
            arr:      []int{1, 2, 3, 4, 5},
            target:   5,
            expected: 4,
        },
        {
            arr:      []int{3},
            target:   1,
            expected: -1,
        },
    }

    for _, test := range tests {
        t.Run("test", func(t *testing.T) {
            result := Task(test.arr, test.target)
            if result != test.expected {
                t.Errorf("For array %v and target %d, expected %d but got %d", test.arr, test.target, test.expected, result)
            }
        })
    }
}
