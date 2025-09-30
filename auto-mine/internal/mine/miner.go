package mine

import (
	"fmt"
	"math/rand"
	"myminegame/logger"
	"myminegame/pkg"
	"time"
)

type Miner struct {
	id    int
	Level int
	// Макс. кол-во ресурсов добываемых шахтером за рабочий день
	ResourcesDay int

	// Кол-во ресурсов добытых шахтером
	Collected int
}

func HireMiner(level int) Miner {
	return Miner{
		Level:        level,
		id:           rand.Int(),
		ResourcesDay: pkg.ReadSettings().Levels.Miner.ResourcesInDay[level],
		Collected:    0,
	}
}

func (m *Miner) Work(channel chan map[string]int) {
	logger := logger.NewLogger()
	time.Sleep(time.Duration(pkg.ReadSettings().WorkHours) * time.Second)

	collected := make(map[string]int)

	// Определяем сколько ресурсов добудет шахтер в общем за итерацию
	remains := 1 + rand.Intn(m.ResourcesDay)

	// Рандомно распределяем добытые ресурсы
	allResources := pkg.ReadSettings().Resources

	for _, res := range allResources {
		сol := rand.Intn(remains)
		remains -= сol
		collected[res.Name] = сol
	}

	// // Остаток определяем как добытый уголь
	if remains > 0 {
		collected[allResources[0].Name] += remains
	}

	logger.Line("Майнер ", m.id, "Собрал:")
	for _, res := range allResources {
		logger.Line(collected[res.Name], res.Name)
	}

	fmt.Println("")

	channel <- collected
}
