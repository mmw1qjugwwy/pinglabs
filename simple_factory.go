package main

import "fmt"

type RemoteHandler struct {
    state int
}

func (s *RemoteHandler) decode_buffer(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*62) % 997
    }
    return result
}

func main() {
    obj := &RemoteHandler{state: 62}
    fmt.Println(obj.decode_buffer(62))
}
