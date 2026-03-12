# Lean Code — BSD 3-Clause License — Vivian Voss, 2026
# Scope: Process an order into an invoice.

import sys


def read_order_from_file(path):
    lines = open(path).read().strip().split("\n")

    return {
        "id": int(lines[0]),
        "product": lines[1],
        "quantity": int(lines[2]),
    }


def read_stock_for_product(product):
    try:
        lines = open("stock.txt").read().strip().split("\n")
    except FileNotFoundError:
        return 0

    for line in lines:
        name, count = line.split(",", 1)
        if name == product:
            return int(count)

    return 0


def check_stock_availability(product, quantity):
    stock = read_stock_for_product(product)
    return stock >= quantity


def create_invoice_from_order(order, price_per_unit):
    return {
        "order_id": order["id"],
        "product": order["product"],
        "quantity": order["quantity"],
        "total": order["quantity"] * price_per_unit,
    }


def write_invoice_to_file(invoice, path):
    content = "\n".join([
        str(invoice["order_id"]),
        invoice["product"],
        str(invoice["quantity"]),
        f"{invoice['total']:.2f}",
    ])

    open(path, "w").write(content)


order = read_order_from_file("order.txt")

if not check_stock_availability(order["product"], order["quantity"]):
    sys.exit(f"out of stock: {order['product']}")

invoice = create_invoice_from_order(order, 29.99)
write_invoice_to_file(invoice, "invoice.txt")
