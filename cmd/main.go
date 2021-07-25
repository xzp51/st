package main

import (
	"flag"
	"fmt"
	"github.com/xzp51/st"
	"io/ioutil"
	"os"
	"path/filepath"
)

var s2t = flag.String("st", "s2t", "简体与繁体转换 \ns2t 简体=>繁体 \nt2s 繁体=>简体\n")

func main() {
	flag.Usage = usage
	flag.Parse()
	paths := flag.Args()
	if len(paths) == 0 {
		fmt.Printf("no path provide, see usage")
		os.Exit(2)
	}
	for _, path := range paths {
		switch dir, err := os.Stat(path); {
		case err != nil:
			fmt.Printf("path: %s error: %s\n", path, err)
			os.Exit(2)
		case dir.IsDir():
			if err = filepath.Walk(path, processDirFile); err != nil {
				fmt.Printf("path: %s error: %s\n", path, err)
				os.Exit(2)
			}
		default:
			if err = processFile(path); err != nil {
				fmt.Printf("path: %s error: %s\n", path, err)
				os.Exit(2)
			}
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: st [flags] [path ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func processDirFile(path string, info os.FileInfo, err error) error {
	if err == nil && !info.IsDir() {
		err = processFile(path)
	}
	return err
}

func processFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	ret := []rune(string(data))
	if *s2t == "s2t" {
		st.S2T(ret)
	} else {
		st.T2S(ret)
	}
	// 寫入文件
	return ioutil.WriteFile(path, []byte(string(ret)), 0666)
}
