package main

import (
    "fmt"
    "math/big"
)

type BigInt struct {
    a, b *big.Int
}

func NewBigInt(a, b int64) *BigInt {
    return &BigInt{
        a: big.NewInt(a),
        b: big.NewInt(b),
    }
}

func (b BigInt) Add() *big.Int {
    return new(big.Int).Add(b.a, b.b)
}

func (b BigInt) Sub() *big.Int {
    return new(big.Int).Sub(b.a, b.b)
}

func (b BigInt) Mul() *big.Int {
    return new(big.Int).Mul(b.a, b.b)
}

func (b BigInt) Div() *big.Int {
    return new(big.Int).Div(b.a, b.b)
}

func main() {
    b := NewBigInt(2, 2)
    fmt.Println("Сумма (a + b):", b.Add())
    fmt.Println("Разность (a - b):", b.Sub())
    fmt.Println("Произведение (a * b):", b.Mul())
    fmt.Println("Частное (b / a):", b.Div())
}
