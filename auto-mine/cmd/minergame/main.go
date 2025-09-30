package main

import (
	"myminegame/internal/mine"
	"myminegame/internal/warehouse"
	"myminegame/logger"
	"myminegame/pkg"
)

type Game struct {
}

func main() {

	newGame := Game{}

	newGame.Go()
}

func (game *Game) Go() {
	logger := logger.NewLogger()

	collector := warehouse.CreateController()
	mine := mine.InitMine()

	collectedMoney := 0
	days := 0

	// Цикл бесконечно крутит программу. Пересоздавая каналы и подсчитывая результаты.
	for {

		channel := collector.CollectResources(pkg.ReadSettings().WorkDays)

		for i := pkg.ReadSettings().WorkDays; i > 0; i-- {
			// Когда рабочий день на шахте закончился, результат отправляется в Контроллер Склада
			channel <- mine.Work()

			collector.Info(days)

			// Считаем прошедшие дни
			days += 1
		}

		collectedMoney += collector.Sell()

		// Закрываем канал, чтобы не было лишних утечек.
		close(channel)
		logger.Line("Заработано за", days, "дней: ", collectedMoney, "руб.")

	}

}
