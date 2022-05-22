# Blockchain Funding System
Blockchain based, severely over-engineered funding system. 
Built in Go and Solidity to explore capabilities of both technologies in creating decentralized apps oriented in solving a particular problem.

Quick specs:
- Funding is being transfered in 3 tiers (30%, 40%, 30%)
- 1 tranche after specified backing is recieved in a designated time
- 2 other ones after a vote, if the investors have deemed the progress in a milestone sufficient
- For each milestone there's a deadline specified by the fundraise creator
- If the progress won't be accepted by that deadline, the remaining funds are being transfered back to the investors
- For progress to be approved by the investors, it requires acceptance of the owners of 70% of the funds gathered
