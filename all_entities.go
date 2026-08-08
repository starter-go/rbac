package rbac

import "github.com/starter-go/rbac/api/dxo"

var theTableNamer dxo.RbacTableNamer

func ListEntities(prefix string) []any {

	all := []any{}

	all = append(all, new(AuthenticationEntity))
	all = append(all, new(TableEntity))
	all = append(all, new(PermissionEntity))
	all = append(all, new(SessionEntity))
	all = append(all, new(RoleEntity))
	all = append(all, new(UserEntity))

	// all = append(all, new( Phonenj  ))
	// all = append(all, new(UserEmail  ))

	theTableNamer.UpdatePrefix(prefix)
	return all
}
