CREATE table fd_agentMaster (
agentCode integer PRIMARY KEY , 
agentCodeName VARCHAR(24),
agentFirstName VARCHAR(28),
agentLastName VARCHAR(28),
agentBranch VARCHAR(16),
agentCategory int,
agentActive bool,
agentEmpType VARCHAR(16),
otherFlags VARCHAR(8),
dtUpdate timestamp
)
select * from fd_agentMaster

--drop table fd_agentMaster
--update fd_agentMas3ter set interestRate = 6.95 where transId = 105704
--delete from fd_agenwtMaster

insert into fd_agentMaster (agentCode, agentCodeName, agentFirstName, agentLastName, agentBranch, agentCategory, agentActive, agentEmpType,otherFlags, dtUpdate)
values 
(5001, 'AGA05', 'Murli', 'Iyer', 'Chennai', 3, true, 'Permanent', '', '2021-01-10 15:03:05' ),
(5002, 'AGB01', 'Rohit', 'Patil', 'Mumbai', 2, true, 'Permanent', '', '2021-11-20 11:32:01' ),
(5003, 'AGC03', 'Joy', 'Chatarjee', 'Kolkata', 2, true, 'Permanent', '', '2022-08-12 10:31:07' ),
(5004, 'AGD07', 'Ajit', 'Singh', 'Delhi', 1, true, 'Permanent', '', '2022-05-22 09:59:15' ),
(5005, 'AGE02', 'Mansukh', 'Patel', 'Surat', 4, true, 'Probation', '', '2023-06-01 16:41:12' ),
(5006, 'AGF04', 'Mansi', 'Shetty', 'Banglore', 4, false, 'Probation', '', '2023-07-02 17:01:12' ),
(5007, 'AGG06', 'Tanmay', 'Kulkarni', 'Mumbai', 4, true, 'Probation', '', '2023-12-28 11:38:32' )

