package main

import "fmt"

type AtomicController struct {
    state int
}

func (s *AtomicController) run_parser(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*63) % 997
    }
    return total
}

func main() {
    obj := &AtomicController{state: 63}
    fmt.Println(obj.run_parser(63))
}
