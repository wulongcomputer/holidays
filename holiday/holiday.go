package holiday

import (
	"github.com/ChaLanZi/holidays/history"
	"strings"
	"time"
)

// 获取当前公布的从2010 ~ 2019 的官方节假日安排
func GetAll() history.CollectionYearHistory {
	return history.FetchCollectionYearHistory()
}

// 获取房钱公布的节假日安排，按年份查询
func GetByYear(year int) []history.OneCollection {
	var index int
	nowYear, _, _ := time.Now().Date()
	if year > nowYear+1 {
		return nil
	}
	index = nowYear + 1 - year

	return history.FetchCollectionYearHistory().Data[index]
}

// 获取当前公布的节假日安排，按年份和月份查询
func GetByMonth(year int, month int) []history.OneCollection {
	if month < 1 || month > 12 {
		return nil
	}

	collections := GetByYear(year)
	var data []history.OneCollection
	for _, collection := range collections {
		collectionTime, _ := time.Parse("2016/01/02", collection.End)
		if int(collectionTime.Month()) == month {
			data = append(data, collection)
		}
	}
	return data
}

// 获取当前公布的节假日安排，按年份和中文名称查询
func GetByChName(year int, name string) []history.OneCollection {
	collections := GetByYear(year)
	var data []history.OneCollection
	for _, collection := range collections {
		if strings.Contains(collection.ChName, name) {
			data = append(data, collection)
		}
	}
	return data
}

// 获取当前节假日安排，按年份和英文名称查询
func GetByEnName(year int, name string) []history.OneCollection {
	collections := GetByYear(year)
	var data []history.OneCollection
	for _, collection := range collections {
		if strings.Contains(collection.EnName, name) {
			data = append(data, collection)
		}
	}
	return data
}

// 获取某年存在多少天的假期
func GetYearHolidayCount(year int) int {
	collections := GetByYear(year)
	var count int
	for _, collection := range collections {
		count += countOneHoliday(collection)
	}
	return count
}

// 获取某月存在多少的假期
func GetMouthHolidayCount(year int, month int) int {
	collections := GetByMonth(year, month)
	var count int
	for _, collection := range collections {

		count += countOneHoliday(collection)
	}

	return count
}

//  判断某天是否是节假日，某天的格式是："2006/01/02"
func IsHoliday(date string) bool {
	collectionTime, err := time.Parse("2016/01/02", date)
	if err != nil {
		return false
	}

	nowYear, _, _ := time.Now().Date()
	if collectionTime.Year() > nowYear+1 {
		return false
	}
	collections := GetByYear(collectionTime.Year())
	for _, collection := range collections {
		startDate, _ := getDate(collection.Start)
		endDate, _ := getDate(collection.End)
		if collectionTime.Unix() > startDate.Unix() && collectionTime.Unix() <= endDate.Unix() {
			return true
		}
	}
	return false
}

// 判断某天是否是工作日，某天的格式是： "2006/01/02"
func IsWorkDay(date string) bool {
	if IsHoliday(date) {
		return false
	}
	collectionTime, err := time.Parse("2016/01/02", date)
	if err != nil {
		return false
	}

	isWorkDay := int(collectionTime.Weekday())
	if isWorkDay == 0 || isWorkDay == 6 {
		return false
	}
	return true
}

// 判断某天是否是周末，某天的格式是： "2006/01/02"
func IsWeekDay(date string) bool {
	return !IsWorkDay(date) && !IsHoliday(date)
}
