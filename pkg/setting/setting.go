package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HttpPort int

	ReadTimeout time.Duration

	WriteTimeout time.Duration

	PageSize int

	JwtSecret string

)

func init()  {
	var err error
	Cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}


	LoadBase(Cfg)
	LoadServe(Cfg)
	LoadApp(Cfg)
}

func LoadBase(Cfg *ini.File)  {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServe(Cfg *ini.File)  {
	serve, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HttpPort = serve.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(serve.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(serve.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp(Cfg *ini.File)  {
	app, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	PageSize = app.Key("PAGE_SIZE").MustInt(10)

	JwtSecret = app.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
}
