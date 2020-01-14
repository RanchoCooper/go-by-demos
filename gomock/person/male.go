package person

// 在接口中直接声明go generate
//go:generate mockgen -destination=../mock/male_mock.go -package=male_mock github.com/RanchoCooper/go-by-demos/gomock/person Male

type Male interface {
	Get(id int64) error
}