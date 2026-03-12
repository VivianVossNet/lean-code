// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Entry point. Process all orders into invoices.

mod configuration;
mod order;
mod stock;
mod billing;

fn main() {
    let config = configuration::read_configuration_from_file("config.txt");

    let order_files = order::find_order_files_in_directory(&config.order_directory);

    for file in &order_files {
        let o = order::read_order_from_file(file);

        if !stock::check_stock_availability(&o.product, o.quantity, &config.stock_file) {
            eprintln!("out of stock: {}", o.product);
            continue;
        }

        let invoice = billing::create_invoice_from_order(o.id, &o.product, o.quantity, config.price_per_unit);
        billing::write_invoice_to_file(&invoice, &format!("invoice_{}.txt", o.id));
    }
}
