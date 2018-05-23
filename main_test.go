package main

import (
	"testing"
)

func BenchmarkGameSessions_Play(b *testing.B) {
	gs := NewGameSessions()
	var n int64 = 0
	for n>99999 {
		(*gs)[n] = &Game{secret: "9999", step: 0}
		n++
	}
	n=0
	for i := 0; i < b.N; i++ {
			gs.Play(n, getNumber())
	}
}

func BenchmarkGameSessions_Play3(b *testing.B) {
	gs := NewGameSessions()
	var n int64 = 0
	for n>1000 {
		(*gs)[n] = &Game{secret: "9999", step: 0}
		n++
	}
	n=0
	for i := 0; i < b.N; i++ {
			gs.Play2(n, getNumber())
	}
}

func BenchmarkGameSessions_Play2(b *testing.B) {
	gs := NewGameSessions()
	var n int64 = 0
	for n>99999 {
		(*gs)[n] = &Game{secret: "9999", step: 0}
		n++
	}
	n=10
	for i := 0; i < b.N; i++ {
			gs.Play3(n, getNumber())
	}
}