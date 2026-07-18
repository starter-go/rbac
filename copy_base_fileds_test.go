package rbac

import (
	"crypto/rand"
	"encoding/json"

	"testing"

	"github.com/starter-go/base/lang"
)

func TestCopyBaseFields(t *testing.T) {

	d1 := new(GroupDTO)
	e1 := new(GroupEntity)
	d2 := new(GroupDTO)
	now := lang.Now()
	uuid := make([]byte, 16)

	errlist := []error{}
	objlist := []any{}

	objlist = append(objlist, d1)
	objlist = append(objlist, e1)
	objlist = append(objlist, d2)

	rand.Read(uuid)

	d1.Creator = 3
	d1.Updater = 44
	d1.Owner = 555
	d1.Group = 6666
	d1.CreatedAt = now
	d1.UpdatedAt = 0
	d1.UUID = lang.NewUUID(uuid)

	err := CopyBaseFieldsD2E(d1, e1)
	if err != nil {
		errlist = append(errlist, err)
	}

	err = CopyBaseFieldsE2D(e1, d2)
	if err != nil {
		errlist = append(errlist, err)
	}

	for _, err := range errlist {
		if err != nil {
			t.Error(err)
			return
		}
	}

	for idx, ob := range objlist {
		bin, err := json.MarshalIndent(ob, "....", "\t")
		if err != nil {
			errlist = append(errlist, err)
		} else {
			str := string(bin)
			t.Logf("object[%d] = %s", idx, str)
		}
	}

}
