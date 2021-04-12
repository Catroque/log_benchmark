# Benchmark Log

Código para benchmark de implementações de log.

São utilizadas, fundamentalmente, 2 (duas) bibliotecas:

- **logrus** : github.com/sirupsen/logrus
- **zerolog** : github.com/rs/zerolog

Com essas bibliotecas são contruídas duas implementação de wrapper para generalização das intefaces. Cada uma delas implementa interfaces distintas simulando o padrão utilizado entre as bibliotecas envolvidas:

- github.com/ademirfteo/gologwrp/pkg/logger
- github.com/Catroque/loggo/pkg/logger

A interface *logrus* testa a condição de log ao final de sua cadeia de funções ao passo que a *zerolog* testa a possibildade de execução no início da cadeia. Isso faz com que não seja necessária a formação de dados para um evento de log que não será registrado na aplicação.

## Testes

Foram realizados 6 (seis) testes para cada tipo de implmentação:

- *Debug2Info_Simple* : Solicitado registro de evento de DEBUG com apenas uma mensagem texto para uma configuração que permite eventos a partir de INFO. Queremos ver o comportamento da implementação quando **não haverá** registro do evento de log.

- *Debug2Info_WithFields* : Conforme o teste anterior, porém o registro do evento conta com a de 2 (dois) campos do tipo string.

- *Info2Info_Simple* : Solicitado registro de evento de INFO com formatação de 2 campos para uma configuração que permite eventos a partir de INFO. Queremos ver o comportamento da implementação quando **haverá** registro do evento de log.

- *Info2Info_WithFields* : Conforme o teste anterior, porém o registro do evento conta com a de 2 (dois) campos do tipo string.

- *Debug2Info_G_Simple* : Solicitado registro de evento de DEBUG com apenas uma mensagem texto para uma configuração que permite eventos a partir de INFO. A configuração do logger é estabelecida globalmente. Comportamento da implementação quando **não haverá** registro do evento de log.

- *Info2Info_G_Simple* : Solicitado registro de evento de INFO com apenas uma mensagem texto para uma configuração que permite eventos a partir de INFO. A configuração do logger é estabelecida globalmente. Comportamento da implementação quando **haverá** registro do evento de log.

## Resultados

- logrus

| Test | Quantity | ns/op | Bytes/op | allocs/op |
| :--- | ---: | -----------: | -----------: | ----------------------: |
| BenchmarkLog_Debug2Info_Simple-4     | 317207274 |     3.79 |    0 |  0 |
| BenchmarkLog_Debug2Info_WithFields-4 |   2108133 |   561.00 |  496 |  4 |
| BenchmarkLog_Info2Info_Simple-4      |     77180 | 14624.00 | 1193 | 24 |
| BenchmarkLog_Info2Info_WithFields-4  |     62114 | 20834.00 | 2122 | 32 |
| BenchmarkLog_Debug2Info_G_Simple-4   | 318490276 |     3.86 |    0 |  0 |
| BenchmarkLog_Info2Info_G_Simple-4    |     80445 | 14798.00 | 1193 | 24 |

- zerolog

| Test | Quantity | ns/op | Bytes/op | allocs/op |
| :--- | ---: | -----------: | -----------: | ----------------------: |
| BenchmarkLog_Debug2Info_Simple-4     | 232308735 |     5.25 |  0 | 0 |
| BenchmarkLog_Debug2Info_WithFields-4 |  22466348 |    49.10 |  0 | 0 |
| BenchmarkLog_Info2Info_Simple-4      |    178635 |  9565.00 |  0 | 0 |
| BenchmarkLog_Info2Info_WithFields-4  |     81573 | 12459.00 | 64 | 2 |
| BenchmarkLog_Debug2Info_G_Simple-4   | 208164003 |     5.22 |  0 | 0 |
| BenchmarkLog_Info2Info_G_Simple-4    |    124914 |  9572.00 |  0 | 0 |

- gologwrp

| Test | Quantity | ns/op | Bytes/op | allocs/op |
| :--- | ---: | -----------: | -----------: | ----------------------: |
| BenchmarkLog_Debug2Info_Simple-4     | 151531218 |     7.85 |    0 | 0 |
| BenchmarkLog_Debug2Info_WithFields-4 |   1387777 |   858.00 | 1008 | 6 |
| BenchmarkLog_Info2Info_Simple-4      |    189288 |  9712.00 |    0 | 0 |
| BenchmarkLog_Info2Info_WithFields-4  |     84388 | 13757.00 | 1008 | 6 |
| BenchmarkLog_Debug2Info_G_Simple-4   | 152163613 |     8.14 |    0 | 0 |
| BenchmarkLog_Info2Info_G_Simple-4    |    207320 |  9229.00 |    0 | 0 |

- loggo

| Test | Quantity | ns/op | Bytes/op | allocs/op |
| :--- | ---: | -----------: | -----------: | ----------------------: |
| BenchmarkLog_Debug2Info_Simple-4     | 18701166 |     55.5 |  16 | 1 |
| BenchmarkLog_Debug2Info_WithFields-4 |  4078516 |   289.00 | 352 | 3 |
| BenchmarkLog_Info2Info_Simple-4      |   186774 |  9624.00 |  16 | 1 |
| BenchmarkLog_Info2Info_WithFields-4  |    85340 | 13256.00 | 416 | 5 |
| BenchmarkLog_Debug2Info_G_Simple-4   | 18295395 |     55.6 |  16 | 1 |
| BenchmarkLog_Info2Info_G_Simple-4    |   119865 |  9338.00 |  16 | 1 |

