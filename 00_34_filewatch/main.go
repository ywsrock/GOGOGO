package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"filewatcher/utils"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
)

const CONFIG_FILENAME = "config.yml"

// 設定情報
type Config struct {
	Src      string `yaml:"src"`
	Dst      string `yaml:"dst"`
	WatchDst string `yaml:"watch"`
}

func (config *Config) InitCfg(cfgPath string) {
	f, err := os.Open(cfgPath)
	if err != nil {
		log.Fatalf("%s->%v", "configファイル見つかりません。", err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("%s->%v", "読み込みエラー。", err)
	}

	if err := yaml.Unmarshal(content, config); err != nil {
		log.Fatal(err)
	}
}

func main() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cfg := filepath.Join(dir, CONFIG_FILENAME)

	// 設定ファイル
	f := new(Config)
	f.InitCfg(cfg)

	done := make(chan bool)
	// 監視開始
	startWatch(f, done)
}

// フォルダ監視
// fsnotify はサブフォルダの監視はできません。
// サブフォルダは再帰的に監視必要
func startWatch(cfg *Config, done chan bool) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()
	fsUtil, err := utils.NewFS()
	if err != nil {
		log.Fatalln(err)
	}

	f := func(fs *utils.FS) {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				// 名前変更の場合
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					fileCopy(fs, event.Name, cfg.Dst, "RE")
				}

				// 新規作成
				if event.Op&fsnotify.Write == fsnotify.Write ||
					event.Op&fsnotify.Create == fsnotify.Create {
					fileCopy(fs, event.Name, cfg.Dst, "CP")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}
	go f(fsUtil)

	err = watcher.Add(cfg.WatchDst)
	if err != nil {
		log.Fatal(err)
	}

	<-done
}

// File Copy
func fileCopy(fs *utils.FS, srcFile string, dstFileName string, mode string) error {

	fileName := filepath.Base(srcFile)
	dstFullName := filepath.Join(dstFileName, fileName)
	log.Println(srcFile, "--->", dstFullName)

	switch mode {
	case "CP":
		info, err := os.Stat(srcFile)
		if err != nil {
			log.Fatalln(err)
		}
		if info.IsDir() {
			return fs.DirCopy(srcFile, dstFullName)
		} else {
			return fs.FileCopy(srcFile, dstFullName)
		}
	case "RE":
		return fs.FileReName(dstFullName)
	default:
		return nil
	}

	// cmd := exec.Command("cmd.exe", "/C", "COPY", srcFile, dstFullName)
	//	stderr, err := cmd.StderrPipe()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	err = cmd.Run()
	//	if errmsg, _ := ioutil.ReadAll(stderr); len(errmsg) != 0 {
	//		log.Println("error-------------->", errmsg)
	//	}
	//
	//	if err != nil {
	//		return err
	//	}
	//	return nil
}
