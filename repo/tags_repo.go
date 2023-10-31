package repo

import "crud-gin-gorm/model"

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(tagsId int)
	FindById(tagsId int) (tars model.Tags, err error)
	FindAll() []model.Tags
}
