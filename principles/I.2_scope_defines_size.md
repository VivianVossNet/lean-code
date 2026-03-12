# I.2 Scope defines size

**I.2.A** The scope of the job defines the size of the unit. Not a line count. Thirty readable lines for one job is lean. Four lines for a fragment ripped out of context is not.

**I.2.B** The question is never "how many lines?" but "how many jobs?"

**Why:** Your head loses context when it has to jump.

---

## Violation

A function split into fragments that mean nothing alone. Clean Code says: extract. Lean Code says: why?

### JavaScript

```javascript
function getLines(path) {
  return readFileSync(path, "utf-8").split("\n");
}

function parseId(lines) {
  return parseInt(lines[0], 10);
}

function parseProduct(lines) {
  return lines[1];
}

function parseQuantity(lines) {
  return parseInt(lines[2], 10);
}

function buildOrder(lines) {
  return {
    id: parseId(lines),
    product: parseProduct(lines),
    quantity: parseQuantity(lines),
  };
}

function readOrderFromFile(path) {
  const lines = getLines(path);
  return buildOrder(lines);
}
```

Six functions. One job. Five of them are noise. `parseId` does not exist as a concept. It is a line of code pretending to be a function.

### Go

```go
func getLines(path string) []string {
    data, _ := os.ReadFile(path)
    return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func parseID(lines []string) int64 {
    id, _ := strconv.ParseInt(lines[0], 10, 64)
    return id
}

func parseProduct(lines []string) string {
    return lines[1]
}

func parseQuantity(lines []string) int {
    q, _ := strconv.Atoi(lines[2])
    return q
}

func buildOrder(lines []string) Order {
    return Order{ID: parseID(lines), Product: parseProduct(lines), Quantity: parseQuantity(lines)}
}

func ReadOrderFromFile(path string) Order {
    lines := getLines(path)
    return buildOrder(lines)
}
```

---

## Lean

One function. One job. The size follows from the scope.

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
```

### Rust

```rust
fn read_order_from_file(path: &str) -> Order {
    let content = fs::read_to_string(path).expect("order file not found");
    let lines: Vec<&str> = content.trim().lines().collect();

    Order {
        id: lines[0].parse().unwrap(),
        product: lines[1].to_string(),
        quantity: lines[2].parse().unwrap(),
    }
}
```

Ten lines. One job. Nothing to extract. The scope defines the size.
