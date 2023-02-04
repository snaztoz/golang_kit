package respository

import (
	"fmt"
	"template/database"
	"template/internal/entities"
	"template/internal/params"
	"template/utils/pagination/gorm_pagination"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepo struct {
	db   *gorm.DB
	name string
}

type UserRepository interface {
	Create(entities.User) (entities.User, error)
	List([]entities.User, params.UserListParams) ([]entities.User, int, error)
}

func NewUserRepository() UserRepository {
	return &userRepo{
		db:   database.ORM(),
		name: "USER REPOSITORY",
	}
}

func (u *userRepo) Create(user entities.User) (entities.User, error) {
	log.Info(fmt.Sprintf("[%s][Create] is executed", u.name))

	if err := u.db.Create(&user).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		return user, err
	}

	return user, nil
}

func (u *userRepo) List(users []entities.User, param params.UserListParams) ([]entities.User, int, error) {
	log.Info(fmt.Sprintf("[%s][List] is executed", u.name))

	var count int64
	u.db.Find(&users).Count(&count)

	db := u.db
	if param.Q != "" {
		db = db.Where("name LIKE ?", param.Q+"%")
		db.Find(&users).Count(&count)
	}

	if err := db.Debug().Scopes(gorm_pagination.Paginate(param.Page, param.Limit)).Find(&users).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][List] %s", u.name, err.Error()))
		return users, int(count), err
	}

	return users, int(count), nil
}
