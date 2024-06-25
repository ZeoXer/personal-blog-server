package service

import (
	"go-server/global"
	article_model "go-server/model"

	"github.com/gin-gonic/gin"
)

type ArticleService struct{}

func makeArticle(c *gin.Context) (*article_model.Article, error) {
	username, _, err := Utils.GetUserInfo(c)
	if err != nil {
		return nil, err
	}

	var RequestBody struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		CategoryID uint   `json:"category_id"`
	}

	if err := c.ShouldBindJSON(&RequestBody); err != nil {
		return nil, err
	}

	article := &article_model.Article{
		Title:      RequestBody.Title,
		Content:    RequestBody.Content,
		Username:   username,
		CategoryID: RequestBody.CategoryID,
	}

	return article, nil
}

func (a *ArticleService) CreateArticle(c *gin.Context) error {
	article, err := makeArticle(c)
	if err != nil {
		return err
	}

	err = global.DB.Create(article).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) GetArticle(c *gin.Context) (*article_model.Article, error) {
	var RequestBody struct {
		ArticleId uint `json:"article_id"`
	}

	if err := c.ShouldBindJSON(&RequestBody); err != nil {
		return nil, err
	}

	var article article_model.Article
	err := global.DB.First(&article, RequestBody.ArticleId).Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (a *ArticleService) UpdateArticle(c *gin.Context) error {
	article, err := makeArticle(c)
	if err != nil {
		return err
	}

	err = global.DB.Save(article).Error
	if err != nil {
		return err
	}

	return nil
}

// TODO
func (a *ArticleService) DeleteArticle(articleId uint) error {
	var article article_model.Article

	err := global.DB.Delete(&article, articleId).Error
	if err != nil {
		return err
	}

	return nil
}

// TODO
func (a *ArticleService) GetArticlesByCategory(category string) ([]article_model.Article, error) {
	var articleList []article_model.Article

	err := global.DB.Where("category = ?", category).Find(&articleList).Error
	if err != nil {
		return nil, err
	}

	return articleList, nil
}

func (a *ArticleService) CreateArticleCategory(c *gin.Context) error {
	username, _, err := Utils.GetUserInfo(c)
	if err != nil {
		return err
	}

	var RequestBody struct {
		CategoryName string `json:"category_name"`
	}

	if err := c.ShouldBindJSON(&RequestBody); err != nil {
		return err
	}

	category := article_model.ArticleCategory{
		Username:     username,
		CategoryName: RequestBody.CategoryName,
	}

	err = global.DB.Create(&category).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) GetArticleCategoryList(c *gin.Context) ([]article_model.ArticleCategory, error) {
	var articleCategoryList []article_model.ArticleCategory
	username, _, err := Utils.GetUserInfo(c)
	if err != nil {
		return nil, err
	}

	err = global.DB.Select("id, username, category_name").Where("username = ?", username).Find(&articleCategoryList).Error
	if err != nil {
		return nil, err
	}

	return articleCategoryList, nil
}

func (a *ArticleService) UpdateArticleCategory(c *gin.Context) error {
	username, _, err := Utils.GetUserInfo(c)
	if err != nil {
		return err
	}

	var RequestBody struct {
		Id           uint   `json:"category_id"`
		CategoryName string `json:"category_name"`
	}
	if err := c.ShouldBindJSON(&RequestBody); err != nil {
		return err
	}

	err = global.DB.Model(&article_model.ArticleCategory{}).Where("id = ? AND username = ?", RequestBody.Id, username).Update("category_name", RequestBody.CategoryName).Error
	if err != nil {
		return err
	}

	return nil
}

// TODO
func (a *ArticleService) DeleteArticleCategory(categoryId uint) error {
	var articleCategory article_model.ArticleCategory

	err := global.DB.Delete(&articleCategory, categoryId).Error
	if err != nil {
		return err
	}

	return nil
}

var ArticleServiceGroup = new(ArticleService)
