package main

import (
	"fmt"
	"os"
)

type Storage interface {

	Save(data string) error
	Get() (string, error)

}

type MemoryStorage struct {

	value string

}

func (storage *MemoryStorage)Save(data string) error {

	storage.value = data
	return nil
	
}

func (storage MemoryStorage)Get() (string, error) {

	return storage.value, nil

}

type FileStorage struct {

	fileName string

}

func (storage FileStorage)Save(data string) error {

	err := os.WriteFile(storage.fileName, []byte(data), 0644)
  if err != nil {
      return err
  }
	return nil

}

func (storage FileStorage)Get() (string, error){

	bytes, err := os.ReadFile(storage.fileName)
	return string(bytes), err

}

func ExecuteTask(s Storage) {

	s.Save("Hello, Go Mentor")
	fmt.Println(s.Get())

}

func main() {

	s1 := MemoryStorage{""}

	fmt.Println(s1.Get())
	s1.Save("zalupa")
	fmt.Println(s1.Get())


	s2 := FileStorage{"example.txt"}

	fmt.Println(s2.Get())
	s2.Save("zaluiiipa")
	fmt.Println(s2.Get())

	ExecuteTask(&s1)
	ExecuteTask(s2)



	
}