# Lean Code — BSD 3-Clause License — Vivian Voss, 2026
# Scope: Read and check stock availability.


def read_stock_for_product(product):
    try:
        lines = open("stock.txt").read().strip().split("\n")
    except FileNotFoundError:
        return 0

    for line in lines:
        name, count = line.split(",", 1)
        if name == product:
            return int(count)

    return 0


def check_stock_availability(product, quantity):
    stock = read_stock_for_product(product)
    return stock >= quantity
