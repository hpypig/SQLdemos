USE sql_demos;
create table `user` (
-- id BIGINT COMMENT '主键',
-- user_id BIGINT,
`id` BIGINT(20) COMMENT '主键',
`name` VARCHAR(20) DEFAULT '' COMMENT '名字', --什么情况设置default，什么情况非空
`age` INT(11) COMMENT '年龄', -- 为什么 int 喜欢用 11 的宽度？
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
