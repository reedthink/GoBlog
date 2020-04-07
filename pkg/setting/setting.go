package setting

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

var (
	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	viper.AddConfigPath(workDir + "./conf") //引入目录
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	err = viper.ReadInConfig() //从上述的路径中读取配置文件

	if err != nil {
		log.Printf("没找到配置文件")
	}
	RunMode = viper.GetString("RUN_MODE")
	HTTPPort = viper.GetInt("server.HTTP_PORT")
	ReadTimeout = time.Duration(viper.GetInt("server.READ_TIMEOUT")) * time.Second
	WriteTimeout = time.Duration(viper.GetInt("server.WRITE_TIMEOUT")) * time.Second
	JwtSecret = viper.GetString("app.JWT_SECRET")
	PageSize = viper.GetInt("app.PAGE_SIZE")
}
