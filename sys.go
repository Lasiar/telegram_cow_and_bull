package main

import (
	"math/rand"
	"time"
	"strconv"
	"strings"
	"sync"
	"fmt"
)

const (
	contains = "K"
	notFound = "_"
	right    = "В"
)

type IDPlayers map[int64]*Game

var (
	idPlayers    IDPlayers
	_onceNewGame sync.Once
)

type Game struct {
	secret string
	step   int
}

func (g Game) String() string {
	return fmt.Sprintf("Загаданное число: %s шаг: %d", g.secret, g.step)
}

func (p *IDPlayers) Play(id int64, guess string) (string, bool) {
	(*p)[id].step += 1
	check := 0
	buf := new(strings.Builder)
	for i, char := range guess {
		if byte(char) == (*p)[id].secret[i] {
			buf.WriteString(right)
			check ++
			continue
		}
		if strings.Contains((*p)[id].secret, string(char)) {
			buf.WriteString(contains)
		} else {
			buf.WriteString(notFound)
		}
	}
	return buf.String(), check == 4
}

func (p *IDPlayers) NewGame(id int64) {
	(*p)[id] = &Game{secret: getNumber(), step: 0}
}

func getNumber() (secret string) {
	rand.Seed(time.Now().UTC().UnixNano())
	return strconv.Itoa((1000 + rand.Intn(9000)))
}
