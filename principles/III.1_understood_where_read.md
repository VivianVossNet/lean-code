# III.1 Understood where it is read

**III.1.A** Code is read far more often than it is written. If you need to jump between eight files to understand one operation, the structure failed. If you need to hold a call graph in your head to follow the logic, the abstraction failed.

**III.1.B** Lean code is understood where it is read.

**Why:** Every context switch has cognitive cost.

---

## Violation

An order flow scattered across layers, files, and indirections.

### Java

```java
// OrderController.java
public class OrderController {
    @Inject private OrderService orderService;

    public Response handle(Request request) {
        return orderService.process(request.getOrderId());
    }
}

// OrderService.java
public class OrderService {
    @Inject private OrderRepository orderRepository;
    @Inject private StockValidator stockValidator;
    @Inject private InvoiceFactory invoiceFactory;
    @Inject private InvoiceWriter invoiceWriter;

    public Response process(long orderId) {
        Order order = orderRepository.findById(orderId);
        stockValidator.validate(order);
        Invoice invoice = invoiceFactory.create(order);
        invoiceWriter.write(invoice);
        return Response.ok();
    }
}

// StockValidator.java
public class StockValidator {
    @Inject private StockRepository stockRepository;

    public void validate(Order order) {
        int stock = stockRepository.findByProduct(order.getProduct());
        if (stock < order.getQuantity()) {
            throw new OutOfStockException(order.getProduct());
        }
    }
}

// InvoiceFactory.java
public class InvoiceFactory {
    public Invoice create(Order order) {
        return new Invoice(order.getId(), order.getProduct(),
            order.getQuantity(), order.getQuantity() * 29.99);
    }
}

// InvoiceWriter.java — yet another file for three lines of logic
```

Five files. Five classes. One operation. To understand what happens when an order arrives, you open five files and hold the call graph in your head. Every jump costs cognitive capacity.

### TypeScript

```typescript
// orderController.ts
export function handleOrder(orderId: number) {
  return orderService.process(orderId);
}

// orderService.ts
export function process(orderId: number) {
  const order = orderRepository.findById(orderId);
  stockValidator.validate(order);
  const invoice = invoiceFactory.create(order);
  invoiceWriter.write(invoice);
}

// stockValidator.ts — one function, one file
// invoiceFactory.ts — one function, one file
// invoiceWriter.ts — one function, one file
```

Each file wraps a single function in a module. The file is not the unit of work. The function is. The files are ceremony.

---

## Lean

One file. The operation is understood where it is read.

### JavaScript

```javascript
import { readFileSync, writeFileSync } from "node:fs";

function readOrderFromFile(path) {
  const lines = readFileSync(path, "utf-8").split("\n");

  return {
    id: parseInt(lines[0], 10),
    product: lines[1],
    quantity: parseInt(lines[2], 10),
  };
}

function readStockForProduct(product) {
  const lines = readFileSync("stock.txt", "utf-8").split("\n");

  for (const line of lines) {
    const [name, count] = line.split(",", 2);
    if (name === product) {
      return parseInt(count, 10);
    }
  }

  return 0;
}

function checkStockAvailability(product, quantity) {
  const stock = readStockForProduct(product);
  return stock >= quantity;
}

function createInvoiceFromOrder(order, pricePerUnit) {
  return {
    orderId: order.id,
    product: order.product,
    quantity: order.quantity,
    total: order.quantity * pricePerUnit,
  };
}

function writeInvoiceToFile(invoice, path) {
  const content = [
    invoice.orderId,
    invoice.product,
    invoice.quantity,
    invoice.total.toFixed(2),
  ].join("\n");

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

### Java

```java
package billing;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;

public class Order {

    record OrderData(long id, String product, int quantity) {}
    record Invoice(long orderId, String product, int quantity, double total) {}

    static OrderData readOrderFromFile(String path) throws Exception {
        List<String> lines = Files.readAllLines(Path.of(path));

        return new OrderData(
            Long.parseLong(lines.get(0)),
            lines.get(1),
            Integer.parseInt(lines.get(2))
        );
    }

    static int readStockForProduct(String product) throws Exception {
        if (!Files.exists(Path.of("stock.txt"))) return 0;

        for (String line : Files.readAllLines(Path.of("stock.txt"))) {
            String[] parts = line.split(",", 2);
            if (parts[0].equals(product)) {
                return Integer.parseInt(parts[1]);
            }
        }

        return 0;
    }

    static boolean checkStockAvailability(String product, int quantity) throws Exception {
        int stock = readStockForProduct(product);
        return stock >= quantity;
    }

    static Invoice createInvoiceFromOrder(OrderData order, double pricePerUnit) {
        return new Invoice(
            order.id(), order.product(), order.quantity(),
            order.quantity() * pricePerUnit
        );
    }

    static void writeInvoiceToFile(Invoice invoice, String path) throws Exception {
        String content = String.format("%d\n%s\n%d\n%.2f",
            invoice.orderId(), invoice.product(), invoice.quantity(), invoice.total());
        Files.writeString(Path.of(path), content);
    }

    public static void main(String[] args) throws Exception {
        OrderData order = readOrderFromFile("order.txt");

        if (!checkStockAvailability(order.product(), order.quantity())) {
            System.err.println("out of stock: " + order.product());
            System.exit(1);
        }

        Invoice invoice = createInvoiceFromOrder(order, 29.99);
        writeInvoiceToFile(invoice, "invoice.txt");
    }
}
```

One file. Five functions. Each function does one thing. The main flow reads top to bottom. No jumps. No indirection. Understood where it is read.
