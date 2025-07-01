package makepackages

import "fmt"

type MakePackagesService interface {
	CreatePackage() error
}

func (m MakePackages) CreatePackage() error {
	fmt.Println("Inside create package")
	return nil
}
