package weloop

// 公共返回
type CommonResponse struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

// 登录返回结构体(只关注关心的部分)
type LoginResponse struct {
	CommonResponse
	AccessToken string `json:"accessToken"`
	Account     string `json:"account"`
	UserId      int    `json:"userId"`
	RegTime     int64  `json:"regTime"`
}

// token是否过期接口
type TokenValidResponse struct {
	CommonResponse
	Valid string `json:"tokenIsValid"`
}

const (
	TokenValid   = "Y"
	TokenExpired = "N"
)

type DailySourceResponse struct {
	UserId int `json:"userId"`
	Data   struct {
		Activity   []DailyActivity `json:"activity"`
		Gain       []DailyGain     `json:"dailyGain"`
		HeartRates []HeartRate     `json:"heartRates"`
	}
}

type ActivityMode int

const (
	// 摘下
	TakeOff ActivityMode = iota + 1
	// 睡眠
	Sleep
	// 散步
	Walk
	// 健走
	FastWalk
	// 低运动量
	LowRun
	// 运动
	Run
)

// 详情
type DailyActivity struct {
	StartTime        int64        `json:"startTimestamp"`
	EndTime          int64        `json:"endTimestamp"`
	Mode             ActivityMode `json:"mode"`
	Calorie          int          `json:"calorie"`
	Distance         int          `json:"distance"`
	DeepSleepSecond  int          `json:"dsTimes"`
	LightSleepSecond int          `json:"lsTimes"`
	WakeSecond       int          `json:"wakeTimes"`
	WakeNum          int          `json:"wakNum"`
	SleepGraph       string       `json:"graphValue"`
	SleepGraphArr    []int        `json:"-"`
	Step             int          `json:"stepCount"`
}

// 每日活动
type DailyGain struct {
	AvgHeartRate  int    `json:"avgHeartRate"`
	MaxHeartRate  int    `json:"maxHeartRate"`
	MinHeartRate  int    `json:"minHeartRate"`
	Calorie       int    `json:"calorie"`
	Distance      int    `json:"distance"`
	Goal          int64  `json:"goal"`
	Date          int64  `json:"happenDate"`
	Step          int    `json:"step"`
	StepInHour    string `json:"stepInHour"`
	StepInHourArr []int  `json:"-"`
}

// 心率
type HeartRate struct {
	Date  int64 `json:"happenDate"`
	Times int64 `json:"times"`
}
