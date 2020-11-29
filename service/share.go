package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"local.packages/models"
)

const (
	rs4Letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rs4LetterIdxBits = 6
	rs4LetterIdxMask = 1<<rs4LetterIdxBits - 1
)

type ShareService struct {
	ctx context.Context
	db  boil.ContextExecutor
}

func NewShareService() *ShareService {
	ctx := context.Background()
	// DB作成
	db := boil.GetContextDB()

	return &ShareService{ctx, db}
}

// CreateShareToken 限定公開用URLに使用するトークンを生成する
func (ss *ShareService) CreateShareToken(userID int) string {

	// キー生成
	return genPublicKey(userID)
}

// RegisteShareToken 限定公開用URLのトークンを登録する
func (ss *ShareService) RegisteShareToken(userID int, token string) error {

	s := &models.ShareToken{}
	s.UserID = userID
	s.Token = token
	s.CreatedBy = userID
	s.CreatedAt = time.Now()
	err := s.Insert(ss.ctx, ss.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

// FindPublicUser 公開用URLトークンからユーザー🆔を取得する
func (ss *ShareService) FindPublicUser(token string) (int, error) {

	fmt.Println(token)

	// DBに登録されているキーを取得
	shareToken, err := models.ShareTokens(qm.Where("token=?", token)).One(ss.ctx, ss.db)

	if err != nil {
		return 0, err
	}
	return shareToken.UserID, nil
}

func genPublicKey(userID int) string {

	// DBに登録されているキーを取得
	existingKeys, _ := models.ShareTokens(qm.Where("user_id=?", userID)).All(context.Background(), boil.GetContextDB())

	// 重複した公開URLを生成しないように、ループしてぶつからないキーが確実に生成されるようにする
	for {
		publicKey := randString(200)
		notExist := true
		for _, k := range existingKeys {
			if k.Token == publicKey {
				notExist = false
			}
		}
		if notExist {
			return publicKey
		}
	}
}

func randString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		idx := int(rand.Int63() & rs4LetterIdxMask)
		if idx < len(rs4Letters) {
			b[i] = rs4Letters[idx]
			i++
		}
	}
	return string(b)
}
