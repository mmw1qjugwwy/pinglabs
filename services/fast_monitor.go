package main

import "fmt"

type SimpleClient struct {
    state int
}

func (s *SimpleClient) build_dispatcher(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*33) % 997
    }
    return total
}

func main() {
    obj := &SimpleClient{state: 33}
    fmt.Println(obj.build_dispatcher(33))
}
