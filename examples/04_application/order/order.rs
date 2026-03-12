// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read orders from file and find order files.

use std::fs;

pub struct Order {
    pub id: i64,
    pub product: String,
    pub quantity: i32,
}

pub fn read_order_from_file(path: &str) -> Order {
    let content = fs::read_to_string(path).expect("order file not found");
    let lines: Vec<&str> = content.trim().lines().collect();

    Order {
        id: lines[0].parse().unwrap(),
        product: lines[1].to_string(),
        quantity: lines[2].parse().unwrap(),
    }
}

pub fn find_order_files_in_directory(directory: &str) -> Vec<String> {
    let mut files: Vec<String> = fs::read_dir(directory)
        .expect("order directory not found")
        .filter_map(|entry| {
            let path = entry.ok()?.path();
            let name = path.file_name()?.to_str()?.to_string();
            if name.starts_with("order_") && name.ends_with(".txt") {
                Some(path.to_str()?.to_string())
            } else {
                None
            }
        })
        .collect();

    files.sort();
    files
}
