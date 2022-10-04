package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func createConfig() *Config {
	c := Config{}
	return &c
}

func getPaht() string {
	p, _ := os.Getwd()
	return filepath.Join(p, "config.yml")
}

//beforeALL 前処理、afterAll 後処理
func TestMain(m *testing.M) {
	fmt.Println("Test------START")
	m.Run()
	fmt.Println("Test------END")
}

func TestSetConfgi(t *testing.T) {
	t.Run("case01", func(t *testing.T) {
		cfg := createConfig()
		path := getPaht()
		cfg.InitCfg(path)
		if cfg.Dst != `C:\Users\en.b\Desktop\test` {
			t.Errorf("watn-->%v\n", cfg.Dst)
		}

	})
}

func TestStartWatch(t *testing.T) {
	cfg := Config{
		Src:      `C:\Users\en.b\Desktop\src\`,
		Dst:      `C:\Users\en.b\Desktop\test`,
		WatchDst: `C:\Users\en.b\Desktop\src\`,
	}
	ch := make(chan bool)
	quit := make(chan struct{})
	go startWatch(&cfg, ch)

	f := func() {
		//	ch <- true
		//	quit <- struct{}{}
		close(ch)
		close(quit)
	}

	time.AfterFunc(time.Second*10, f)
	<-quit
}

func TestFileCopy(t *testing.T) {
	cases := []struct {
		name string
		data string
		dst  string
		want bool
	}{
		{"No1", "aa", "vv", false},
		{"No2", "bb", "bb", false},
		{"No3", "cc", "tt", false},
		{"No4", `C:\Users\en.b\Desktop\src\aa.txt`, `C:\Users\en.b\Desktop\test`, false},
	}
	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			e := fileCopy(v.data, v.dst)
			if e != nil {
				t.Errorf("error-->%v", e)
			}
		})
	}
}
