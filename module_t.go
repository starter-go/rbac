package rbac

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName    = "github.com/starter-go/rbac"
	theModuleVersion = "v0.10.8"
	theModuleRev     = 25
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

func BuildModuleForLib() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)
	mb.Version(theModuleVersion).Revision(theModuleRev).Name(theModuleName + "#lib")
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	return mb
}

func BuildModuleForTest() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)
	mb.Version(theModuleVersion).Revision(theModuleRev).Name(theModuleName + "#test")
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	return mb
}
