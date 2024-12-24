package main

import "testing"

func TestTask(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {
            input:    "главрыба",
            expected: "абырвалг",
        },
        {
            input:    "hello",
            expected: "olleh",
        },
    }

    for _, test := range tests {
        t.Run("test", func(t *testing.T) {
            result := Task(test.input)
            if result != test.expected {
                t.Errorf("for input '%s', expected '%s' but got '%s'", test.input, test.expected, result)
            }
        })
    }
}
