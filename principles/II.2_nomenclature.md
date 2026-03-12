# II.2 Nomenclature

**II.2.A** Spell out completely. `readConfiguration`, not `readConfig`. A name is a sentence. Sentences do not contain abbreviations.

**II.2.B** Choose the most naive word. `read`, not `fetch`. `check`, not `validate`. `write`, not `persist`. The simpler the word, the less interpretation it requires.

**II.2.C** No synonyms. No `handle` or `process` that hide what actually happens. One verb, one meaning, everywhere.

**Why:** Every translation in your head consumes capacity.

---

## Violation: abbreviations

### Python

```python
def rd_ord(p):
    pass

def chk_stk(prod):
    pass

def mk_inv(o, ppu):
    pass

def wr_inv(inv, p):
    pass
```

Every abbreviation requires translation. `rd_ord` could be `read_order`, `redirect_ordinary`, or anything else. The reader must guess. Guessing is interpretation. Interpretation consumes capacity.

### JavaScript

```javascript
function rdOrdFromFile(p) { /* ... */ }
function chkStkAvail(prod, qty) { /* ... */ }
function mkInvFromOrd(ord, ppu) { /* ... */ }
function wrInvToFile(inv, p) { /* ... */ }
```

---

## Violation: fancy synonyms

### TypeScript

```typescript
function fetchOrder(path: string) { /* ... */ }
function validateStockAvailability(product: string) { /* ... */ }
function persistInvoice(invoice: Invoice) { /* ... */ }
function destroyExpiredOrders() { /* ... */ }
function acquireConfiguration() { /* ... */ }
```

`fetch` instead of `read`. `validate` instead of `check`. `persist` instead of `write`. `destroy` instead of `delete`. `acquire` instead of `read`. Each synonym forces the reader to map it back to the simple concept.

The naive word eliminates that step. `read` is understood without thinking. `acquire` requires a mental detour.

### Go

```go
func FetchOrder(path string) Order { /* ... */ }
func ValidateStockAvailability(product string) bool { /* ... */ }
func PersistInvoice(invoice Invoice) { /* ... */ }
func DestroyExpiredOrders() { /* ... */ }
func AcquireConfiguration() Config { /* ... */ }
```

---

## Violation: synonyms across the codebase

### JavaScript

```javascript
function readOrder(path) { /* ... */ }
function fetchStock(product) { /* ... */ }
function loadConfiguration(path) { /* ... */ }
function getUser(id) { /* ... */ }
function retrieveInvoice(id) { /* ... */ }
```

Five functions. Five different verbs. All do the same thing: retrieve data. The reader must remember that `read`, `fetch`, `load`, `get`, and `retrieve` mean the same operation. That is five translations for one concept.

---

## Lean

### JavaScript

```javascript
function readOrderFromFile(path) { /* ... */ }
function readStockForProduct(product) { /* ... */ }
function readConfigurationFromFile(path) { /* ... */ }
function readUserFromDatabase(id) { /* ... */ }
function readInvoiceFromDatabase(id) { /* ... */ }
```

One verb for one operation. Everywhere. No translation required.

### Python

```python
def read_order_from_file(path):          pass
def read_stock_for_product(product):     pass
def read_configuration_from_file(path):  pass
def read_user_from_database(user_id):    pass
def read_invoice_from_database(invoice_id): pass
```

### Go

```go
func ReadOrderFromFile(path string) Order { /* ... */ }
func ReadStockForProduct(product string) int { /* ... */ }
func ReadConfigurationFromFile(path string) Config { /* ... */ }
func ReadUserFromDatabase(id int64) User { /* ... */ }
func ReadInvoiceFromDatabase(id int64) Invoice { /* ... */ }
```

Same verb. Same meaning. Zero cognitive overhead.
