# I.1 One unit, one job

**I.1.A** A function does one thing. A file has one responsibility. A module serves one purpose.

**Why:** Your head can follow one clear thing immediately.

---

## Violation

A function that does three jobs: read, validate, and write.

### JavaScript

```javascript
function processOrder(path) {
  const data = readFileSync(path, "utf-8");
  const lines = data.split("\n");
  const order = {
    id: parseInt(lines[0], 10),
    product: lines[1],
    quantity: parseInt(lines[2], 10),
  };

  if (!order.product || order.quantity <= 0) {
    throw new Error("invalid order");
  }

  const stock = readFileSync("stock.txt", "utf-8");
  for (const line of stock.split("\n")) {
    const [name, count] = line.split(",");
    if (name === order.product && parseInt(count, 10) >= order.quantity) {
      writeFileSync("invoice.txt", `${order.id}\n${order.product}\n${order.quantity}`);
      return;
    }
  }

  throw new Error("out of stock");
}
```

One function. Three jobs. Reading, checking, and writing are tangled. Change the stock format and the invoice breaks.

### Go

```go
func ProcessOrder(path string) error {
    data, _ := os.ReadFile(path)
    lines := strings.Split(string(data), "\n")
    id, _ := strconv.ParseInt(lines[0], 10, 64)
    quantity, _ := strconv.Atoi(lines[2])

    if lines[1] == "" || quantity <= 0 {
        return fmt.Errorf("invalid order")
    }

    stock, _ := os.ReadFile("stock.txt")
    for _, line := range strings.Split(string(stock), "\n") {
        parts := strings.SplitN(line, ",", 2)
        if parts[0] == lines[1] {
            count, _ := strconv.Atoi(parts[1])
            if count >= quantity {
                content := fmt.Sprintf("%d\n%s\n%d", id, lines[1], quantity)
                return os.WriteFile("invoice.txt", []byte(content), 0644)
            }
        }
    }

    return fmt.Errorf("out of stock")
}
```

---

## Lean

Each function does exactly one thing. The main flow reads like a sentence.

### JavaScript

```javascript
function readOrderFromFile(path) {
  const lines = readFileSync(path, "utf-8").split("\n");

  return {
    id: parseInt(lines[0], 10),
    product: lines[1],
    quantity: parseInt(lines[2], 10),
  };
}

function checkStockAvailability(product, quantity) {
  const stock = readStockForProduct(product);
  return stock >= quantity;
}

function writeInvoiceToFile(invoice, path) {
  const content = [invoice.orderId, invoice.product, invoice.quantity].join("\n");
  writeFileSync(path, content);
}

const order = readOrderFromFile("order.txt");

if (!checkStockAvailability(order.product, order.quantity)) {
  process.stderr.write(`out of stock: ${order.product}\n`);
  process.exit(1);
}

const invoice = createInvoiceFromOrder(order, 29.99);
writeInvoiceToFile(invoice, "invoice.txt");
```

### Go

```go
func ReadOrderFromFile(path string) Order {
    data, _ := os.ReadFile(path)
    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    id, _ := strconv.ParseInt(lines[0], 10, 64)
    quantity, _ := strconv.Atoi(lines[2])

    return Order{ID: id, Product: lines[1], Quantity: quantity}
}

func CheckStockAvailability(product string, quantity int) bool {
    stock := ReadStockForProduct(product)
    return stock >= quantity
}

func WriteInvoiceToFile(invoice Invoice, path string) {
    content := fmt.Sprintf("%d\n%s\n%d\n%.2f",
        invoice.OrderID, invoice.Product, invoice.Quantity, invoice.Total)
    os.WriteFile(path, []byte(content), 0644)
}

func main() {
    order := ReadOrderFromFile("order.txt")

    if !CheckStockAvailability(order.Product, order.Quantity) {
        fmt.Fprintf(os.Stderr, "out of stock: %s\n", order.Product)
        os.Exit(1)
    }

    invoice := CreateInvoiceFromOrder(order, 29.99)
    WriteInvoiceToFile(invoice, "invoice.txt")
}
```

### Python

```python
def read_order_from_file(path):
    lines = open(path).read().split("\n")

    return {
        "id": int(lines[0]),
        "product": lines[1],
        "quantity": int(lines[2]),
    }

def check_stock_availability(product, quantity):
    stock = read_stock_for_product(product)
    return stock >= quantity

def write_invoice_to_file(invoice, path):
    content = f"{invoice['order_id']}\n{invoice['product']}\n{invoice['quantity']}"
    open(path, "w").write(content)

order = read_order_from_file("order.txt")

if not check_stock_availability(order["product"], order["quantity"]):
    sys.exit(f"out of stock: {order['product']}")

invoice = create_invoice_from_order(order, 29.99)
write_invoice_to_file(invoice, "invoice.txt")
```

---

## The manifesto as its own example

This document follows I.1. Each category (I, II, III) has one job. Each topic (I.1, I.2, II.1...) has one job. Each rule (A, B, C) has one job. The pattern holds even when a topic has only one rule. An established pattern is maintained regardless of how many items fulfil it. The structure must always be immediately recognisable.
