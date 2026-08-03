package dxo

import (
	"time"

	"github.com/starter-go/base/lang"
	"gorm.io/gorm"
)

////////////////////////////////////////////////////////////////////////////////
// DTO - struct

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

////////////////////////////////////////////////////////////////////////////////
// Entity - struct

type Entity struct {
	UUID lang.UUID `gorm:"unique"`

	Table TableID

	CreatedAt time.Time
	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"index"`

	Group   GroupID
	Owner   UserID
	Creator UserID
	Updater UserID
}

////////////////////////////////////////////////////////////////////////////////
// VO - struct

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

////////////////////////////////////////////////////////////////////////////////
// DTO - methods

// GetTarget implements DTORef.
func (inst *DTO) GetTarget() *DTO {
	return inst
}

func (inst *DTO) _impl() DTORef {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// Entity - methods

// GetTarget implements EntityRef.
func (inst *Entity) GetTarget() *Entity {
	return inst
}

func (inst *Entity) _impl() EntityRef {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// VO - methods

// GetTarget implements VORef.
func (inst *VO) GetTarget() *VO {
	return inst
}

func (inst *VO) _impl() VORef {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type BaseDTO = DTO

type BaseVO = VO

type BaseEntity = Entity

////////////////////////////////////////////////////////////////////////////////
// EOF
