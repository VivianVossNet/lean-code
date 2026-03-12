# Le Manifeste Lean Code

## Préambule

Clean Code a résolu un vrai problème en 2008. Les monolithes Java enterprise avaient des méthodes de 3 000 lignes et aucune structure. La prescription était justifiée.

L'industrie l'a copiée comme loi universelle.

Lean Code n'est pas le contraire de Clean Code. C'est ce que Clean Code aurait dû être : un ensemble minimal de principes qui passent à l'échelle, d'une seule fonction à un système entier.

Pas de règles sur le nombre de lignes. Pas de dogme d'extraction. Pas de cérémonie.
KISS dans la tête, pas des règles sur un mur.

---

## L'axiome

**Préserver la capacité cognitive du développeur pour la logique.**

La capacité cognitive est finie. Chaque interprétation inutile, chaque changement de contexte, chaque saut vers un autre fichier est de la capacité qui ne va pas à la compréhension de la logique.

Tous les principes de ce manifeste servent un seul objectif : garder la tête du développeur libre pour le travail réel.

---

## I. Structure

Comment le code est découpé.

### I.1 Une unité, une tâche

**I.1.A** Une fonction fait une seule chose. Un fichier a une seule responsabilité. Un module sert un seul objectif.

**Pourquoi :** Votre tête peut suivre immédiatement une chose claire.

### I.2 Le périmètre définit la taille

**I.2.A** Le périmètre de la tâche définit la taille de l'unité. Pas un nombre de lignes. Trente lignes lisibles pour une tâche, c'est lean. Quatre lignes pour un fragment arraché de son contexte, non.

**I.2.B** La question n'est jamais "combien de lignes ?" mais "combien de tâches ?"

**Pourquoi :** Votre tête perd le contexte quand elle doit sauter.

---

## II. Nommage

Comment le code est étiqueté.

### II.1 Verbe et périmètre issus d'un vocabulaire fermé

**II.1.A** Un nom de fonction suit le schéma `{verb}{scopeInMaxThreeWords}`. Le verbe provient d'un ensemble fini :

| Verbe | Signification |
|-------|---------------|
| `read` | Récupérer des données |
| `write` | Stocker des données |
| `create` | Faire exister |
| `delete` | Supprimer de l'existence |
| `update` | Modifier des données existantes |
| `find` | Chercher selon des critères |
| `check` | Vérifier une condition |
| `parse` | Transformer un format |
| `render` | Produire une sortie |

**II.1.B** Le périmètre décrit le domaine en trois mots maximum. Ensemble, verbe et périmètre se lisent comme une phrase courte. Si vous avez besoin de plus de trois mots, le périmètre est trop large — découpez la tâche.

`readOrderFromDatabase`. `checkStockAvailability`. `createInvoiceFromOrder`. `writeInvoiceToFile`. `parseConfigFromYaml`. Aucun commentaire nécessaire. Le nom est le contrat.

**Pourquoi :** Un verbe, un sens, partout. Aucune interprétation requise.

### II.2 Nomenclature

**II.2.A** Écrire en toutes lettres. `readConfiguration`, pas `readConfig`. Un nom est une phrase. Les phrases ne contiennent pas d'abréviations.

**II.2.B** Choisir le mot le plus simple. `read`, pas `fetch`. `check`, pas `validate`. `write`, pas `persist`. Plus le mot est simple, moins il demande d'interprétation.

**II.2.C** Pas de synonymes. Pas de `handle` ou `process` qui masquent ce qui se passe réellement. Un verbe, un sens, partout.

**Pourquoi :** Chaque traduction dans votre tête consomme de la capacité.

### II.3 Le namespace porte le périmètre, le nom porte le substantif

**II.3.A** Les structures de données portent le substantif. Le namespace porte le périmètre. `billing.Order`, pas `BillingOrder`. `stock.Item`, pas `StockItem`. Le namespace fait une tâche. Le nom de type en fait une autre. Aucun des deux ne fait le travail de l'autre.

**II.3.B** Suivre les conventions de nommage du langage dans lequel vous écrivez. `read_order_from_database` en Python. `readOrderFromDatabase` en TypeScript. `ReadOrderFromDatabase` en Go. Le principe reste. La casse s'adapte.

**Pourquoi :** Chaque partie du nom a une tâche. C'est le principe I.1 appliqué au nommage.

---

## III. Périmètre

Comment le code reste séparé.

### III.1 Compris là où il est lu

**III.1.A** Le code est lu bien plus souvent qu'il n'est écrit. Si vous devez sauter entre huit fichiers pour comprendre une opération, la structure a échoué. Si vous devez garder un graphe d'appels en tête pour suivre la logique, l'abstraction a échoué.

**III.1.B** Le code lean se comprend là où il est lu.

**Pourquoi :** Chaque changement de contexte a un coût cognitif.

### III.2 Même code, périmètre différent, reste séparé

**III.2.A** Deux fonctions peuvent contenir un code identique. Si elles servent des préoccupations distinctes qui peuvent évoluer indépendamment, elles restent séparées. Les fusionner couple deux périmètres. Quand l'un change, l'abstraction casse.

**III.2.B** La duplication n'est pas l'ennemi. Le couplage prématuré, si.

**Pourquoi :** Des préoccupations couplées forcent à penser à tout en même temps.

### III.3 Commentaires

**III.3.A** Un fichier commence par un en-tête : la licence et le périmètre du fichier. C'est le seul commentaire qui existe par convention.

**III.3.B** Tout le reste est auto-documenté. Si du code a besoin d'un commentaire pour être compris, un principe a été violé. Corriger le nom (II), corriger la structure (I), ou corriger le périmètre (III). Ne pas corriger avec un commentaire.

**Pourquoi :** Un commentaire est l'aveu que le code ne parle pas de lui-même.

---

## Le fractal

Ces principes ne sont pas sept règles séparées. C'est une seule règle appliquée récursivement :

**Chaque unité fait une seule chose, est dimensionnée par son périmètre, est nommée à partir d'un vocabulaire fermé, est lisible sur place, et est séparée par préoccupation.**

Cela vaut pour une fonction. Cela vaut pour un fichier. Cela vaut pour un module. Cela vaut pour un système.

Quand la base de code grandit, les principes ne changent pas. Ils s'extrapolent. Le système passe à l'échelle parce que les règles sont fractales, pas parce que de nouvelles règles s'ajoutent.

---

## Ce que Lean Code n'est pas

**Pas du code sale.** Le code lean est structuré, nommé et intentionnel.

**Pas une excuse pour des fonctions longues.** Une fonction de 500 lignes qui fait douze choses viole I.1.

**Pas anti-abstraction.** Les abstractions sont des outils. Les utiliser quand elles réduisent la complexité. Pas quand elles ne font que la déplacer.

**Pas un dogme.** Si une situation exige de briser un principe, brisez-le. Puis documentez pourquoi.

---

## Le test

Ouvrir n'importe quel fichier. Lire n'importe quelle fonction. Si vous comprenez ce qu'elle fait sans sauter ailleurs, c'est lean.

Sinon, ce ne l'est pas.

---

*Vivian Voss, 2026*
