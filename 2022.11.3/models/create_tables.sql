-- CREATE DATABASE sql_demos;
-- USE sql_demos;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
-- id BIGINT COMMENT '主键',
-- user_id BIGINT,
`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT'主键',
`name` VARCHAR(20) DEFAULT '' COMMENT '名字', -- 什么情况设置default，什么情况非空
`age` INT(11) DEFAULT 0 COMMENT '年龄', -- 为什么 int 喜欢用 11 的宽度？
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

# INSERT INTO user(name, age) VALUES ('qwe', 1);
SELECT name, age FROM user WHERE id IN (3, 2, 4, 5, 6) ORDER BY FIND_IN_SET(id, '3,2,4,5,6');
