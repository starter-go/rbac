package dxo

var theRbacTableNamePrefix string

type RbacTableNamer struct {
}

func (inst *RbacTableNamer) GetFullTableName(simpleName string) string {
	prefix := theRbacTableNamePrefix
	if prefix == "" {
		prefix = "rbac_dg_"
	}
	return prefix + simpleName
}

func (inst *RbacTableNamer) UpdatePrefix(prefix string) {
	older := theRbacTableNamePrefix
	if older == "" {
		theRbacTableNamePrefix = prefix
	}
}
