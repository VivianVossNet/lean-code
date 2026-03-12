# The Lean Code Manifesto

## Preamble

Clean Code solved a real problem in 2008. Java enterprise monoliths had 3,000-line methods and no structure. The prescription was genuine.

The industry copied it as universal law.

Lean Code is not the opposite of Clean Code. It is what Clean Code should have been: a minimal set of principles that scale from a single function to an entire system.

No line count rules. No extraction dogma. No ceremony.
KISS in your head, not rules on a wall.

---

## The Axiom

**Preserve the developer's cognitive capacity for logic.**

Cognitive capacity is finite. Every unnecessary interpretation, every context switch, every jump to another file is capacity that does not go into understanding the logic.

All principles in this manifesto serve one purpose: keep the developer's head free for the actual work.

---

## I. Structure

How code is cut.

### I.1 One unit, one job

**I.1.A** A function does one thing. A file has one responsibility. A module serves one purpose.

**Why:** Your head can follow one clear thing immediately.

### I.2 Scope defines size

**I.2.A** The scope of the job defines the size of the unit. Not a line count. Thirty readable lines for one job is lean. Four lines for a fragment ripped out of context is not.

**I.2.B** The question is never "how many lines?" but "how many jobs?"

**Why:** Your head loses context when it has to jump.

---

## II. Naming

How code is labelled.

### II.1 Verb and scope from a closed vocabulary

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

`readOrderFromDatabase`. `checkStockAvailability`. `createInvoiceFromOrder`. `writeInvoiceToFile`. `parseConfigFromYaml`. No comment needed. The name is the contract.

**Why:** One verb, one meaning, everywhere. No interpretation required.

### II.2 Nomenclature

**II.2.A** Spell out completely. `readConfiguration`, not `readConfig`. A name is a sentence. Sentences do not contain abbreviations.

**II.2.B** Choose the most naive word. `read`, not `fetch`. `check`, not `validate`. `write`, not `persist`. The simpler the word, the less interpretation it requires.

**II.2.C** No synonyms. No `handle` or `process` that hide what actually happens. One verb, one meaning, everywhere.

**Why:** Every translation in your head consumes capacity.

### II.3 Namespace carries scope, name carries noun

**II.3.A** Data structures carry the noun. The namespace carries the scope. `billing.Order`, not `BillingOrder`. `stock.Item`, not `StockItem`. The namespace does one job. The type name does another. Neither does the work of the other.

**II.3.B** Follow the naming conventions of the language you write in. `read_order_from_database` in Python. `readOrderFromDatabase` in TypeScript. `ReadOrderFromDatabase` in Go. The principle stays. The casing adapts.

**Why:** Each part of the name has one job. That is principle I.1 applied to naming.

---

## III. Scope

How code stays separated.

### III.1 Understood where it is read

**III.1.A** Code is read far more often than it is written. If you need to jump between eight files to understand one operation, the structure failed. If you need to hold a call graph in your head to follow the logic, the abstraction failed.

**III.1.B** Lean code is understood where it is read.

**Why:** Every context switch has cognitive cost.

### III.2 Same code, different scope, stays separate

**III.2.A** Two functions may contain identical code. If they serve different concerns that can change independently, they remain separate. Merging them couples two scopes. When one changes, the abstraction breaks.

**III.2.B** Duplication is not the enemy. Premature coupling is.

**Why:** Coupled concerns force simultaneous thinking.

### III.3 Comments

**III.3.A** A file begins with a header: the licence and the scope of the file. This is the only comment that exists by convention.

**III.3.B** Everything else is self-documenting. If code needs a comment to be understood, a principle has been violated. Fix the name (II), fix the structure (I), or fix the scope (III). Do not fix it with a comment.

**Why:** A comment is an admission that the code does not speak for itself.

---

## The Fractal

These principles are not seven separate rules. They are one rule applied recursively:

**Every unit does one thing, is sized by its scope, is named from a closed vocabulary, is readable in place, and is separated by concern.**

This holds for a function. It holds for a file. It holds for a module. It holds for a system.

When the codebase grows, the principles do not change. They extrapolate. The system scales because the rules are fractal, not because new rules are added.

---

## What Lean Code is not

**Not dirty code.** Lean code is structured, named, and intentional.

**Not an excuse for long functions.** A 500-line function doing twelve things violates I.1.

**Not anti-abstraction.** Abstractions are tools. Use them when they reduce complexity. Not when they merely relocate it.

**Not dogma.** If a situation demands breaking a principle, break it. Then document why.

---

## The test

Open any file. Read any function. If you understand what it does without jumping elsewhere, it is lean.

If you cannot, it is not.

---

*Vivian Voss, 2026*
