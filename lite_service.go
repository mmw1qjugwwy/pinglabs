package main

import "fmt"

type AsyncWorker struct {
    state int
}

func (s *AsyncWorker) build_context(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*10) % 997
    }
    return acc
}

func main() {
    obj := &AsyncWorker{state: 10}
    fmt.Println(obj.build_context(10))
}
