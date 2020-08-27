# Go-Counting

## ->develop

## Формулировка задачи:
Процессу на stdin приходят строки, содержащие URL или названия файлов. Каждый такой
URL нужно запросить, каждый файл нужно прочитать, и посчитать кол-во вхождений строки
"Go" в ответе. В конце работы приложение выводит на экран общее кол-во найденных строк
"Go" во всех источниках данных, например:
- $ echo -e 'https://golang.org\n/etc/passwd\nhttps://golang.org\nhttps://golang.org' |
- go run 1.go
- Count for https://golang.org: 9
- Count for /etc/passwd: 0
- Count for https://golang.org: 9
- Count for https://golang.org: 9
- Total: 27
Каждый источник данных должен начать обрабатываться сразу после вычитывания и
параллельно с вычитыванием следующего. Источники должны обрабатываться параллельно,но
не более k=5 одновременно. Обработчики данных не должны порождать лишних горутин, т.е.
если k=1000 а обрабатываемых источников нет, не должно создаваться 1000 горутин. Нужно
обойтись без глобальных переменных и использовать только стандартные библиотеки.

## Описание:
Работа одновременно 5 параллельных обрабатываемых горутин осуществляется с помощью буфферизированного канала queueCh. Есть несколько небольших модулей
1) Open - для понимания что попалось на инпут - файл или url-ссылка
2) Counter - основной обработчик (менеджер) для инпута. Синхронизация горутин (в моем случае воркеров) осуществляется с помощью sync.WaitGroup, а ограничение до 5 с помощью буфферизированного канала. 
3) Workers - воркер, который парсит данные и считает вхождения слова Go. Парсер считывает данные по чанкам, т.к. есть вероятность, что файл будет огромный и памяти просто не хватит. 
