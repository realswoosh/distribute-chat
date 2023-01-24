package main

import (
	"log"
	"testing"
)

func TestSliceAppend(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{100, 200, 300}
	s1 = append(s1[:1], s2[:3]...)

	log.Println("s1 = ", s1)

	s10 := []int{1, 2, 3}
	s10 = s10[1:]
	log.Println("s10 = ", s10)
}
