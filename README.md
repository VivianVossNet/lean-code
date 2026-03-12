# Lean Code

A minimal set of coding principles that scale fractally, from a single function to an entire system.

Lean Code is not the opposite of Clean Code. It is what Clean Code should have been.

## The Axiom

**Preserve the developer's cognitive capacity for logic.**

## The Principles

**I. Structure** — How code is cut.

- **I.1** One unit, one job.
- **I.2** Scope defines size.

**II. Naming** — How code is labelled.

- **II.1** Verb and scope from a closed vocabulary.
- **II.2** Nomenclature: spell out, choose the most naive word, no synonyms.
- **II.3** Namespace carries scope, name carries noun.

**III. Scope** — How code stays separated.

- **III.1** Understood where it is read.
- **III.2** Same code, different scope, stays separate.
- **III.3** Comments: file header only. Everything else is self-documenting.

Read the full manifesto: [manifesto.md](manifesto.md)

Translations: [Deutsch](i18n/de/manifesto.md) · [Español](i18n/es/manifesto.md) · [Français](i18n/fr/manifesto.md) · [Italiano](i18n/it/manifesto.md) · [Português](i18n/pt/manifesto.md) · [中文](i18n/zh/manifesto.md) · [日本語](i18n/ja/manifesto.md) · [Русский](i18n/ru/manifesto.md) · [العربية](i18n/ar/manifesto.md)

## Examples

The `examples/` directory demonstrates Lean Code at four fractal levels:

| Level | Directory | What it shows |
|-------|-----------|---------------|
| Function | `01_function/` | One function, one job. Naming, scope, size. |
| File | `02_file/` | Multiple functions in one file. Composition and flow. |
| Module | `03_module/` | Multiple files with namespaces. Separation of concerns. |
| Application | `04_application/` | Complete small app. The fractal at system level. |

Each level is implemented in five languages: JavaScript, Go, Python, Rust, and PHP.

All examples use the same domain — processing orders into invoices — so you can trace the same principles across languages and scales.

## Principles

The `principles/` directory contains one file per principle with code examples showing violations and lean alternatives.

## The Test

Open any file. Read any function. If you understand what it does without jumping elsewhere, it is lean. If you cannot, it is not.

## License

BSD 3-Clause — Vivian Voss, 2026
