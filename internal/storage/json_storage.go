package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/EricRider/student-manager/internal/model"
)

type JSONStorage struct {
	FileName string
}

func NewJSONStorage(fileName string) *JSONStorage {
	return &JSONStorage{
		FileName: fileName,
	}
}

// Save 保存学生数据到 JSON 文件
func (j *JSONStorage) Save(students []model.Student) error {

	fmt.Println("Writing to:", j.FileName)

	data, err := json.MarshalIndent(students, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return os.WriteFile(j.FileName, data, 0644)
}

// Load 从 JSON 文件读取学生数据
func (j *JSONStorage) Load() ([]model.Student, error) {

	var students []model.Student

	data, err := os.ReadFile(j.FileName)

	if err != nil {

		// 文件不存在时返回空切片
		if os.IsNotExist(err) {
			return []model.Student{}, nil
		}

		return nil, err
	}

	if len(data) == 0 {
		return []model.Student{}, nil
	}

	err = json.Unmarshal(data, &students)

	if err != nil {
		return nil, err
	}

	return students, nil
}
