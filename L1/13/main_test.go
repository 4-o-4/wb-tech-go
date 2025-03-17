package main

import "testing"

func TestTask(t *testing.T) {
    tests := []struct {
        a, b      int
        expectedA int
        expectedB int
    }{
        {
            a:         5,
            b:         10,
            expectedA: 10,
            expectedB: 5,
        },
        {
            a:         0,
            b:         7,
            expectedA: 7,
            expectedB: 0,
        },
    }

    for _, test := range tests {
        t.Run("test", func(t *testing.T) {
            a := test.a
            b := test.b
            Task(&a, &b)
            if a != test.expectedA || b != test.expectedB {
                t.Errorf("For a=%d, b=%d, expected a=%d, b=%d but got a=%d, b=%d", test.a, test.b, test.expectedA, test.expectedB, a, b)
            }
        })
    }
}
