package makepackages

import (
	"cli-boilerplate-generator/pkg/utils/file_utils"
	"fmt"
	"os"
	"path/filepath"
)

type MakePackagesService interface {
	CreatePackage() error
}

func (m MakePackages) CreatePackage() error {

	fullFile := filepath.Join(m.CurrentPath, m.PackageName)

	if !file_utils.FileExists(fullFile) {
		err := os.MkdirAll(fullFile, os.ModePerm)

		if err != nil {
			return err
		}
	}

	//moduleName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(m.PackageName), " ", "-"))

	//subFolders := folder_utils.ListFoldersInDir(m.CurrentPath)
	fmt.Println("Inside create %s", m.CurrentPath)
	return nil
}
