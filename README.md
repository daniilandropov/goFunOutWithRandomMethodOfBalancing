# FanOut Example in Go

## Русский

Этот пример демонстрирует функцию **FanOut**, которая распределяет значения из одного канала (`chIn`) на несколько каналов-получателей.  
Основные моменты:

- `FanOut[T any](chIn <-chan T, n int) []<-chan T` — дженерик-функция, создаёт `n` выходных каналов.
- Внутри горутины происходит распределение значений по трём стратегиям:
  1. **Случайный с шагом** — берётся следующий канал по кругу, если он забит, пробуем следующий.
  2. **Пока не удастся** — пытаемся отправить значение в канал до успешной записи.
  3. **По наименее загруженному** — ищем канал с минимальной длиной буфера.
- После закрытия входного канала все выходные каналы закрываются.
- В `main` создаём 5000 значений, отправляем в канал, закрываем его.
- Используем `WaitGroup` для ожидания завершения всех горутин-получателей, которые просто выводят значения на экран.

Особенности:
- Буферизация каналов помогает избежать блокировок.
- `rand.Intn(3)` выбирает случайную стратегию распределения для каждого значения.

---

## English

This example demonstrates the **FanOut** function, which distributes values from a single channel (`chIn`) to multiple output channels.  
Key points:

- `FanOut[T any](chIn <-chan T, n int) []<-chan T` is a generic function that creates `n` output channels.
- A goroutine handles the distribution of values using three strategies:
  1. **Step-by-step random** — the next channel in a round-robin fashion; if full, try the next one.
  2. **Try until successful** — keeps attempting to send the value until it succeeds.
  3. **Least loaded channel** — finds the channel with the smallest buffer length.
- When the input channel is closed, all output channels are closed as well.
- In `main`, we create 5000 values, send them into the input channel, and close it.
- `WaitGroup` is used to wait for all receiver goroutines to finish, which simply print values to the console.

Notes:
- Buffered channels help prevent blocking.
- `rand.Intn(3)` chooses a random distribution strategy for each value.
