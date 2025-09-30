package pkg

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	Tax       int `json:"tax"`
	WorkDays  int `json:"work_days"`
	WorkHours int `json:"work_hours"`
	Resources []struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	} `json:"resources"`
	Levels struct {
		Warehouse struct {
			Capacity []int `json:"capacity"`
		} `json:"warehouse"`
		Mine struct {
			MinersCount []int `json:"miners_count"`
		} `json:"mine"`
		Miner struct {
			ResourcesInDay []int `json:"resources_in_day"`
		} `json:"miner"`
	} `json:"levels"`
	Logs bool `json:"logs"`
}

func ReadSettings() Settings {
	file, err := os.ReadFile("./configs/settings.json")

	if err != nil {
		fmt.Println(err.Error())
		return Settings{}
	}

	var data Settings
	json.Unmarshal(file, &data)

	return data
}
