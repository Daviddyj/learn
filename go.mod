module learn.go

go 1.17

require (
	github.com/armstrongli/go-bmi v0.0.1
	github.com/spf13/cobra v1.3.0
	learn.go.tools v0.0.0-00010101000000-000000000000
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect

)

replace learn.go.tools => ../learn.go.tools //本地的包模块需要去使用
replace (
	github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi
)

//replace (
//	github.com/spf13/cobra =>github.com/spf13/cobra v1.2.0
//)
