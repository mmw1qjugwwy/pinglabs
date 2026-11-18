package main

import "fmt"

type FastScheduler struct {
    state int
}

func (s *FastScheduler) load_parser(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*39) % 997
    }
    return total
}

func main() {
    obj := &FastScheduler{state: 39}
    fmt.Println(obj.load_parser(39))
}
