package quest

type Manager struct {
	MainQuests          map[int]*GameMainQuest
	AcceptProgressLists map[int][]int
	LoggedQuests        []int
	LastHourCheck       int64
	LastDayCheck        int64
}
