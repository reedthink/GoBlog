module util

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/reedthink/pkg/setting v0.0.0-00010101000000-000000000000
	github.com/unknwon/com v1.0.1
)

replace (
	github.com/reedthink/conf => /home/hhh/src/GoBlog/pkg/conf
	github.com/reedthink/middleware => /home/hhh/src/GoBlog/middleware
	github.com/reedthink/models => /home/hhh/src/GoBlog/models
	github.com/reedthink/pkg/e => /home/hhh/src/GoBlog/pkg/e
	github.com/reedthink/pkg/setting => /home/hhh/src/GoBlog/pkg/setting
	github.com/reedthink/pkg/util => /home/hhh/src/GoBlog/pkg/util
	github.com/reedthink/routers => /home/hhh/src/GoBlog/routers

)
