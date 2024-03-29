package holiday

import (
	"github.com/ChaLanZi/holidays/history"
	"time"
)

func getDate(value string) (time.Time, error) {
	valueTime, err := time.Parse("2016/01/02", value)
	if err != nil {
		return time.Time{}, err
	}
	return valueTime, nil
}

func countOneHoliday(value history.OneCollection) int {
	start, _ := time.Parse("2016/01/02", value.Start)
	end, _ := time.Parse("2016/01/02", value.End)
	return int(end.Sub(start).Hours()/24) + 1
}
