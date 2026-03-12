# Das Lean-Code-Manifest

## Vorwort

Clean Code hat 2008 ein echtes Problem gelöst. Java-Enterprise-Monolithen hatten 3.000-Zeilen-Methoden und keine Struktur. Die Verordnung war berechtigt.

Die Branche hat sie als universelles Gesetz kopiert.

Lean Code ist nicht das Gegenteil von Clean Code. Es ist das, was Clean Code hätte sein sollen: ein minimaler Satz von Prinzipien, der von einer einzelnen Funktion bis zu einem ganzen System skaliert.

Keine Zeilenlimit-Regeln. Kein Extraktionsdogma. Keine Zeremonie.
KISS im Kopf, nicht Regeln an der Wand.

---

## Das Axiom

**Bewahre die kognitive Kapazität des Entwicklers für die Logik.**

Kognitive Kapazität ist endlich. Jede unnötige Interpretation, jeder Kontextwechsel, jeder Sprung in eine andere Datei ist Kapazität, die nicht ins Verstehen der Logik fließt.

Alle Prinzipien in diesem Manifest dienen einem Zweck: den Kopf des Entwicklers frei halten für die eigentliche Arbeit.

---

## I. Struktur

Wie Code geschnitten wird.

### I.1 Eine Einheit, eine Aufgabe

**I.1.A** Eine Funktion tut eine Sache. Eine Datei hat eine Zuständigkeit. Ein Modul dient einem Zweck.

**Warum:** Dein Kopf kann einer klaren Sache sofort folgen.

### I.2 Der Scope bestimmt die Größe

**I.2.A** Der Scope der Aufgabe bestimmt die Größe der Einheit. Nicht eine Zeilenanzahl. Dreißig lesbare Zeilen für eine Aufgabe sind lean. Vier Zeilen für ein Fragment, das aus dem Kontext gerissen wurde, sind es nicht.

**I.2.B** Die Frage ist nie "wie viele Zeilen?", sondern "wie viele Aufgaben?"

**Warum:** Dein Kopf verliert den Kontext, wenn er springen muss.

---

## II. Benennung

Wie Code beschriftet wird.

### II.1 Verb und Scope aus einem geschlossenen Vokabular

**II.1.A** Ein Funktionsname folgt dem Muster `{verb}{scopeInMaxThreeWords}`. Das Verb stammt aus einer endlichen Menge:

| Verb | Bedeutung |
|------|---------|
| `read` | Daten abrufen |
| `write` | Daten speichern |
| `create` | Erzeugen |
| `delete` | Löschen |
| `update` | Bestehende Daten ändern |
| `find` | Nach Kriterien suchen |
| `check` | Eine Bedingung prüfen |
| `parse` | Ein Format transformieren |
| `render` | Ausgabe erzeugen |

**II.1.B** Der Scope beschreibt die Domäne in maximal drei Wörtern. Zusammen lesen sich Verb und Scope wie ein kurzer Satz. Wenn du mehr als drei Wörter brauchst, ist der Scope zu breit -- teile die Aufgabe.

`readOrderFromDatabase`. `checkStockAvailability`. `createInvoiceFromOrder`. `writeInvoiceToFile`. `parseConfigFromYaml`. Kein Kommentar nötig. Der Name ist der Vertrag.

**Warum:** Ein Verb, eine Bedeutung, überall. Keine Interpretation nötig.

### II.2 Nomenklatur

**II.2.A** Vollständig ausschreiben. `readConfiguration`, nicht `readConfig`. Ein Name ist ein Satz. Sätze enthalten keine Abkürzungen.

**II.2.B** Wähle das naivste Wort. `read`, nicht `fetch`. `check`, nicht `validate`. `write`, nicht `persist`. Je einfacher das Wort, desto weniger Interpretation erfordert es.

**II.2.C** Keine Synonyme. Kein `handle` oder `process`, die verschleiern, was tatsächlich passiert. Ein Verb, eine Bedeutung, überall.

**Warum:** Jede Übersetzung im Kopf verbraucht Kapazität.

### II.3 Namespace trägt den Scope, Name trägt das Substantiv

