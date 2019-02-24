# holidays
用来学习Go的Demo ，编写一个可查询假期的API

## API List
---

GetAll

> 获取当前公布的从2010 ~ 2019 的官方节假日安排

GetByYear(year int)

> 获取房钱公布的节假日安排，按年份查询

GetByMonth(year int, month int)

> 获取当前公布的节假日安排，按年份和月份查询

GetByChName(year int, name string)

> 获取当前公布的节假日安排，按年份和中文名称查询

GetByEnName(year int, name string)

> 获取当前节假日安排，按年份和英文名称查询

GetYearHolidayCount(year int)

> 获取某年存在多少天的假期

GetMonthHolidayCount(year int, mouth int)

> 获取某月存在多少假期

IsHoliday(date string)

> 判断某天是否是节假日，某天的格式是："2006/01/02"

IsWorkDay(date string)

> 判断某天是否是工作日，某天的格式是： "2006/01/02"

IsWeekDay(date string)

> 判断某天是否是周末，某天的格式是： "2006/01/02"