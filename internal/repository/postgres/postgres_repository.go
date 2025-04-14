package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/aAmer0neee/comments-service-test-task/internal/config"
	"github.com/aAmer0neee/comments-service-test-task/internal/domain"
	"github.com/google/uuid"
)

type PostgresRepository struct {
	Db *gorm.DB
}

func ConnectPostgres(cfg config.Cfg) (*PostgresRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Name,
		cfg.Postgres.Port,
		cfg.Postgres.Sslmode)

	r, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	fmt.Printf("[Repository] [INFO] Open Data Base %s\n", r.Name())

	if cfg.Postgres.Migrate {
		if err = r.AutoMigrate(&Article{}, &Comment{}); err != nil {
			return nil, err
		}
		fmt.Printf("[Repository] [INFO] migrate %s\n", r.Name())
	}
	return &PostgresRepository{Db: r}, nil
}

func (r *PostgresRepository) CreateArticle(article domain.Article) (domain.Article, error) {
	gormArticle := convertArticleToGorm(article)
	if err := r.Db.Create(&gormArticle).Error; err != nil {
		return domain.Article{}, err
	}
	return convertGormToArticle(gormArticle), nil
}

func (r *PostgresRepository) GetArticle(id uuid.UUID) (domain.Article, error) {
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
	article := []domain.Article{}
	for _, record := range dst {
		article = append(article, convertGormToArticle(record))
	}
	return article, nil
}

func (r *PostgresRepository) ArticleRecordsCount(article domain.Article) (int32, error) {
	var count int64
	return int32(count),
		r.Db.Model(convertArticleToGorm(article)).
			Count(&count).
			Error
}

func (r *PostgresRepository) CreateComment(comment domain.Comment) (domain.Comment, error) {

	gormComment := convertCommentToGorm(comment)

	err := r.Db.Create(&gormComment).Error
	if err != nil {
		return domain.Comment{}, err
	}
	return convertGormToComment(gormComment), nil
}

func (r *PostgresRepository) GetComments(articleId uuid.UUID, page, limit int) ([]domain.Comment, error) {
	offset := (page - 1) * limit
	dst := []Comment{}

	query := `WITH RECURSIVE root_comments AS (
    SELECT * FROM comments
    WHERE parent_id IS NULL AND article_id = $3
    ORDER BY created_at DESC
    OFFSET $1 LIMIT $2
	),
	comment_tree AS (
		SELECT * FROM root_comments
		UNION ALL
		SELECT c.* FROM comments c
		INNER JOIN comment_tree ct ON c.parent_id = ct.id
	)
	SELECT * FROM comment_tree
	ORDER BY parent_id DESC, created_at DESC;`

	err := r.Db.Raw(query, offset, limit, articleId).Scan(&dst).Error /* 	Order("created_at DESC").Where("parent_id is null").Offset(offset).Limit(limit).Find(&dst).Error */
	if err != nil {
		return nil, err
	}
	comment := []domain.Comment{}
	for _, record := range dst {
		if record.ID != uuid.Nil {
			comment = append(comment, convertGormToComment(record))
		}

	}
	return comment, nil
}

func (r *PostgresRepository) CommentsRecordCount(comment domain.Comment) (int32, error) {
	var count int64
	return int32(count),
		r.Db.Model(convertCommentToGorm(comment)).
			Where("parent_id is null").
			Count(&count).
			Error
}
