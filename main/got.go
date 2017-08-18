package main
//
//import (
//	"fmt"
//	"gopkg.in/ini.v1"
//	"flag"
//	"log"
//)
//
//var updateConfigFlag *string = flag.String("uc", "update.ini", "升级工具配置文件")
//
//var (
//	gtype string
//	registrydb string
//	middledb string
//)
//
//func main() {
//	flag.Parse()
//	fmt.Println(*updateConfigFlag)
//
//	cfg, err := ini.Load(*updateConfigFlag)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	def_section := cfg.Section("")
//
//	if def_section.HasKey("type") {
//		gtype = def_section.Key("type").String()
//	} else {
//		gtype = "develop"
//	}
//
//	registrydb = def_section.Key("registry").String()
//	middledb = def_section.Key("platform").String()
//
//	fmt.Println(registrydb)
//	fmt.Println(middledb)
//}
