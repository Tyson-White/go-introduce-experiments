package mine

import (
	"maps"
	"myminegame/pkg"
)

type Mine struct {
	Level        int
	Miners       []Miner
	ResourcesDay int
}

func InitMine() *Mine {
	mine := &Mine{
		Level: 1,
	}

	for i := 0; i < len(pkg.ReadSettings().Levels.Mine.MinersCount); i++ {
		miner := HireMiner(1)
		// Суммируем общее добываемое кол-во ресурсов
		mine.ResourcesDay += miner.ResourcesDay
		// Добавляем первых шахтеров
		mine.Miners = append(mine.Miners, miner)
	}

	return mine
}

func (m *Mine) Work() map[string]int {

	// Список добытых за рабочий день ресурсов.
	resources := make(map[string]int)

	// Принимает мапу из добытых ресурсов
	channel := make(chan map[string]int)

	for _, miner := range m.Miners {
		// Запускаем майнеров параллельно
		// Передаем майнеру канал, куда он будет отдавать ресурсы
		go miner.Work(channel)
	}

	for i := 0; i < len(m.Miners); i++ {
		collectedResources := <-channel

		// Получаем названия всех добытых ресурсов
		names := maps.Keys(collectedResources)

		for n := range names {
			// обновляем кол-во добытых ресурсов в шахте за рабочий день
			resources[n] += collectedResources[n]
		}
	}

	close(channel)
	return resources
}
