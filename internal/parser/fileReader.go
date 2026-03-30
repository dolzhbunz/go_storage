package parser

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"productStorage/internal/models"
)

func ParseProductsFromFile(path string)([]*models.Product,error){ //читает файл с продуктами и парсит его в слайс продуктов
	file, err := os.Open(path) //открытие файла по указанному пути
	if err != nil{
		return nil, fmt.Errorf("Error: %w", err)
	}

	defer file.Close() //откладываем закрытие файла до выхода из функции

	var Products []*models.Product //иницилиазация слайса для хранения продуктов
	lineNum := 1 //счетчик строк
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){// посторочное чтение файла
		line := scanner.Text()

		if strings.TrimSpace(line) == ""{//пропуск пустых строк
			lineNum++
			continue
		}

		parts := strings.Split(line, ";") //разделить строку ;

		if len(parts) != 3{ //проверка кол-во полей
			fmt.Printf("[Строка: %d] Недостаточно данных для сохранения продукта", lineNum)
			lineNum++
			continue
		}

		//удаление лишних пробелов
		name := strings.TrimSpace(parts[0])
		sbin := strings.TrimSpace(parts[1])
		date := strings.TrimSpace(parts[2])

		product, err := models.NewProduct(name, sbin, date)//создание нового продукта через конструктор

		if err != nil{ //вывод предупреждения если не создалось
			fmt.Printf("Товар %s на строке %d отклонён. Error: %w\n", name, lineNum, err)
		} else{//если успешно создался, дбавляем в слайс
			Products = append(Products, product)
		}
		lineNum++

	}

	if err:= scanner.Err(); err != nil{ //сканирование на наличие ошибка
		return nil, fmt.Errorf("Ошибка чтения файла %w",err)
	}

	return Products, nil //возращаем успешные обработанные продукты
}