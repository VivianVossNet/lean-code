// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read application configuration.

import { readFileSync } from "node:fs";

function readConfigurationFromFile(path) {
  const lines = readFileSync(path, "utf-8").trim().split("\n");
  const configuration = {};

  for (const line of lines) {
    const [key, value] = line.split("=", 2);
    configuration[key.trim()] = value.trim();
  }

  return configuration;
}

export { readConfigurationFromFile };
