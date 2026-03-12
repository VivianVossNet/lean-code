# II.1 Verb and scope from a closed vocabulary

**II.1.A** A function name follows the pattern `{verb}{scopeInMaxThreeWords}`. The verb comes from a finite set:

| Verb | Meaning |
|------|---------|
| `read` | Retrieve data |
| `write` | Store data |
| `create` | Bring into existence |
| `delete` | Remove from existence |
| `update` | Modify existing data |
| `find` | Search by criteria |
| `check` | Verify a condition |
| `parse` | Transform a format |
| `render` | Produce output |

**II.1.B** The scope describes the domain in a maximum of three words. Together, verb and scope read like a short sentence. If you need more than three words, the scope is too broad — split the job.

**Why:** One verb, one meaning, everywhere. No interpretation required.

---

## Violation

Verbs that hide what actually happens.

### JavaScript

```javascript
function handleOrder(data) { /* read? write? check? */ }
function processPayment(info) { /* create? update? delete? */ }
function manageStock(items) { /* read? update? check? */ }
function doInvoice(order) { /* create? write? render? */ }
```

Every name requires opening the function to understand it. The name does not tell you what happens. It tells you something happens.

### Go

```go
func HandleOrder(data []byte) error { /* ... */ }
func ProcessPayment(info PaymentInfo) error { /* ... */ }
func ManageStock(items []Item) error { /* ... */ }
func DoInvoice(order Order) error { /* ... */ }
```

---

## Lean

Each name is a sentence. Verb tells you the operation. Scope tells you the domain.

### JavaScript

```javascript
function readOrderFromFile(path) { /* ... */ }
function checkStockAvailability(product, quantity) { /* ... */ }
function createInvoiceFromOrder(order, pricePerUnit) { /* ... */ }
function writeInvoiceToFile(invoice, path) { /* ... */ }
function parseConfigFromYaml(content) { /* ... */ }
function findOrderByCustomer(customerId) { /* ... */ }
function updateStockQuantity(product, delta) { /* ... */ }
function deleteExpiredInvoices(cutoffDate) { /* ... */ }
function renderInvoiceDocument(invoice) { /* ... */ }
```

### Go

```go
func ReadOrderFromFile(path string) Order { /* ... */ }
func CheckStockAvailability(product string, quantity int) bool { /* ... */ }
func CreateInvoiceFromOrder(order Order, pricePerUnit float64) Invoice { /* ... */ }
func WriteInvoiceToFile(invoice Invoice, path string) { /* ... */ }
func ParseConfigFromYaml(content string) Config { /* ... */ }
func FindOrderByCustomer(customerID int64) []Order { /* ... */ }
func UpdateStockQuantity(product string, delta int) { /* ... */ }
func DeleteExpiredInvoices(cutoff time.Time) { /* ... */ }
func RenderInvoiceDocument(invoice Invoice) string { /* ... */ }
```

### Python

```python
def read_order_from_file(path):              pass
def check_stock_availability(product, qty):  pass
def create_invoice_from_order(order, ppu):   pass
def write_invoice_to_file(invoice, path):    pass
def parse_config_from_yaml(content):         pass
def find_order_by_customer(customer_id):     pass
def update_stock_quantity(product, delta):    pass
def delete_expired_invoices(cutoff_date):    pass
def render_invoice_document(invoice):        pass
```

Nine verbs. Every function in the codebase uses one of them. A new developer reads any function name and knows what it does. No comment needed. The name is the contract.
