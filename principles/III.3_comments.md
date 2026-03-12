# III.3 Comments

**III.3.A** A file begins with a header: the licence and the scope of the file. This is the only comment that exists by convention.

**III.3.B** Everything else is self-documenting. If code needs a comment to be understood, a principle has been violated. Fix the name (II), fix the structure (I), or fix the scope (III). Do not fix it with a comment.

**Why:** A comment is an admission that the code does not speak for itself.

---

## Violation: comments as crutches

### JavaScript

```javascript
// Get order from file and parse it
function getOrder(p) {
  // Read the file
  const d = readFileSync(p, "utf-8");
  // Split by newlines
  const l = d.split("\n");
  // Return order object with parsed fields
  return {
    id: parseInt(l[0], 10),  // order ID
    product: l[1],            // product name
    quantity: parseInt(l[2], 10), // quantity ordered
  };
}
```

Seven comments. Every single one exists because a name failed. `p` needs a comment because it should be `path`. `d` needs a comment because it should be the return value of a clearly named operation. `getOrder` needs a comment because it should be `readOrderFromFile`.

### Python

```python
def proc(p, ppu):
    """Process an order: read it, check stock, create invoice, write it."""
    o = _rd(p)           # read order
    if not _chk(o):      # check stock
        return None
    inv = _mk(o, ppu)    # create invoice
    _wr(inv)             # write invoice
    return inv
```

A docstring and four inline comments to explain what four function calls do. The function names are the problem, not the missing documentation.

---

## Lean: the file header is the only comment

### JavaScript

```javascript
// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read orders from file.

import { readFileSync } from "node:fs";

function readOrderFromFile(path) {
  const lines = readFileSync(path, "utf-8").trim().split("\n");

  return {
    id: parseInt(lines[0], 10),
    product: lines[1],
    quantity: parseInt(lines[2], 10),
  };
}

export { readOrderFromFile };
```

### Go

```go
// Lean Code — BSD 3-Clause License — Vivian Voss, 2026
// Scope: Read orders from file.

package order

import (
    "os"
    "strconv"
    "strings"
)

type Order struct {
    ID       int64
    Product  string
    Quantity int
}

func ReadOrderFromFile(path string) Order {
    data, _ := os.ReadFile(path)
    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    id, _ := strconv.ParseInt(lines[0], 10, 64)
    quantity, _ := strconv.Atoi(lines[2])

    return Order{ID: id, Product: lines[1], Quantity: quantity}
}
```

### Python

```python
# Lean Code — BSD 3-Clause License — Vivian Voss, 2026
# Scope: Read orders from file.


def read_order_from_file(path):
    lines = open(path).read().strip().split("\n")

    return {
        "id": int(lines[0]),
        "product": lines[1],
        "quantity": int(lines[2]),
    }
```

Two lines at the top: licence and scope. Nothing else. The code explains itself because the names follow II.1, II.2, and II.3. If it did not, the answer would be to fix the names — not to add comments.
