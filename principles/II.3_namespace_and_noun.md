# II.3 Namespace carries scope, name carries noun

**II.3.A** Data structures carry the noun. The namespace carries the scope. `billing.Order`, not `BillingOrder`. `stock.Item`, not `StockItem`. The namespace does one job. The type name does another. Neither does the work of the other.

**II.3.B** Follow the naming conventions of the language you write in. `read_order_from_database` in Python. `readOrderFromDatabase` in TypeScript. `ReadOrderFromDatabase` in Go. The principle stays. The casing adapts.

**Why:** Each part of the name has one job. That is principle I.1 applied to naming.

---

## Violation: scope baked into the name

### Java

```java
public class BillingOrder { /* ... */ }
public class BillingInvoice { /* ... */ }
public class BillingCustomer { /* ... */ }
public class StockItem { /* ... */ }
public class StockWarehouse { /* ... */ }
```

Every class name repeats the scope. `BillingOrder` says "billing" twice: once in the package, once in the name. That is redundancy. Redundancy is noise. Noise consumes capacity.

### Go

```go
package billing

type BillingOrder struct { /* ... */ }
type BillingInvoice struct { /* ... */ }
```

`billing.BillingOrder`. The stutter is the symptom. The violation is that the name does the namespace's job.

---

## Violation: framework suffixes

### Java

```java
public class OrderDTO { /* ... */ }
public class OrderEntity { /* ... */ }
public class OrderModel { /* ... */ }
public class OrderVO { /* ... */ }
```

Four classes. One concept. The suffixes describe the technical layer, not the domain. `DTO`, `Entity`, `Model`, `VO` are framework vocabulary, not domain vocabulary. The reader must know the framework to understand the name.

### TypeScript

```typescript
interface OrderInterface { /* ... */ }
type OrderType = { /* ... */ };
class OrderModel { /* ... */ }
```

The suffix tells you what the language construct is. The language already tells you that.

---

## Lean

### Go

```go
package billing

type Order struct {
    ID       int64
    Product  string
    Quantity int
}

type Invoice struct {
    OrderID  int64
    Product  string
    Quantity int
    Total    float64
}
```

`billing.Order`. `billing.Invoice`. The namespace carries the scope. The name carries the noun. Each does one job.

### Python

```python
# billing/order.py

class Order:
    def __init__(self, id, product, quantity):
        self.id = id
        self.product = product
        self.quantity = quantity

class Invoice:
    def __init__(self, order_id, product, quantity, total):
        self.order_id = order_id
        self.product = product
        self.quantity = quantity
        self.total = total
```

`billing.Order`. Not `BillingOrderModel`.

### PHP

```php
namespace Billing;

class Order {
    public function __construct(
        public readonly int $id,
        public readonly string $product,
        public readonly int $quantity,
    ) {}
}

class Invoice {
    public function __construct(
        public readonly int $orderId,
        public readonly string $product,
        public readonly int $quantity,
        public readonly float $total,
    ) {}
}
```

`Billing\Order`. Not `Billing\BillingOrderEntity`.

### Rust

```rust
mod billing {
    pub struct Order {
        pub id: i64,
        pub product: String,
        pub quantity: i32,
    }

    pub struct Invoice {
        pub order_id: i64,
        pub product: String,
        pub quantity: i32,
        pub total: f64,
    }
}
```

`billing::Order`. Not `BillingOrder`.

### C

C has no namespace construct. The scope becomes a prefix. This is a language constraint, not a principle violation.

```c
typedef struct {
    int64_t id;
    char product[128];
    int quantity;
} billing_order;

typedef struct {
    int64_t order_id;
    char product[128];
    int quantity;
    double total;
} billing_invoice;
```

`billing_order`. The prefix is the namespace. The language demands it. The principle adapts.
