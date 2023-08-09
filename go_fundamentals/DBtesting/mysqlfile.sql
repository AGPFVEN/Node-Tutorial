use testDB;

DROP TABLE IF EXISTS testDB1;

CREATE TABLE testDB1 (
    pnumber CHAR(9) NOT NULL,
    pzip BLOB
);

INSERT INTO testDB1(pnumber, pzip) VALUE ("000000001", LOAD_FILE("/DBtesting/testzip.zip"));

/*
mysql> show variables like 'secure_file_priv'
    -> ;
+------------------+-----------------------+
| Variable_name    | Value                 |
+------------------+-----------------------+
| secure_file_priv | /var/lib/mysql-files/ |
+------------------+-----------------------+
1 row in set (0,00 sec)