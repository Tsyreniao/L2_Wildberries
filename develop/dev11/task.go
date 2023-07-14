package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API:
	POST /create_event
	POST /update_event
	POST /delete_event
	GET /events_for_day
	GET /events_for_week
	GET /events_for_month

Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400.
	   В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

// Структура конфиг
type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// Чтение и парсинг из данных из конфига
func readConfig() Config {
	f, err := os.ReadFile("config.json")
	if err != nil {
		log.Println(err)
	}
	var config Config
	json.Unmarshal([]byte(f), &config)
	return config
}

// Структура ивент
type Event struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Date     time.Time `json:"date"`
	Duration int       `json:"duration"`
}

// Глобальные переменные
var (
	// Слайс ивентов
	events []Event
	// Ошибка
	err error
)

func main() {
	// Инициализируем конфиг
	config := readConfig()
	// Получаем адрес
	addr := config.Host + ":" + config.Port
	log.Print("Starting the service on address: ", addr)

	// Методы API
	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)
	http.HandleFunc("/events_for_day", eventsForDayHandler)
	http.HandleFunc("/events_for_week", eventsForWeekHandler)
	http.HandleFunc("/events_for_month", eventsForMonthHandler)

	// Запуск сервера
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(addr, nil))
}

// Функция для парсинга строки www-url-form-encoded
func parseURL(r *http.Request) (Event, error) {
	var event Event

	r.ParseForm()

	event.ID, err = strconv.Atoi(r.Form["id"][0])
	if err != nil {
		log.Print("ERROR Сouldn't convert id to int")
		return event, err
	}

	event.Title = r.Form["title"][0]

	event.Date, err = time.Parse("2006-01-02", r.Form["date"][0])
	if err != nil {
		log.Print("ERROR Сouldn't convert date to time.Time")
		return event, err
	}

	event.Duration, err = strconv.Atoi(r.Form["duration"][0])
	if err != nil {
		log.Print("ERROR Сouldn't convert duration to int")
		return event, err
	}

	return event, err
}

