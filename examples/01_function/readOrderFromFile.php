<?php
// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read an order from a text file.

function readOrderFromFile(string $path): array
{
    $lines = explode("\n", trim(file_get_contents($path)));

    return [
        'id'       => (int) $lines[0],
        'product'  => $lines[1],
        'quantity' => (int) $lines[2],
    ];
}
