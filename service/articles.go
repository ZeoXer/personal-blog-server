package service

import (
	"go-server/global"
	article_model "go-server/model"
)

type ArticleService struct{}

func (a *ArticleService) CreateArticle(article *article_model.Article) error {
	err := global.DB.Create(article).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) GetArticle(articleId uint) (*article_model.Article, error) {
	var article article_model.Article

	err := global.DB.First(&article, articleId).Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (a *ArticleService) GetArticlesByCategory(category string) ([]article_model.Article, error) {
	var articleList []article_model.Article

	err := global.DB.Where("category = ?", category).Find(&articleList).Error
	if err != nil {
		return nil, err
	}

	return articleList, nil
}

func (a *ArticleService) UpdateArticle(article *article_model.Article) error {
	err := global.DB.Save(article).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) DeleteArticle(articleId uint) error {
	var article article_model.Article

	err := global.DB.Delete(&article, articleId).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) CreateArticleCategory(category *article_model.ArticleCategory) error {
	err := global.DB.Create(category).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) GetArticleCategoryList(author string) ([]article_model.ArticleCategory, error) {
	var articleCategoryList []article_model.ArticleCategory

	err := global.DB.Where("author = ?", author).Find(&articleCategoryList).Error
	if err != nil {
		return nil, err
	}

	return articleCategoryList, nil
}

func (a *ArticleService) UpdateArticleCategory(article *article_model.ArticleCategory) error {
	err := global.DB.Save(article).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) DeleteArticleCategory(categoryId uint) error {
	var articleCategory article_model.ArticleCategory

	err := global.DB.Delete(&articleCategory, categoryId).Error
	if err != nil {
		return err
	}

	return nil
}

var ArticleServiceGroup = new(ArticleService)
