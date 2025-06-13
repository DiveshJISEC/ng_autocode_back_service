CREATE table fd_trans (
transId integer PRIMARY KEY , 
userId VARCHAR(24),
matchAccount integer,
fdSchemeName VARCHAR(64),
interestRate float,
amount integer,
maturityDays integer,
autoRenew	bool,
fdAgentCode VARCHAR(12),
otherFlags VARCHAR(8),
dtUpdate timestamp
)

select * from fd_trans

--drop table fd_tra1ns
--update fd_tran1s set interestRate = 6.95 where transId = 105704
--delete from fd_t2rans

insert into fd_trans (transId, userId, matchAccount, fdSchemeName, interestRate, amount, maturityDays,autoRenew, fdAgentCode, otherFlags, dtUpdate)
values
(105701, 'PRANBH', '8589429', 'ICICI LT SECURE 5.7 %', '5.70', 50000, 180, true, 'AGB01', '', '2025-05-10 15:02:05' ),
 (105702, 'AJAYSN', '8589714', 'ICICI HOME FIN 7.1 %', '7.10', 200000, 390, false, 'AGA05', '', '2025-05-10 16:52:07' ),
 (105704, 'VIKRLA', '8589001', 'ICICI LT SECURE 6.95 %', '6.95', 350000, 440, true, 'AGC03', '', '2025-05-12 10:02:47' ),
 (105707, 'PRANBH', '8589429', 'ICICI HOME FIN 7.1 %', '7.10', 38000, 390, true, 'AGB01', '', '2025-05-19 09:52:01' ),
 (105708, 'VIKRLA', '8589001', 'ICICI LT SHORT 5.55 %', '5.55', 2200000, 60, false, 'AGB01', '', '2025-05-19 18:07:10' ),
 (105712, 'AMITKU', '8589071', 'ICICI LT SHORT 5.55 %', '5.55', 1600000, 45, false, 'AGC03', '', '2025-05-25 14:12:41' ),
 (105713, 'AMITKU', '8589071', 'ICICI HOME FIN 7.1 %','7.10', 90000, 390, true, 'AGC03', '', '2025-05-26 18:32:15' ),
 (105715, 'VIKRLA', '8589001', 'ICICI LT SECURE 6.75 %', '6.75', 180000, 390, true, 'AGB01', '', '2025-05-31 16:42:01' ),
 (105718, 'PRANBH', '8589429', 'ICICI TAX saver 6.1 %', '6.10', 700000, 1900, false, 'AGA05', '', '2025-06-01 11:12:51' ),
 (105719, 'AMITKU', '8589429', 'ICICI TAX saver 6.1 %', '6.10', 300000, 1900, false, 'AGA05', '', '2025-06-01 14:52:01' )
