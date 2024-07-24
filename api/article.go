package api

import "github.com/gin-gonic/gin"

type ArticleAPI struct{}

func (a *ArticleAPI) AddArticleCategory(c *gin.Context) {
	err := ArticleService.CreateArticleCategory(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "新增文章分類成功", nil, 1, c)
}

func (a *ArticleAPI) GetAllArticleCategory(c *gin.Context) {
	categories, err := ArticleService.GetArticleCategoryList(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "取得所有文章分類成功", categories, 1, c)
}

func (a *ArticleAPI) UpdateArticleCategory(c *gin.Context) {
	err := ArticleService.UpdateArticleCategory(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "更新文章分類成功", nil, 1, c)
}

func (a *ArticleAPI) DeleteArticleCategory(c *gin.Context) {
	err := ArticleService.DeleteArticleCategory(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "刪除文章分類成功", nil, 1, c)
}

func (a *ArticleAPI) AddArticle(c *gin.Context) {
	err := ArticleService.CreateArticle(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "建立文章成功", nil, 1, c)
}

func (a *ArticleAPI) GetArticle(c *gin.Context) {
	article, err := ArticleService.GetArticle(c)

	if err != nil {
		Utils.CJSON(404, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "讀取文章成功", article, 1, c)
}

func (a *ArticleAPI) UpdateArticle(c *gin.Context) {
	err := ArticleService.UpdateArticle(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "更新文章成功", nil, 1, c)
}

func (a *ArticleAPI) DeleteArticle(c *gin.Context) {
	err := ArticleService.DeleteArticle(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "刪除文章成功", nil, 1, c)
}

func (a *ArticleAPI) GetArticlesByCategory(c *gin.Context) {
	articles, err := ArticleService.GetArticlesByCategory(c)

	if err != nil {
		Utils.CJSON(404, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "取得文章列表成功", articles, 1, c)
}

func (a *ArticleAPI) GetArticleAnalysis(c *gin.Context) {
	analysis, err := ArticleService.GetArticleAnalysis(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "取得統計數據成功", analysis, 1, c)
}

func (a *ArticleAPI) SearchArticleByKeyword(c *gin.Context) {
	articles, err := ArticleService.SearchArticleByKeyword(c)

	if err != nil {
		Utils.CJSON(404, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "搜尋文章結果成功", articles, 1, c)
}

var ArticleAPIGroup = new(ArticleAPI)
