// ----------------------------------------------------------------------------
// The MIT License
// Simple image resizer https://github.com/Leopotam/ecs
// Copyright (c) 2020 Leopotam <leopotam@gmail.com>
// ----------------------------------------------------------------------------

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"sync"
	"time"

	"github.com/disintegration/imaging"
)

func main() {
	var scale int
	var srcDir string
	var dstDir string
	var verbose bool
	flag.IntVar(&scale, "scale", 100, "scale in percent, from 1 to 99.")
	flag.StringVar(&srcDir, "src", "", "source path, current folder by default.")
	flag.StringVar(&dstDir, "dst", "", "destination path, same as source path by default.")
	flag.BoolVar(&verbose, "v", false, "verbose output, false by default.")
	flag.Parse()
	if scale < 1 || scale >= 100 {
		flag.PrintDefaults()
		return
	}
	if len(srcDir) == 0 {
		srcDir = "."
	}
	if len(dstDir) == 0 {
		dstDir = srcDir
	}

	// save start time.
	var startTime time.Time
	if verbose {
		startTime = time.Now()
	}

	// init resizer library.
	imaging.AutoOrientation(true)

	// unwrap for ioutil.ReadDir() - no need to sort files.
	srcDirFile, err := os.Open(srcDir)
	if err != nil {
		log.Fatal(err)
	}
	files, err := srcDirFile.Readdir(-1)
	srcDirFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	for _, f := range files {
		if !f.IsDir() {
			fName := f.Name()
			if _, err := imaging.FormatFromFilename(fName); err == nil {
				wg.Add(1)
				go processFile(&wg, fName, srcDir, dstDir, scale, verbose)
			}
		}
	}
	wg.Wait()
	// print elapsed time.
	if verbose {
		fmt.Printf("elapsed time: %v\n", time.Since(startTime))
	}
}

func processFile(wg *sync.WaitGroup, fileName, srcPath, dstPath string, scale int, verbose bool) {
	defer wg.Done()
	img, err := imaging.Open(path.Join(srcPath, fileName))
	if err != nil {
		log.Println(err)
		return
	}
	srcSize := img.Bounds().Max
	dstImage := imaging.Resize(img, srcSize.X*scale/100, 0, imaging.Lanczos)
	if verbose {
		dstSize := dstImage.Bounds().Max
		fmt.Printf("[%v] (%vx%v) -> (%vx%v)\n", fileName, srcSize.X, srcSize.Y, dstSize.X, dstSize.Y)
	}
	if err := imaging.Save(dstImage, path.Join(dstPath, fileName)); err != nil {
		log.Println(err)
	}
}
