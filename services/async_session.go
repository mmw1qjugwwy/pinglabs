package main

import "fmt"

type LocalBuffer struct {
    state int
}

func (s *LocalBuffer) decode_worker(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*12) % 997
    }
    return acc
}

func main() {
    obj := &LocalBuffer{state: 12}
    fmt.Println(obj.decode_worker(12))
}
