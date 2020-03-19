module middleware

go 1.13

require github.com/gin-gonic/gin v1.5.0

replace (
	github.com/reedthink/conf => /home/hhh/src/GoBlog/pkg/conf
	github.com/reedthink/middleware => /home/hhh/src/GoBlog/middleware
	github.com/reedthink/models => /home/hhh/src/GoBlog/models
	github.com/reedthink/pkg/e => /home/hhh/src/GoBlog/pkg/e
	github.com/reedthink/pkg/setting => /home/hhh/src/GoBlog/pkg/setting
	github.com/reedthink/pkg/util => /home/hhh/src/GoBlog/pkg/util
	github.com/reedthink/routers => /home/hhh/src/GoBlog/routers
)
