package main

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"net/http"
	"regexp"
	"time"
)

// Сетевой адрес.
//
// Служба будет слушать запросы на всех IP-адресах
// компьютера на порту 20002.
const addr = "0.0.0.0:27002"

// Протокол сетевой службы.
const proto = "tcp4"

func main() {
	// Запуск сетевой службы.
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// Подключения обрабатываются в бесконечном цикле.
	for {
		// Принимаем подключение.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Вызов обработчика подключения.
		go handleConn(conn)
	}
}

// Обработчик. Вызывается для каждого соединения.
func handleConn(conn net.Conn) {

	var proverb []string // Для хранения найденных поговорок.

	// Получаем с сайта исходную страницу с поговорками.
	content, err := http.Get("https://go-proverbs.github.io")
	if err != nil {
		log.Fatal(err)
	}
	defer content.Body.Close()

	// Регулярное выражение для поиска.
	re := regexp.MustCompile(`">(.*)\.</a`)

	scanner := bufio.NewScanner(content.Body)

	// Сканируем полученную страницу для поиска строк, удовлетворяющих регулярному выражению.
	for i := 0; scanner.Scan(); i++ {
		submatches := re.FindAllStringSubmatch(string(scanner.Text()), -1)
		for _, s := range submatches {
			proverb = append(proverb, s[1]) // Добавляем поговорки в слайс.
		}
	}

	for {
		conn.Write([]byte(proverb[rand.Intn(len(proverb))] + "." + "\n" + "\r")) // Вывод поговорки на новой строке с возвратом корректки.
		time.Sleep(3 * time.Second)
	}
}