// Функция создания ивента
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	// Убедимся что прилетел POST запрос
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		response := map[string]string{"error": "Method is not allowed"}
		jsonResponse(w, http.StatusMethodNotAllowed, response)
		return
	}

	// Убедимся что верный тип контента
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.Header().Set("Allow Content-Type", "application/x-www-form-urlencoded")
		response := map[string]string{"error": "Content type is not allowed"}
		jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	// Парсим тело запроса
	event, err := parseURL(r)
	if err != nil {
		log.Print("ERROR parse URL")
		response := map[string]string{"error": "Incorrect syntax"}
		jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	// Добавляем ивент в слайс
	events = append(events, event)

	log.Print("Event created ", event)
	response := map[string]string{"result": "Event created"}
	jsonResponse(w, http.StatusOK, response)
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	// Убедимся что прилетел POST запрос
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		response := map[string]string{"error": "Method is not allowed"}
		jsonResponse(w, http.StatusMethodNotAllowed, response)
		return
	}

	// Убедимся что верный тип контента
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.Header().Set("Allow Content-Type", "application/x-www-form-urlencoded")
		response := map[string]string{"error": "Content type is not allowed"}
		jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	// Парсим тело запроса
	event, err := parseURL(r)
	if err != nil {
		log.Print("ERROR parse URL")
		response := map[string]string{"error": "Incorrect syntax"}
		jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	// Ищем ивент с нужным id
	for k := range events {
		if events[k].ID == event.ID {
			// Обновляем ивент
			events[k].Date = event.Date
			events[k].Duration = event.Duration
			events[k].Title = event.Title

			log.Print("Event updated ", event)
			response := map[string]string{"result": "Event updated"}
			jsonResponse(w, http.StatusOK, response)
			return
		}
	}

	response := map[string]string{"error": "ID not found"}
	jsonResponse(w, http.StatusBadRequest, response)
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	// Убедимся что прилетел POST запрос
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		response := map[string]string{"error": "Method is not allowed"}
		jsonResponse(w, http.StatusMethodNotAllowed, response)
		return
	}

	// Убедимся что верный тип контента
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.Header().Set("Allow Content-Type", "application/x-www-form-urlencoded")
		response := map[string]string{"error": "Content type is not allowed"}
		jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	// Парсим тело запроса
	r.ParseForm()
	var eventID int
	eventID, err = strconv.Atoi(r.Form["id"][0])
	if err != nil {
		log.Print("ERROR Сouldn't convert id to int")
		response := map[string]string{"error": "Incorrect syntax"}
		jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	// Ищем ивент с нужным id
	for k := range events {
		if events[k].ID == eventID {
			// Удаляем ивент из слайса
			events = append(events[:k], events[k+1:]...)

			log.Print("Event with ID = ", eventID, " was deleted")
			response := map[string]string{"result": "Event deleted"}
			jsonResponse(w, http.StatusOK, response)
			return
		}
	}

	response := map[string]string{"error": "ID not found"}
	jsonResponse(w, http.StatusBadRequest, response)
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	// Убедимся что прилетел GET запрос
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		response := map[string]string{"error": "Method is not allowed"}
		jsonResponse(w, http.StatusMethodNotAllowed, response)
		return
	}

	// Парсим запрос и получаем дату
	var eventDate time.Time
	eventDate, err = time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		log.Print("ERROR Сouldn't convert date to time.Time")
		response := map[string]string{"error": "Incorrect syntax"}
		jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	// Слайс ивентов в запрошенный день
	var eventsJSON []map[string]interface{}
	var isFound bool
	// Ищем ивенты в запрошенный день
	for k := range events {
		if events[k].Date == eventDate {
			isFound = true

			// Структуру Event переводим в []byte
			var eventJSON map[string]interface{}
			b, err := json.Marshal(events[k])
			if err != nil {
				log.Print("ERROR Сouldn't convert events[k] to []byte")
				response := map[string]string{"error": "Convert error"}
				jsonResponse(w, http.StatusInternalServerError, response)
				return
			}
			// []byte переводим в map[string]interface{} и добавляем в слайс
			json.Unmarshal(b, &eventJSON)
			eventsJSON = append(eventsJSON, eventJSON)
		}
	}

	// Если ивенты найдены, то выводим их
	if isFound {
		response := map[string]interface{}{"result": "Events are shown"}
		eventsJSON = append(eventsJSON, response)
		jsonResponse(w, http.StatusOK, eventsJSON)
	} else {
		response := map[string]string{"error": "Events not found"}
		jsonResponse(w, http.StatusInternalServerError, response)
	}
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	// Убедимся что прилетел GET запрос
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		response := map[string]string{"error": "Method is not allowed"}
		jsonResponse(w, http.StatusMethodNotAllowed, response)
		return
	}

	// Парсим запрос и получаем дату
	var eventDate time.Time
	eventDate, err = time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		log.Print("ERROR Сouldn't convert date to time.Time")
		response := map[string]string{"error": "Incorrect syntax"}
		jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	// Слайс ивентов за неделю с полученного дня
	var eventsJSON []map[string]interface{}
	var isFound bool
	// Дата через неделю с полученной даты
	eventPlusWeek := eventDate.AddDate(0, 0, 7)
	// Ищем ивенты за неделю с полученного дня
	for k := range events {
		if (events[k].Date.After(eventDate) && events[k].Date.Before(eventPlusWeek)) || events[k].Date == eventDate {
			isFound = true

			// Структуру Event переводим в []byte
			var eventJSON map[string]interface{}
			b, err := json.Marshal(events[k])
			if err != nil {
				log.Print("ERROR Сouldn't convert events[k] to []byte")
				response := map[string]string{"error": "Convert error"}
				jsonResponse(w, http.StatusInternalServerError, response)
				return
			}
			// []byte переводим в map[string]interface{} и добавляем в слайс
			json.Unmarshal(b, &eventJSON)
			eventsJSON = append(eventsJSON, eventJSON)
		}
	}

	// Если ивенты найдены, то выводим их
	if isFound {
		response := map[string]interface{}{"result": "Events are shown"}
		eventsJSON = append(eventsJSON, response)
		jsonResponse(w, http.StatusOK, eventsJSON)
	} else {
		response := map[string]string{"error": "Events not found"}
		jsonResponse(w, http.StatusInternalServerError, response)
	}
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	// Убедимся что прилетел GET запрос
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		response := map[string]string{"error": "Method is not allowed"}
		jsonResponse(w, http.StatusMethodNotAllowed, response)
		return
	}

	// Парсим запрос и получаем дату
	var eventDate time.Time
	eventDate, err = time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		log.Print("ERROR Сouldn't convert date to time.Time")
		response := map[string]string{"error": "Incorrect syntax"}
		jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	// Слайс ивентов за месяц с полученного дня
	var eventsJSON []map[string]interface{}
	var isFound bool
	// Дата через месяц с полученной даты
	eventPlusMonth := eventDate.AddDate(0, 1, 0)
	// Ищем ивенты за месяц с полученного дня
	for k := range events {
		if (events[k].Date.After(eventDate) && events[k].Date.Before(eventPlusMonth)) || events[k].Date == eventDate {
			isFound = true

			// Структуру Event переводим в []byte
			var eventJSON map[string]interface{}
			b, err := json.Marshal(events[k])
			if err != nil {
				log.Print("ERROR Сouldn't convert events[k] to []byte")
				response := map[string]string{"error": "Convert error"}
				jsonResponse(w, http.StatusInternalServerError, response)
				return
			}
			// []byte переводим в map[string]interface{} и добавляем в слайс
			json.Unmarshal(b, &eventJSON)
			eventsJSON = append(eventsJSON, eventJSON)
		}
	}

	// Если ивенты найдены, то выводим их
	if isFound {
		response := map[string]interface{}{"result": "Events are shown"}
		eventsJSON = append(eventsJSON, response)
		jsonResponse(w, http.StatusOK, eventsJSON)
	} else {
		response := map[string]string{"error": "Events not found"}
		jsonResponse(w, http.StatusInternalServerError, response)
	}
}

// Вспомогательная функция, которая устанавливает заголовок, статус код и выводит данные в формате json
func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
