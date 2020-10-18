package service

import (
	"strings"
	"time"

	// "github.com/kzpolicy/user/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"
	"local.packages/models"

	"context"
)

type UserService struct {
	ctx context.Context
	db  boil.ContextExecutor
}

func NewUserService() *UserService {
	ctx := context.Background()
	// DB作成
	db := boil.GetContextDB()

	return &UserService{ctx, db}
}

func (UserService) FindUser(email string) (*models.User, error) {
	// データ取得
	return models.Users(qm.Where("email=?", email)).One(context.Background(), boil.GetContextDB())
}

func (UserService) FindUserById(userID int) (*models.User, error) {
	// データ取得
	return models.Users(qm.Where("user_id=?", userID)).One(context.Background(), boil.GetContextDB())
}

func FindAllUser() (models.UserSlice, error) {
	// データ取得
	return models.Users().All(context.Background(), boil.GetContextDB())
}

func (us *UserService) AddUser(user *models.User) (*models.User, error) {
	now := time.Now()

	// User登録
	u := &models.User{}
	u.Email = user.Email
	// 暗号化
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return &models.User{}, err
	}
	emailSlice := strings.Split(user.Email, "@")
	u.AccountName = emailSlice[0]
	u.Password = string(hashed)
	// u.AccountImg = user.AccountImg
	u.CreatedBy = 1
	u.CreatedAt = now
	err2 := u.Insert(us.ctx, us.db, boil.Infer())

	return u, err2
}

func (u *UserService) UpdateUser(user models.User) error {

	var updCols map[string]interface{}

	if user.AccountImg != nil {
		updCols = map[string]interface{}{
			models.UserColumns.AccountName:  user.AccountName,
			models.UserColumns.Email:        user.Email,
			models.UserColumns.Introduction: user.Introduction,
			models.UserColumns.AccountImg:   user.AccountImg,
			models.UserColumns.ContentType:  user.ContentType.String,
			models.UserColumns.ModifiedAt:   time.Now(),
		}
	} else {
		updCols = map[string]interface{}{
			models.UserColumns.AccountName:  user.AccountName,
			models.UserColumns.Email:        user.Email,
			models.UserColumns.Introduction: user.Introduction,
			models.UserColumns.ContentType:  user.ContentType.String,
			models.UserColumns.ModifiedAt:   time.Now(),
		}
	}

	query := qm.WhereIn(models.UserColumns.Email+" = ?", user.Email)

	_, err := models.Users(query).UpdateAll(u.ctx, u.db, updCols)

	return err
}

func (u *UserService) FindRegisteredUser(email string, password string) (*models.User, error) {
	user, err1 := u.FindUser(email)
	if err1 != nil {
		return nil, err1
	}
	err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err2 != nil {
		return nil, err2
	}
	return user, nil
}
