package rbac

const (
	theModuleName    = "github.com/starter-go/rbac"
	theModuleVersion = "v0.0.14"
	theModuleRev     = 14
)

////////////////////////////////////////////////////////////////////////////////

type ModuleInfo interface {
	Name() string
	Version() string
	Revision() int
}

func GetModuleInfo() ModuleInfo {
	return &theModuleInfo
}

////////////////////////////////////////////////////////////////////////////////

var theModuleInfo innerModuleInfo

type innerModuleInfo struct {
}

func (inst *innerModuleInfo) Name() string {
	return theModuleName
}
func (inst *innerModuleInfo) Version() string {
	return theModuleVersion
}
func (inst *innerModuleInfo) Revision() int {
	return theModuleRev
}
