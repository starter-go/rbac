package test4rbac
import (
    p24287f458 "github.com/starter-go/rbac"
    pad1f96441 "github.com/starter-go/rbac/src/test/golang/testcom"
     "github.com/starter-go/application"
)

// type pad1f96441.TryAuth in package:github.com/starter-go/rbac/src/test/golang/testcom
//
// id:com-ad1f9644153d1e68-testcom-TryAuth
// class:class-0dc072ed44b3563882bff4e657a52e62-Unit
// alias:
// scope:singleton
//
type pad1f964415_testcom_TryAuth struct {
}

func (inst* pad1f964415_testcom_TryAuth) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ad1f9644153d1e68-testcom-TryAuth"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Unit"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pad1f964415_testcom_TryAuth) new() any {
    return &pad1f96441.TryAuth{}
}

func (inst* pad1f964415_testcom_TryAuth) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pad1f96441.TryAuth)
	nop(ie, com)

	
    com.AuthService = inst.getAuthService(ie)


    return nil
}


func (inst*pad1f964415_testcom_TryAuth) getAuthService(ie application.InjectionExt)p24287f458.AuthService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-AuthService").(p24287f458.AuthService)
}



// type pad1f96441.TryDaoForAuthent in package:github.com/starter-go/rbac/src/test/golang/testcom
//
// id:com-ad1f9644153d1e68-testcom-TryDaoForAuthent
// class:class-0dc072ed44b3563882bff4e657a52e62-Unit
// alias:
// scope:singleton
//
type pad1f964415_testcom_TryDaoForAuthent struct {
}

func (inst* pad1f964415_testcom_TryDaoForAuthent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ad1f9644153d1e68-testcom-TryDaoForAuthent"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Unit"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pad1f964415_testcom_TryDaoForAuthent) new() any {
    return &pad1f96441.TryDaoForAuthent{}
}

func (inst* pad1f964415_testcom_TryDaoForAuthent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pad1f96441.TryDaoForAuthent)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pad1f964415_testcom_TryDaoForAuthent) getDao(ie application.InjectionExt)p24287f458.AuthenticationDAO{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-AuthenticationDAO").(p24287f458.AuthenticationDAO)
}



// type pad1f96441.TryDaoForPerm in package:github.com/starter-go/rbac/src/test/golang/testcom
//
// id:com-ad1f9644153d1e68-testcom-TryDaoForPerm
// class:class-0dc072ed44b3563882bff4e657a52e62-Unit
// alias:
// scope:singleton
//
type pad1f964415_testcom_TryDaoForPerm struct {
}

func (inst* pad1f964415_testcom_TryDaoForPerm) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ad1f9644153d1e68-testcom-TryDaoForPerm"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Unit"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pad1f964415_testcom_TryDaoForPerm) new() any {
    return &pad1f96441.TryDaoForPerm{}
}

func (inst* pad1f964415_testcom_TryDaoForPerm) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pad1f96441.TryDaoForPerm)
	nop(ie, com)

	
    com.PermissionDao = inst.getPermissionDao(ie)


    return nil
}


func (inst*pad1f964415_testcom_TryDaoForPerm) getPermissionDao(ie application.InjectionExt)p24287f458.PermissionDAO{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-PermissionDAO").(p24287f458.PermissionDAO)
}



// type pad1f96441.TryDaoForRole in package:github.com/starter-go/rbac/src/test/golang/testcom
//
// id:com-ad1f9644153d1e68-testcom-TryDaoForRole
// class:class-0dc072ed44b3563882bff4e657a52e62-Unit
// alias:
// scope:singleton
//
type pad1f964415_testcom_TryDaoForRole struct {
}

func (inst* pad1f964415_testcom_TryDaoForRole) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ad1f9644153d1e68-testcom-TryDaoForRole"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Unit"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pad1f964415_testcom_TryDaoForRole) new() any {
    return &pad1f96441.TryDaoForRole{}
}

func (inst* pad1f964415_testcom_TryDaoForRole) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pad1f96441.TryDaoForRole)
	nop(ie, com)

	
    com.RoleDao = inst.getRoleDao(ie)


    return nil
}


func (inst*pad1f964415_testcom_TryDaoForRole) getRoleDao(ie application.InjectionExt)p24287f458.RoleDAO{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-RoleDAO").(p24287f458.RoleDAO)
}



// type pad1f96441.TryDaoForSession in package:github.com/starter-go/rbac/src/test/golang/testcom
//
// id:com-ad1f9644153d1e68-testcom-TryDaoForSession
// class:class-0dc072ed44b3563882bff4e657a52e62-Unit
// alias:
// scope:singleton
//
type pad1f964415_testcom_TryDaoForSession struct {
}

func (inst* pad1f964415_testcom_TryDaoForSession) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ad1f9644153d1e68-testcom-TryDaoForSession"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Unit"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pad1f964415_testcom_TryDaoForSession) new() any {
    return &pad1f96441.TryDaoForSession{}
}

func (inst* pad1f964415_testcom_TryDaoForSession) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pad1f96441.TryDaoForSession)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pad1f964415_testcom_TryDaoForSession) getDao(ie application.InjectionExt)p24287f458.SessionDAO{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-SessionDAO").(p24287f458.SessionDAO)
}



// type pad1f96441.TryDaoForUser in package:github.com/starter-go/rbac/src/test/golang/testcom
//
// id:com-ad1f9644153d1e68-testcom-TryDaoForUser
// class:class-0dc072ed44b3563882bff4e657a52e62-Unit
// alias:
// scope:singleton
//
type pad1f964415_testcom_TryDaoForUser struct {
}

func (inst* pad1f964415_testcom_TryDaoForUser) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ad1f9644153d1e68-testcom-TryDaoForUser"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Unit"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pad1f964415_testcom_TryDaoForUser) new() any {
    return &pad1f96441.TryDaoForUser{}
}

func (inst* pad1f964415_testcom_TryDaoForUser) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pad1f96441.TryDaoForUser)
	nop(ie, com)

	
    com.UserDao = inst.getUserDao(ie)


    return nil
}


func (inst*pad1f964415_testcom_TryDaoForUser) getUserDao(ie application.InjectionExt)p24287f458.UserDAO{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-UserDAO").(p24287f458.UserDAO)
}



// type pad1f96441.TryDaoSet in package:github.com/starter-go/rbac/src/test/golang/testcom
//
// id:com-ad1f9644153d1e68-testcom-TryDaoSet
// class:class-0dc072ed44b3563882bff4e657a52e62-Unit
// alias:
// scope:singleton
//
type pad1f964415_testcom_TryDaoSet struct {
}

func (inst* pad1f964415_testcom_TryDaoSet) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ad1f9644153d1e68-testcom-TryDaoSet"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Unit"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pad1f964415_testcom_TryDaoSet) new() any {
    return &pad1f96441.TryDaoSet{}
}

func (inst* pad1f964415_testcom_TryDaoSet) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pad1f96441.TryDaoSet)
	nop(ie, com)

	
    com.DaoSetService = inst.getDaoSetService(ie)
    com.AuthService = inst.getAuthService(ie)


    return nil
}


func (inst*pad1f964415_testcom_TryDaoSet) getDaoSetService(ie application.InjectionExt)p24287f458.DaoSetService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-DaoSetService").(p24287f458.DaoSetService)
}


func (inst*pad1f964415_testcom_TryDaoSet) getAuthService(ie application.InjectionExt)p24287f458.AuthService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-AuthService").(p24287f458.AuthService)
}


