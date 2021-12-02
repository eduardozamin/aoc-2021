#!/usr/bin/env python3

import numpy as np
from pathlib import Path

with open(Path(__file__).parent / "input-day01") as f:
    data = f.read()

all_numbers = [int(x) for x in data.split()]
arr1 = np.array(all_numbers[:-1])
arr2 = np.array(all_numbers[1:])

arr_diff = arr2 - arr1

results = [x for x in arr_diff if x > 0]

print(len(results))
