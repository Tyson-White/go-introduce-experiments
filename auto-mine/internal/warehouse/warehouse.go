package warehouse

import (
	"math"
	"math/rand"
	"myminegame/pkg"
)

type Warehouse struct {
	id    int
	Level int

	// Общая вместимость склада зависит от уровня
	Capacity int

	// Добытые ресурсы: уголь, серебро, золото
	Collected map[string]int
}

func CreateWarehouse(level int) Warehouse {
	return Warehouse{
		id:        rand.Int(),
		Level:     level,
		Capacity:  pkg.ReadSettings().Levels.Warehouse.Capacity[level],
		Collected: make(map[string]int),
	}
}

// Возвращает остаток ресурсов, которые не поместились на склад.
// Возвращает false если не все ресурсы поместились
func (w *Warehouse) PutResources(resources map[string]int) (map[string]int, bool) {

	freeSpace := w.AvailableSpace()

	for name, amount := range resources {
		remain := freeSpace - amount

		if remain <= 0 {
			w.Collected[name] += resources[name] - int(math.Abs(float64(remain)))
			resources[name] = int(math.Abs(float64(remain)))

			return resources, false
		}

		w.Collected[name] += resources[name]
	}

	return resources, true
}

func (w *Warehouse) AvailableSpace() int {
	completed := 0

	for _, amount := range w.Collected {
		completed += amount
	}

	return w.Capacity - completed
}

func (w *Warehouse) CalculateIncome() int {

	summary := 0

	// Map - название/цена

	for _, res := range pkg.ReadSettings().Resources {
		summary += w.Collected[res.Name] * res.Price
	}

	w.Collected = make(map[string]int)
	return summary
}
