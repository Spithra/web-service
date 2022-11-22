package main //Обьявление пакета main

//Импорт библиотек
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Обьявление структуры flavor, будет использоваться для хранения данных (flavors) в памяти
type flavor struct {
	ID     int8
	Flavor string
	Price  float32
}

// Инициализируем структуру и передвем ей данные
var flavors = []flavor{

	{ID: 1, Flavor: "Клубника", Price: 115.99},
	{ID: 2, Flavor: "Малина", Price: 116.99},
	{ID: 3, Flavor: "Апельсин", Price: 117.99},
}

// Определяем для функции параметр gin.Context, она возвращает все вкусы (flavors) в формате JSON
func getFlavors(c *gin.Context) {
	// Вызов Context.JSON это возьмет структуру в формате JSON и добавит в ответ. StatusOK сообщает код 200 OK
	c.JSON(http.StatusOK, flavors)
}

func main() {
	// Инициализация роутера
	router := gin.Default()
	// Функция GET связывает метод GET http и путь /flavors с функцией getFlavors
	router.GET("/flavors", getFlavors) // здесь передано имя функции getFlavors
	//а вот 3 строчки ниже передается результат функции
	//Вообще POST, PUT и DELETE должны были выполнять свои функции а именно добавлять обновлять и удалять
	//данные соответсвенно, но это пока в разработке и добавлено для возможности простестировать
	// работоспособность методов
	router.POST("/post", func(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"test": "test"}) })
	router.PUT("/put", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"test": "test"}) })
	router.DELETE("/delete", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"test": "test"}) })
	// Запускаем сервер
	router.Run("localhost:8080")
}
