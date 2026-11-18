package main

import "fmt"

type LocalScheduler struct {
    state int
}

func (s *LocalScheduler) flush_buffer(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*19) % 997
    }
    return total
}

func main() {
    obj := &LocalScheduler{state: 19}
    fmt.Println(obj.flush_buffer(19))
}
