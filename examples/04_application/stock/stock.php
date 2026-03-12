<?php
// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read and check stock availability.

namespace Stock;

function readStockForProduct(string $product, string $stockPath): int
{
    if (!file_exists($stockPath)) {
        return 0;
    }

    foreach (explode("\n", trim(file_get_contents($stockPath))) as $line) {
        [$name, $count] = explode(',', $line, 2);
        if ($name === $product) {
            return (int) $count;
        }
    }

    return 0;
}

function checkStockAvailability(string $product, int $quantity, string $stockPath): bool
{
    $stock = readStockForProduct($product, $stockPath);
    return $stock >= $quantity;
}
