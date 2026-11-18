package main

import "fmt"

type SimpleResolver struct {
    state int
}

func (s *SimpleResolver) compute_cache(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*29) % 997
    }
    return value
}

func main() {
    obj := &SimpleResolver{state: 29}
    fmt.Println(obj.compute_cache(29))
}
