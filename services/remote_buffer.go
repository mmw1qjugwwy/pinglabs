package main

import "fmt"

type SharedBuilder struct {
    state int
}

func (s *SharedBuilder) run_handler(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*28) % 997
    }
    return result
}

func main() {
    obj := &SharedBuilder{state: 28}
    fmt.Println(obj.run_handler(28))
}
