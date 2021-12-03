#!/usr/bin/env python3

from pathlib import Path


def get_rating(data, mode):
    assert mode in ("o2", "co2")
    binary_length = len(data[0])
    pool = data
    for pos in range(binary_length):
        if len(pool) == 1:
            break
        break_even = len(pool) / 2
        bit_criteria = int(sum([int(x[pos]) for x in pool]) // break_even)
        if mode == "co2":
            bit_criteria = abs(bit_criteria - 1)
        new_pool = [x for x in pool if int(x[pos]) == bit_criteria]
        pool = new_pool
    assert len(pool) == 1
    return pool[0]


def main():
    # Load Data
    with open(Path(__file__).parent / "input-day03") as f:
        data = f.read().split()
    break_even = len(data) / 2
    binary_length = len(data[0])

    # Part 1
    γ_final = ''
    Ε_final = ''
    for pos in range(binary_length):
        sum_digits = sum([int(x[pos]) for x in data])
        γ = int(sum_digits // break_even)
        γ_final += str(γ)
        Ε_final += str(abs(γ - 1))
    print("Gamma (γ) rate:", γ_final, int(γ_final, 2))
    print("Epsilon (Ε) rate:", Ε_final, int(Ε_final, 2))
    print("Power Consumption:", int(γ_final, 2) * int(Ε_final, 2))

    # Part 2
    o2_rating = get_rating(data, "o2")
    co2_rating = get_rating(data, "co2")
    print("--------")
    print("Oxygen Rating:", o2_rating, int(o2_rating, 2))
    print("CO2 Rating:", co2_rating, int(co2_rating, 2))
    print("Life Support Rating:", int(o2_rating, 2) * int(co2_rating, 2))


if __name__ == '__main__':
    main()
