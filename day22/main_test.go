package main

import (
	"fmt"
	"testing"

	"github.com/gonutz/check"
)

func TestCubeSide(t *testing.T) {
	s := newCubeSide([]string{
		"123",
		"456",
		"789",
	}, 1000, 2000)

	want := func(x, y int, expect byte) {
		t.Helper()
		check.Eq(t, string(s.at(x, y)), string(expect))
	}
	wantOriginal := func(x, y, wantX, wantY int) {
		t.Helper()
		haveX, haveY := s.originalXY(x, y)
		check.Eq(t, [2]int{haveX, haveY}, [2]int{wantX, wantY})
	}

	want(0, 0, '1')
	want(1, 0, '2')
	want(2, 1, '6')
	want(1, 2, '8')
	want(2, 2, '9')
	wantOriginal(0, 0, 1000, 2000)
	wantOriginal(1, 2, 1001, 2002)

	// 369
	// 258
	// 147
	s.rotateLeft()
	want(0, 0, '3')
	want(1, 0, '6')
	want(2, 1, '8')
	want(1, 2, '4')
	want(2, 2, '7')
	fmt.Println("now")
	wantOriginal(0, 0, 1002, 2000)
	fmt.Println("done")
	wantOriginal(1, 2, 1000, 2001)

	// 123
	// 456
	// 789
	s.rotateRight()
	want(0, 0, '1')
	want(1, 0, '2')
	want(2, 1, '6')
	want(1, 2, '8')
	want(2, 2, '9')
	wantOriginal(0, 0, 1000, 2000)
	wantOriginal(1, 2, 1001, 2002)

	// 741
	// 852
	// 963
	s.rotateRight()
	want(0, 0, '7')
	want(1, 0, '4')
	want(2, 1, '2')
	want(1, 2, '6')
	want(2, 2, '3')
	wantOriginal(0, 0, 1000, 2002)
	wantOriginal(1, 2, 1002, 2001)
}
