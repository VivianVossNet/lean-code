<?php
// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read and check stock availability.

namespace Stock;

function readStockForProduct(string $product): int
{
    if (!file_exists('stock.txt')) {
        return 0;
    }

    $lines = explode("\n", trim(file_get_contents('stock.txt')));

    foreach ($lines as $line) {
        [$name, $count] = explode(',', $line, 2);
        if ($name === $product) {
            return (int) $count;
        }
    }

    return 0;
}

function checkStockAvailability(string $product, int $quantity): bool
{
    $stock = readStockForProduct($product);
    return $stock >= $quantity;
}
