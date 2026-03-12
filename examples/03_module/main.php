<?php
// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Entry point. Orchestrate order to invoice.

require_once __DIR__ . '/order/order.php';
require_once __DIR__ . '/stock/stock.php';
require_once __DIR__ . '/billing/billing.php';

$order = Order\readOrderFromFile('order.txt');

if (!Stock\checkStockAvailability($order['product'], $order['quantity'])) {
    fwrite(STDERR, "out of stock: {$order['product']}\n");
    exit(1);
}

$invoice = Billing\createInvoiceFromOrder($order, 29.99);
Billing\writeInvoiceToFile($invoice, 'invoice.txt');
