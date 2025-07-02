package filepathconstants

import (
	"os"
	"path/filepath"
	"runtime"
)

// FilePathConstants defines constants related to file paths.
var (
	_, b, _, _           = runtime.Caller(0)
	FileSeparator string = string(filepath.Separator)
	ProjectPath   string = filepath.Dir(filepath.Join(filepath.Dir(b), ".."))
	ProjectName   string = os.Getenv("PROJECT_NAME")
	PresentDir    string = filepath.Join(filepath.Dir(ProjectPath), "")
	UploadDir     string = filepath.Join(PresentDir, "?same-document", "?same", FileSeparator)

	ResourcesFolderAbsolutePath        = filepath.Join(ProjectPath, "resources")
	StorageFolderAbsolutePath          = filepath.Join(ResourcesFolderAbsolutePath, "storage")
	GolangFolderAbsolutePath           = filepath.Join(StorageFolderAbsolutePath, "golang")
	ProjectNameStorageFileAbsolutePath = filepath.Join(StorageFolderAbsolutePath, "project_name")
	//AngularFolderAbsolutePath            = filepath.Join(StorageFolderAbsolutePath, "angular")

)
