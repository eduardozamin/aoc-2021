#!/usr/bin/env python3

import numpy as np
from pathlib import Path

def solve_part1(all_numbers):
    arr1 = np.array(all_numbers[:-1])
    arr2 = np.array(all_numbers[1:])
    arr_diff = arr2 - arr1
    return len([x for x in arr_diff if x > 0])

def main():
    with open(Path(__file__).parent / "input-day01") as f:
        data = f.read()
    all_numbers = [int(x) for x in data.split()]

    # Part 1
    result_part1 = solve_part1(all_numbers)
    print("Part 1:", result_part1)

    # Part 2
    arr1 = np.array(all_numbers[0:-2])
    arr2 = np.array(all_numbers[1:-1])
    arr3 = np.array(all_numbers[2:])
    arr_sum = arr1 + arr2 + arr3
    result_part2 = solve_part1(arr_sum)
    print("Part 2:", result_part2)



if __name__ == '__main__':
    main()
