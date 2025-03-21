package controller

import (
	"gin_demo/common"
	"gin_demo/model"
	"gin_demo/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 定义接口
type ICategoryController interface {
	RestController
}

// CategoryController 是 ICategoryController 的实现
type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{DB: db}
}

// CategoryController 是 ICategoryController 的实现
// type CategoryController struct {
// }

// ? 由于 CategoryController 实现了 Create 方法，
// ? 符合 ICategoryController 接口的定义，
// ? 所以它被视为实现了该接口（Go 是隐式接口实现）

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误,分类名称必填")
	}

	if len(requestCategory.Name) > 50 {
		response.Fail(ctx, nil, "参数不合法")
	}

	c.DB.Create(&requestCategory)

	response.Success(ctx, gin.H{"Category": requestCategory}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory model.Category

	if err := ctx.Bind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据绑定错误，"+err.Error())
		return
	}

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误,分类名称必填")
		return
	}

	// 获取path中的参数
	categoryID, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		response.Fail(ctx, nil, "无效的ID")
		return
	}

	var updatecategory model.Category
	result := c.DB.First(&updatecategory, categoryID)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// 如果没有找到记录
			response.Fail(ctx, nil, "分类不存在")
			return
		}
		// 其他错误
		response.Fail(ctx, nil, "数据库错误："+result.Error.Error())
		return
	}

	// map struct name value
	c.DB.Model(&updatecategory).Update("name", requestCategory.Name)

	response.Success(ctx, gin.H{"category": updatecategory}, "修改成功")

}

func (c CategoryController) Show(ctx *gin.Context) {

	// 获取path中的参数
	categoryID, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		response.Fail(ctx, nil, "无效的ID")
		return
	}

	var category model.Category
	result := c.DB.First(&category, categoryID)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// 如果没有找到记录
			response.Fail(ctx, nil, "分类不存在")
			return
		}
		// 其他错误
		response.Fail(ctx, nil, "数据库错误："+result.Error.Error())
		return
	}

	response.Success(ctx, gin.H{"category": category}, "")

}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryID, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		response.Fail(ctx, nil, "无效的ID")
		return
	}

	// var category model.Category
	// result := c.DB.First(&category, categoryID)

	result := c.DB.Delete(model.Category{}, categoryID)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// 如果没有找到记录
			response.Fail(ctx, nil, "删除失败,请重试")
			return
		}
		// 其他错误
		response.Fail(ctx, nil, "数据库错误："+result.Error.Error())
		return
	}

	response.Success(ctx, nil, "")

}
