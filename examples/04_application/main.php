<?php
// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Entry point. Process all orders into invoices.

require_once __DIR__ . '/configuration/configuration.php';
require_once __DIR__ . '/order/order.php';
require_once __DIR__ . '/stock/stock.php';
require_once __DIR__ . '/billing/billing.php';

$configuration = Configuration\readConfigurationFromFile('config.txt');
$pricePerUnit = (float) $configuration['price_per_unit'];

$orderFiles = Order\findOrderFilesInDirectory($configuration['order_directory']);

foreach ($orderFiles as $file) {
    $order = Order\readOrderFromFile($file);

    if (!Stock\checkStockAvailability($order['product'], $order['quantity'], $configuration['stock_file'])) {
        fwrite(STDERR, "out of stock: {$order['product']}\n");
        continue;
    }

    $invoice = Billing\createInvoiceFromOrder($order, $pricePerUnit);
    Billing\writeInvoiceToFile($invoice, "invoice_{$order['id']}.txt");
}
