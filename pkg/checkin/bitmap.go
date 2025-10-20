package checkin

import (
	"math/bits"
	"time"
)

// 签到记录管理器（基于位图）
type CheckInManager struct {
	year       int      // 记录的年份
	bitMap     []uint64 // 位图存储（每个uint64存储64天的状态）
	daysInYear int      // 当年的总天数（考虑闰年）
}

// 初始化指定年份的签到管理器
func NewCheckInManager(year int) *CheckInManager {
	// 判断是否为闰年（2月有29天）
	isLeap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)
	days := 365
	if isLeap {
		days = 366
	}

	// 计算需要多少个uint64（每个存储64天）
	numUint64 := (days + 63) / 64 // 向上取整
	return &CheckInManager{
		year:       year,
		bitMap:     make([]uint64, numUint64),
		daysInYear: days,
	}
}

// 获取某天在当年的第几天（1-365/366）
func getDayOfYear(t time.Time) int {
	return t.YearDay()
}

// 签到：标记某天为已签到
func (m *CheckInManager) CheckIn(t time.Time) bool {
	// 校验年份是否匹配
	if t.Year() != m.year {
		return false
	}

	day := getDayOfYear(t)
	if day < 1 || day > m.daysInYear {
		return false
	}

	// 计算该天在哪个uint64和哪个位
	index := (day - 1) / 64  // 第几个uint64（从0开始）
	bitPos := (day - 1) % 64 // 该uint64中的第几位（0-63）

	// 用位或操作标记为1（已签到）
	m.bitMap[index] |= 1 << bitPos
	return true
}

// 查询某天是否签到
func (m *CheckInManager) IsChecked(t time.Time) bool {
	if t.Year() != m.year {
		return false
	}

	day := getDayOfYear(t)
	if day < 1 || day > m.daysInYear {
		return false
	}

	index := (day - 1) / 64
	bitPos := (day - 1) % 64

	// 用位与操作判断该位是否为1
	return (m.bitMap[index] & (1 << bitPos)) != 0
}

// 统计当年总签到天数
func (m *CheckInManager) TotalCheckIns() int {
	total := 0
	// 统计每个uint64中1的个数（内置函数快速计算）
	for _, u := range m.bitMap {
		total += bits.OnesCount64(u)
	}
	return total
}

// 获取当前连续签到天数（截至今天）
func (m *CheckInManager) ContinuousCheckInsToday() int {
	today := time.Now()
	if today.Year() != m.year {
		return 0
	}

	todayDay := getDayOfYear(today)
	count := 0

	// 从今天往前遍历，直到遇到未签到的日子
	for day := todayDay; day >= 1; day-- {
		index := (day - 1) / 64
		bitPos := (day - 1) % 64

		if (m.bitMap[index] & (1 << bitPos)) == 0 {
			break // 遇到未签到，终止统计
		}
		count++
	}
	return count
}
