package constants

//时间格式化模板
const (
	DateFormatLayout       = "2006-01-02"
	TimeFormatLayout       = "15:04:05"
	DatetimeFormatLayout   = DateFormatLayout + " " + TimeFormatLayout
	DatetimeTZFormatLayout = DateFormatLayout + "T" + TimeFormatLayout + ".000Z+08:00"
)
