# Lean Code — BSD 3-Clause License — Vivian Voss, 2026
# Scope: Read application configuration.


def read_configuration_from_file(path):
    configuration = {}

    for line in open(path).read().strip().split("\n"):
        key, value = line.split("=", 1)
        configuration[key.strip()] = value.strip()

    return configuration
