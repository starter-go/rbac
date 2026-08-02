package core

import "github.com/starter-go/rbac"

type RbacDaoSetServiceImpl struct {

	//starter:component

	_as func(rbac.DaoSetService) //starter:as("#")

	Loader rbac.DaoSetLoader //starter:inject("#")

	holder rbac.DaoSetHolder
}

// GetDaoSet implements [rbac.DaoSetService].
func (inst *RbacDaoSetServiceImpl) GetDaoSet() (*rbac.DaoSet, error) {
	h := &inst.holder
	return h.Get(inst)
}

// GetLoader implements [rbac.DaoSetService].
func (inst *RbacDaoSetServiceImpl) GetLoader() rbac.DaoSetLoader {
	return inst.Loader
}

func (inst *RbacDaoSetServiceImpl) _impl() rbac.DaoSetService {
	return inst
}
