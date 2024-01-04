CREATE TABLE movies
(
    id                  INT             NOT NULL AUTO_INCREMENT,
    title               VARCHAR(255)    NULL,
    description         TEXT            NULL,
    rating              DECIMAL(2,1)    NULL,
    image               TEXT            NULL,
    created_at          DATETIME        NULL,
    updated_at          DATETIME        NULL,
    PRIMARY KEY (id)
 ) ENGINE = InnoDB;