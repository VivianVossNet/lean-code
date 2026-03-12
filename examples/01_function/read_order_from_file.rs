// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read an order from a text file.

use std::fs;

struct Order {
    id: i64,
    product: String,
    quantity: i32,
}

fn read_order_from_file(path: &str) -> Order {
    let content = fs::read_to_string(path).expect("order file not found");
    let lines: Vec<&str> = content.trim().lines().collect();

    Order {
        id: lines[0].parse().unwrap(),
        product: lines[1].to_string(),
        quantity: lines[2].parse().unwrap(),
    }
}
