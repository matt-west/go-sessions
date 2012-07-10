package main

import (
	"./sessions"
	"fmt"
)

func main() {
	s := sessions.Session{}
	fmt.Println(s.Get("boo"))
}
