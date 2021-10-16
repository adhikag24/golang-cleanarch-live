package migration

import (
	"log"

	"bitbucket.org/adhikag24/golang-api/app/utils"
)

func RunMigration() error {
	db, err := utils.CreateConnection()

	if err != nil {
		log.Fatalf("Unable to connect to database %v", err)
	}

	defer db.Close()
	sqlStatement := `
	
	CREATE TABLE IF NOT EXISTS users(
		Userid VARCHAR(50),
		Name   VARCHAR(50)
	);
	
	with data(Userid, Name)  as (
		values
		   ( '01', 'Budi'),
		   ( '02', 'Nano')
	 ) 
	 insert into users (Userid, Name) 
	 select d.Userid, d.Name
	 from data d
	 where not exists (select * from users);`

	_, err = db.Exec(sqlStatement)

	if err != nil {
		log.Fatalf("Error in Migration: %v", err)
	}

	return err

}
