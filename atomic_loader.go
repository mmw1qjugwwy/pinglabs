package main

import "fmt"

type FastCache struct {
    state int
}

func (s *FastCache) decode_cache(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*48) % 997
    }
    return total
}

func main() {
    obj := &FastCache{state: 48}
    fmt.Println(obj.decode_cache(48))
}
