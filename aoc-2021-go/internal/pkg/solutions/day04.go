package solutions

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type BingoNumber struct {
	number int
	marked bool
}

type BingoBoard struct {
	cells    [5][5]*BingoNumber
	IsWinner bool
	LastDraw int
}

func (b *BingoBoard) getScore() (int, error) {
	if !b.IsWinner {
		return 0, fmt.Errorf("board did not win yet")
	}
	unmarked := b.getUnmarkedSum()
	return unmarked * b.LastDraw, nil
}

func (b *BingoBoard) getUnmarkedSum() int {
	result := 0
	for i := 0; i < len(b.cells); i++ {
		for j := 0; j < len(b.cells[i]); j++ {
			if !b.cells[i][j].marked {
				result += b.cells[i][j].number
			}
		}
	}
	return result
}

func (b *BingoBoard) play(draw int) {
	if b.IsWinner {
		return
	}
	bingoRows := [5]bool{true, true, true, true, true}
	bingoCols := [5]bool{true, true, true, true, true}
	for i := 0; i < len(b.cells); i++ {
		for j := 0; j < len(b.cells[i]); j++ {
			if draw == b.cells[i][j].number {
				b.cells[i][j].marked = true
				b.LastDraw = draw
			}
			if !b.cells[i][j].marked {
				bingoRows[i] = false
				bingoCols[j] = false
			}
		}
	}
	for i := 0; i < 5; i++ {
		b.IsWinner = b.IsWinner || bingoRows[i] || bingoCols[i]
	}
}

func NewBingoBoard(s string) (*BingoBoard, error) {
	var cells [5][5]*BingoNumber
	for i, row := range strings.Split(strings.Trim(s, "\n"), "\n") {
		if len(row) == 0 {
			if i >= 5 {
				continue
			} else {
				return nil, fmt.Errorf("invalid data: pos=%v data='%v'", i, s)
			}
		}
		for j, entry := range strings.Fields(row) {
			n, err := strconv.Atoi(entry)
			if err != nil {
				return nil, fmt.Errorf("%v: char '%v'", err, entry)
			}
			cells[i][j] = &BingoNumber{number: n, marked: false}
		}
	}
	return &BingoBoard{cells: cells, IsWinner: false}, nil
}

type BingoGame struct {
	draws  []int
	boards []BingoBoard
}

func NewBingoGame(s string) (*BingoGame, error) {
	// 1 - Load the bingo draws
	draws, err := csvRow2intArray(strings.Split(s, "\n")[0])
	if err != nil {
		return nil, err
	}
	// parse the boards
	var boards []BingoBoard
	for i, boardData := range strings.Split(s, "\n\n") {
		// the first entry is the list of draws - ignore it
		if i == 0 {
			continue
		}
		b, err := NewBingoBoard(boardData)
		if err != nil {
			return nil, err
		}
		boards = append(boards, *b)
	}
	return &BingoGame{draws: draws, boards: boards}, nil
}

func (g *BingoGame) findFirstWinner() {
	for _, draw := range g.draws {
		for i, b := range g.boards {
			b.play(draw)
			if b.IsWinner {
				g.showGameResults(i, b, draw)
				return
			}
		}
	}
}

func (g *BingoGame) findLastWinner() {
	var lastBoard BingoBoard
	lastIndex := 0
	lastDraw := 0
	for _, draw := range g.draws {
		for i := 0; i < len(g.boards); i++ {
			if g.boards[i].IsWinner {
				continue
			}
			g.boards[i].play(draw)
			if g.boards[i].IsWinner {
				lastBoard = g.boards[i]
				lastIndex = i
				lastDraw = draw
			}
		}
	}
	g.showGameResults(lastIndex, lastBoard, lastDraw)
}

func (g *BingoGame) showGameResults(boardNumber int, b BingoBoard, draw int) {
	log.Printf("Board IsWinner the game.......: %v", boardNumber)
	log.Printf("Sum of unmarked numbers..: %v", b.getUnmarkedSum())
	log.Printf("Last drawn number........: %v (%v)", b.LastDraw, draw)
	score, err := b.getScore()
	if err == nil {
		log.Printf("Board score..............: %v", score)
	} else {
		log.Printf("Board score..............: %v", err)
	}
}

func SolveDay04() error {
	log.Printf("❄ ❄ ❄ Advent of Code 2021 - Day 04 ❄ ❄ ❄\n")
	// Load the data
	input, err := ioutil.ReadFile("input/input-day04")
	if err != nil {
		return err
	}
	// Getting started
	// Part 1
	log.Println("❄ ❄ * PART 1")
	game1, err := NewBingoGame(string(input))
	if err != nil {
		return err
	}
	game1.findFirstWinner()
	// Part 2
	log.Println("❄ ❄ * PART 2")
	game2, err := NewBingoGame(string(input))
	if err != nil {
		return err
	}
	game2.findLastWinner()

	return nil
}
