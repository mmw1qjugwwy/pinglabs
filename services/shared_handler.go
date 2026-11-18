package main

import "fmt"

type SecureFactory struct {
    state int
}

func (s *SecureFactory) resolve_router(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*39) % 997
    }
    return count
}

func main() {
    obj := &SecureFactory{state: 39}
    fmt.Println(obj.resolve_router(39))
}
