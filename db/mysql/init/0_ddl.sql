-- SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
-- SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
-- SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

CREATE SCHEMA IF NOT EXISTS `time_management_app` DEFAULT CHARACTER SET utf8;
USE `time_management_app`;

 --
 -- table
 --
-- -----------------------------------------------------
-- Table `dojo_api`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `time_management_app`.`users` (
`studentNumber` VARCHAR(64) NOT NULL COMMENT '学籍番号',
`auth_token` VARCHAR(128) NOT NULL COMMENT '認証トークン',
`name` VARCHAR(64) NOT NULL COMMENT 'ユーザ名',
PRIMARY KEY (`studentNumber`),
INDEX `idx_auth_token` (`auth_token` ASC))
ENGINE = InnoDB
COMMENT = 'ユーザ';

CREATE TABLE IF NOT EXISTS `time_management_app`.`timeManagement` (
   `time_id` VARCHAR(128) NOT NULL COMMENT '時間ID',
  `studentNumber` VARCHAR(128) NOT NULL COMMENT '学籍番号',
  `Status` VARCHAR(128) NOT NULL COMMENT'active rest leave',
  `year` INT NOT NULL COMMENT '年',
  `month` INT NOT NULL COMMENT '月',
  `day` INT NOT NULL COMMENT '日',
  `hour` INT NOT NULL COMMENT '時間',
  `minute` INT NOT NULL COMMENT '分',
  PRIMARY KEY (`time_id`)
)
ENGINE = InnoDB
COMMENT = '時間管理';

CREATE TABLE IF NOT EXISTS `time_management_app`.`Time zone data`(
    `time_id` VARCHAR(128) NOT NULL COMMENT '時間ID',
    `working_hours` INT COMMENT '働いた時間(時)',
    `working_minutes` INT COMMENT '働いた時間(分)'
)
ENGINE = InnoDB

--  CREATE TABLE IF NOT EXISTS `time_management_app`.`timeDetails` (
--      `time_id` VARCHAR(128) NOT NULL COMMENT '時間ID',
--      `yearDate` VARCHAR(64) NOT NULL COMMENT '年日付',
--      `hourMinute` VARCHAR(64) NOT NULL COMMENT '時分',
--      PRIMARY KEY (`time_id`)
--  )