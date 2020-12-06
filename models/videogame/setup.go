package videogame

import "gorm.io/gorm"

func AutoMigrate(database *gorm.DB) error {
	if err := database.AutoMigrate(&VideoGameType{}); err != nil {
		return err
	}
	if err := database.AutoMigrate(&VideoGamePlateform{}); err != nil {
		return err
	}
	return database.AutoMigrate(&VideoGame{})
}
