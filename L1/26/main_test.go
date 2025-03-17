package main

import "testing"

func TestTask(t *testing.T) {
    tests := []struct {
        input    string
        expected bool
    }{
        {
            input:    "abcd",
            expected: true,
        },
        {
            input:    "abCdefAaf",
            expected: false,
        },
        {
            input:    "aabcd",
            expected: false,
        },
        {
            input:    "aA",
            expected: true,
        },
    }

    for _, test := range tests {
        t.Run("test", func(t *testing.T) {
            result := Task(test.input)
            if result != test.expected {
                t.Errorf("for input '%s', expected %v but got %v", test.input, test.expected, result)
            }
        })
    }
}
