CREATE TABLE IF NOT EXISTS todos (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    Title varchar(255) NOT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
    ) ENGINE=InnoDB;