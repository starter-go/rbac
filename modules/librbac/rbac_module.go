package librbac

import (
	"github.com/starter-go/application"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/gen/main4rbac"
	"github.com/starter-go/rbac/gen/test4rbac"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
)

func Module() application.Module {
	return ModuleForLib()
}

func ModuleForTest() application.Module {

	mb := rbac.BuildModuleForTest()

	mb.Components(test4rbac.ExportComponents)

	mb.Depend(ModuleForLib())
	mb.Depend(units.Module())

	return mb.Create()
}

func ModuleForLib() application.Module {

	mb := rbac.BuildModuleForLib()

	mb.Components(main4rbac.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}
