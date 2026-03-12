<?php
// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read application configuration.

namespace Configuration;

function readConfigurationFromFile(string $path): array
{
    $configuration = [];

    foreach (explode("\n", trim(file_get_contents($path))) as $line) {
        [$key, $value] = explode('=', $line, 2);
        $configuration[trim($key)] = trim($value);
    }

    return $configuration;
}
