package rbac

import (
	"time"

	"github.com/starter-go/base/lang"
	"gorm.io/gorm"
)

////////////////////////////////////////////////////////////////////////////////

// DTO 是基本的 DTO
type DTO struct {
	UUID lang.UUID `json:"uuid"`

	CreatedAt lang.Time `json:"created_at"`
	UpdatedAt lang.Time `json:"updated_at"`
	DeletedAt lang.Time `json:"deleted_at"`

	Group   GroupID `json:"group"` // 该对象的默认权限分组
	Owner   UserID  `json:"owner"`
	Creator UserID  `json:"creator"`
	Updater UserID  `json:"updater"`
}

// GetTarget implements DTORef.
func (inst *BaseDTO) GetTarget() *DTO {
	return inst
}

func (inst *BaseDTO) _impl() DTORef {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

// VO 是通用的基本 VO 结构
type VO struct {
	Status     int               `json:"status"`
	Message    string            `json:"message"`
	Error      string            `json:"error"`
	Time       time.Time         `json:"time"`
	Timestamp  lang.Time         `json:"timestamp"`
	Pagination *Pagination       `json:"pagination"`
	Properties map[string]string `json:"properties"`
}

// GetTarget implements VORef.
func (inst *BaseVO) GetTarget() *VO {
	return inst
}

func (inst *BaseVO) _impl() VORef {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type Entity struct {
	UUID lang.UUID `gorm:"unique"`

	CreatedAt time.Time
	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"index"` // 这个字段需要在扩展结构中定义

	Group   GroupID
	Owner   UserID
	Creator UserID
	Updater UserID
}

// GetTarget implements EntityRef.
func (inst *Entity) GetTarget() *Entity {
	return inst
}

func (inst *Entity) _impl() EntityRef {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
