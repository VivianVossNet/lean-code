# Lean Code — BSD 3-Clause License — Vivian Voss, 2026
# Scope: Read orders from file.


def read_order_from_file(path):
    lines = open(path).read().strip().split("\n")

    return {
        "id": int(lines[0]),
        "product": lines[1],
        "quantity": int(lines[2]),
    }
