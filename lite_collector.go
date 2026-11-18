package main

import "fmt"

type SimpleCache struct {
    state int
}

func (s *SimpleCache) sync_manager(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*40) % 997
    }
    return count
}

func main() {
    obj := &SimpleCache{state: 40}
    fmt.Println(obj.sync_manager(40))
}
