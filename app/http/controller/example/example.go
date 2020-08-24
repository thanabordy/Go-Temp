package example

import (
	"app/app/model"
	"app/lib/database"
	"github.com/gin-gonic/gin"
	"time"
)

type ResponseExample struct {
	Status  bool
	Code    int64
	Message string
}

type RequestCreateUser struct {
	Username        string `form:"username"`
	Password        string `form:"password"`
}

type RequestUpdateUser struct {
	ID        int64 `form:"id"`
	Username        string `form:"username"`
	Password        string `form:"password"`
}


type RequestDeleteUser struct {
	ID        int64 `form:"id"`
}

func FuncTest(ctx *gin.Context) {

	res := ResponseExample{}
	res.Status = true
	res.Code = 200
	res.Message = "example"
	ctx.JSON(200, res)
	return
}

// SelelectAll
func SelectUser(ctx *gin.Context) {
	res := ResponseExample{}
	//Bind Data

	//Connect DB
	conDB, err := database.DB().Begin()
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	defer conDB.Rollback()
	
	// //Select data to db
	user := []model.User{}
	err = conDB.Model(&user).Select()


	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	//commit db and return
	conDB.Commit()
	res.Status = true
	res.Code = 200
	res.Message = "Success"
	ctx.JSON(200, user)
	return
}

// Select By ID
func SelectUserByid(ctx *gin.Context) {
	res := ResponseExample{}
	//Bind Data

	//Connect DB
	conDB, err := database.DB().Begin()
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	defer conDB.Rollback()
	
	// //Select data to db
	user := model.User{}
	err = conDB.Model(&user).Where("id = ?",  ctx.Param("id")).Select()


	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	//commit db and return
	conDB.Commit()
	res.Status = true
	res.Code = 200
	res.Message = "Success"
	ctx.JSON(200, user)
	return
}

// Create
func CreateUser(ctx *gin.Context) {
	res := ResponseExample{}
	//Bind Data
	input := RequestCreateUser{}
	err := ctx.Bind(&input)
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	//Connect DB
	conDB, err := database.DB().Begin()
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	defer conDB.Rollback()
	//Create data to db
	user := model.User{
		Username:  input.Username,
		Password:  input.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = conDB.Insert(&user)
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	//commit db and return
	conDB.Commit()
	res.Status = true
	res.Code = 200
	res.Message = "Success"
	ctx.JSON(200, res)
	return
}

// Delete with body id Soft Delete
func DeleteUser(ctx *gin.Context) {
	res := ResponseExample{}
	//Bind Data
	input := RequestDeleteUser{}
	err := ctx.Bind(&input)
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	//Connect DB
	conDB, err := database.DB().Begin()
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	defer conDB.Rollback()
	//Delete data to db

	user := model.User{
		ID:  input.ID,
	}
	err = conDB.Delete(&user)

	// user := model.User{
	// 	ID:  input.ID,
	// }
	// _, err = conDB.Model(&user).WherePK().Delete()

	

	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	//commit db and return
	conDB.Commit()
	res.Status = true
	res.Code = 200
	res.Message = "Success"
	ctx.JSON(200, user)
	return
}

// Delete with Param Soft Delete
func DeleteUserByid(ctx *gin.Context) {
	res := ResponseExample{}
	//Bind Data
	//Connect DB
	conDB, err := database.DB().Begin()
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	defer conDB.Rollback()
	//Delete data to db

	user := model.User{}
	err = conDB.Model(&user).Where("id = ?", ctx.Param("id")).First() 
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	conDB.Delete(&user)
	

	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	//commit db and return
	conDB.Commit()
	res.Status = true
	res.Code = 200
	res.Message = "Success"
	ctx.JSON(200, user)
	return
}

// Update body
func UpdateUser(ctx *gin.Context) {
	res := ResponseExample{}
	//Bind Data
	input := RequestUpdateUser{}
	err := ctx.Bind(&input)
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	//Connect DB
	conDB, err := database.DB().Begin()
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	defer conDB.Rollback()
	//Select & Update data to db

	user := model.User{ID: input.ID}
	// user := model.User{
	// 	ID: input.ID,
	// 	Username: input.Username,
	// 	Password: input.Password,
	// }
	err = conDB.Select(&user)
	user.Username = input.Username
	user.Password = input.Password
	user.UpdatedAt = time.Now()
	err = conDB.Update(&user)	


	
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	//commit db and return
	conDB.Commit()
	res.Status = true
	res.Code = 200
	res.Message = "Success"
	ctx.JSON(200, user)
	return
}

// Update Param
func UpdateUserByid(ctx *gin.Context) {
	res := ResponseExample{}
	//Bind Data
	input := RequestUpdateUser{}
	err := ctx.Bind(&input)
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	//Connect DB
	conDB, err := database.DB().Begin()
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	defer conDB.Rollback()
	//Select & Update data to db

	user := model.User{}

	err = conDB.Model(&user).Where("id = ?", ctx.Param("id")).First() 
	user.Username = input.Username
	user.Password = input.Password
	user.UpdatedAt = time.Now()
	conDB.Update(&user)	


	
	if err != nil {
		res.Status = false
		res.Code = 412
		res.Message = err.Error()
		ctx.JSON(412, res)
		return
	}
	//commit db and return
	conDB.Commit()
	res.Status = true
	res.Code = 200
	res.Message = "Success"
	ctx.JSON(200, user)
	return
}