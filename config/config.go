package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigInstance Config

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0"

type Config struct {
	PushPlusToken  string `yaml:"pushplus_token"`
	RefreshToken   string `yaml:"refresh_token"`
	BilibiliCookie string `yaml:"bilibili_cookie"`
	KKCookie       string `yaml:"kk_cookie"`
	JdCookie       string `yaml:"jd_cookie"`
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	confFIle, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err.Error())
	}
	config := Config{}
	yaml.Unmarshal(confFIle, &config)
	ConfigInstance = config
}
_UP_A4A_11_=wb962196611548ec87860c5b4da5dc75; __pus=2a905393e939b466cfc67d63fa568c11AAQ/OlpUygPF55swJUGlJZL/TwUxSIXaQmpHwlbhtCs2KrbWMa9Gji6aNjdNpVfSrGKaOBlbzc/7XfxKEUQAuUQa; __kp=6a5e4bc0-f879-11ee-9a35-89faed88f81b; __kps=AAQ17f8XeulhE0ubkegikhWo; __ktd=394oqig8ASSlg+xzLDENZg==; __uid=AAQ17f8XeulhE0ubkegikhWo; __puus=2543f488799746c140ddbe7ab096eebeAAS5RTZ47xcyElubQpiC+0SdP2oAAEkuWuv7LM6ekkXVECWWXO/s+SmBAw4uWqJSrpjstv2hrrmOEj6HUe79QEGE3mgkTm2dgF00FpXhXCOcFHxdhd8Um6JSz2EPIHsWsapOY+YE5VI7kmK40T/G3W3+PCHgsC4AVnzewmFdoZLE0nJZMOUiWrONBXDXEnY/vB1txBSLkzEx9MCapXdXph8m
