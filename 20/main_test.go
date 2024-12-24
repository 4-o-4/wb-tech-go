package main

import "testing"

func TestTask(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {
            input:    "snow dog sun",
            expected: "sun dog snow",
        },
        {
            input:    "one two",
            expected: "two one",
        },
        {
            input:    "",
            expected: "",
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
