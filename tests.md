- тестовые файлы должны иметь паттерн \*\_test.go, пакет должен импортировать "testing", функция внутри этого пакета должна начинаться с Test
- go test -v для отображения информации о тестах
- go test -coverprofile=c.out + go tool cover -html=c.out для вывода покрытия

- функции-бенчмарки начинаются со слова Benchmark
- флаг -benchmem добавляет в отчет информацию об использовании памяти
- флаг -cpu определяет количество задействованных потоков

- профилирование cpu - go test -cpuprofile=cpu.out
- профилирование памяти - go test -memprofile.mem.out
- профилирование блокировок - go test -blockprofile=block.out

- функции-примеры начинаются с Example и исполняются по обычным правилам. Создаются в целях документирования. Если в функции есть "// Output:", то данные будут проверены