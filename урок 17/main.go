package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	menu := map[string]float64{
		"Эспрессо": 800,
		"Латте":    1200,
		"Капучино": 1100,
		"Сэндвич":  1500,
		"Круассан": 900,
	}

	order := []string{}

	fmt.Println("Добро пожаловать в наше кафе!")
	fmt.Println("Меню:")

	for dish, price := range menu {
		fmt.Printf("%s - %.0f тенге\n", dish, price)
	}

	fmt.Println("\nВведите название блюда для заказа.")
	fmt.Println("Для завершения введите exit.")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Print("Ваш заказ: ")
		dish := strings.TrimSpace(scanner.Text())

		if strings.ToLower(dish) == "exit" {
			break
		}

		if price, exists := menu[dish]; exists {
			order = append(order, dish)
			fmt.Printf("%s добавлено в корзину (%.0f тенге)\n", dish, price)
		} else {
			fmt.Println("К сожалению, этого блюда нет в меню")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения ввода: %v\n", err)
	}

	fmt.Println("\n========== ФИНАЛЬНЫЙ ЧЕК ==========")

	var total float64

	if len(order) == 0 {
		fmt.Println("Вы ничего не заказали.")
	}

	for i, dish := range order {
		price := menu[dish]
		fmt.Printf("%d. %s - %.0f тенге\n", i+1, dish, price)
		total += price
	}

	fmt.Printf("\nСумма без скидки: %.0f тенге\n", total)

	discount := float64(0)

	if total > 5000 {
		discount = total * 0.10
		fmt.Printf("Скидка 10%%: %.0f тенге\n", discount)
	} else {
		fmt.Println("Скидка: 0 тенге")
	}

	fmt.Printf("Итого к оплате: %.0f тенге\n", total-discount)
	fmt.Println("====================================")
}
