// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Entry point. Orchestrate order to invoice.

mod order;
mod stock;
mod billing;

use std::process;

fn main() {
    let o = order::read_order_from_file("order.txt");

    if !stock::check_stock_availability(&o.product, o.quantity) {
        eprintln!("out of stock: {}", o.product);
        process::exit(1);
    }

    let invoice = billing::create_invoice_from_order(o.id, &o.product, o.quantity, 29.99);
    billing::write_invoice_to_file(&invoice, "invoice.txt");
}
