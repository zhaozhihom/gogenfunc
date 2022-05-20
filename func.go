package main

import (
	"fmt"
	"strings"

	"C"
	"github.com/zhaozhihom/gen"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//export generate
func generate(outPath, url, tables, models string) int {
	g := gen.NewGenerator(gen.Config{
		OutPath: outPath,
		/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
		//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		/* FieldCoverable: true,*/
		//if you want to generate index tags from database, set FieldWithIndexTag true
		/* FieldWithIndexTag: true,*/
		//if you want to generate type tags from database, set FieldWithTypeTag true
		/* FieldWithTypeTag: true,*/
		//if you need unit tests for query code, set WithUnitTest true
		/* WithUnitTest: true, */
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		fmt.Println("db connect err: " + err.Error())
		return -1
	}

	fmt.Println("db connect success!")

	g.UseDB(db)

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	tableSlice := strings.Split(tables, ",")
	modelSlice := strings.Split(models, ",")
	if len(tableSlice) != len(modelSlice) {
		return -1
	}
	for i, table := range tableSlice {
		if model := modelSlice[i]; model == "" {
			g.ApplyBasic(g.GenerateModel(table))
		} else {
			g.ApplyBasic(g.GenerateModelAs(table, models))
		}
	}

	// g.ApplyInterface(&model.User{})

	// execute the action of code generation
	fmt.Println("generate start.")
	g.Execute()
	fmt.Println("generate end.")

	return 0
}

func main() {
	generate("dal/query", "root:@(127.0.0.1:3306)/gorm_test", "user_role,users", ",")
}
