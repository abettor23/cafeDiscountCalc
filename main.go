package main

import (
	"fmt"
)

func main() {
	fmt.Println("Мы были рады обслужить Вас в нашем ресторане, давайте посчитаем Ваш чек.")
	fmt.Println("Введите, пожалуйста, день недели от 1 (понедельник) до 7 (воскресенье):")
	var dayWeek int
	fmt.Scan(&dayWeek)

	fmt.Println("На какую сумму вы наели? (будьте честны с нами!)")
	var totalCheck int
	fmt.Scan(&totalCheck)

	if dayWeek > 7 {
		fmt.Println("Вы ввели какую то ерунду. Попробуйте попасть по цифрам правильно.")
	} else if dayWeek > 5 {
		if totalCheck >= 10000 {
			fmt.Print("Вы наели на ", totalCheck, "руб., получите скидку 10%!\n")
			totalCheck = totalCheck - (totalCheck*10)/100
			fmt.Print("Сумма к оплате: ", totalCheck, "руб.\n")
		} else if totalCheck > 0 {
			fmt.Print("Вы наели на ", totalCheck, "руб., скидка вам не положена! (надо больле есть)\n")
			fmt.Print("Сумма к оплате: ", totalCheck, "руб.\n")
		} else {
			fmt.Println("Вы ввели какую то ерунду. Попробуйте попасть по цифрам правильно.")
		}
	} else if dayWeek >= 1 {
		fmt.Print("Вы наели на ", totalCheck, "руб., получите скидку 10% в будний день!\n")
		totalCheck = totalCheck - (totalCheck*10)/100
		fmt.Print("Сумма к оплате: ", totalCheck, "руб.\n")
	} else {
		fmt.Println("Вы ввели какую то ерунду. Попробуйте попасть по цифрам правильно.")
	}
}
