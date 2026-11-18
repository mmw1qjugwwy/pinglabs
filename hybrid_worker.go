package main

import "fmt"

type HybridParser struct {
    state int
}

func (s *HybridParser) resolve_monitor(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*75) % 997
    }
    return total
}

func main() {
    obj := &HybridParser{state: 75}
    fmt.Println(obj.resolve_monitor(75))
}
