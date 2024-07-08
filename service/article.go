package service

import (
	"go-server/global"
	article_model "go-server/model"
	"strconv"

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
	authorName := c.Param("authorName")

	articleIdParam := c.Param("articleId")
	articleId, err := strconv.Atoi(articleIdParam)
	if err != nil {
		return nil, err
	}

	var article article_model.Article
	if authorName != "" {
		err = global.DB.Where("id = ? AND is_published = ?", articleId, true).First(&article).Error
	} else {
		err = global.DB.First(&article, articleId).Error
	}

	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (a *ArticleService) UpdateArticle(c *gin.Context) error {
	username, _, err := Utils.GetUserInfo(c)
	if err != nil {
		return err
	}

	articleIdParam := c.Param("articleId")
	articleId, err := strconv.Atoi(articleIdParam)
	if err != nil {
		return err
	}

	var originArticle article_model.Article
	err = global.DB.Where("id = ? AND username = ?", articleId, username).First(&originArticle).Error
	if err != nil {
		return err
	}

	article, err := makeArticle(c)
	if err != nil {
		return err
	}

	err = global.DB.Model(&article_model.Article{}).Where("id = ?", articleId).Updates(article).Error

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

func (a *ArticleService) GetArticlesByCategory(c *gin.Context) ([]article_model.Article, error) {
	authorName := c.Param("authorName")
	username, _, err := Utils.GetUserInfo(c)
	if err != nil && authorName == "" {
		return nil, err
	}

	categoryIdParam := c.Param("categoryId")
	categoryId, err := strconv.Atoi(categoryIdParam)
	if err != nil {
		return nil, err
	}

	var articleList []article_model.Article
	if authorName != "" {
		err = global.DB.Where("username = ? AND category_id = ? AND is_published = ?", authorName, categoryId, true).Find(&articleList).Error
	} else {
		err = global.DB.Where("username = ? AND category_id = ?", username, categoryId).Find(&articleList).Error
	}

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
	authorName := c.Param("authorName")
	username, _, err := Utils.GetUserInfo(c)
	if err != nil && authorName == "" {
		return nil, err
	}

	if authorName != "" {
		err = global.DB.Where("username = ?", authorName).Find(&articleCategoryList).Error
	} else {
		err = global.DB.Where("username = ?", username).Find(&articleCategoryList).Error
	}

	if err != nil {
		return nil, err
	}

	return articleCategoryList, nil
}

func (a *ArticleService) UpdateArticleCategory(c *gin.Context) error {
	categoryIdParam := c.Param("categoryId")
	categoryId, err := strconv.Atoi(categoryIdParam)
	if err != nil {
		return err
	}

	username, _, err := Utils.GetUserInfo(c)
	if err != nil {
		return err
	}

	var originCategory article_model.ArticleCategory
	err = global.DB.Where("id = ? AND username = ?", categoryId, username).First(&originCategory).Error
	if err != nil {
		return err
	}

	var RequestBody struct {
		CategoryName string `json:"category_name"`
	}
	if err := c.ShouldBindJSON(&RequestBody); err != nil {
		return err
	}

	err = global.DB.Model(&article_model.ArticleCategory{}).Where("id = ?", categoryId).Update("category_name", RequestBody.CategoryName).Error
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
