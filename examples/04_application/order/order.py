# Lean Code — BSD 3-Clause License — Vivian Voss, 2026
# Scope: Read orders from file and find order files.

import os


def read_order_from_file(path):
    lines = open(path).read().strip().split("\n")

    return {
        "id": int(lines[0]),
        "product": lines[1],
        "quantity": int(lines[2]),
    }


def find_order_files_in_directory(directory):
    return [
        os.path.join(directory, f)
        for f in sorted(os.listdir(directory))
        if f.startswith("order_") and f.endswith(".txt")
    ]
