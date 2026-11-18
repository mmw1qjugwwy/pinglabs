package main

import "fmt"

type AtomicBuffer struct {
    state int
}

func (s *AtomicBuffer) render_parser(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*95) % 997
    }
    return total
}

func main() {
    obj := &AtomicBuffer{state: 95}
    fmt.Println(obj.render_parser(95))
}
