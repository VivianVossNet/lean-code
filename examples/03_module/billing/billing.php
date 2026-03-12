<?php
// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Create and write invoices.

namespace Billing;

function createInvoiceFromOrder(array $order, float $pricePerUnit): array
{
    return [
        'orderId'  => $order['id'],
        'product'  => $order['product'],
        'quantity' => $order['quantity'],
        'total'    => $order['quantity'] * $pricePerUnit,
    ];
}

function writeInvoiceToFile(array $invoice, string $path): void
{
    $content = implode("\n", [
        $invoice['orderId'],
        $invoice['product'],
        $invoice['quantity'],
        number_format($invoice['total'], 2, '.', ''),
    ]);

    file_put_contents($path, $content);
}
