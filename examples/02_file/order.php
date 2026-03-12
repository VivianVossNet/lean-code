<?php
// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Process an order into an invoice.

function readOrderFromFile(string $path): array
{
    $lines = explode("\n", trim(file_get_contents($path)));

    return [
        'id'       => (int) $lines[0],
        'product'  => $lines[1],
        'quantity' => (int) $lines[2],
    ];
}

function readStockForProduct(string $product): int
{
    if (!file_exists('stock.txt')) {
        return 0;
    }

    $lines = explode("\n", trim(file_get_contents('stock.txt')));

    foreach ($lines as $line) {
        [$name, $count] = explode(',', $line, 2);
        if ($name === $product) {
            return (int) $count;
        }
    }

    return 0;
}

function checkStockAvailability(string $product, int $quantity): bool
{
    $stock = readStockForProduct($product);
    return $stock >= $quantity;
}

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

$order = readOrderFromFile('order.txt');

if (!checkStockAvailability($order['product'], $order['quantity'])) {
    fwrite(STDERR, "out of stock: {$order['product']}\n");
    exit(1);
}

$invoice = createInvoiceFromOrder($order, 29.99);
writeInvoiceToFile($invoice, 'invoice.txt');
