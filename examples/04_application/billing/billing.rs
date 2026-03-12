// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Create and write invoices.

use std::fs;

pub struct Invoice {
    pub order_id: i64,
    pub product: String,
    pub quantity: i32,
    pub total: f64,
}

pub fn create_invoice_from_order(order_id: i64, product: &str, quantity: i32, price_per_unit: f64) -> Invoice {
    Invoice {
        order_id,
        product: product.to_string(),
        quantity,
        total: quantity as f64 * price_per_unit,
    }
}

pub fn write_invoice_to_file(invoice: &Invoice, path: &str) {
    let content = format!(
        "{}\n{}\n{}\n{:.2}",
        invoice.order_id, invoice.product, invoice.quantity, invoice.total
    );

    fs::write(path, content).expect("failed to write invoice");
}
