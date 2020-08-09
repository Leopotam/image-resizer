// ----------------------------------------------------------------------------
// The MIT License
// Simple image resizer https://github.com/Leopotam/ecs
// Copyright (c) 2020 Leopotam <leopotam@gmail.com>
// ----------------------------------------------------------------------------

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"sync"

	"github.com/disintegration/imaging"
)

func main() {
	var scale int
	var srcDir string
	var dstDir string
	flag.IntVar(&scale, "scale", 100, "scale in percent, from 1 to 99.")
	flag.StringVar(&srcDir, "src", "", "source path, current folder by default.")
	flag.StringVar(&dstDir, "dst", "", "destination path, same as source path by default.")
	flag.Parse()
	if scale < 1 || scale >= 100 {
		flag.PrintDefaults()
	}
	if len(srcDir) == 0 {
		srcDir = "."
	}
	if len(dstDir) == 0 {
		dstDir = srcDir
	}

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	for _, f := range files {
		if !f.IsDir() {
			if _, err := imaging.FormatFromFilename(f.Name()); err == nil {
				wg.Add(1)
				go processFile(&wg, f.Name(), srcDir, dstDir, scale)
			}
		}
	}
	wg.Wait()
}

func processFile(wg *sync.WaitGroup, fileName, srcPath, dstPath string, scale int) {
	defer wg.Done()
	img, err := imaging.Open(path.Join(srcPath, fileName))
	if err != nil {
		log.Println(err)
		return
	}
	srcSize := img.Bounds().Max
	dstImage := imaging.Resize(img, srcSize.X*scale/100, 0, imaging.Lanczos)
	dstSize := dstImage.Bounds().Max
	fmt.Printf("[%v] (%vx%v) -> (%vx%v)\n", fileName, srcSize.X, srcSize.Y, dstSize.X, dstSize.Y)
	if err := imaging.Save(dstImage, path.Join(dstPath, fileName)); err != nil {
		log.Println(err)
	}
}
