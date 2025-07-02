package makeenums

import (
	"cli-boilerplate-generator/pkg/utils/file_utils"
	"fmt"
	"os"
	"path/filepath"
)

type MakeEnumsService interface {
	CreateEnums() error
}

func (m MakeEnums) CreateEnums() error {

	enumName := m.EnumName + "-" + "enums"
	fullFile := filepath.Join(m.CurrentPath, enumName)
	if !file_utils.FileExists(fullFile) {
		err := os.MkdirAll(fullFile, os.ModePerm)

		if err != nil {
			return err
		}
	}

	//enumFileSnakeCase := filepath.Join(fullFile, strings.ToLower(strings.ReplaceAll(enumName, "-", "_"))+".go")

	fmt.Println("Inside create %s", m.CurrentPath)
	return nil
}
