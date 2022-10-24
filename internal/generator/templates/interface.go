package templates

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util/file"
)

const InterfaceFileName = "interface.go"

func GenerateInterface(module string, com config.Component, file file.Interface) error {
	path := fmt.Sprintf("%s/%s/%s", com.Path, com.Package, InterfaceFileName)
	if file.Exists(path) {
		return nil
	}

	// Create file
	j := jen.NewFile(com.Package)

	// Add header
	j.PackageComment(fileGeneratedSafeEditHeader())

	// Gomock comment
	j.Comment(mockGenComment(module, com.Path, com.Package))

	// Add interface
	j.Type().Id("Interface").Interface(
		jen.Comment("TODO: Add methods to interface"),
	)

	// Write file
	c := fmt.Sprintf("%#v", j)
	err := file.WriteFile(path, []byte(c), 0644)
	if err != nil {
		return fmt.Errorf("failed to write interface file: %w", err)
	}

	return nil
}

func mockGenComment(module, path, pkg string) string {
	return "//go:generate mockgen --build_flags=--mod=mod --package=" + pkg + " " +
		"--destination=mock/" + InterfaceFileName + " " +
		module + "/" + path + "/" + pkg + " " +
		"Interface\n"
}
