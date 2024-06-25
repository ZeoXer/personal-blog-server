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

var ArticleAPIGroup = new(ArticleAPI)

func (a *ArticleAPI) UpdateArticleCategory(c *gin.Context) {
	err := ArticleService.UpdateArticleCategory(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "更新文章分類成功", nil, 1, c)
}
