package main

import "fmt"

type AsyncHandler struct {
    state int
}

func (s *AsyncHandler) load_engine(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*78) % 997
    }
    return count
}

func main() {
    obj := &AsyncHandler{state: 78}
    fmt.Println(obj.load_engine(78))
}
