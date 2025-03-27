package controller

import (
	"errors"
	"gin_demo/common"
	"gin_demo/model"
	"gin_demo/response"
	"gin_demo/vo"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IPostController interface {
	RestController
	PageList(c *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

// ? 创建
func (p PostController) Create(c *gin.Context) {
	var requestPost vo.CreatePostRequest

	// 数据验证
	if err := c.ShouldBind(&requestPost); err != nil {
		response.Fail(c, nil, "参数不合法")
		return
	}

	// 获取登录用户 user
	user, _ := c.Get("user")

	// 创建post

	post := model.Post{
		UserID:     user.(model.User).ID,
		CategoryID: requestPost.CategoryID,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	// 插入数据

	err := p.DB.Create(&post).Error

	if err != nil {
		panic(err)
		return
	}

	response.Success(c, nil, "创建成功")

}

// ? 更新
func (p PostController) Update(c *gin.Context) {
	var requestPost vo.CreatePostRequest

	// 数据验证
	if err := c.ShouldBind(&requestPost); err != nil {
		response.Fail(c, nil, "参数不合法")
		return
	}

	// 获取path 中的id参数

	postID := c.Params.ByName("id")

	var post model.Post

	if err := p.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail(c, nil, "文章不存在")
			return
		}
	}

	// 获取登录用户 user
	user, _ := c.Get("user")

	userID := user.(model.User).ID

	if userID != post.UserID {
		response.Fail(c, nil, "文章不属于您,请勿非法操作!")
		return
	}

	err := p.DB.Model(&post).Updates(requestPost).Error

	if err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}

	response.Success(c, gin.H{"post": post}, "更新成功")

}

// TODO 继续编写文章的controller

// ? 查看
func (p PostController) Show(c *gin.Context) {
	// var requestPost vo.CreatePostRequest

	// // 数据验证
	// if err := c.ShouldBind(&requestPost); err != nil {
	// 	response.Fail(c, nil, "参数不合法")
	// 	return
	// }

	// 获取path 中的id参数

	postID := c.Params.ByName("id")

	var post model.Post

	if err := p.DB.Preload("Category").Where("id = ?", postID).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail(c, nil, "文章不存在")
			return
		}
	}

	response.Success(c, gin.H{"post": post}, "成功")
}

// ? 删除
func (p PostController) Delete(c *gin.Context) {
	// var requestPost vo.CreatePostRequest

	// // 数据验证
	// if err := c.ShouldBind(&requestPost); err != nil {
	// 	response.Fail(c, nil, "参数不合法")
	// 	return
	// }

	// 获取path 中的id参数

	postID := c.Params.ByName("id")

	var post model.Post

	if err := p.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail(c, nil, "文章不存在")
			return
		}
	}

	// 获取登录用户 user
	user, _ := c.Get("user")

	userID := user.(model.User).ID

	if userID != post.UserID {
		response.Fail(c, nil, "文章不属于您,请勿非法操作!")
		return
	}

	p.DB.Delete(&post)

	response.Success(c, nil, "删除成功")

}

func (p PostController) PageList(c *gin.Context) {
	// 获取分页参数

	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	// 分页

	var posts []model.Post

	p.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 总记录查询

	var total int64

	p.DB.Model(model.Post{}).Count(&total)

	response.Success(c, gin.H{"data": posts, "total": total}, "")
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return PostController{DB: db}
}
