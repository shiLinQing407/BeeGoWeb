/*2019年04月01日23:30:13*/
CREATE TABLE `go_news` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `class_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分类id',
  `tags` varchar(255) NOT NULL COMMENT '标签',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `content` text NOT NULL,
  `create_time` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='资讯表'

CREATE TABLE `go_news_class` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `class_name_zh` varchar(10) NOT NULL COMMENT '分类名称 中文',
  `class_name_en` varchar(30) DEFAULT NULL COMMENT '分类名称 英文',
  `create_time` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='资讯分类表'

insert into `go_web`.`go_menu` ( `parent_id`, `title_cn`, `title_en`, `icon`, `url_for`, `type`, `logic_delete`, `update_time`, `create_time`, `sort`) values ( '0', '资讯管理', null, '', null, '1', '0', null, null, '0');
insert into `go_web`.`go_menu` ( `parent_id`, `title_cn`, `title_en`, `icon`, `url_for`, `type`, `logic_delete`, `update_time`, `create_time`, `sort`) values ( '400', '资讯列表', null, '', 'NewsController.Index', '1', '0', null, null, '0');
insert into `go_web`.`go_menu` ( `parent_id`, `title_cn`, `title_en`, `icon`, `url_for`, `type`, `logic_delete`, `update_time`, `create_time`, `sort`) values ( '400', '资讯分类', null, '', 'NewsClassController.Index', '1', '0', null, null, '0');

insert into `go_web`.`go_role_menu_rel` ( `role_id`, `menu_id`, `logic_delete`, `create_time`) values ( '1', '400', '0', null);
insert into `go_web`.`go_role_menu_rel` ( `role_id`, `menu_id`, `logic_delete`, `create_time`) values ( '1', '401', '0', null);
insert into `go_web`.`go_role_menu_rel` ( `role_id`, `menu_id`, `logic_delete`, `create_time`) values ( '0', '402', '0', null);