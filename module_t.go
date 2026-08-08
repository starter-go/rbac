package rbac

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName    = "github.com/starter-go/rbac"
	theModuleVersion = "v0.10.11"
	theModuleRev     = 28
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
	theCoreModuleResPath = "src/core/resources"
	theExtModuleResPath  = "src/extension/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

//go:embed "src/core/resources"
var theCoreModuleResFS embed.FS

//go:embed "src/extension/resources"
var theExtModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

func BuildModuleForMain() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)
	mb.Version(theModuleVersion).Revision(theModuleRev).Name(theModuleName + "#main")
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	return mb
}

func BuildModuleForCore() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)
	mb.Version(theModuleVersion).Revision(theModuleRev).Name(theModuleName + "#core")
	mb.EmbedResources(theCoreModuleResFS, theCoreModuleResPath)

	return mb
}

func BuildModuleForExtension() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)
	mb.Version(theModuleVersion).Revision(theModuleRev).Name(theModuleName + "#extension")
	mb.EmbedResources(theExtModuleResFS, theExtModuleResPath)

	return mb
}

func BuildModuleForTest() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)
	mb.Version(theModuleVersion).Revision(theModuleRev).Name(theModuleName + "#test")
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	return mb
}

////////////////////////////////////////////////////////////////////////////////
// EOF
