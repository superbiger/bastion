-- ------------------------admin------------------------------------
create table admin
(
	id int(11) unsigned auto_increment
		primary key,
	username varchar(100) null comment '用户名',
	password varchar(255) null comment '密码',
	avatar varchar(255) null comment '头像',
	nickname varchar(255) null comment '昵称',
	email varchar(255) null comment '邮箱',
	phone int null comment '手机',
	created_at timestamp null,
	updated_at timestamp null,
	deleted_at timestamp null
)
charset=utf8mb4;

create index deleted_at
	on admin (deleted_at);

create index username
	on admin (username);

-- -----------------------alarm-------------------------------------
create table alarm
(
	id int(11) unsigned auto_increment
		primary key,
	strategy varchar(11) null comment '策略',
	metrics int null comment '指标',
	message varchar(11) null comment '信息',
	adminId int null comment '提醒用户',
	created_at timestamp null,
	updated_at timestamp null,
	deleted_at timestamp null
)
charset=utf8mb4;

create index deleted_at
	on alarm (deleted_at);
-- ------------------------breadcrumb------------------------------------
create table breadcrumb
(
	id int(11) unsigned auto_increment
		primary key,
	data text null,
	error_id int null,
		created_at timestamp null,
	updated_at timestamp null,
	deleted_at timestamp null
)
charset=utf8mb4;

create index deleted_at
	on breadcrumb (deleted_at);

-- -----------------------device-------------------------------------
create table device
(
	id int(11) unsigned auto_increment
		primary key,
	uid varchar(50) null,
	browser_ua varchar(255) null,
	browser_result varchar(255) null,
	created_at timestamp null,
	updated_at timestamp null,
	deleted_at timestamp null,
	constraint uid
		unique (uid)
)
charset=utf8mb4;


create index deleted_at
	on device (deleted_at);

-- --------------------------error----------------------------------
create table error
(
	id int(11) unsigned auto_increment
		primary key,
	app_id varchar(255) null,
	tag varchar(100) null,
	path varchar(255) null,
	error_msg varchar(255) null,
	error_string varchar(255) null,
	file_url varchar(200) null,
	lineno varchar(11) null,
	colno varchar(11) null,
	uid varchar(50) null,
	biz_user_id varchar(50) null,
	created_at timestamp null,
	updated_at timestamp null,
	deleted_at timestamp null
)
charset=utf8mb4;

create index deleted_at
	on error (deleted_at);

-- ---------------------organization---------------------------------------
create table `organization`
(
	id int(11) unsigned auto_increment
		primary key,
	name int null comment '组织名字',
	remark int null comment '备注',
	admin_id int null comment '归属人',
	created_at timestamp null,
	updated_at timestamp null,
	deleted_at timestamp null
)
charset=utf8mb4;

create index deleted_at
	on organization (deleted_at);

-- ----------------------project--------------------------------------
create table project
(
	id int(11) unsigned auto_increment
		primary key,
	app_id varchar(255) null comment '应用id',
	name varchar(255) null comment '项目id',
	type varchar(255) null comment '类型',
	group_id int null comment '组织id',
	created_at timestamp null,
	updated_at timestamp null,
	deleted_at timestamp null
)
charset=utf8mb4;

create index deleted_at
	on project (deleted_at);
