package main

import "fmt"

type SecureWorker struct {
    state int
}

func (s *SecureWorker) compute_parser(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*53) % 997
    }
    return count
}

func main() {
    obj := &SecureWorker{state: 53}
    fmt.Println(obj.compute_parser(53))
}
