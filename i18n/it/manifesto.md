# Il Manifesto del Lean Code

## Premessa

Clean Code risolveva un problema reale nel 2008. I monoliti enterprise in Java avevano metodi da 3.000 righe e nessuna struttura. La ricetta era genuina.

L'industria l'ha copiata come legge universale.

Lean Code non è il contrario di Clean Code. È ciò che Clean Code avrebbe dovuto essere: un insieme minimale di principi che scala da una singola funzione a un intero sistema.

Nessuna regola sul numero di righe. Nessun dogma sull'estrazione. Nessuna cerimonia.
KISS nella testa, non regole appese al muro.

---

## L'Assioma

**Preserva la capacità cognitiva dello sviluppatore per la logica.**

La capacità cognitiva è finita. Ogni interpretazione superflua, ogni cambio di contesto, ogni salto a un altro file è capacità che non va nella comprensione della logica.

Tutti i principi di questo manifesto servono un unico scopo: tenere la testa dello sviluppatore libera per il lavoro vero.

---

## I. Struttura

Come il codice viene suddiviso.

### I.1 Un'unità, un compito

**I.1.A** Una funzione fa una cosa. Un file ha una responsabilità. Un modulo serve uno scopo.

**Perché:** La tua testa riesce a seguire immediatamente una cosa chiara.

### I.2 Lo scope definisce la dimensione

**I.2.A** Lo scope del compito definisce la dimensione dell'unità. Non un conteggio di righe. Trenta righe leggibili per un compito è lean. Quattro righe per un frammento strappato dal contesto non lo è.

**I.2.B** La domanda non è mai "quante righe?" ma "quanti compiti?"

**Perché:** La tua testa perde il contesto quando deve saltare.

---

## II. Naming

Come il codice viene etichettato.

### II.1 Verbo e scope da un vocabolario chiuso

**II.1.A** Il nome di una funzione segue il pattern `{verb}{scopeInMaxThreeWords}`. Il verbo proviene da un insieme finito:

| Verb | Meaning |
|------|---------|
| `read` | Recuperare dati |
| `write` | Salvare dati |
| `create` | Portare in esistenza |
| `delete` | Rimuovere dall'esistenza |
| `update` | Modificare dati esistenti |
| `find` | Cercare per criteri |
| `check` | Verificare una condizione |
| `parse` | Trasformare un formato |
| `render` | Produrre output |

**II.1.B** Lo scope descrive il dominio in un massimo di tre parole. Insieme, verbo e scope si leggono come una frase breve. Se servono più di tre parole, lo scope è troppo ampio — dividi il compito.

`readOrderFromDatabase`. `checkStockAvailability`. `createInvoiceFromOrder`. `writeInvoiceToFile`. `parseConfigFromYaml`. Nessun commento necessario. Il nome è il contratto.

**Perché:** Un verbo, un significato, ovunque. Nessuna interpretazione richiesta.

### II.2 Nomenclatura

**II.2.A** Scrivi per esteso. `readConfiguration`, non `readConfig`. Un nome è una frase. Le frasi non contengono abbreviazioni.

**II.2.B** Scegli la parola più semplice. `read`, non `fetch`. `check`, non `validate`. `write`, non `persist`. Più la parola è semplice, meno interpretazione richiede.

**II.2.C** Niente sinonimi. Niente `handle` o `process` che nascondono cosa succede davvero. Un verbo, un significato, ovunque.

**Perché:** Ogni traduzione nella tua testa consuma capacità.

### II.3 Il namespace porta lo scope, il nome porta il sostantivo

**II.3.A** Le strutture dati portano il sostantivo. Il namespace porta lo scope. `billing.Order`, non `BillingOrder`. `stock.Item`, non `StockItem`. Il namespace fa un compito. Il nome del tipo ne fa un altro. Nessuno dei due fa il lavoro dell'altro.

**II.3.B** Segui le convenzioni di naming del linguaggio in cui scrivi. `read_order_from_database` in Python. `readOrderFromDatabase` in TypeScript. `ReadOrderFromDatabase` in Go. Il principio resta. Il casing si adatta.

**Perché:** Ogni parte del nome ha un compito. Questo è il principio I.1 applicato al naming.

---

## III. Scope

Come il codice resta separato.

### III.1 Compreso dove si legge

**III.1.A** Il codice viene letto molto più spesso di quanto venga scritto. Se devi saltare tra otto file per capire un'operazione, la struttura ha fallito. Se devi tenere un grafo di chiamate in testa per seguire la logica, l'astrazione ha fallito.

**III.1.B** Il codice lean si comprende dove si legge.

**Perché:** Ogni cambio di contesto ha un costo cognitivo.

### III.2 Stesso codice, scope diverso, resta separato

**III.2.A** Due funzioni possono contenere codice identico. Se servono concern diversi che possono cambiare indipendentemente, restano separate. Unirle accoppia due scope. Quando uno cambia, l'astrazione si rompe.

**III.2.B** La duplicazione non è il nemico. L'accoppiamento prematuro lo è.

**Perché:** Concern accoppiati costringono a pensare simultaneamente.

### III.3 Commenti

**III.3.A** Un file inizia con un header: la licenza e lo scope del file. Questo è l'unico commento che esiste per convenzione.

**III.3.B** Tutto il resto è auto-documentante. Se il codice ha bisogno di un commento per essere capito, un principio è stato violato. Correggi il nome (II), correggi la struttura (I), o correggi lo scope (III). Non correggerlo con un commento.

**Perché:** Un commento è un'ammissione che il codice non parla da solo.

---

## Il Frattale

Questi principi non sono sette regole separate. Sono una regola applicata ricorsivamente:

**Ogni unità fa una cosa, è dimensionata dal suo scope, è nominata da un vocabolario chiuso, è leggibile sul posto, ed è separata per concern.**

Questo vale per una funzione. Vale per un file. Vale per un modulo. Vale per un sistema.

Quando la codebase cresce, i principi non cambiano. Si estrapolano. Il sistema scala perché le regole sono frattali, non perché se ne aggiungono di nuove.

---

## Cosa Lean Code non è

**Non è codice sporco.** Il codice lean è strutturato, nominato e intenzionale.

**Non è una scusa per funzioni lunghe.** Una funzione da 500 righe che fa dodici cose viola I.1.

**Non è anti-astrazione.** Le astrazioni sono strumenti. Usale quando riducono la complessità. Non quando la spostano soltanto.

**Non è dogma.** Se una situazione richiede di infrangere un principio, infrangilo. Poi documenta il perché.

---

## Il test

Apri un file qualsiasi. Leggi una funzione qualsiasi. Se capisci cosa fa senza saltare altrove, è lean.

Se non ci riesci, non lo è.

---

*Vivian Voss, 2026*
