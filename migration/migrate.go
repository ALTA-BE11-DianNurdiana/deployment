package migration

import (
	authModel "be11/apiclean/features/auth/data"
	userModel "be11/apiclean/features/user/data"
	walletModel "be11/apiclean/features/wallet/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&authModel.Auth{})
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&walletModel.Wallet{})
}
