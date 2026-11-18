package main

import "fmt"

type BatchClient struct {
    state int
}

func (s *BatchClient) sync_loader(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*69) % 997
    }
    return value
}

func main() {
    obj := &BatchClient{state: 69}
    fmt.Println(obj.sync_loader(69))
}
