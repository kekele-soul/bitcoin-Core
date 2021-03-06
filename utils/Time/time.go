package Time

/**
该包主要是对时间戳进行数据格式化工具
*/

import "time"

const TIME_FORMAT_ONE = "2006年01月02日 15:04:05"
const TIME_FORMAT_TWO = "2006/01/02 15:04:05"
const TIME_FORMAT_THREE = "2006-01-02 15:04:05"
const TIME_FORMAT_FOUR = "2006.01.02 15:04:05"

//根据所需要的格式,生成相应格式的当前日期
func TimeNow(format string) string {
	return time.Now().Format(format)
}

// 将int64整形时间格式转换成相应的格式
func TimeFormat(sec int64, nsec int64, format string) string {
	return time.Unix(sec, nsec).Format(format)
}
