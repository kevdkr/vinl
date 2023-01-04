CREATE USER vinl WITH PASSWORD 'test';
CREATE DATABASE vinl OWNER vinl;
CREATE TABLE IF NOT EXISTS transactions (
       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
       date VARCHAR(10) NOT NULL,
       payee VARCHAR(100) NOT NULL,
       comment VARCHAR(200),
       payee_comment VARCHAR(200),
       is_comment BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS accounts (
       -- id SERIAL PRIMARY KEY,
       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
       transactionid UUID NOT NULL,
       name VARCHAR(100) NOT NULL,
       amount VARCHAR(100),
       comment VARCHAR(200),
       is_comment BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS transactions_accounts (
       transaction_id UUID,
       account_id UUID,
       PRIMARY KEY (transaction_id, account_id),
       CONSTRAINT fk_transaction FOREIGN KEY(transaction_id) REFERENCES transactions(id),
       CONSTRAINT fk_account FOREIGN KEY(account_id) REFERENCES accounts(id)
);
