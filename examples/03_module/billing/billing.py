# Lean Code — BSD 3-Clause License — Vivian Voss, 2026
# Scope: Create and write invoices.


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
