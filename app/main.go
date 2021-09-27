package main

import (
	_userService "user-crud-ca/business/users"
	_userController "user-crud-ca/controllers/users"
	_userRepo "user-crud-ca/drivers/databases/users"

	_driverFactory "user-crud-ca/drivers"
	_dbDriver "user-crud-ca/drivers/mysql"

	_routes "user-crud-ca/app/routes"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB){
	db.AutoMigrate(&_userRepo.User{})
}

func main(){
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitDB()
	dbMigrate(db)

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userService.NewUserService(userRepo)
	userCtrl := _userController.NewUserController(userUsecase)

	routesInit := _routes.ControllerList{
		UserController: *userCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}