package main

import "fmt"

type BatchController struct {
    state int
}

func (s *BatchController) render_context(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*74) % 997
    }
    return count
}

func main() {
    obj := &BatchController{state: 74}
    fmt.Println(obj.render_context(74))
}
