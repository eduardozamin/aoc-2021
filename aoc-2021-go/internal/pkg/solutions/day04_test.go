package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBingoBoard(t *testing.T) {
	b, err := NewBingoBoard(`
24 75 59 41 17
58 74 64 92 39
68  8 78 85 72
18  3 22  4 34
11 76  6 28 50
`)
	assert.Nil(t, err)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			assert.False(t, b.cells[i][j].marked)
		}
	}
}

func TestBingoBoard_simple_play(t *testing.T) {
	b, err := NewBingoBoard(`
24 75 59 41 17
58 74 64 92 39
68  8 78 85 72
18  3 22  4 34
11 76  6 28 50
`)
	assert.Nil(t, err)
	b.play(24)
	b.play(17)
	b.play(3)
	b.play(4)
	b.play(50)
	assert.True(t, b.cells[0][0].marked)
	assert.True(t, b.cells[0][4].marked)
	assert.True(t, b.cells[3][1].marked)
	assert.True(t, b.cells[3][3].marked)
	assert.True(t, b.cells[4][4].marked)
	assert.False(t, b.IsWinner)
}

func TestBingoBoard_simple_play_won(t *testing.T) {
	b, err := NewBingoBoard(`
24 75 59 41 17
58 74 64 92 39
68  8 78 85 72
18  3 22  4 34
11 76  6 28 50
`)
	assert.Nil(t, err)
	for _, n := range []int{24, 58, 68, 18} {
		b.play(n)
		assert.False(t, b.IsWinner)
	}
	b.play(11)
	assert.True(t, b.IsWinner)
}

func TestBingoBoard_play(t *testing.T) {
	type args struct {
		draws []int
	}
	boardData := `
24 75 59 41 17
58 74 64 92 39
68  8 78 85 72
18  3 22  4 34
11 76  6 28 50
`
	tests := []struct {
		name    string
		args    args
		wantWon bool
	}{
		{"win by first column", args{[]int{24, 58, 68, 18, 11}}, true},
		{"did not win by last column", args{[]int{24, 58, 68, 18}}, false},
		{"win by last column", args{[]int{17, 39, 72, 34, 50}}, true},
		{"did not win by last column", args{[]int{17, 39, 34, 50}}, false},
		{"win by first row", args{[]int{24, 75, 59, 41, 17}}, true},
		{"did not win by last row", args{[]int{75, 59, 41, 17}}, false},
		{"win by last row", args{[]int{11, 76, 6, 28, 50}}, true},
		{"did not win by last row", args{[]int{11, 76, 6, 28}}, false},
	}
	for _, tt := range tests {
		b, err := NewBingoBoard(boardData)
		assert.Nil(t, err)
		assert.NotNil(t, b)
		t.Run(tt.name, func(t *testing.T) {
			for _, n := range tt.args.draws {
				b.play(n)
			}
			assert.Equal(t, tt.wantWon, b.IsWinner)
		})
	}
}
