CREATE table fd_category (
id integer PRIMARY KEY , 
categoryId integer,
categoryDesc VARCHAR(32),
otherFlags VARCHAR(8),
dtUpdate timestamp
)
select * from fd_category

--drop table fd_cate3gory
--update fd_cat3egory set interestRate = 6.95 where transId = 105704
--delete from fd_ca5tegory

insert into fd_category (id, categoryId, categoryDesc, otherFlags, dtUpdate)
values 
(101, 1, 'Privilege', '','2022-01-03 10:00:01' ),
(102, 2, 'High', '','2022-01-03 10:00:01' ),
(103, 3, 'Medium', '','2022-01-03 10:00:01' ),
(104, 4, 'Low', '','2022-01-03 10:00:01' ),


