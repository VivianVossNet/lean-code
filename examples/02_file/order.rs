// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Process an order into an invoice.

use std::fs;
use std::process;

struct Order {
    id: i64,
    product: String,
    quantity: i32,
}

struct Invoice {
    order_id: i64,
    product: String,
    quantity: i32,
    total: f64,
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

fn read_stock_for_product(product: &str) -> i32 {
    let content = match fs::read_to_string("stock.txt") {
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

fn check_stock_availability(product: &str, quantity: i32) -> bool {
    let stock = read_stock_for_product(product);
    stock >= quantity
}

fn create_invoice_from_order(order: &Order, price_per_unit: f64) -> Invoice {
    Invoice {
        order_id: order.id,
        product: order.product.clone(),
        quantity: order.quantity,
        total: order.quantity as f64 * price_per_unit,
    }
}

fn write_invoice_to_file(invoice: &Invoice, path: &str) {
    let content = format!(
        "{}\n{}\n{}\n{:.2}",
        invoice.order_id, invoice.product, invoice.quantity, invoice.total
    );

    fs::write(path, content).expect("failed to write invoice");
}

fn main() {
    let order = read_order_from_file("order.txt");

    if !check_stock_availability(&order.product, order.quantity) {
        eprintln!("out of stock: {}", order.product);
        process::exit(1);
    }

    let invoice = create_invoice_from_order(&order, 29.99);
    write_invoice_to_file(&invoice, "invoice.txt");
}
