module learn

go 1.17

require (
	github.com/armstrongli/go-bmi v0.0.1
	github.com/spf13/cobra v1.3.0
	learn.go.tools v0.0.0-00010101000000-000000000000
)

require (
	github.com/favadi/protoc-go-inject-tag v1.3.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.3.2 // indirect
	gorm.io/gorm v1.23.2 // indirect

)

replace learn.go.tools => ../learn.go.tools //本地的包模块需要去使用

replace github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi

//replace (
//	github.com/spf13/cobra =>github.com/spf13/cobra v1.2.0
//)
