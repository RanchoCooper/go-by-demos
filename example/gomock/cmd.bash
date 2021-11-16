# 手动生成mock文件
mockgen -source=./person/male.go -destination=./mock/male_mock.go -package=mock

# 自动生成mock文件
go generate ./...

# 测试
go test ./user

## 测试覆盖率
go test -cover ./user

## 可视化界面
go test ./... -coverprofile=cover.out   # 生成profile文件
go tool cover -html=cover.go            # 生成可视化界面

