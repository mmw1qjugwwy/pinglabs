package main

import "fmt"

type BatchCollector struct {
    state int
}

func (s *BatchCollector) run_processor(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*24) % 997
    }
    return result
}

func main() {
    obj := &BatchCollector{state: 24}
    fmt.Println(obj.run_processor(24))
}
