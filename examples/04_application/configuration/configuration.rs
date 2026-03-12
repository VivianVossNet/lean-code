// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read application configuration.

use std::collections::HashMap;
use std::fs;

pub struct Configuration {
    pub price_per_unit: f64,
    pub stock_file: String,
    pub order_directory: String,
}

pub fn read_configuration_from_file(path: &str) -> Configuration {
    let content = fs::read_to_string(path).expect("configuration file not found");
    let mut values = HashMap::new();

    for line in content.trim().lines() {
        let parts: Vec<&str> = line.splitn(2, '=').collect();
        values.insert(parts[0].trim().to_string(), parts[1].trim().to_string());
    }

    Configuration {
        price_per_unit: values["price_per_unit"].parse().unwrap(),
        stock_file: values["stock_file"].clone(),
        order_directory: values["order_directory"].clone(),
    }
}
