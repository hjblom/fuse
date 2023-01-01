package generator

import (
	"github.com/hjblom/fuse/internal/generator/modules"
	"github.com/hjblom/fuse/internal/generator/packages"
)

var ModuleGenerators []ModuleGenerator = []ModuleGenerator{
	modules.FuseGenerator,
	modules.ConfigGenerator,
	modules.ModInitGenerator,
	modules.ModTidyGenerator,
}

var PackageGenerators []PackageGenerator = []PackageGenerator{
	packages.ConfigGenerator,
	packages.PackageGenerator,
	packages.InterfaceGenerator,
	packages.ServiceGenerator,
	packages.SetupGenerator,
}
