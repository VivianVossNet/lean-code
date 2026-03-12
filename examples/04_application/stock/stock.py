# Lean Code — BSD 3-Clause License — Vivian Voss, 2026
# Scope: Read and check stock availability.


def read_stock_for_product(product, stock_path):
    try:
        lines = open(stock_path).read().strip().split("\n")
    except FileNotFoundError:
        return 0

    for line in lines:
        name, count = line.split(",", 1)
        if name == product:
            return int(count)

    return 0


def check_stock_availability(product, quantity, stock_path):
    stock = read_stock_for_product(product, stock_path)
    return stock >= quantity
