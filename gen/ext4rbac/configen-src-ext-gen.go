package ext4rbac
import (
    p24287f458 "github.com/starter-go/rbac"
    p52f05ffa6 "github.com/starter-go/rbac/src/extension/golang/extcom/iauth"
    p6a368d170 "github.com/starter-go/rbac/src/extension/golang/extcom/iservices"
     "github.com/starter-go/application"
)

// type p52f05ffa6.PasswordAuthenticator in package:github.com/starter-go/rbac/src/extension/golang/extcom/iauth
//
// id:com-52f05ffa69ecfffd-iauth-PasswordAuthenticator
// class:class-a9adc578e4b7f0bd1de2909d6ab953e5-Registry
// alias:
// scope:singleton
//
type p52f05ffa69_iauth_PasswordAuthenticator struct {
}

func (inst* p52f05ffa69_iauth_PasswordAuthenticator) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-52f05ffa69ecfffd-iauth-PasswordAuthenticator"
	r.Classes = "class-a9adc578e4b7f0bd1de2909d6ab953e5-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p52f05ffa69_iauth_PasswordAuthenticator) new() any {
    return &p52f05ffa6.PasswordAuthenticator{}
}

func (inst* p52f05ffa69_iauth_PasswordAuthenticator) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p52f05ffa6.PasswordAuthenticator)
	nop(ie, com)

	
    com.UserDao = inst.getUserDao(ie)


    return nil
}


func (inst*p52f05ffa69_iauth_PasswordAuthenticator) getUserDao(ie application.InjectionExt)p24287f458.UserDAO{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-UserDAO").(p24287f458.UserDAO)
}



// type p6a368d170.CheckerServiceImpl in package:github.com/starter-go/rbac/src/extension/golang/extcom/iservices
//
// id:com-6a368d1706307cfd-iservices-CheckerServiceImpl
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-CheckerService
// scope:singleton
//
type p6a368d1706_iservices_CheckerServiceImpl struct {
}

func (inst* p6a368d1706_iservices_CheckerServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6a368d1706307cfd-iservices-CheckerServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-CheckerService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6a368d1706_iservices_CheckerServiceImpl) new() any {
    return &p6a368d170.CheckerServiceImpl{}
}

func (inst* p6a368d1706_iservices_CheckerServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6a368d170.CheckerServiceImpl)
	nop(ie, com)

	
    com.RegList = inst.getRegList(ie)


    return nil
}


func (inst*p6a368d1706_iservices_CheckerServiceImpl) getRegList(ie application.InjectionExt)[]p24287f458.CheckerRegistry{
    dst := make([]p24287f458.CheckerRegistry, 0)
    src := ie.ListComponents(".class-24287f4589fe5add27fb48a88d706565-CheckerRegistry")
    for _, item1 := range src {
        item2 := item1.(p24287f458.CheckerRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p6a368d170.TableServiceImpl in package:github.com/starter-go/rbac/src/extension/golang/extcom/iservices
//
// id:com-6a368d1706307cfd-iservices-TableServiceImpl
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-TableService
// scope:singleton
//
type p6a368d1706_iservices_TableServiceImpl struct {
}

func (inst* p6a368d1706_iservices_TableServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6a368d1706307cfd-iservices-TableServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-TableService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6a368d1706_iservices_TableServiceImpl) new() any {
    return &p6a368d170.TableServiceImpl{}
}

func (inst* p6a368d1706_iservices_TableServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6a368d170.TableServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p6a368d1706_iservices_TableServiceImpl) getDao(ie application.InjectionExt)p24287f458.TableDAO{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-TableDAO").(p24287f458.TableDAO)
}



// type p6a368d170.UserServiceImpl in package:github.com/starter-go/rbac/src/extension/golang/extcom/iservices
//
// id:com-6a368d1706307cfd-iservices-UserServiceImpl
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-UserService
// scope:singleton
//
type p6a368d1706_iservices_UserServiceImpl struct {
}

func (inst* p6a368d1706_iservices_UserServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6a368d1706307cfd-iservices-UserServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-UserService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6a368d1706_iservices_UserServiceImpl) new() any {
    return &p6a368d170.UserServiceImpl{}
}

func (inst* p6a368d1706_iservices_UserServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6a368d170.UserServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.CheckerSer = inst.getCheckerSer(ie)


    return nil
}


func (inst*p6a368d1706_iservices_UserServiceImpl) getDao(ie application.InjectionExt)p24287f458.UserDAO{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-UserDAO").(p24287f458.UserDAO)
}


func (inst*p6a368d1706_iservices_UserServiceImpl) getCheckerSer(ie application.InjectionExt)p24287f458.CheckerService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-CheckerService").(p24287f458.CheckerService)
}


