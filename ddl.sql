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
where not exists (select * from users);