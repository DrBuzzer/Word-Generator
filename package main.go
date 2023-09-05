package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var themes = []string{"ФИО", "Товары для магазина", "Названия магазинов либо складов"}

func main() {
	var wordCount int
	var theme int

	// Get the desired word count from the user
	// Получение желаемого количества слов от пользователя
	fmt.Print("Enter the desired word count for generation: ")
	_, err := fmt.Scan(&wordCount)
	if err != nil {
		log.Fatal(err)
	}

	// Get the selected theme from the user
	// Получение выбранной темы от пользователя
	fmt.Println("Choose a word theme:")
	for i, t := range themes {
		fmt.Printf("%d. %s\n", i+1, t)
	}
	fmt.Print("Enter the number of the selected theme: ")
	_, err = fmt.Scan(&theme)
	if err != nil {
		log.Fatal(err)
	}

	// Generate a list of words
	// Генерация списка слов
	words := generateWords(wordCount, themes[theme-1])

	// Export the generated word list to Excel
	// Экспорт сгенерированного списка слов в Excel
	err = exportToExcel(words)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generation and export completed successfully!")
}

func generateWords(count int, theme string) []string {
	rand.Seed(time.Now().UnixNano())

	words := make([]string, count)
	for i := 0; i < count; i++ {
		words[i] = generateWord(theme)
	}

	return words
}

func generateWord(theme string) string {
	switch theme {
	case "ФИО":
		return generateFullName()
	case "Товары для магазина":
		return generateProduct()
	case "Названия магазинов либо складов":
		return generateStoreName()
	default:
		return "Unknown"
	}
}

func generateFullName() string {
	firstNames := []string{"Иван", "Александр", "Сергей", "Мария", "Елена", "Анна"}
	lastNames := []string{"Иванов", "Петров", "Сидоров", "Смирнова", "Кузнецова", "Васильева"}

	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]

	return fmt.Sprintf("%s %s", firstName, lastName)
}

func generateProduct() string {
	products := []string{"Телефон", "Ноутбук", "Телевизор", "Холодильник", "Планшет", "Кофемашина"}

	return products[rand.Intn(len(products))]
}

func generateStoreName() string {
	storeNames := []string{"Магазин №1", "Склад А", "Магазин ABC", "Супермаркет XYZ"}

	return storeNames[rand.Intn(len(storeNames))]
}

func exportToExcel(words []string) error {
	file, err := os.Create("word_list.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Запись слов в файл CSV
	for _, word := range words {
		err := writer.Write([]string{word})
		if err != nil {
			return err
		}
	}

	return nil
}
