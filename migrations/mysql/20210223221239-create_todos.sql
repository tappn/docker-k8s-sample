
-- +migrate Up
CREATE TABLE todos(
  id              INT           UNSIGNED  NOT NULL  AUTO_INCREMENT,
  title           VARCHAR(50)             NOT NULL,
  description     VARCHAR(50)             NOT NULL,
  deadline        DATETIME                NOT NULL,
  created_at      DATETIME                NOT NULL,
  updated_at      DATETIME                NOT NULL,
  deleted_at      DATETIME,
  PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE todos;
