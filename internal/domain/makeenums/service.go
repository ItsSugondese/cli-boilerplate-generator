package makeenums

import (
	filepathconstants "cli-boilerplate-generator/constants/file_path_constants"
	"cli-boilerplate-generator/pkg/utils/file_utils"
	"cli-boilerplate-generator/pkg/utils/file_writer"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type MakeEnumsService interface {
	CreateEnums() error
}

func (m MakeEnums) CreateEnums() error {

	enumName := strings.ToLower(m.EnumName + "-" + "enums")
	fullFile := filepath.Join(m.CurrentPath, enumName)
	if !file_utils.FileExists(fullFile) {
		err := os.MkdirAll(fullFile, os.ModePerm)

		if err != nil {
			return err
		}
	}

	enumFileSnakeCase := filepath.Join(fullFile, strings.ToLower(strings.ReplaceAll(enumName, "-", "_"))+".go")

	file_writer.ReadAndWriteFromStorageFileToEnumFile(filepathconstants.GolangEnumGeneratorFileAbsolutePath, enumFileSnakeCase, enumName)
	fmt.Println("Inside create ", m.CurrentPath)
	return nil
}
