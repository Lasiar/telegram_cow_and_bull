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

func (gs *GameSessions) Play(key int64, guess string) (string, bool) {

	game, ok := gs.GetGame(key)

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

func (gs *GameSessions) Play2(key int64, guess string) (string, bool) {

	game, ok := gs.GetGame(key)

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

func (gs *GameSessions) Play3(id int64, guess string) (string, bool) {
	(*gs)[id].step += 1
	check := 0
	buf := new(strings.Builder)
	for i, char := range guess {
		if byte(char) == (*gs)[id].secret[i] {
			buf.WriteString(right)
			check ++
			continue
		}
		if strings.Contains((*gs)[id].secret, string(char)) {
			buf.WriteString(contains)
		} else {
			buf.WriteString(notFound)
		}
	}
	return buf.String(), check == 4
}

func (gs *GameSessions) NewGame(id int64) *Game {
	game := &Game{secret: getNumber(), step: 0}
	(*gs)[id] = game
	return game
}

func (gs *GameSessions) GetGame(key int64) (*Game, bool) {
	game, ok := (*gs)[key]
	return game, ok
}

func (gs *GameSessions) DeleteGame(key int64) {
	delete(*gs, key)
}

type Game struct {
	secret string
	step   int
}

func (g Game) GetStingStepLine() string {
	if g.step == 1 {
		return "шаг"
	}
	if g.step < 5 {
		return "шага"
	}
	return "шагов"
}

func (g Game) String() string {
	return fmt.Sprintf("Загаданное число: %s шаг: %d", g.secret, g.step)
}

func NewGameSessions() *GameSessions {
	g := GameSessions(make(map[int64]*Game))
	return &g
}

func getNumber() (secret string) {
	rand.Seed(time.Now().UTC().UnixNano())
	return strconv.Itoa((1000 + rand.Intn(9000)))
}
