CREATE table fd_matchAccount (
id integer PRIMARY KEY , 
matchAccount integer,
userId VARCHAR(24),
userFirstName VARCHAR(28),
userLastName VARCHAR(28),
userBranch VARCHAR(16),
userAddress VARCHAR(255),
userCategory integer,
userActive bool,
otherFlags VARCHAR(8),
dtUpdate timestamp
)
select * from fd_matchAccount

--drop table fd_match1Account
--update fd_match3Account set interestRate = 6.95 where transId = 105704
--delete from fd_matchA4ccount

insert into fd_matchAccount (id, matchAccount, userId, userFirstName, userLastName, userBranch, userAddress, userCategory,userActive, otherFlags, dtUpdate)
values 
(1001, '8589429', 'PRANBH', 'Pravin', 'Anbhor', 'Mumbai', '204, Sunshine society, Near Mahim church, Mumbai-400016', 1, true, '', '2024-11-10 15:03:05' ),
(1002, '8589714', 'AJAYSN', 'Ajay', 'Shrnivas', 'Chennai', '71, Multipie Bunglow, Near Marina Beach, Chennai-601071', 2, true, '', '2024-12-21 10:08:08' ),
(1003, '8589001', 'VIKRLA', 'Viram', 'Lalwani', 'Mumbai', '32, Summer Plaza, Ulhasnagar, Thane-421001', 1, true, '', '2025-01-08 16:17:10' ),
(1004, '8589001', 'AMITKU', 'Amit', 'Kutty', 'Chennai', '7, High Street, Kolkatta, K-701001', 1, true, '', '2025-02-28 08:17:10' ),

