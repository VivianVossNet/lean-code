// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read and check stock availability.

use std::fs;

pub fn read_stock_for_product(product: &str, stock_path: &str) -> i32 {
    let content = match fs::read_to_string(stock_path) {
        Ok(c) => c,
        Err(_) => return 0,
    };

    for line in content.trim().lines() {
        let parts: Vec<&str> = line.splitn(2, ',').collect();
        if parts[0] == product {
            return parts[1].parse().unwrap_or(0);
        }
    }

    0
}

pub fn check_stock_availability(product: &str, quantity: i32, stock_path: &str) -> bool {
    let stock = read_stock_for_product(product, stock_path);
    stock >= quantity
}
