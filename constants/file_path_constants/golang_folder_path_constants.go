package filepathconstants

import "path/filepath"

var (
	GolangCodeFolderAbsolutePath         = filepath.Join(GolangFolderAbsolutePath, "code")
	GolangProjectSetupFolderAbsolutePath = filepath.Join(GolangFolderAbsolutePath, "project-setup")

	GolangProjectSetupCodesFolderAbsolutePath = filepath.Join(GolangProjectSetupFolderAbsolutePath, "codes")

	GolangAPIFolderAbsolutePath              = filepath.Join(GolangCodeFolderAbsolutePath, "api")
	GolangCustomValidationFolderAbsolutePath = filepath.Join(GolangCodeFolderAbsolutePath, "custom-validation")
	GolangEnumFolderAbsolutePath             = filepath.Join(GolangCodeFolderAbsolutePath, "enum")

	GolangAPIGetFolderAbsolutePath      = filepath.Join(GolangAPIFolderAbsolutePath, "get")
	GolangAPIPaginateFolderAbsolutePath = filepath.Join(GolangAPIFolderAbsolutePath, "paginate")
	GolangAPISaveFolderAbsolutePath     = filepath.Join(GolangAPIFolderAbsolutePath, "save")
)
