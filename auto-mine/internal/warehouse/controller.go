package warehouse

import (
	"myminegame/logger"
	"myminegame/pkg"
)

type WarehouseController struct {
	// Список складов
	List [1]*Warehouse
}

func CreateController() *WarehouseController {
	c := &WarehouseController{}

	// TODO: Нужно ограничение по уровням.
	for i := 0; i < 1; i++ {
		w := CreateWarehouse(1)
		c.List[i] = &w
	}

	return c
}

// Возвращает канал, на который передаются добытые ресурсы.
func (wc *WarehouseController) CollectResources(times int) chan map[string]int {
	// TODO: Нужно сделать равномерное распределение ресурсов по складам

	channel := make(chan map[string]int)

	go func() {
		// Цикл крутится кол-во рабочих дней, чтобы сохранить результат для каждого дня
		for i := 0; i < times; i++ {
			resources := <-channel

			for _, w := range wc.List {
				remains, ok := w.PutResources(resources)

				if !ok {
					// Передаем оставшиеся запасы на след. склад
					resources = remains

					continue
				}

				break
			}
		}
	}()

	return channel
}

// Выводит лог для всех складов
func (w *WarehouseController) Info(day int) {
	logger := logger.NewLogger()

	logger.Line("")
	logger.Line("Информация по складам за", day, "день: ")

	for index, warehouse := range w.List {
		logger.Line("")
		logger.Line("__ Склад", index+1, "__")
		logger.Line("")
		logger.Line("Заполненность:", warehouse.Capacity-warehouse.AvailableSpace(), "/", warehouse.Capacity)
		logger.Line("Ресурсы:", warehouse.Collected)
	}
}

// Возвращает заработанную сумму с вычетом налога.
func (w *WarehouseController) Sell() int {
	summary := 0

	for _, warehouse := range w.List {
		summary += warehouse.CalculateIncome()
	}

	summaryWithTax := summary - (summary / 100 * pkg.ReadSettings().Tax)

	return summaryWithTax
}
