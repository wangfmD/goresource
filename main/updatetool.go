package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "gopkg.in/ini.v1"
    "log"
    "net/http"
	 "fmt"
	 "flag"
	 "strconv"
	 "strings"
	 "encoding/json"
	 "regexp"
)

var bootStrapConfigFlag * string = flag.String("c", "config.example.ini", "bootstrap工具配置文件")
var updateConfigFlag * string = flag.String("uc", "update.ini", "升级工具配置文件")
var portFlag * int = flag.Int("p", 0, "程序http服务端口")
var restartall * bool = flag.Bool("a", true, "默认全部重启")
var gtype string

var registrydb string
var middledb string
var registryurl string

func main() {
	flag.Parse()
	
	cfg1, _ := ini.Load(*bootStrapConfigFlag)
	def_section1 := cfg1.Section("")
	registryurl = def_section1.Key("registry").Value()
	
	cfg, err := ini.Load(*updateConfigFlag)
    if err != nil {
        log.Fatalln(err)
    }
    def_section := cfg.Section("")

    if def_section.HasKey("type") {
        gtype = def_section.Key("type").String()
    } else {
    	gtype = "develop"
    }
    
    registrydb = def_section.Key("registry").String()
    middledb = def_section.Key("platform").String()
    
	if *portFlag == 0 {
		db, err := sql.Open("mysql", registrydb)
		checkErr(err)
	    tags := getImageTags(db)
	    result := updateConfig(*bootStrapConfigFlag, gtype, tags, cfg , nil)
	    if * restartall {
	    	runcmd("./bootstrap -c " + *bootStrapConfigFlag)
	    } else {
	  	    restartContainer(*bootStrapConfigFlag, result)
	    }
	    db.Close()
	} else {
	    mainMux.m = make(map[string]func(*http.Request, map[string]interface{})(map[string]interface{}))
	    HandleFunc("update", webupdate)
	    HandleFunc("restart", webRestartContainer)
	    HandleFunc("restartall", webRestartAll)
	    HandleFunc("initdb", reInitDb)
	    HandleFunc("containerstatus", containerStatus)
		http.HandleFunc("/update.do", Handle)
	    http.ListenAndServe(":" + strconv.Itoa(*portFlag) , nil)
	}
}

func containerStatus(req *http.Request, m map[string]interface{})(map[string]interface{}) {
	fmt.Println("-------  get container status---------------")
	re := make(map[string]interface{})
	v, err := runcmd("docker ps -a | grep " + registryurl)
	if err != nil {
		re["status"] = "fail"
		re["msg"] = err.Error()
		return re
	}
	reg := regexp.MustCompile(` +`)
	r := reg.ReplaceAllString(v, " ")
	contianers := strings.Split(r, "\n")
	containerMap := make(map[string](map[string]string))
	
	for _, v := range contianers {
		if v=="" {
            continue
        }
		inf := make(map[string]string)
		
		infos := strings.Split(string(v), " ")
        imageversion := infos[1][(strings.Index(infos[1], "/")+1):]
        contianername := infos[len(infos)-1]
        status := "running"
		if ! strings.Contains(v, "ago Up") {
			status = "stop"
		}
        inf["image"] = imageversion
        inf["status"] = status
        
        containerMap[contianername] = inf
	}
	
	cfg, _ := ini.Load(*bootStrapConfigFlag)
	startover := true
	for _, section := range cfg.Sections() {
		contianername := section.Name()
		if contianername == "DEFAULT" {
			continue
		}
		imagetag := section.Key("image").String()
		
		if data, ok := containerMap[contianername] ; ok {
			if imagetag != data["image"] {
				startover = false
			}
			data["cfg"] = imagetag
			if data["status"] == "stop" && contianername != "middledatabase" {
				startover = false
			}
		} else {
			inf := make(map[string]string)
	        inf["status"] = "none"
	        inf["cfg"] = imagetag
	        containerMap[contianername] = inf
			startover = false
		}
	}
	
	v, err = runcmd("ps -ef | grep bootstrap | grep -v grep | grep -v java")
	fmt.Printf("bootstrap status:" + v)
	if v != "" {
		startover = false
	}
	
	re["status"] = "suc"
	re["containerinfo"] = containerMap
	re["startover"] = startover
	return re
}

func reInitDb(req *http.Request, m map[string]interface{})(map[string]interface{}) {
	fmt.Println("------- init db:" + middledb + "---------------")
	re := make(map[string]interface{})
	db, err := sql.Open("mysql", middledb)
	checkErr(err)
	_, err = db.Exec("drop database middle")
	checkErr(err)
	db.Close()
	runcmd("./bootstrap -c " + *bootStrapConfigFlag + " -r middledatabase")
	re["status"] = "suc"
	return re
}

func webRestartAll(req *http.Request, m map[string]interface{})(map[string]interface{}) {
	re := make(map[string]interface{})
	
	block := false
	if data, ok := m["block"]; ok {
		block, _ = data.(bool)
	}
	if block {
		runcmd("./bootstrap -c " + *bootStrapConfigFlag)
	} else {
		go runcmd("./bootstrap -c " + *bootStrapConfigFlag)
	}
	re["status"] = "suc"
	return re
}



func webRestartContainer(req *http.Request, m map[string]interface{})(map[string]interface{}) {
	fmt.Println("-------------restart containers--------------" )
	re := make(map[string]interface{})
	var containers []string
	if data, ok := m["data"]; ok {
		fmt.Println(data )
		cs , _ := data.([]interface{})
		containers = make([]string, len(cs), len(cs))
		for index, c := range cs {
			cc , _ := c.(string)
			containers[index] = cc
		} 
	}
	fmt.Println(containers )
	block := false
	if data, ok := m["block"]; ok {
		block, _ = data.(bool)
	}
	if block {
		restartContainers(*bootStrapConfigFlag, containers)
	} else {
		go restartContainers(*bootStrapConfigFlag, containers)
	}
	re["status"] = "suc"
	return re
}

func webupdate(req *http.Request, m map[string]interface{})(map[string]interface{}) {
 	re := make(map[string]interface{})
 	defaulttype := gtype
 	var conf map[string](map[string]string)
 	if data, ok := m["data"]; ok {
 		fmt.Println(data)
 		datamap, _ := data.(map[string]interface{})
 		if t, ok := datamap["type"]; ok {
 			defaulttype, _ = t.(string)
 		}
 		if t, ok := datamap["images"]; ok {
 			d, _ := json.Marshal(t)
 			json.Unmarshal(d, &conf)
 	//		conf, _ = t.(map[string](map[string]string))
 		}
 	}
 	fmt.Println( " -----------update data-------------- ")
	fmt.Println( conf)
	cfg, _ := ini.Load(*updateConfigFlag)
	
	db, err := sql.Open("mysql", registrydb)
	checkErr(err)
    tags := getImageTags(db)
    db.Close()
    
 	result := updateConfig(*bootStrapConfigFlag, defaulttype, tags, cfg , conf)
 	
 	block := false
	if data, ok := m["block"]; ok {
		block, _ = data.(bool)
	}
	restart:=false
 	if data, ok := m["restart"]; ok {
		restart, _ = data.(bool)
	}
 	
 	if restart {
 		if block {
		 	restartContainer(*bootStrapConfigFlag, result)
 		} else {
 			go restartContainer(*bootStrapConfigFlag, result)
 		}
 	}
 	
 	re["status"] = "suc"
 	re["result"] = result
 	
 	return re
 }




