- поток ОС имеет блок памяти фиксированного размера (зачастую до 2мб) для стека, в котором он хранит локальные переменные вызовов функций, находящихся в работе или приостановленных в ожидании связанных функций
- горутина начинает работу со стеком в 2кб и может вырасти до 1гб
- т.к. потоки ОС планируются в ее ядре, передача управления от одного потока другому требует полного переключения контекста, т.е. сохранения состояния одного потока в памяти, восстановление состояния другого потока и обновление структур данных планировщика. Это медленная операция из-за слабой локальности и необходимого количества обащений к памяти
- язык имеет свой планировщик, который использует M:N планирование, потому что мультиплексирует выполнение M-горутин на N-потоках
- основная суть планировщика заключается в том, что он пытается управлять G, M и P: горутинами, машинами (потоками) и процессорами.

  - «G» - просто горутина Golang
  - «M» - поток ОС, который может выполнять что-либо или же бездействовать
  - «P» можно рассматривать как ЦП в планировщике ОС; он представляет ресурсы, необходимые для выполнения нашего Go кода, такие как планировщик или состояние распределителя памяти
  - В рантайме они представлены как структуры: type g, type m или type p

- основная задача планировщика состоит в том, чтобы сопоставить каждую G (код, который мы хотим выполнить) с M (где его выполнять) и P (права и ресурсы для выполнения)

- задания планировщика аналогичны заданиям планировщика ядра, но связаны только с горутинами одной программы
- при необходимости, планировщик переводит горутины в спящий режим и активурет новые
- переменная GOMAXPROCS по дефолту равна количеству ядер процессора, но ее можно изменить
- если переменная GOMAXPROCS = 0, то конкурентности не будет.
- сначала горутина попадает в одноэлементный стек, из него в очередь. Сначала исполнится горутина из стека, потом из очереди
- если горутина делает системный вызов рантайм передает управление ОС, где из тред пула берется отдельный тред для исполнения этой горутины. После исполнения эта горутина попадает в глобальную очередь
- сначала горутины берутся из локальной очереди, потом крадутся у соседа (других локальных очередей), потом берутся из глобальной очереди
- каждый 61-й такт берем (ищем) горутину в глобальной очереди
- для сетевых вызовов используется абстракция над epoll - netpoller
- пока горутина ждет сетевого ответа она пассивно хранится в epoll, после сетевого ответа она пробуждается и перемещается в глобальную очередь
- sysmon - одтельная горутина следит за триггером сборщика мусора и подвигает застрявшие горутины, которые долго исполняются, давая время поработать другим