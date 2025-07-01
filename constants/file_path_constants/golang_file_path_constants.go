package filepathconstants

import "path/filepath"

var (
	GolangControllerFileAbsolutePath           = filepath.Join(GolangCodeFolderAbsolutePath, "controller", "controller")
	GolangDTOPaginationRequestFileAbsolutePath = filepath.Join(GolangCodeFolderAbsolutePath, "dto_pagination_request", "dto_pagination_request")
	GolangDTORequestFileAbsolutePath           = filepath.Join(GolangCodeFolderAbsolutePath, "dto_request", "dto_request")
	GolangDTOResponseFileAbsolutePath          = filepath.Join(GolangCodeFolderAbsolutePath, "dto_response", "dto_response")
	GolangModelFileAbsolutePath                = filepath.Join(GolangCodeFolderAbsolutePath, "model", "model")
	GolangNavigatorFileAbsolutePath            = filepath.Join(GolangCodeFolderAbsolutePath, "navigator", "navigator")
	GolangRepositoryFileAbsolutePath           = filepath.Join(GolangCodeFolderAbsolutePath, "repository", "repository")
	GolangRouteFileAbsolutePath                = filepath.Join(GolangCodeFolderAbsolutePath, "route", "route")
	GolangServiceFileAbsolutePath              = filepath.Join(GolangCodeFolderAbsolutePath, "service", "service")

	GolangCustomValidationGeneratorFileAbsolutePath = filepath.Join(GolangCustomValidationFolderAbsolutePath, "validation_generator")

	GolangEnumGeneratorFileAbsolutePath = filepath.Join(GolangEnumFolderAbsolutePath, "enum_generator")

	GolangAPIGetAllControllerFileAbsolutePath = filepath.Join(GolangAPIGetFolderAbsolutePath, "get_all_controller")
	GolangAPIGetAllRepoFileAbsolutePath       = filepath.Join(GolangAPIGetFolderAbsolutePath, "get_all_repo")
	GolangAPIGetAllRouteFileAbsolutePath      = filepath.Join(GolangAPIGetFolderAbsolutePath, "get_all_route")
	GolangAPIGetAllServiceFileAbsolutePath    = filepath.Join(GolangAPIGetFolderAbsolutePath, "get_all_service")
	GolangAPIGetByIDRepoFileAbsolutePath      = filepath.Join(GolangAPIGetFolderAbsolutePath, "get_by_id_repo")
	GolangAPIGetByIDServiceFileAbsolutePath   = filepath.Join(GolangAPIGetFolderAbsolutePath, "get_by_id_service")

	GolangAPIPaginateControllerFileAbsolutePath    = filepath.Join(GolangAPIPaginateFolderAbsolutePath, "paginate_controller")
	GolangAPIPaginateRepoFileAbsolutePath          = filepath.Join(GolangAPIPaginateFolderAbsolutePath, "paginate_repo")
	GolangAPIPaginateRouteFileAbsolutePath         = filepath.Join(GolangAPIPaginateFolderAbsolutePath, "paginate_route")
	GolangAPIPaginateServiceFileAbsolutePath       = filepath.Join(GolangAPIPaginateFolderAbsolutePath, "paginate_service")
	GolangAPIPaginateServiceModifyFileAbsolutePath = filepath.Join(GolangAPIPaginateFolderAbsolutePath, "paginate_service_modify")

	GolangAPISaveAPIFileAbsolutePath                = filepath.Join(GolangAPISaveFolderAbsolutePath, "save_api")
	GolangAPISaveAPIWithImageFileAbsolutePath       = filepath.Join(GolangAPISaveFolderAbsolutePath, "save_api_with_image")
	GolangAPISaveAPIWithUpdateFileAbsolutePath      = filepath.Join(GolangAPISaveFolderAbsolutePath, "save_api_with_update")
	GolangAPISaveAPIWithUpdateImageFileAbsolutePath = filepath.Join(GolangAPISaveFolderAbsolutePath, "save_api_with_update_image")
)
