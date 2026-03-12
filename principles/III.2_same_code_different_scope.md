# III.2 Same code, different scope, stays separate

**III.2.A** Two functions may contain identical code. If they serve different concerns that can change independently, they remain separate. Merging them couples two scopes. When one changes, the abstraction breaks.

**III.2.B** Duplication is not the enemy. Premature coupling is.

**Why:** Coupled concerns force simultaneous thinking.

---

## Violation: context parameter

Two concerns merged because the code looked similar.

### JavaScript

```javascript
function formatAmount(amount, context) {
  if (context === "invoice") {
    return `${amount.toFixed(2)} EUR`;
  }

  if (context === "report") {
    return `${amount.toFixed(2)} EUR`;
  }

  return amount.toFixed(2);
}
```

Today the formatting is identical. Tomorrow the report needs a thousands separator. Or the invoice needs a tax line. The `context` parameter couples two scopes. When one changes, the other must be tested.

### Go

```go
func FormatAmount(amount float64, context string) string {
    switch context {
    case "invoice":
        return fmt.Sprintf("%.2f EUR", amount)
    case "report":
        return fmt.Sprintf("%.2f EUR", amount)
    default:
        return fmt.Sprintf("%.2f", amount)
    }
}
```

---

## Violation: shared helper with flags

### Python

```python
def format_amount(amount, include_currency=True, thousands_separator=False):
    result = f"{amount:.2f}"
    if thousands_separator:
        parts = result.split(".")
        parts[0] = f"{int(parts[0]):,}"
        result = ".".join(parts)
    if include_currency:
        result += " EUR"
    return result
```

A function that started as two lines. Flags were added to serve both callers. Now neither caller understands it without reading the flags. The abstraction serves no one.

---

## Lean

Two functions. Same code. Different scopes. Each changes independently.

### JavaScript

```javascript
function renderInvoiceAmount(amount) {
  return `${amount.toFixed(2)} EUR`;
}

function renderReportAmount(amount) {
  return `${amount.toFixed(2)} EUR`;
}
```

### Go

```go
func RenderInvoiceAmount(amount float64) string {
    return fmt.Sprintf("%.2f EUR", amount)
}

func RenderReportAmount(amount float64) string {
    return fmt.Sprintf("%.2f EUR", amount)
}
```

### Python

```python
def render_invoice_amount(amount):
    return f"{amount:.2f} EUR"

def render_report_amount(amount):
    return f"{amount:.2f} EUR"
```

### Rust

```rust
fn render_invoice_amount(amount: f64) -> String {
    format!("{:.2} EUR", amount)
}

fn render_report_amount(amount: f64) -> String {
    format!("{:.2} EUR", amount)
}
```

### PHP

```php
function renderInvoiceAmount(float $amount): string {
    return number_format($amount, 2, '.', '') . ' EUR';
}

function renderReportAmount(float $amount): string {
    return number_format($amount, 2, '.', '') . ' EUR';
}
```

Identical today. When the report changes, the invoice is not touched. When the invoice changes, the report is not tested. No coupling. No flags. No `context` parameter.

A linter will flag this as duplication. The linter is wrong. These are two scopes that happen to look the same today. The cost of separation is two identical functions. The cost of coupling is a broken abstraction when requirements diverge.
