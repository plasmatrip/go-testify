package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//проверка, что сервер вернул код ответа 200
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	//прооверяем, что тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=murmansk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//проверка, что сервер вернул код ответа 400
	//если ответ не Bad Requset, то дальше проверять нечего
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	//ожидаемое сообщение от сервера
	expected := "wrong city value"
	//проверка, что при неправильном городе сервер вернул в теле ответа сообщение "wrong city value"
	assert.Equal(t, expected, responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := len(cafeList["moscow"])
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=5", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//проверка, что сервер вернул код ответа 200
	//если ответ не ОК, то дальше проверять нечего
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	//преобразуем ответ сервера с кафе из строки в слайс
	listOfCafe := strings.Split(responseRecorder.Body.String(), ",")
	//проряем, что количество кафе равно ожидаемому
	assert.Len(t, listOfCafe, totalCount)
}
