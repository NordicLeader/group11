package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Car struct {
	Mark  string `json:"Mark"`
	Model string `json:"Model"`
	Photo string `json:"Photo"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		markCar := r.URL.Query().Get("mark")
		modelCar := r.URL.Query().Get("model")
		photoCar := r.URL.Query().Get("photo")

		b := Car{Mark: markCar, Model: modelCar, Photo: photoCar}

		if markCar != "" {

			//считываем то что было в файле
			dataFromFile, _ := ioutil.ReadFile("data.json")

			//создаем контейнер под данные из json
			addCar := []Car{}

			//преоразовать json из файла в срез машин
			json.Unmarshal(dataFromFile, &addCar)

			//добавляем собранную машину к списку тех что были
			addCar = append(addCar, b)

			//собираем json из обновленного списка
			jsonData, _ := json.Marshal(addCar)

			//переводим json данные из среза байт в строку
			jsonstring := string(jsonData)

			//открываем файл для работы с ним
			file, _ := os.Create("data.json")

			//отложенное закрытие файла
			defer file.Close()

			//записываем в файл данные
			file.WriteString(jsonstring)

			//выводим на экран
			fmt.Fprintf(w, jsonstring)

		}

	})
	http.ListenAndServe(":8080", nil)
}
