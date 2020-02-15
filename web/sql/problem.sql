DROP DATABASE if exists `question`;
CREATE DATABASE `question`
    default character set utf8 default collate utf8_general_ci;
use `question`;

DROP TABLE if exists `user`;
CREATE TABLE user(
    `id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '用户的id',
    `login_name` VARCHAR(50) NOT NULL COMMENT '用户的登录名',
    `pwd` VARCHAR(50) NOT NULL COMMENT '用户的登陆密码',
    `name` VARCHAR(50) NOT NULL COMMENT '用户的真实姓名',
    `phone_number` VARCHAR(13) NOT NULL COMMENT '用户的手机号码',
    `job_number` VARCHAR(20) NOT NULL COMMENT '用户的工号',
    `permission` INT NOT NULL COMMENT '管理权限 0是普通用户 1是管理员 2是超级管理员',
    `gender` INT NOT NULL COMMENT '0 男, 1 女',
    `deleted` BOOLEAN NOT NULL COMMENT '表示员工已经离职'
)ENGINE = Innodb default charset utf8 comment '用户信息';

DROP TABLE if exists `event`;
CREATE TABLE event(
    `event_id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '事件的id',
    `manage_id` INT NOT NULL COMMENT '本事件的管理员id',
    `event_title` VARCHAR(50) NOT NULL COMMENT '事件的标题',
    `event_description` TEXT NOT NULL COMMENT '事件的描述',
    `event_time` JSON COMMENT '事件时间,保存为json形式,包括开始时间start_time、结束时间end_time、答题时间time',
    `event_num` JSON COMMENT '题目数量,保存为json形式,包括单选题数量single、多选题数量multiple、填空题数量fill、判断题数量judge',
    `event_type` VARCHAR(50) NOT NULL COMMENT '事件的种类',
    `problem_random` BOOLEAN NOT NULL COMMENT '是否控制题目随机排序,0--否,1--是',
    `option_random` BOOLEAN NOT NULL COMMENT '是否控制选项随机排序,0--否,1--是',
    `answer_time` VARCHAR(20) COMMENT '答题时间配置,答题时的规定完成时间',
    `participant_num` INT NOT NULL COMMENT '参加比赛的小组人数',
    `credit_rule` JSON NOT NULL COMMENT '积分规则,保存为json形式,包括单选题分数single_score、多选题分数multi_score、填空题分数fill_score、判断题分数judge_score、当日本人全对额外加分person_score、当日团队全对额外加分team_score、团队总积分上限team_score_up、个人总积分上限person_score_up'
)ENGINE = Innodb default charset utf8 comment '事件表,保存发起的比赛信息';

DROP TABLE if exists `event_problem`;
CREATE TABLE event_problem(
    `refer_event_id` INT NOT NULL COMMENT '关联的事件的id',
    `problem_id` INT NOT NULL COMMENT '使用的题目的id'
)ENGINE = Innodb default charset utf8 comment '比赛的题目';

DROP TABLE if exists `team`;
CREATE TABLE team(
    `team_id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '组id',
    `team_name` VARCHAR(50) NOT NULL COMMENT '组名',
    `refer_event_id` INT NOT NULL COMMENT '参见的事件id,关联的事件的id',
    `team_credit` FLOAT NOT NULL COMMENT '本组在事件中的积分'
)ENGINE = Innodb default charset utf8 comment '组信息表';

DROP TABLE if exists `credit_log`;
CREATE TABLE credit_log(
    `credit_log_id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '日志id',
    `refer_event_id` INT NOT NULL COMMENT '关联的事件的id',
    `refer_participant_id`  INT NOT NULL comment '参赛者id',
    `refer_team_id` INT NOT NULL COMMENT '进行操作的组id',
    `change_time` TIMESTAMP NOT NULL COMMENT '操作时间',
    `change_value` FLOAT NOT NULL COMMENT '操作值,更改的值,正为加分，负为减分',
    `change_type` INT NOT NULL COMMENT '分数操作类型：1答题加分，2当日全部答对额外加分，3当日小组全部答对额外加分',
    `change_reason` VARCHAR(100) NOT NULL COMMENT '更改原因'
)ENGINE = Innodb default charset utf8 comment '积分详细信息表';

DROP TABLE if exists `participant`;
CREATE TABLE participant(
	  `participant_id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '参赛者id',
    `refer_event_id` INT NOT NULL COMMENT '参见的事件id,关联的事件的id',
    `user_id` INT NOT NULL COMMENT '参赛人id,参加这次比赛的用户的id',
    `team_id` INT NOT NULL COMMENT '组id,所属组id',
    `credit` FLOAT NOT NULL COMMENT '该用户在比赛中的个人积分',
    `leader` BOOLEAN NOT NULL COMMENT '是否为组长,0--否,1--是',
    `waited_answer` JSON comment '保存参加比赛要答的题的答案'
)ENGINE = Innodb default charset utf8 comment '保存参加比赛的用户信息';

DROP TABLE if exists `participant_haved_answer`;
CREATE TABLE participant_haved_answer(
   `refer_participant_id`  INT NOT NULL comment '参赛者id',
   `refer_problem_id`   INT NOT NULL comment '题id',
   `refer_team_id` INT NOT NULL comment '关联的组id',
   `answer_date` Date not null comment '用户答题日期',
   `user_answer` VARCHAR(60) comment '用户答题结果',
   `true_or_false` boolean comment '用户答题是否正确'
)ENGINE = Innodb default charset utf8 comment '保存参加比赛的用户已经答的题';

DROP TABLE if exists `problem`;
CREATE TABLE problem(
    `problem_id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '题目的id',
    `problem_class` VARCHAR(20) NOT NULL COMMENT '题目的分类,如业务型,技术型等',
    `problem_type` INT NOT NULL COMMENT '题目的类型：0--填空题,1--单选题,2--多选题,3--判断题',
    `problem_content` VARCHAR(200) NOT NULL COMMENT '题目的内容',
    `problem_answer` VARCHAR(100) NOT NULL COMMENT '题目的答案',
    `problem_option` VARCHAR(500) NOT NULL COMMENT '题目的选项'
)ENGINE = Innodb default charset utf8 comment '题目表,包含所有题目';

INSERT INTO user (login_name,pwd,name,phone_number,job_number,permission,gender,deleted) VALUES ("admin","123321","admin","13808771234","1",1,0,0);
INSERT INTO user (login_name,pwd,name,phone_number,job_number,permission,gender,deleted) VALUES ("user","123321","user","13808771234","2",0,0,0);