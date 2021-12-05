#!/usr/bin/env python3

import numpy as np
from pathlib import Path
from pprint import pprint

def parse_board(board_data):
    data = board_data.split("\n")
    bingo_boards = []
    bingo_numbers = []
    current_bingo_board = []
    for row in data:
        if not row.strip():
            continue
        if "," in row:
            if not bingo_numbers:
                bingo_numbers = [int(x) for x in row.strip().split(",")]
                continue
        current_bingo_board.append([[int(x), 0] for x in row.strip().split()])
        if len(current_bingo_board) == 5:
            bingo_boards.append(current_bingo_board)
            current_bingo_board = []

    return bingo_numbers, bingo_boards


def play_bingo(n, board):
    assert len(board), 5
    assert len(board[0]), 5
    for i in range(5):
        for j in range(5):
            if board[i][j][0] == n:
                board[i][j][1] = 1


def calculate_score(n, board):
    unmarked = []
    for i in range(5):
        for j in range(5):
            if board[i][j][1] == 0:
                unmarked.append(board[i][j][0])
    print("Sum of score:", sum(unmarked))
    return n * sum(unmarked)


def check_win(board):
    for i in range(5):
        check_row = board[i]
        check_col = [row[i] for row in board]
        if all(x[1] == 1 for x in check_row):
            return True
        if all(x[1] == 1 for x in check_col):
            return True
    return False


def solve_part1(data):
    b_numbers, b_boards = parse_board(data)
    for n in b_numbers:
        for board in b_boards:
            play_bingo(n, board)
            if check_win(board):
                return n, board


def solve_part2(data):
    b_numbers, b_boards = parse_board(data)
    won_boards = []
    for n in b_numbers:
        for board in b_boards:
            if board in won_boards:
                continue
            play_bingo(n, board)
            if check_win(board):
                won_boards.append(board)
                last_won_n = n
    return last_won_n, won_boards[-1]

def main():
    # Load Data
    with open(Path(__file__).parent / "input-day04") as f:
        data = f.read()

    print("----- Part 1 -----")
    part1_n, part1_board = solve_part1(data)
    pprint(part1_board)
    print("Part 1 winning number:", part1_n)
    print("Part 1 Score:", calculate_score(part1_n, part1_board))

    print("----- Part 2 -----")
    part2_n, part2_board = solve_part2(data)
    pprint(part2_board)
    print("Part 2 winning number:", part2_n)
    print("Part 2 Score:", calculate_score(part2_n, part2_board))



if __name__ == '__main__':
    main()
