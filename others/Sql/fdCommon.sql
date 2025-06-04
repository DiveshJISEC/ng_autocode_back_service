
select * from fd_agentMaster

select * from fd_trans


SELECT agentCode FROM fd_agentMaster

SELECT agentCode, agentCodeName, agentFirstName, agentLastName, agentBranch, agentCategory, 
agentActive, agentEmpType, otherFlags, dtUpdate FROM fd_agentMaster

SELECT agentCode, agentCodeName, agentFirstName, agentLastName,
		       agentBranch, agentCategory, agentActive, agentEmpType,
		       otherFlags, dtUpdate
		FROM fd_agentMaster


			SELECT transId, userId, matchAccount, fdSchemeName, 
		       interestRate, amount, maturityDays, autoRenew,
		       fdagentCode, otherFlags, dtUpdate
		FROM fd_trans