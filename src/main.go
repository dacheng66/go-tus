package main

import (
	"./tus"
	"github.com/fsnotify/fsnotify"
	"log"
	"io/ioutil"
	"encoding/json"
	"os"
	"mime"
	"path"
	"fmt"
)

func main() {
	//创建一个监控对象
	watch, err := fsnotify.NewWatcher();
	if err != nil {
		log.Fatal(err);
	}
	defer watch.Close();
	//添加要监控的对象，文件或文件夹
	err = watch.Add("c:/sunmoon/program/Git/sunmoon/go/tus/src");
	if err != nil {
		log.Fatal(err);
	}
	//我们另启一个goroutine来处理监控对象的事件
	go func() {
		for {
			select {
			case ev := <-watch.Events:
				{
					//判断事件发生的类型，如下5种
					if ev.Op&fsnotify.Create == fsnotify.Create {
						log.Println("创建文件 : ", ev.Name);
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						log.Println("写入文件 : ", ev.Name);
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						log.Println("删除文件 : ", ev.Name);
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						log.Println("重命名文件 : ", ev.Name);
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						log.Println("修改权限 : ", ev.Name);
					}
				}
			case err := <-watch.Errors:
				{
					log.Println("error : ", err);
					return;
				}
			}
		}
	}();




	data, e := ioutil.ReadFile("conf/conf.json")
	if e != nil {
		panic(e)
	}
	conf := tus.Conf{}
	if e = json.Unmarshal(data,&conf); e != nil {
		panic(e)
	}

	filepath := "c:/sunmoon/program/Git/sunmoon/go/tus/src/tus.png"

	mimetype := mime.TypeByExtension(path.Ext(filepath))
	fmt.Println("mimetype",mimetype)
	f, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	// create the tus client.
	client, _ := tus.NewClient(conf.UploadUrl, nil)

	// create an upload from a file.
	upload, _ := tus.NewUploadFromFile(f,conf.Meta,mimetype)

	// create the uploader.
	uploader, _ := client.CreateUpload(upload)

	// start the uploading process.
	uploader.Upload()

	//阻塞
	select {};
}