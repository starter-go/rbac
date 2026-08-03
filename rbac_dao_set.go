package rbac

type DaoSet struct {
	Authentications AuthenticationDAO

	Permissions PermissionDAO

	Roles RoleDAO

	Sessions SessionDAO

	Tables TableDAO

	Users UserDAO
}

type DaoSetProvider interface {
	Provide(dst *DaoSet) *DaoSet
}

type DaoSetRegistration struct {
	Label string

	Enabled bool

	Priority int

	Provider DaoSetProvider
}

type DaoSetRegistry interface {
	Registration() *DaoSetRegistration
}

////////////////////////////////////////////////////////////////////////////////

type DaoSetService interface {
	GetLoader() DaoSetLoader

	GetDaoSet() (*DaoSet, error)
}

////////////////////////////////////////////////////////////////////////////////

type DaoSetLoader interface {
	Load() (*DaoSet, error)
}

////////////////////////////////////////////////////////////////////////////////

type DaoSetHolder struct {
	cached *DaoSet
}

func (inst *DaoSetHolder) Get(ser DaoSetService) (*DaoSet, error) {

	ds := inst.cached
	if ds != nil {
		return ds, nil
	}

	// do load
	ldr := ser.GetLoader()
	ds2, err := ldr.Load()
	if err != nil {
		return nil, err
	}

	inst.cached = ds2
	ds = ds2
	return ds, nil
}

func NewDaoSetHolder(ds *DaoSet) *DaoSetHolder {
	return &DaoSetHolder{cached: ds}
}

////////////////////////////////////////////////////////////////////////////////
// EOF
