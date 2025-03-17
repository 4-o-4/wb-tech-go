package main

import "fmt"

type Human struct {
    Name string
    Age  int
}

func (h Human) Info() {
    fmt.Printf("Привет! Меня зовут %s, мне %d лет.\n", h.Name, h.Age)
}

type Action struct {
    Human
    Skill string
}

func (a Action) PerformAction() {
    fmt.Printf("%s выполняет действие с навыком: %s\n", a.Name, a.Skill)
}

func main() {
    a := Action{
        Human: Human{
            Name: "Иван",
            Age:  30,
        },
        Skill: "Программирование",
    }
    a.Info()
    a.PerformAction()
}
