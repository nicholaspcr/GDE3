// package writer is responsible for writing data related
// to the Elem struct that can be found on the models pkg

package writer

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/nicholaspcr/gde3/pkg/problems/models"
)

// Writer is the custom writer provided by the gode package, it contains the
// methods used to write the information regarding Elements and Elem
type Writer struct {
	*csv.Writer
}

// NewWriter returns a Writer pointer that contains the methods to write
// Elements and Elem into a file with a specific path
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Writer: csv.NewWriter(w),
	}
}

// CheckFilePath checks the existance of filePath
func CheckFilePath(basePath, filePath string) {
	folders := strings.Split(filePath, "/")
	for _, folder := range folders {
		basePath += "/" + folder
		if _, err := os.Stat(basePath); os.IsNotExist(err) {
			err = os.Mkdir(basePath, os.ModePerm)
			if err != nil {
				fmt.Println(basePath, folder)
				log.Fatal(err)
			}
		}
	}
}

// writeHeader writes the header of the csv writer file
func (w *Writer) WriteHeader(sz int) {
	tmpData := []string{}
	for i := 0; i < sz; i++ {
		tmpData = append(tmpData, fmt.Sprintf("elem[%d]", i))
	}
	err := w.Write(tmpData)
	if err != nil {
		log.Fatal("Couldn't write file")
	}
	w.Flush()
}

func incrementHeader(h string) string {
	pos := len(h) - 1
	for ; h[pos] == 'Z'; pos-- {
		h[pos] = byte('A')
		if pos == 0 {
			h = 'A' + h
		} else {
			h[pos-1]++
		}
	}
}

// WriteGeneration writes the objectives in the csv writer file
func (w *Writer) ElementsObjectives(elems models.Elements) error {
	if len(elems) == 0 {
		return errors.New("custom error")
	}
	data := [][]string{}
	objs := len(elems[0].Objs)
	for i := 0; i < objs; i++ {
		tmpData := []string{}
		for _, p := range elems {
			tmpData = append(tmpData, fmt.Sprintf("%5.3f", p.Objs[i]))
		}
		data = append(data, tmpData)
	}
	err := w.WriteAll(data)
	if err != nil {
		log.Fatal("Couldn't write file")
	}
	w.Flush()
	return nil
}

func (w *Writer) ElementsVectors(elems models.Elements) error {
	if len(elems) == 0 {
		return errors.New("custom error")
	}
	data := [][]string{}
	objs := len(elems[0].Objs)
	for i := 0; i < objs; i++ {
		tmpData := []string{}
		for _, p := range elems {
			tmpData = append(tmpData, fmt.Sprintf("%5.3f", p.Objs[i]))
		}
		data = append(data, tmpData)
	}
	err := w.WriteAll(data)
	if err != nil {
		log.Fatal("Couldn't write file")
	}
	w.Flush()
	return nil
}

// WriteGeneration writes the objectives in the csv writer file
func (w *Writer) ElemObjs(e models.Elem) error {
	data := []string{}
	objs := len(e.Objs)
	for i := 0; i < objs; i++ {
		data = append(data, fmt.Sprintf("%5.3f", e.Objs[i]))
	}
	err := w.Write(data)
	if err != nil {
		log.Fatal("Couldn't write file")
	}
	w.Flush()
	return nil
}

func (w *Writer) ElemVectors(e models.Elem) error {
	data := []string{}
	objs := len(e.Objs)
	for i := 0; i < objs; i++ {
		data = append(data, fmt.Sprintf("%5.3f", e.Objs[i]))
	}
	err := w.Write(data)
	if err != nil {
		log.Fatal("Couldn't write file")
	}
	w.Flush()
	return nil
}

// writeResult creates a file and writes all the elements in it
// it should be used to write a single time a specific result
// in the given path
func writeResult(path string, elems models.Elements) {
	f, _ := os.Create(path)
	writer := csv.NewWriter(f)
	writer.Comma = '\t'

	// header
	headerData := []string{"elems"}
	collumn := 'A'
	for range elems[0].Objs {
		headerData = append(headerData, string(collumn))
		collumn++
	}
	err := writer.Write(headerData)
	if err != nil {
		log.Fatal("Couldn't write file")
	}
	writer.Flush()

	bodyData := [][]string{}
	for i := range elems {
		tmpData := []string{}
		tmpData = append(tmpData, fmt.Sprintf("elem[%d]", i))
		for _, p := range elems[i].Objs {
			tmpData = append(tmpData, fmt.Sprint(p))
		}
		bodyData = append(bodyData, tmpData)
	}
	err = writer.WriteAll(bodyData)
	if err != nil {
		log.Fatal("Couldn't write file")
	}
	writer.Flush()
	f.Close()
}
