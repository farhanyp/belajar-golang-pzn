package main

import (
	"fmt"
	_ "package-init/config" //ini adalah blank indetifier karna diawal ada char _
	"package-init/database"
)


func main() {

	fmt.Println(database.GetDatabase())

}