package main

import (
	"database/sql"  
    "fmt"  
    _ "github.com/go-sql-driver/mysql"
    "container/list"
    "gopkg.in/ini.v1"
    "regexp"
    "strings"
    "bytes"
     "os/exec"
     "time"
)


func getImageTags(db *sql.DB) map[string]*list.List {
	rows, err := db.Query("select namespaces.name, repositories.name, tags.name from namespaces, repositories , tags where namespaces.id=repositories.namespace_id and tags.repository_id=repositories.id order by tags.id desc")
	checkErr(err)
	re := make(map[string]*list.List)
	for rows.Next() {
		var namespace string
		var imagename string
		var tagname string
		err = rows.Scan(&namespace, &imagename, &tagname)
		image := namespace + "/" + imagename
//		fmt.Println(image + ":" + tagname)
		
		if v, ok := re[image]; ok {
			v.PushBack(tagname)
        }else {
        	l := list.New()
        	l.PushBack(tagname)
        	re[image] = l
        }
	}
	return re
}

func checkErr(err error) {  
    if err != nil {  
        panic(err)  
    }  
}

func getImageVersion(imageName string, defaulttype string, registryVersion map[string]*list.List,  defaultConf *ini.File,  userConf map[string](map[string]string)) string {

	if v, ok := registryVersion[imageName]; ok {
		tagtype := defaulttype
		tagregex := `.*`
		if c, ok := userConf[imageName]; ok {
			if t, ok := c["type"]; ok {
				tagtype = t
			}
			if t, ok := c["regex"]; ok {
				tagregex = t
			}
			
		} else {
			imagesection := defaultConf.Section(imageName)
			if imagesection != nil {
				if imagesection.HasKey("type") {
					tagtype = imagesection.Key("type").String()
				}
				if imagesection.HasKey("regex") {
					tagregex = imagesection.Key("regex").String()
				}
			}
		}
		
		if tagtype == "none" {
			return ""
		}
		reg := regexp.MustCompile(tagregex)
		for  e := v.Front(); e != nil; e = e.Next() {
			tag, _ :=  e.Value.(string)
			if tagtype == "release" && (strings.Contains(tag, ".A.") || strings.Contains(tag, ".B.")){
				continue
			} 
			if tagtype == "test"  && strings.Contains(tag, ".B.") {
				continue
			}
			if reg.MatchString(tag) {
				return tag
			}
		}
		
    }
	return ""
}

func updateConfig(confname string, defaulttype string,  registryVersion map[string]*list.List,  defaultConf *ini.File,  userConf map[string](map[string]string)) [](map[string]string) {
	cfg, _ := ini.Load(confname)
	size := len(cfg.Sections())-1
	re := make([](map[string]string), size, size)
	index := 0
	for _, section := range cfg.Sections() {
		if section.Name() == "DEFAULT" {
			continue
		}
		imageinfo := make(map[string]string)
		oldimagetag := section.Key("image").String()
		info := strings.Split(oldimagetag, ":")
		imageName := info[0]
		
		tag := getImageVersion(imageName, defaulttype, registryVersion, defaultConf , userConf)
		newimagetag := imageName +":"+tag
		imageinfo["container"] = section.Name()
		imageinfo["old"] = oldimagetag
		if tag!="" && newimagetag != oldimagetag {
			section.Key("image").SetValue(newimagetag)
			imageinfo["new"] = newimagetag
			fmt.Println("container:" + section.Name() + " ( " + oldimagetag + " -> " + newimagetag + " )" )
		}else {
			fmt.Println("container:" + section.Name() + " ( " + oldimagetag + " )" )
		}
		re[index] = imageinfo
		index ++
	}
	cfg.SaveTo(confname)
	return re
}

func restartContainer(confpath string, containers [](map[string]string) ){
	for _, container := range containers {
		containername := container["container"]
		if _, ok := container["new"] ; ok {
			for i:=0; i<30; i++ {
				_, err := runcmd("./bootstrap -c " + confpath + " -r " + containername)
				if err == nil {
					break
				} else {
					time.Sleep(10e9)
				}
			}
		}
	}
}

func restartContainers(confpath string, containers []string ){
	for _, container := range containers {
		fmt.Println(" restart container:" + container + " ( " + "./bootstrap -c " + confpath + " -r " + container + " )" )
		for i:=0; i<30; i++ {
			_, err := runcmd("./bootstrap -c " + confpath + " -r " + container)
			if err == nil {
					break
				} else {
					time.Sleep(10e9)
				}
		}
	}
}
	
func runcmd(c string) (string, error){
    cmd := exec.Command("bash", "-c", c)
    var out_buf, err_buf bytes.Buffer
    cmd.Stdout = &out_buf
    cmd.Stderr = &err_buf
    if !strings.Contains(c, "docker login") {
        fmt.Printf("[Debug] %v\n", cmd.Args)
    } else {
        fmt.Printf("[Debug] docker login")
    }
    err := cmd.Run()
    if err != nil {
        fmt.Println("[Error] Excute cmd failed! cmd:%s \n%s\n%s\nerror: %v\n", c, out_buf.String(), err_buf.String(), err)
    }
    return out_buf.String(), err
}

