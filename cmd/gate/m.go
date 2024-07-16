package main

import (
	"cmd/gate/main.go/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//g := gen.NewGenerator(gen.Config{
	//	OutPath: "./internal/ym/dal",
	//	Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	//})

	conn, _ := gorm.Open(postgres.Open("host=62.109.5.203 user=akaletr password=adelaida2011 dbname=gate port=5432"), &gorm.Config{})
	//connAuth, _ := gorm.Open(postgres.Open("host=62.109.5.203 user=akaletr password=adelaida2011 dbname=auth port=5432"), &gorm.Config{})
	//conn2, _ := gorm.Open(postgres.Open("host=45.9.27.162 user=root password=LfhTG7T7dzwmkXh dbname=sldb_dev port=5432"), &gorm.Config{})

	//connYM, _ := gorm.Open(postgres.Open("host=45.9.27.162 user=root password=LfhTG7T7dzwmkXh dbname=sldb_ym port=5432"), &gorm.Config{})

	conn.Migrator().AutoMigrate(&entity.Link{}, entity.Group{}, entity.User{})

	//
	//dal.SetDefault(connYM)
	//
	//g.UseDB(connYM) // reuse your gorm db
	//g.ApplyBasic(g.GenerateAllTable()...)
	//g.Execute()

	//u := query.User
	//fmt.Println(u)
	//
	//user, _ := query.User.First()
	//fmt.Println(user)

	//users, err := u.WithContext(context.Background()).First()
	//if err != nil {
	//	fmt.Println(err)
	//}

	//users, err := dal.Q.User.Find()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(users)
	//
	//usersNew := make([]entity.User, 0)
	//a := make([]auth.UserAuth, 0)
	//
	//id := 3
	//
	//for _, user := range users {
	//	u2 := entity.User{
	//		ID:        id,
	//		CreatedAt: user.CreatedAt,
	//		UpdatedAt: user.UpdatedAt,
	//		DeletedAt: gorm.DeletedAt{},
	//		FirstName: "",
	//		LastName:  "",
	//		Email:     user.Email,
	//		Role:      user.Role,
	//	}
	//
	//	a2 := auth.UserAuth{
	//		ID:    id,
	//		Hash:  user.Password,
	//		Salt:  user.Salt,
	//		Email: user.Email,
	//	}
	//
	//	usersNew = append(usersNew, u2)
	//	a = append(a, a2)
	//	id++
	//}
	//
	//connAuth.Model(auth.UserAuth{}).Create(a)
	//conn.Model(entity.User{}).Create(usersNew)
}
