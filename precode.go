package main

import (
	"net/http"
	"strconv"
	"strings"
)

// Список кафе по городу Москва
var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count") // Извлекаем параметр 'count' из GET-запроса
	if countStr == "" {                      // Если параметр 'count' не предоставлен, отправляем ошибку клиенту
		w.WriteHeader(http.StatusBadRequest) // Устанавливаем HTTP статус-код 400
		w.Write([]byte("count missing"))     // Отправляем сообщение об ошибке в теле ответа
		return                               // Прекращаем дальнейшую обработку функции
	}

	count, err := strconv.Atoi(countStr) // Преобразуем 'count' из строки в целое число
	if err != nil {                      // Если 'count' не может быть преобразован, сообщаем об ошибке
		w.WriteHeader(http.StatusBadRequest) // Устанавливаем HTTP статус-код 400
		w.Write([]byte("wrong count value")) // Отправляем сообщение об ошибке в теле ответа
		return                               // Прекращаем дальнейшую обработку функции
	}

	city := req.URL.Query().Get("city") // Извлекаем параметр 'city' из запроса

	cafe, ok := cafeList[city] // Пытаемся получить список кафе для данного города
	if !ok {                   // Если город не найден в списке, сообщаем об ошибке
		w.WriteHeader(http.StatusBadRequest) // Устанавливаем HTTP статус-код 400
		w.Write([]byte("wrong city value"))  // Отправляем сообщение об ошибке в теле ответа
		return                               // Прекращаем дальнейшую обработку функции
	}

	if count > len(cafe) { // Если запрашиваемое количество кафе больше, чем доступно, корректируем это количество
		count = len(cafe) // Присваиваем 'count' количество доступных кафе, чтобы вернуть их все
	}

	answer := strings.Join(cafe[:count], ",") // Объединяем имена кафе в строку, разделяя их запятой

	w.WriteHeader(http.StatusOK) // Устанавливаем HTTP статус-код 200, так как все обработано корректно
	w.Write([]byte(answer))      // Отправляем ответ с перечнем кафе
}

// Тесты перенес в отдельный файл main_test.go
