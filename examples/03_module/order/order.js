// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read orders from file.

import { readFileSync } from "node:fs";

function readOrderFromFile(path) {
  const lines = readFileSync(path, "utf-8").trim().split("\n");

  return {
    id: parseInt(lines[0], 10),
    product: lines[1],
    quantity: parseInt(lines[2], 10),
  };
}

export { readOrderFromFile };
