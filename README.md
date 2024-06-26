# Решение задачи сортировки шаров по контейнерам

## Описание задачи

Дана задача сортировки шаров по контейнерам. У нас есть `n` контейнеров и `n` различных цветов шаров, причем шары распределены по контейнерам случайным образом. Допускается одна операция: обменять два шара в разных контейнерах. Необходимо выяснить, можно ли с помощью этой операции рассортировать шары так, чтобы в каждом контейнере лежали только шары одного цвета и каждый цвет присутствовал только в одном контейнере.

## Принцип решения

1. **Подсчет суммарного количества шаров каждого цвета и суммарного количества шаров в каждом контейнере**:
    - Создаем два массива `colorSums` и `containerSums` размером `n`, где `colorSums[i]` хранит общее количество шаров `i`-го цвета во всех контейнерах, а `containerSums[i]` хранит общее количество шаров во всех контейнерах `i`.
2. **Сравнение массивов**:
    - Если отсортированные массивы `colorSums` и `containerSums` совпадают, то можно рассортировать шары, иначе нельзя.

## Запуск решения

```bash
go run main.go
```

## Запуск тестов

```bash
go test -v
```