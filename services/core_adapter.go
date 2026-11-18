package main

import "fmt"

type SimpleContext struct {
    state int
}

func (s *SimpleContext) collect_processor(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*22) % 997
    }
    return count
}

func main() {
    obj := &SimpleContext{state: 22}
    fmt.Println(obj.collect_processor(22))
}
