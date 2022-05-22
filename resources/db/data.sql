PRAGMA foreign_keys=ON;
BEGIN TRANSACTION;
CREATE TABLE [users] (
id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
login TEXT NOT NULL,
password TEXT NOT NULL,
first_name TEXT NOT NULL,
last_name TEXT NOT NULL,
email TEXT NOT NULL,
wallet_address TEXT NOT NULL
);
CREATE TABLE [projects] (
id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
owner_id INTEGER NOT NULL,
title TEXT NOT NULL,
description TEXT NOT NULL,
goal REAL NOT NULL,
milestone_1_date DATETIME NOT NULL,
milestone_2_date DATETIME NOT NULL,
milestone_3_date DATETIME NOT NULL,
smart_contract_address TEXT NOT NULL,
FOREIGN KEY(owner_id) REFERENCES users(id)
);
CREATE TABLE [backing] (
backer_id INTEGER NOT NULL,
project_id INTEGER NOT NULL,
amount REAL NOT NULL,
PRIMARY KEY (backer_id, project_id, amount),
FOREIGN KEY(backer_id) REFERENCES users(id),
FOREIGN KEY(project_id) REFERENCES projects(id)
);
COMMIT;
