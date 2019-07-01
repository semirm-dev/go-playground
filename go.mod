module github.com/semirm-dev/go-playground

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/semirm-dev/go-common v0.0.0-20190626103215-5558ddd5ff73
	github.com/sirupsen/logrus v1.4.2
)

replace github.com/semirm-dev/go-common v0.0.0-20190626103215-5558ddd5ff73 => ../go-common
