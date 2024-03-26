package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMainHandlerWhenRequestIsCorrect проверяет, что запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder() // Инициализируем ResponseRecorder, который будет записывать ответ сервера
	handler := http.HandlerFunc(mainHandle)    // Оборачиваем наш обработчик в http.HandlerFunc и обрабатываем запрос
	handler.ServeHTTP(responseRecorder, req)   // Запускаем обработку HTTP запроса 'req'

	assert.Equal(t, http.StatusOK, responseRecorder.Code) // Проверяем, что сервер вернул статус-код 200
	assert.NotEmpty(t, responseRecorder.Body.String())    //Проверяет, что тело HTTP-ответа не является пустой строкой
}

// TestMainHandlerWhenCityNotSupported проверяет, что если в параметре city, передается город, которого нет в нашем списке, то сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=newyork", nil) // Создаем запрос с количеством кафе в городе "newyork"

	responseRecorder := httptest.NewRecorder() // Инициализируем ResponseRecorder, который будет записывать ответ сервера
	handler := http.HandlerFunc(mainHandle)    // Оборачиваем наш обработчик в http.HandlerFunc и обрабатываем запрос
	handler.ServeHTTP(responseRecorder, req)   // Запускаем обработку HTTP запроса 'req'

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)       // Проверяем, что сервер вернул статус-код 400
	assert.Equal(t, "wrong city value", responseRecorder.Body.String()) // Проверяем, что фактическое тело HTTP-ответа, полученное из `responseRecorder.Body.String()`, соответствует ожидаемой строке `"wrong city value"`.
}

// TestMainHandlerWhenCountMoreThanTotal проверяет, что если в параметре count указано больше кафе, чем есть всего в списке, то должны вернуться все доступные кафе для данного города.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4 // Количество кафе в Москве

	req := httptest.NewRequest("GET", "/cafe?count="+strconv.Itoa(totalCount+1)+"&city=moscow", nil) // Создаем запрос с количеством кафе большим, чем общее количество

	responseRecorder := httptest.NewRecorder() // Инициализируем ResponseRecorder, который будет записывать ответ сервера
	handler := http.HandlerFunc(mainHandle)    // Оборачиваем наш обработчик в http.HandlerFunc и обрабатываем запрос
	handler.ServeHTTP(responseRecorder, req)   // Запускаем обработку HTTP запроса 'req'

	expectedCafes := strings.Join(cafeList["moscow"], ",")         // Ожидаемый ответ - список всех кафе, ставим запятую между названиями
	assert.Equal(t, http.StatusOK, responseRecorder.Code)          // Проверяем, что сервер вернул статус-код 200
	assert.Equal(t, expectedCafes, responseRecorder.Body.String()) // Проверяем, что тело ответа совпадает с ожидаемым списком кафе,
	// что подтверждает, что в случае запроса на количество больше доступного, сервер вернет полный список кафе
}
