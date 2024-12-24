package main

import (
    "math"
    "testing"
)

func TestDistance(t *testing.T) {
    tests := []struct {
        p1, p2   *Point
        expected float64
    }{
        {
            p1:       NewPoint(0, 0),
            p2:       NewPoint(3, 0),
            expected: 3,
        },
        {
            p1:       NewPoint(0, 0),
            p2:       NewPoint(3, 4),
            expected: 5,
        },
        {
            p1:       NewPoint(-1, -1),
            p2:       NewPoint(2, 3),
            expected: 5,
        },
    }

    for _, test := range tests {
        t.Run("test", func(t *testing.T) {
            result := Distance(*test.p1, *test.p2)
            if math.Abs(result-test.expected) > 1e-9 {
                t.Errorf("for points %v and %v, expected %v but got %v", test.p1, test.p2, test.expected, result)
            }
        })
    }
}
