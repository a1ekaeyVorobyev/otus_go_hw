package hw10

import (
	"errors"
	"fmt"
	"github.com/cheggaaa/pb"
	"io"
	"os"
)

func GetFileSize(f *os.File) (int64, error) {

	fi, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func FileExists(filename string) bool {

	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Copy(from string, to string, limit int, offset int, rewriteFile bool) error {

	fileSource, err := os.Open(from)
	bufferSize := 64
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("File not found")
		}
		return err
	}

	if !rewriteFile {
		if FileExists(to) {
			return errors.New("File exist")
		}

	}

	fileDestination, err := os.Create(to)

	if err != nil {
		return fmt.Errorf(err.Error())
	} else {
		errors.New("File exists")
	}

	defer fileSource.Close()
	defer fileDestination.Close()
	data := make([]byte, bufferSize)
	var sizeFile int64
	sizeFile, err = GetFileSize(fileSource)
	if int64(limit) < sizeFile{
		sizeFile = int64(limit)
	}
	if err != nil {
		return err
	}

	if int(sizeFile) < offset {
		return errors.New("sizeFile <offset ")
	}
	if int(sizeFile) <= bufferSize{
		bufferSize = bufferSize/int(sizeFile)/100zzz
	}
	bar := pb.New64(sizeFile - int64(bufferSize))
	bar.Add(offset)
	bar.Start()

	for {
		n, err := fileSource.ReadAt(data, int64(bufferSize))
		fileDestination.Write(data[:n])

		offset += limit
		if err == io.EOF {
			//time.Sleep(time.Millisecond*6)
			bar.Add(int(sizeFile-bar.Current()) - bufferSize)
			break
		}
		//time.Sleep(time.Millisecond*6)
		bar.Add(bufferSize)
	}
	bar.Finish()
	return nil
}
