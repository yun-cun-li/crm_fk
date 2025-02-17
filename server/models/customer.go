package models

type Customer struct {
	Id       int64  `gorm:"primaryKey"`
	Name     string `gorm:"name"`
	Source   string `gorm:"source"`
	Phone    string `gorm:"phone"`
	Email    string `gorm:"email"`
	Industry string `gorm:"industry"`
	Level    string `gorm:"level"`
	Remarks  string `gorm:"remarks"`
	Region   string `gorm:"region"`
	Address  string `gorm:"address"`
	Status   int    `gorm:"status"`
	Creator  int64  `gorm:"creator"`
	Created  int64  `gorm:"created"`
	Updated  int64  `gorm:"updated"`
}

type CustomerCreateParam struct {
	Name     string `json:"name" binding:"required"`
	Source   string `json:"source" binding:"-"`
	Phone    string `json:"phone" binding:"omitempty,len=11"`
	Email    string `json:"email" binding:"omitempty,email"`
	Industry string `json:"industry" binding:"-"`
	Level    string `json:"level" binding:"-"`
	Remarks  string `json:"remarks" binding:"-"`
	Region   string `json:"region" binding:"-"`
	Address  string `json:"address" binding:"-"`
	Status   int    `json:"status" binding:"-"`
	Creator  int64  `json:"creator,omitempty" binding:"-"`
}

type CustomerUpdateParam struct {
	Id       int64  `json:"id" binding:"required"`
	Name     string `json:"name" binding:"-"`
	Source   string `json:"source" binding:"-"`
	Phone    string `json:"phone" binding:"omitempty,len=11"`
	Email    string `json:"email" binding:"omitempty,email"`
	Industry string `json:"industry" binding:"-"`
	Level    string `json:"level" binding:"-"`
	Remarks  string `json:"remarks" binding:"-"`
	Region   string `json:"region" binding:"-"`
	Address  string `json:"address" binding:"-"`
	Status   int    `json:"status" binding:"-"`
}

type CustomerDeleteParam struct {
	Ids []int64 `json:"ids" binding:"required"`
}

type CustomerQueryParam struct {
	Id      int64  `form:"id" binding:"omitempty,gt=0"`
	Name    string `form:"name" binding:"-"`
	Phone   string `form:"phone" binding:"omitempty,len=11"`
	Creator int64  `form:"creator,omitempty" binding:"-"`
	Page    Page
}

type CustomerList struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Source   string `json:"source"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Industry string `json:"industry"`
	Level    string `json:"level"`
	Remarks  string `json:"remarks"`
	Region   string `json:"region"`
	Address  string `json:"address"`
	Status   int    `json:"status"`
	Created  int64  `json:"created"`
	Updated  int64  `json:"updated"`
}

type CustomerInfo struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Source   string `json:"source"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Industry string `json:"industry"`
	Level    string `json:"level"`
	Remarks  string `json:"remarks"`
	Region   string `json:"region"`
	Address  string `json:"address"`
	Status   int    `json:"status"`
}

type CustomerOption struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
