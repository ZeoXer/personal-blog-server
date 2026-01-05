package service

import (
	"fmt"
	"go-server/global"
	article_model "go-server/model"
	"math"
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
		Title       string `json:"title"`
		Content     string `json:"content"`
		IsPublished bool   `json:"is_published"`
		CategoryID  uint   `json:"category_id"`
	}

	if err := c.ShouldBindJSON(&RequestBody); err != nil {
		return nil, err
	}

	article := &article_model.Article{
		Title:       RequestBody.Title,
		Content:     RequestBody.Content,
		IsPublished: RequestBody.IsPublished,
		Username:    username,
		CategoryID:  RequestBody.CategoryID,
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

	var reqMap map[string]interface{}
	if err := c.ShouldBindJSON(&reqMap); err != nil {
		return err
	}

	updates := map[string]interface{}{}
	if v, ok := reqMap["title"]; ok {
		updates["title"] = v
	}
	if v, ok := reqMap["content"]; ok {
		updates["content"] = v
	}
	if v, ok := reqMap["is_published"]; ok {
		updates["is_published"] = v
	}
	if v, ok := reqMap["category_id"]; ok {
		updates["category_id"] = v
	}

	if len(updates) == 0 {
		return nil
	}

	err = global.DB.Model(&article_model.Article{}).Where("id = ?", articleId).Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) DeleteArticle(c *gin.Context) error {
	username, _, err := Utils.GetUserInfo(c)
	articleId := c.Param("articleId")
	if err != nil || articleId == "" {
		return fmt.Errorf("information error")
	}

	var article article_model.Article

	err = global.DB.Where("id = ? AND username = ?", articleId, username).Delete(&article).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) GetAllPublicArticles(c *gin.Context) ([]article_model.Article, error) {
	authorName := c.Param("authorName")
	var articleList []article_model.Article

	err := global.DB.Where("username = ? AND is_published = ?", authorName, true).Order("updated_at DESC").Find(&articleList).Error

	if err != nil {
		return nil, err
	}

	return articleList, nil
}

func (a *ArticleService) GetArticlesByCategory(c *gin.Context) ([]article_model.Article, int64, error) {
	authorName := c.Param("authorName")
	page := c.Query("page")
	username, _, err := Utils.GetUserInfo(c)
	if err != nil && authorName == "" {
		return nil, 0, err
	}

	categoryIdParam := c.Param("categoryId")
	categoryId, err := strconv.Atoi(categoryIdParam)
	if err != nil {
		return nil, 0, err
	}

	pageNum := 1
	if page != "" {
		if p, perr := strconv.Atoi(page); perr == nil && p > 0 {
			pageNum = p
		}
	}
	pageSize := 10
	offset := (pageNum - 1) * pageSize

	var articleList []article_model.Article
	var totalCount int64

	errA := error(nil)
	errC := error(nil)
	if authorName != "" {
		errA = global.DB.Where("username = ? AND category_id = ? AND is_published = ?", authorName, categoryId, true).
			Order("updated_at DESC").
			Limit(pageSize).
			Offset(offset).
			Find(&articleList).Error
		errC = global.DB.Model(&article_model.Article{}).
			Where("username = ? AND category_id = ? AND is_published = ?", authorName, categoryId, true).
			Count(&totalCount).Error
	} else {
		errA = global.DB.Where("username = ? AND category_id = ?", username, categoryId).
			Order("updated_at DESC").
			Limit(pageSize).
			Offset(offset).
			Find(&articleList).Error
		errC = global.DB.Model(&article_model.Article{}).
			Where("username = ? AND category_id = ?", username, categoryId).
			Count(&totalCount).Error
	}

	if errA != nil || errC != nil {
		return nil, 0, fmt.Errorf("failed to get articles: %v, %v", errA, errC)
	}

	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	return articleList, totalPage, nil
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

func (a *ArticleService) GetArticleCategoryById(c *gin.Context) (article_model.ArticleCategory, error) {
	var articleCategory article_model.ArticleCategory
	authorName := c.Param("authorName")
	categoryIdParam := c.Param("categoryId")
	categoryId, err := strconv.Atoi(categoryIdParam)
	if err != nil {
		return articleCategory, err
	}

	err = global.DB.Where("username = ? AND id = ?", authorName, categoryId).First(&articleCategory).Error
	if err != nil {
		return articleCategory, err
	}

	return articleCategory, nil
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

func (a *ArticleService) DeleteArticleCategory(c *gin.Context) error {
	username, _, err := Utils.GetUserInfo(c)
	categoryId := c.Param("categoryId")
	if err != nil || categoryId == "" {
		return fmt.Errorf("information error")
	}

	var articleCategory article_model.ArticleCategory

	err = global.DB.Where("id = ? AND username = ?", categoryId, username).Delete(&articleCategory).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) GetArticleAnalysis(c *gin.Context) (article_model.ArticleAnalysis, error) {
	articleAnalysis := article_model.ArticleAnalysis{}
	authorName := c.Param("authorName")

	var articles []article_model.Article
	var articleCategories []article_model.ArticleCategory

	err := global.DB.Where("username = ?", authorName).Find(&articles).Error
	if err != nil {
		return articleAnalysis, err
	}

	err = global.DB.Where("username = ?", authorName).Find(&articleCategories).Error
	if err != nil {
		return articleAnalysis, err
	}

	articleAnalysis = article_model.ArticleAnalysis{
		ArticleAmount:         uint(len(articles)),
		ArticleCategoryAmount: uint(len(articleCategories)),
	}

	return articleAnalysis, nil
}

func (a *ArticleService) SearchArticleByKeyword(c *gin.Context) ([]article_model.Article, error) {
	var articles []article_model.Article
	authorName := c.Param("authorName")
	username, _, err := Utils.GetUserInfo(c)
	if err != nil && authorName == "" {
		return nil, err
	}
	searchName := username

	if authorName != "" {
		searchName = authorName
	}

	keyword := c.Query("keyword")

	if keyword == "" {
		return nil, fmt.Errorf("找不到輸入的關鍵字")
	}

	if authorName != "" {
		err = global.DB.Where("username = ? AND (title LIKE ? OR content LIKE ?) AND is_published = ?", searchName, "%"+keyword+"%", "%"+keyword+"%", true).Find(&articles).Error
	} else {
		err = global.DB.Where("username = ? AND (title LIKE ? OR content LIKE ?)", searchName, "%"+keyword+"%", "%"+keyword+"%").Find(&articles).Error
	}

	if err != nil {
		return nil, err
	}

	return articles, nil
}

var ArticleServiceGroup = new(ArticleService)
