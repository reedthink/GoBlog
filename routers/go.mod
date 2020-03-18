module routers

go 1.13

require (
	github.com/EDDYCJY/go-gin-example v0.0.0-20200316052038-a2781e62bfc6
	github.com/astaxie/beego v1.12.1
	github.com/gin-gonic/gin v1.5.0
	github.com/jinzhu/gorm v1.9.12
	github.com/reedthink/models v0.0.0-00010101000000-000000000000
	github.com/reedthink/pkg/e v0.0.0-00010101000000-000000000000
	github.com/reedthink/pkg/util v0.0.0-00010101000000-000000000000
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
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
