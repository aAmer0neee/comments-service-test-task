package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/aAmer0neee/comments-service-test-task/internal/config"
	"github.com/aAmer0neee/comments-service-test-task/internal/domain"
)

type PostgresRepository struct {
	Db *gorm.DB
}

func ConnectPostgres(cfg config.Cfg) (*PostgresRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Repository.Host,
		cfg.Repository.User,
		cfg.Repository.Password,
		cfg.Repository.Name,
		cfg.Repository.Port,
		cfg.Repository.Sslmode)

	r, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	fmt.Printf("[Repository] [INFO] Open Data Base %s\n", r.Name())

	if cfg.Repository.Migrate {
		if err = r.AutoMigrate(&Article{}, &Comment{}); err != nil {
			return nil, err
		}
		fmt.Printf("[Repository] [INFO] migrate %s\n", r.Name())
	}
	return &PostgresRepository{Db: r}, nil
}

func (r *PostgresRepository) PostArticle(article domain.Article) error {
	return r.Db.Create(
		convertArticleToGorm(article)).Error
}

func (r *PostgresRepository) GetArticle(id int) (domain.Article, error) {
	dst := Article{}
	if err := r.Db.Find(&dst, id).Error; err != nil {
		return domain.Article{}, err
	}
	return convertGormToArticle(dst), nil
}

func (r *PostgresRepository) GetListArticles(page, limit int) ([]domain.Article, error) {
	offset := (page - 1) * limit
	dst := []Article{}
	err := r.Db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&dst).Error
	if err != nil {
		return nil, err
	}
	article := make([]domain.Article, len(dst))
	for  _, record := range dst {
		article = append(article,convertGormToArticle(record))
	}
	return article, nil
}

func (r *PostgresRepository) RecordsCount(a interface{})(int, error) {
	var count int64
	return int(count), r.Db.Model(&a).Count(&count).Error
}

func (r *PostgresRepository) PostComment(comment domain.Comment) error {
	return r.Db.Create(
		convertCommentToGorm(comment)).Error
}

func (r *PostgresRepository) GetListComments(page, limit int) ([]domain.Comment, error) {
	offset := (page - 1) * limit
	dst := []Comment{}
	err := r.Db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&dst).Error
	if err != nil {
		return nil, err
	}
	comment := make([]domain.Comment, len(dst))
	for  _, record := range dst {
		comment = append(comment,convertGormToComment(record))
	}
	return comment, nil
}