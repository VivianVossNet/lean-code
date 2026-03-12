# El Manifiesto Lean Code

## Preámbulo

Clean Code resolvió un problema real en 2008. Los monolitos enterprise en Java tenían métodos de 3.000 líneas y ninguna estructura. La receta era genuina.

La industria la copió como ley universal.

Lean Code no es lo contrario de Clean Code. Es lo que Clean Code debería haber sido: un conjunto mínimo de principios que escalan desde una sola función hasta un sistema entero.

Sin reglas de conteo de líneas. Sin dogma de extracción. Sin ceremonia.
KISS en tu cabeza, no reglas en una pared.

---

## El Axioma

**Preservar la capacidad cognitiva del desarrollador para la lógica.**

La capacidad cognitiva es finita. Cada interpretación innecesaria, cada cambio de contexto, cada salto a otro archivo es capacidad que no se destina a comprender la lógica.

Todos los principios de este manifiesto sirven a un único propósito: mantener la cabeza del desarrollador libre para el trabajo real.

---

## I. Estructura

Cómo se divide el código.

### I.1 Una unidad, un trabajo

**I.1.A** Una función hace una cosa. Un archivo tiene una responsabilidad. Un módulo sirve a un propósito.

**Por qué:** Tu cabeza puede seguir una cosa clara de inmediato.

### I.2 El alcance define el tamaño

**I.2.A** El alcance del trabajo define el tamaño de la unidad. No un conteo de líneas. Treinta líneas legibles para un trabajo es lean. Cuatro líneas para un fragmento arrancado de su contexto no lo es.

**I.2.B** La pregunta nunca es "¿cuántas líneas?" sino "¿cuántos trabajos?"

**Por qué:** Tu cabeza pierde contexto cuando tiene que saltar.

---

## II. Nomenclatura

Cómo se etiqueta el código.

### II.1 Verbo y alcance de un vocabulario cerrado

**II.1.A** El nombre de una función sigue el patrón `{verb}{scopeInMaxThreeWords}`. El verbo proviene de un conjunto finito:

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

**II.1.B** El alcance describe el dominio en un máximo de tres palabras. Juntos, verbo y alcance se leen como una frase corta. Si necesitas más de tres palabras, el alcance es demasiado amplio — divide el trabajo.

`readOrderFromDatabase`. `checkStockAvailability`. `createInvoiceFromOrder`. `writeInvoiceToFile`. `parseConfigFromYaml`. Sin comentarios necesarios. El nombre es el contrato.

**Por qué:** Un verbo, un significado, en todas partes. Sin interpretación requerida.

### II.2 Nomenclatura

**II.2.A** Escribe completo. `readConfiguration`, no `readConfig`. Un nombre es una oración. Las oraciones no contienen abreviaturas.

**II.2.B** Elige la palabra más sencilla. `read`, no `fetch`. `check`, no `validate`. `write`, no `persist`. Cuanto más simple la palabra, menos interpretación requiere.

**II.2.C** Sin sinónimos. Sin `handle` o `process` que oculten lo que realmente sucede. Un verbo, un significado, en todas partes.

**Por qué:** Cada traducción en tu cabeza consume capacidad.

### II.3 El espacio de nombres lleva el alcance, el nombre lleva el sustantivo

**II.3.A** Las estructuras de datos llevan el sustantivo. El espacio de nombres lleva el alcance. `billing.Order`, no `BillingOrder`. `stock.Item`, no `StockItem`. El espacio de nombres hace un trabajo. El nombre de tipo hace otro. Ninguno hace el trabajo del otro.

**II.3.B** Sigue las convenciones de nomenclatura del lenguaje en el que escribes. `read_order_from_database` en Python. `readOrderFromDatabase` en TypeScript. `ReadOrderFromDatabase` en Go. El principio permanece. El estilo se adapta.

**Por qué:** Cada parte del nombre tiene un trabajo. Eso es el principio I.1 aplicado a la nomenclatura.

---

## III. Alcance

Cómo el código se mantiene separado.

### III.1 Se entiende donde se lee

**III.1.A** El código se lee mucho más de lo que se escribe. Si necesitas saltar entre ocho archivos para entender una operación, la estructura falló. Si necesitas mantener un grafo de llamadas en tu cabeza para seguir la lógica, la abstracción falló.

**III.1.B** El código lean se entiende donde se lee.

**Por qué:** Cada cambio de contexto tiene un coste cognitivo.

### III.2 Mismo código, distinto alcance, permanece separado

**III.2.A** Dos funciones pueden contener código idéntico. Si sirven a distintas responsabilidades que pueden cambiar de forma independiente, permanecen separadas. Fusionarlas acopla dos alcances. Cuando uno cambia, la abstracción se rompe.

**III.2.B** La duplicación no es el enemigo. El acoplamiento prematuro sí.

**Por qué:** Las responsabilidades acopladas obligan a pensar en todo a la vez.

### III.3 Comentarios

**III.3.A** Un archivo comienza con una cabecera: la licencia y el alcance del archivo. Este es el único comentario que existe por convención.

**III.3.B** Todo lo demás se autodocumenta. Si el código necesita un comentario para ser comprendido, se ha violado un principio. Corrige el nombre (II), corrige la estructura (I) o corrige el alcance (III). No lo corrijas con un comentario.

**Por qué:** Un comentario es una admisión de que el código no habla por sí mismo.

---

## El Fractal

Estos principios no son siete reglas separadas. Son una sola regla aplicada recursivamente:

**Cada unidad hace una cosa, su tamaño lo define su alcance, se nombra desde un vocabulario cerrado, se lee en su sitio y se separa por responsabilidad.**

Esto se cumple para una función. Se cumple para un archivo. Se cumple para un módulo. Se cumple para un sistema.

Cuando la base de código crece, los principios no cambian. Se extrapolan. El sistema escala porque las reglas son fractales, no porque se añadan reglas nuevas.

---

## Lo que Lean Code no es

**No es código sucio.** El código lean está estructurado, nombrado y es intencional.

**No es una excusa para funciones largas.** Una función de 500 líneas que hace doce cosas viola I.1.

**No es anti-abstracción.** Las abstracciones son herramientas. Úsalas cuando reducen complejidad. No cuando simplemente la reubican.

**No es dogma.** Si una situación exige romper un principio, rómpelo. Después documenta por qué.

---

## La prueba

Abre cualquier archivo. Lee cualquier función. Si entiendes lo que hace sin saltar a otro sitio, es lean.

Si no puedes, no lo es.

---

*Vivian Voss, 2026*
