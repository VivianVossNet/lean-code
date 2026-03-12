# O Manifesto do Lean Code

## Preâmbulo

O Clean Code resolveu um problema real em 2008. Os monólitos enterprise em Java tinham métodos com 3.000 linhas e nenhuma estrutura. A prescrição era genuína.

A indústria copiou-a como lei universal.

O Lean Code não é o oposto do Clean Code. É aquilo que o Clean Code deveria ter sido: um conjunto mínimo de princípios que escala de uma única função a um sistema inteiro.

Sem regras de contagem de linhas. Sem dogma de extracção. Sem cerimónia.
KISS na cabeça, não regras numa parede.

---

## O Axioma

**Preservar a capacidade cognitiva do programador para a lógica.**

A capacidade cognitiva é finita. Cada interpretação desnecessária, cada mudança de contexto, cada salto para outro ficheiro é capacidade que não vai para a compreensão da lógica.

Todos os princípios neste manifesto servem um único propósito: manter a cabeça do programador livre para o trabalho real.

---

## I. Estrutura

Como o código é dividido.

### I.1 Uma unidade, uma tarefa

**I.1.A** Uma função faz uma coisa. Um ficheiro tem uma responsabilidade. Um módulo serve um propósito.

**Porquê:** A cabeça consegue seguir uma coisa clara de imediato.

### I.2 O âmbito define o tamanho

**I.2.A** O âmbito da tarefa define o tamanho da unidade. Não uma contagem de linhas. Trinta linhas legíveis para uma tarefa é lean. Quatro linhas para um fragmento arrancado do contexto não é.

**I.2.B** A pergunta nunca é "quantas linhas?" mas "quantas tarefas?"

**Porquê:** A cabeça perde contexto quando tem de saltar.

---

## II. Nomenclatura

Como o código é etiquetado.

### II.1 Verbo e âmbito a partir de um vocabulário fechado

**II.1.A** O nome de uma função segue o padrão `{verbo}{âmbitoEmMáximoTrêsPalavras}`. O verbo vem de um conjunto finito:

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

**II.1.B** O âmbito descreve o domínio num máximo de três palavras. Juntos, verbo e âmbito lêem-se como uma frase curta. Se precisar de mais de três palavras, o âmbito é demasiado amplo — divida a tarefa.

`readOrderFromDatabase`. `checkStockAvailability`. `createInvoiceFromOrder`. `writeInvoiceToFile`. `parseConfigFromYaml`. Nenhum comentário necessário. O nome é o contrato.

**Porquê:** Um verbo, um significado, em todo o lado. Nenhuma interpretação necessária.

### II.2 Nomenclatura

**II.2.A** Escreva por extenso. `readConfiguration`, não `readConfig`. Um nome é uma frase. Frases não contêm abreviaturas.

**II.2.B** Escolha a palavra mais simples. `read`, não `fetch`. `check`, não `validate`. `write`, não `persist`. Quanto mais simples a palavra, menos interpretação exige.

**II.2.C** Sem sinónimos. Nada de `handle` ou `process` que escondem o que realmente acontece. Um verbo, um significado, em todo o lado.

**Porquê:** Cada tradução na cabeça consome capacidade.

### II.3 O namespace carrega o âmbito, o nome carrega o substantivo

**II.3.A** As estruturas de dados carregam o substantivo. O namespace carrega o âmbito. `billing.Order`, não `BillingOrder`. `stock.Item`, não `StockItem`. O namespace faz uma tarefa. O nome do tipo faz outra. Nenhum faz o trabalho do outro.

**II.3.B** Siga as convenções de nomenclatura da linguagem em que escreve. `read_order_from_database` em Python. `readOrderFromDatabase` em TypeScript. `ReadOrderFromDatabase` em Go. O princípio mantém-se. A capitalização adapta-se.

**Porquê:** Cada parte do nome tem uma tarefa. É o princípio I.1 aplicado à nomenclatura.

---

## III. Âmbito

Como o código se mantém separado.

### III.1 Compreendido onde é lido

**III.1.A** O código é lido muito mais vezes do que é escrito. Se precisa de saltar entre oito ficheiros para compreender uma operação, a estrutura falhou. Se precisa de manter um grafo de chamadas na cabeça para seguir a lógica, a abstracção falhou.

**III.1.B** Código lean é compreendido onde é lido.

**Porquê:** Cada mudança de contexto tem custo cognitivo.

### III.2 Mesmo código, âmbito diferente, mantém-se separado

**III.2.A** Duas funções podem conter código idêntico. Se servem preocupações diferentes que podem mudar independentemente, mantêm-se separadas. Fundi-las acopla dois âmbitos. Quando um muda, a abstracção quebra.

**III.2.B** A duplicação não é o inimigo. O acoplamento prematuro é.

**Porquê:** Preocupações acopladas forçam pensamento simultâneo.

### III.3 Comentários

**III.3.A** Um ficheiro começa com um cabeçalho: a licença e o âmbito do ficheiro. Este é o único comentário que existe por convenção.

**III.3.B** Tudo o resto é auto-documentado. Se o código precisa de um comentário para ser compreendido, um princípio foi violado. Corrija o nome (II), corrija a estrutura (I), ou corrija o âmbito (III). Não corrija com um comentário.

**Porquê:** Um comentário é uma admissão de que o código não fala por si.

---

## O Fractal

Estes princípios não são sete regras separadas. São uma regra aplicada recursivamente:

**Cada unidade faz uma coisa, é dimensionada pelo seu âmbito, é nomeada a partir de um vocabulário fechado, é legível no local, e é separada por preocupação.**

Isto vale para uma função. Vale para um ficheiro. Vale para um módulo. Vale para um sistema.

Quando a base de código cresce, os princípios não mudam. Extrapolam. O sistema escala porque as regras são fractais, não porque se adicionam regras novas.

---

## O que o Lean Code não é

**Não é código sujo.** Código lean é estruturado, nomeado e intencional.

**Não é desculpa para funções longas.** Uma função de 500 linhas que faz doze coisas viola I.1.

**Não é anti-abstracção.** Abstracções são ferramentas. Use-as quando reduzem complexidade. Não quando apenas a recolocam.

**Não é dogma.** Se uma situação exige quebrar um princípio, quebre-o. Depois documente porquê.

---

## O teste

Abra qualquer ficheiro. Leia qualquer função. Se compreende o que faz sem saltar para outro sítio, é lean.

Se não compreende, não é.

---

*Vivian Voss, 2026*
