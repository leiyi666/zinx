package main

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"io/ioutil"
	"strconv"
)

func main() {
	fileName := "G:\\Golang\\src\\go_code\\fund\\8\\"
	fileNames := make([]string, 0)
	for i := 1; i < 48; i++ {
		temp := fileName + strconv.Itoa(i) + ".jpg"
		fileNames = append(fileNames, temp)
	}
	res, err := makeThumbnails5(fileNames)
	fmt.Println(res)
	fmt.Println(err)
}

func makeThumbnails1(filenames []string) {
	for _, f := range filenames {
		go makeThumbnail(f) // NOTE: ignoring errors
	}
}

func makeThumbnails2(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			makeThumbnail(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := makeThumbnail(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: incorrect: goroutine leak!
		}
	}

	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = makeThumbnail(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func makeThumbnail(filename string) (string, error) {
	//读取本地文件
	imgData, err := ioutil.ReadFile(filename)
	if err != nil {
		return filename, err
	}
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		return filename, err
	}
	//生成缩略图
	image = imaging.Resize(image, 1024, 1024, imaging.Lanczos)
	err = imaging.Save(image, filename)
	if err != nil {
		return filename, err
	}
	return filename, nil
}