**II.3.A** Datenstrukturen tragen das Substantiv. Der Namespace trägt den Scope. `billing.Order`, nicht `BillingOrder`. `stock.Item`, nicht `StockItem`. Der Namespace hat eine Aufgabe. Der Typname hat eine andere. Keiner übernimmt die Arbeit des anderen.

**II.3.B** Folge den Namenskonventionen der Sprache, in der du schreibst. `read_order_from_database` in Python. `readOrderFromDatabase` in TypeScript. `ReadOrderFromDatabase` in Go. Das Prinzip bleibt. Die Schreibweise passt sich an.

**Warum:** Jeder Teil des Namens hat eine Aufgabe. Das ist Prinzip I.1, angewandt auf Benennung.

---

## III. Scope

Wie Code getrennt bleibt.

### III.1 Verstanden, wo es gelesen wird

**III.1.A** Code wird weit öfter gelesen als geschrieben. Wenn du zwischen acht Dateien springen musst, um eine Operation zu verstehen, hat die Struktur versagt. Wenn du einen Call-Graph im Kopf halten musst, um der Logik zu folgen, hat die Abstraktion versagt.

**III.1.B** Lean Code wird verstanden, wo er gelesen wird.

**Warum:** Jeder Kontextwechsel hat kognitive Kosten.

### III.2 Gleicher Code, anderer Scope, bleibt getrennt

**III.2.A** Zwei Funktionen können identischen Code enthalten. Wenn sie unterschiedliche Belange bedienen, die sich unabhängig voneinander ändern können, bleiben sie getrennt. Sie zusammenzuführen koppelt zwei Scopes. Wenn sich einer ändert, bricht die Abstraktion.

**III.2.B** Duplikation ist nicht der Feind. Voreilige Kopplung ist es.

**Warum:** Gekoppelte Belange erzwingen gleichzeitiges Denken.

### III.3 Kommentare

**III.3.A** Eine Datei beginnt mit einem Header: die Lizenz und der Scope der Datei. Das ist der einzige Kommentar, der per Konvention existiert.

**III.3.B** Alles andere ist selbstdokumentierend. Wenn Code einen Kommentar braucht, um verstanden zu werden, wurde ein Prinzip verletzt. Korrigiere den Namen (II), korrigiere die Struktur (I) oder korrigiere den Scope (III). Korrigiere es nicht mit einem Kommentar.

**Warum:** Ein Kommentar ist das Eingeständnis, dass der Code nicht für sich selbst spricht.

---

## Das Fraktal

Diese Prinzipien sind nicht sieben separate Regeln. Sie sind eine Regel, rekursiv angewandt:

**Jede Einheit tut eine Sache, ist durch ihren Scope dimensioniert, ist aus einem geschlossenen Vokabular benannt, ist an Ort und Stelle lesbar und ist nach Belangen getrennt.**

Das gilt für eine Funktion. Es gilt für eine Datei. Es gilt für ein Modul. Es gilt für ein System.

Wenn die Codebasis wächst, ändern sich die Prinzipien nicht. Sie extrapolieren. Das System skaliert, weil die Regeln fraktal sind, nicht weil neue Regeln hinzugefügt werden.

---

## Was Lean Code nicht ist

**Kein schmutziger Code.** Lean Code ist strukturiert, benannt und intentional.

**Keine Ausrede für lange Funktionen.** Eine 500-Zeilen-Funktion, die zwölf Dinge tut, verletzt I.1.

**Nicht anti-Abstraktion.** Abstraktionen sind Werkzeuge. Setze sie ein, wenn sie Komplexität reduzieren. Nicht, wenn sie sie nur verlagern.

**Kein Dogma.** Wenn eine Situation das Brechen eines Prinzips erfordert, brich es. Dann dokumentiere warum.

---

## Der Test

Öffne eine beliebige Datei. Lies eine beliebige Funktion. Wenn du verstehst, was sie tut, ohne woandershin springen zu müssen, ist sie lean.

Wenn nicht, ist sie es nicht.

---

*Vivian Voss, 2026*
