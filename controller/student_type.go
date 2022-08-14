package controller

type Student struct {
	Name        string
	DoB         string
	Sex         sex
	Phone       string
	University  string
	GradelLevel Grade

	GPA            float64
	BestRewardName string

	EnglishScore   int64
	EntryTestScore float64
}

type sex string

const (
	Male   string = "Male"
	Female string = "Female"
)

type Grade string

const (
	Good   string = "good"
	Normal string = "Normal"
)
