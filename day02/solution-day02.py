#!/usr/bin/env python3

import numpy as np
from pathlib import Path
from pprint import pprint
from itertools import *


def main():
    # Load Data
    with open(Path(__file__).parent / "input-day02") as f:
        data = f.read().split()
    pairs = list(zip(data[0::2], [int(x) for x in data[1::2]]))

    # Part 1
    forward = sum([x[1] for x in pairs if x[0] == 'forward'])
    up = sum([-x[1] for x in pairs if x[0] == 'up'])
    down = sum([x[1] for x in pairs if x[0] == 'down'])
    depth = up + down
    print("---")
    print("Result Part 1:")
    print("Forward:", forward)
    print("Depth:", depth)
    print("Forward x Depth:", forward * depth)

    # Part 2
    def forward_delta(command):
        if command[0] == 'forward':
            return command[1]
        return 0

    def aim_delta(command):
        if command[0] == 'up':
            return -command[1]
        if command[0] == 'down':
            return command[1]
        return 0

    forward_deltas = [forward_delta(x) for x in pairs]
    aim_deltas = [aim_delta(x) for x in pairs]
    aim = [sum(aim_deltas[:i+1]) for i in range(len(aim_deltas))]
    depth_deltas = [f * a for (f, a) in zip(forward_deltas, aim)]

    print(len(forward_deltas))
    print(len(aim_deltas))
    print(len(depth_deltas))

    assert(sum(forward_deltas) == forward)
    print("---")
    print("Result Part 2:")
    print("Forward:", sum(forward_deltas))
    print("Depth:", sum(depth_deltas))
    print("Forward x Depth:", sum(forward_deltas) * sum(depth_deltas))


if __name__ == '__main__':
    main()
