package rbac

import (
	"github.com/starter-go/application"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/gen/core4rbac"
	"github.com/starter-go/rbac/gen/ext4rbac"
	"github.com/starter-go/rbac/gen/main4rbac"
	"github.com/starter-go/rbac/gen/test4rbac"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
)

func Module() application.Module {
	return ModuleForMain()
}

////////////////////////////////////////////////////////////////////////////////

func ModuleForMain() application.Module {

	mb := rbac.BuildModuleForMain()
	mb.Components(main4rbac.ExportComponents)

	mb.Depend(ModuleForCore())
	mb.Depend(ModuleForExtension())
	mb.Depend(units.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {

	mb := rbac.BuildModuleForTest()
	mb.Components(test4rbac.ExportComponents)

	mb.Depend(ModuleForMain())
	mb.Depend(units.Module())

	return mb.Create()
}

func ModuleForExtension() application.Module {

	mb := rbac.BuildModuleForExtension()
	mb.Components(ext4rbac.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}

func ModuleForCore() application.Module {

	mb := rbac.BuildModuleForCore()
	mb.Components(core4rbac.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}
