package model

import "github.com/childelins/go-grpc-srv/global"

type Lecturer struct {
	LecturerId int    `gorm:"primaryKey;column:lecturerId" json:"lecturerId"`
	CompanyId  int    `gorm:"column:companyId" json:"companyId"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Avatar     string `json:"avatar"`
	*Model
}

func (l *Lecturer) TableName() string {
	return "lecturers"
}

func (l *Lecturer) GetCount(companyId int, name string) int {
	var count int64
	db := global.DB.Where("companyId = ?", companyId)
	if len(name) > 0 {
		db.Where("name like ?", "%"+name+"%")
	}

	db.Model(l).Count(&count)
	return int(count)
}

func (l *Lecturer) GetList(companyId int, name string, page int, limit int) ([]*Lecturer, error) {
	var lecturers []*Lecturer
	db := global.DB.Scopes(Paginate(page, limit)).Where("companyId = ?", companyId)
	if len(name) > 0 {
		db.Where("name like ?", "%"+name+"%")
	}

	db.Order("createdAt desc").Find(&lecturers)
	return lecturers, nil
}
