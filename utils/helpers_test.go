package utils

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

func TestBetterCopy1(t *testing.T) {
	assert := assert.New(t)
	src := bytes.NewBuffer(make([]byte, 0, 256*1024*1024))
	_tmp := make([]byte, 1024)
	for i := 0; i < src.Cap(); i += 1024 {
		_tmp[0] = byte(i / 1024)
		src.Write(_tmp)
	}
	dst := bytes.NewBuffer(make([]byte, src.Cap()))
	fullLen := src.Len()
	n, err := BetterCopy(dst, src, 256*1024, nil)
	log.Infof("BetterCopy n: %v err: %v", n, err)
	assert.Equal(n, fullLen)
	assert.Nil(err)
}

func TestWriterClose(t *testing.T) {
	//assert := assert.New(t)
	reader, writer := io.Pipe()
	_ = reader
	//go ioutil.ReadAll(reader)
	go func() {
		time.Sleep(time.Second)
		fmt.Printf("closing\n")
		writer.Close()
	}()
	fmt.Printf("writing\n")
	n, err := writer.Write(make([]byte, 10))
	fmt.Printf("writer %d %v\n", n, err)
}

func TestReaderClose(t *testing.T) {
	//assert := assert.New(t)
	reader, writer := io.Pipe()
	_ = writer
	//go ioutil.ReadAll(reader)
	go func() {
		time.Sleep(time.Second)
		fmt.Printf("closing\n")
		reader.Close()
	}()
	fmt.Printf("reading\n")
	buffer := make([]byte, 100*1024)
	n, err := reader.Read(buffer)
	fmt.Printf("read %d %v\n", n, err)
}
