package controller

import (
	"gin_demo/common"
	"gin_demo/model"
	"gin_demo/response"
	"gin_demo/vo"
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

// ? 创建
func (c CategoryController) Create(ctx *gin.Context) {
	// var requestCategory model.Category
	// ctx.Bind(&requestCategory)

	// if requestCategory.Name == "" {
	// 	response.Fail(ctx, nil, "数据验证错误,分类名称必填")
	// }

	// if len(requestCategory.Name) > 50 {
	// 	response.Fail(ctx, nil, "参数不合法")
	// }

	// c.DB.Create(&requestCategory)

	var requestCategory vo.CreateCategoryRequest

	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "参数不合法")
		return
	}

	category := model.Category{Name: requestCategory.Name}

	c.DB.Create(&category)

	response.Success(ctx, gin.H{"Category": requestCategory}, "")
}

// ? 更新
func (c CategoryController) Update(ctx *gin.Context) {

	var requestCategory vo.CreateCategoryRequest

	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "参数不合法")
		return
	}

	// 获取path中的参数
	categoryID, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		response.Fail(ctx, nil, "无效的ID")
		return
	}

	var updateCategory model.Category
	result := c.DB.First(&updateCategory, categoryID)

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

	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")

}

// TODO 继续重构分类请求控制

// ? 查看
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

// ? 删除
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
