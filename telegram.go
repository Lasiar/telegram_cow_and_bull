package main

import (
	"math/rand"
	"time"
	"strconv"
	"strings"
	"fmt"
	"bytes"
)

const (
	contains = "K"
	notFound = "_"
	right    = "В"
)

type GameSessions map[int64]*Game

type Game struct {
	secret string
	step   int
}

func (g Game) String() string {
	return fmt.Sprintf("Загаданное число: %s шаг: %d", g.secret, g.step)
}

func (p *GameSessions) Play(key int64, guess string) (string, bool) {

	game, ok := p.GetGame(key)

	if !ok {
		return "", false
	}

	game.step ++

	check := 0
	buf := new(strings.Builder)
	for i, char := range guess {
		if byte(char) == game.secret[i] {
			buf.WriteString(right)
			check ++
			continue
		}
		if strings.Contains(game.secret, string(char)) {
			buf.WriteString(contains)
		} else {
			buf.WriteString(notFound)
		}
	}
	return buf.String(), check == 4
}

func (p *GameSessions) Play2(key int64, guess string) (string, bool) {

	game, ok := p.GetGame(key)

	if !ok {
		return "", false
	}

	game.step ++

	check := 0
	buf := new(bytes.Buffer)
	for i, char := range guess {
		if byte(char) == game.secret[i] {
			buf.WriteString(right)
			check ++
			continue
		}
		if strings.Contains(game.secret, string(char)) {
			buf.WriteString(contains)
		} else {
			buf.WriteString(notFound)
		}
	}
	return buf.String(), check == 4
}

func (p *GameSessions) Play3(id int64, guess string) (string, bool) {
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


func (p *GameSessions) NewGame(id int64) *Game {
	game := &Game{secret: getNumber(), step: 0}
	(*p)[id] = game
	return game
}

func (g *GameSessions) GetGame(key int64) (*Game, bool) {
	game, ok := (*g)[key]
	return game, ok
}

func (g *GameSessions) DeleteGame(key int64) {
	delete(*g, key)
}

func NewGameSessions() *GameSessions {
	g := GameSessions(make(map[int64]*Game))
	return &g
}

func getNumber() (secret string) {
	rand.Seed(time.Now().UTC().UnixNano())
	return strconv.Itoa((1000 + rand.Intn(9000)))
}
