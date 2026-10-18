
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `casbin_rules` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `ptype` varchar(24) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v0` varchar(24) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v1` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v2` varchar(24) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v3` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v4` varchar(12) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v5` varchar(12) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `casbinrule_ptype_v0_v1_v2_v3_v4_v5` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=497 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `casbin_rules` DISABLE KEYS */;
INSERT IGNORE INTO `casbin_rules` VALUES (441,'p','001','/sys-api/api','POST','','',''),(445,'p','001','/sys-api/api/create','POST','','',''),(443,'p','001','/sys-api/api/delete','POST','','',''),(442,'p','001','/sys-api/api/list','POST','','',''),(444,'p','001','/sys-api/api/update','POST','','',''),(449,'p','001','/sys-api/authority/api/create_or_update','POST','','',''),(448,'p','001','/sys-api/authority/api/role','POST','','',''),(447,'p','001','/sys-api/authority/menu/create_or_update','POST','','',''),(446,'p','001','/sys-api/authority/menu/role','POST','','',''),(450,'p','001','/sys-api/captcha','GET','','',''),(393,'p','001','/sys-api/configuration','POST','','',''),(397,'p','001','/sys-api/configuration/create','POST','','',''),(395,'p','001','/sys-api/configuration/delete','POST','','',''),(394,'p','001','/sys-api/configuration/list','POST','','',''),(396,'p','001','/sys-api/configuration/update','POST','','',''),(413,'p','001','/sys-api/department','POST','','',''),(417,'p','001','/sys-api/department/create','POST','','',''),(415,'p','001','/sys-api/department/delete','POST','','',''),(414,'p','001','/sys-api/department/list','POST','','',''),(416,'p','001','/sys-api/department/update','POST','','',''),(430,'p','001','/sys-api/dict/:name','GET','','',''),(440,'p','001','/sys-api/dictionary','POST','','',''),(439,'p','001','/sys-api/dictionary/create','POST','','',''),(437,'p','001','/sys-api/dictionary/delete','POST','','',''),(431,'p','001','/sys-api/dictionary/list','POST','','',''),(438,'p','001','/sys-api/dictionary/update','POST','','',''),(435,'p','001','/sys-api/dictionary_detail','POST','','',''),(434,'p','001','/sys-api/dictionary_detail/create','POST','','',''),(436,'p','001','/sys-api/dictionary_detail/delete','POST','','',''),(432,'p','001','/sys-api/dictionary_detail/list','POST','','',''),(433,'p','001','/sys-api/dictionary_detail/update','POST','','',''),(451,'p','001','/sys-api/menu','POST','','',''),(461,'p','001','/sys-api/menu/create','POST','','',''),(459,'p','001','/sys-api/menu/delete','POST','','',''),(458,'p','001','/sys-api/menu/list','GET','','',''),(457,'p','001','/sys-api/menu/role/list','GET','','',''),(460,'p','001','/sys-api/menu/update','POST','','',''),(425,'p','001','/sys-api/oauth/login','POST','','',''),(424,'p','001','/sys-api/oauth_provider','POST','','',''),(429,'p','001','/sys-api/oauth_provider/create','POST','','',''),(427,'p','001','/sys-api/oauth_provider/delete','POST','','',''),(426,'p','001','/sys-api/oauth_provider/list','POST','','',''),(428,'p','001','/sys-api/oauth_provider/update','POST','','',''),(408,'p','001','/sys-api/position','POST','','',''),(412,'p','001','/sys-api/position/create','POST','','',''),(410,'p','001','/sys-api/position/delete','POST','','',''),(409,'p','001','/sys-api/position/list','POST','','',''),(411,'p','001','/sys-api/position/update','POST','','',''),(462,'p','001','/sys-api/role','POST','','',''),(466,'p','001','/sys-api/role/create','POST','','',''),(464,'p','001','/sys-api/role/delete','POST','','',''),(463,'p','001','/sys-api/role/list','POST','','',''),(465,'p','001','/sys-api/role/update','POST','','',''),(418,'p','001','/sys-api/token','POST','','',''),(423,'p','001','/sys-api/token/create','POST','','',''),(421,'p','001','/sys-api/token/delete','POST','','',''),(420,'p','001','/sys-api/token/list','POST','','',''),(419,'p','001','/sys-api/token/logout','POST','','',''),(422,'p','001','/sys-api/token/update','POST','','',''),(469,'p','001','/sys-api/user','POST','','',''),(467,'p','001','/sys-api/user/access_token','GET','','',''),(477,'p','001','/sys-api/user/change_password','POST','','',''),(479,'p','001','/sys-api/user/create','POST','','',''),(474,'p','001','/sys-api/user/delete','POST','','',''),(476,'p','001','/sys-api/user/info','GET','','',''),(475,'p','001','/sys-api/user/list','POST','','',''),(481,'p','001','/sys-api/user/login','POST','','',''),(470,'p','001','/sys-api/user/logout','GET','','',''),(473,'p','001','/sys-api/user/perm','GET','','',''),(472,'p','001','/sys-api/user/profile','GET','','',''),(471,'p','001','/sys-api/user/profile','POST','','',''),(468,'p','001','/sys-api/user/refresh_token','GET','','',''),(480,'p','001','/sys-api/user/register','POST','','',''),(478,'p','001','/sys-api/user/update','POST','','',''),(482,'p','001','/sys-api/core/ping','GET','','',''),(483,'p','001','/sys-api/core/init/database','GET','','',''),(484,'p','001','/sys-api/core/init/job_database','GET','','',''),(485,'p','001','/sys-api/core/init/mcms_database','GET','','',''),(486,'p','001','/sys-api/captcha/email','GET','','',''),(487,'p','001','/sys-api/captcha/sms','GET','','',''),(488,'p','001','/sys-api/user/login_by_email','POST','','',''),(489,'p','001','/sys-api/user/login_by_sms','POST','','',''),(490,'p','001','/sys-api/user/register_by_email','POST','','',''),(491,'p','001','/sys-api/user/register_by_sms','POST','','',''),(492,'p','001','/sys-api/user/reset_password_by_email','POST','','',''),(493,'p','001','/sys-api/user/reset_password_by_sms','POST','','',''),(494,'p','001','/sys-api/oauth/login/callback','GET','','',''),(495,'p','001','/sys-api/configuration/system/list','POST','','',''),(496,'p','001','/sys-api/dict/public/:name','GET','','','');
/*!40000 ALTER TABLE `casbin_rules` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `cloud_file_tag_cloud_files` (
  `cloud_file_tag_id` bigint unsigned NOT NULL,
  `cloud_file_id` char(36) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`cloud_file_tag_id`,`cloud_file_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `cloud_file_tag_cloud_files` DISABLE KEYS */;
/*!40000 ALTER TABLE `cloud_file_tag_cloud_files` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `file_tag_files` (
  `file_tag_id` bigint unsigned NOT NULL,
  `file_id` char(36) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`file_tag_id`,`file_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `file_tag_files` DISABLE KEYS */;
/*!40000 ALTER TABLE `file_tag_files` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `fms_cloud_file_tags` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'CloudFileTag''s name | 标签名称',
  `remark` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The remark of tag | 标签的备注',
  PRIMARY KEY (`id`),
  KEY `cloudfiletag_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Cloud File''s Tags Table | 云文件标签表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `fms_cloud_file_tags` DISABLE KEYS */;
/*!40000 ALTER TABLE `fms_cloud_file_tags` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `fms_cloud_files` (
  `id` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `state` tinyint(1) DEFAULT '1' COMMENT 'State true: normal false: ban | 状态 true 正常 false 禁用',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The file''s name | 文件名',
  `url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The file''s url | 文件地址',
  `size` bigint unsigned NOT NULL COMMENT 'The file''s size | 文件大小',
  `file_type` tinyint unsigned NOT NULL COMMENT 'The file''s type | 文件类型',
  `user_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The user who upload the file | 上传用户的 ID',
  `cloud_file_storage_providers` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `cloudfile_name` (`name`),
  KEY `cloudfile_file_type` (`file_type`),
  KEY `cloudfile_url` (`url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Cloud File Table | 云文件表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `fms_cloud_files` DISABLE KEYS */;
/*!40000 ALTER TABLE `fms_cloud_files` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `fms_file_tags` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'FileTag''s name | 标签名称',
  `remark` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The remark of tag | 标签的备注',
  PRIMARY KEY (`id`),
  KEY `filetag_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='File''s Tags Table | 文件标签表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `fms_file_tags` DISABLE KEYS */;
/*!40000 ALTER TABLE `fms_file_tags` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `fms_files` (
  `id` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'File''s name | 文件名称',
  `file_type` tinyint unsigned NOT NULL COMMENT 'File''s type | 文件类型',
  `size` bigint unsigned NOT NULL COMMENT 'File''s size | 文件大小',
  `path` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'File''s path | 文件路径',
  `user_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'User''s UUID | 用户的 UUID',
  `md5` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The md5 of the file | 文件的 md5',
  PRIMARY KEY (`id`),
  KEY `file_user_id` (`user_id`),
  KEY `file_file_type` (`file_type`),
  KEY `file_path` (`path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='File Table | 文件表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `fms_files` DISABLE KEYS */;
/*!40000 ALTER TABLE `fms_files` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `fms_storage_providers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `state` tinyint(1) DEFAULT '1' COMMENT 'State true: normal false: ban | 状态 true 正常 false 禁用',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The cloud storage service name | 服务名称',
  `bucket` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The cloud storage bucket name | 云存储服务的存储桶',
  `secret_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The secret ID | 密钥 ID',
  `secret_key` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The secret key | 密钥 Key',
  `endpoint` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The service URL | 服务器地址',
  `folder` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The folder in cloud | 云服务目标文件夹',
  `region` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The service region | 服务器所在地区',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is it the default provider | 是否为默认提供商',
  `use_cdn` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Does it use CDN | 是否使用 CDN',
  `cdn_url` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'CDN URL | CDN 地址',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Cloud File''s Provider Configuration Table | 云文件服务商配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `fms_storage_providers` DISABLE KEYS */;
/*!40000 ALTER TABLE `fms_storage_providers` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `mcms_email_logs` (
  `id` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `target` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The target email address | 目标邮箱地址',
  `subject` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The subject | 发送的标题',
  `content` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The content | 发送的内容',
  `send_status` tinyint unsigned NOT NULL COMMENT 'The send status, 0 unknown 1 success 2 failed | 发送的状态, 0 未知， 1 成功， 2 失败',
  `provider` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The sms service provider | 短信服务提供商',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Email Log Table | 邮件日志表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mcms_email_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `mcms_email_logs` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `mcms_email_providers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The email provider name | 电子邮件服务的提供商',
  `auth_type` tinyint unsigned NOT NULL COMMENT 'The auth type, supported plain, CRAMMD5 | 鉴权类型, 支持 plain, CRAMMD5',
  `email_addr` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The email address | 邮箱地址',
  `password` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The email''s password | 电子邮件的密码',
  `host_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The host name is the email service''s host address | 电子邮箱服务的服务器地址',
  `identify` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The identify info, for CRAMMD5 | 身份信息, 支持 CRAMMD5',
  `secret` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The secret, for CRAMMD5 | 邮箱密钥, 用于 CRAMMD5',
  `port` int unsigned DEFAULT NULL COMMENT 'The port of the host | 服务器端口',
  `tls` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Whether to use TLS | 是否采用 tls 加密',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is it the default provider | 是否为默认提供商',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Email Provider Configuration Table | 邮件服务商配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mcms_email_providers` DISABLE KEYS */;
/*!40000 ALTER TABLE `mcms_email_providers` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `mcms_sms_logs` (
  `id` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `phone_number` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The target phone number | 目标电话',
  `content` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The content | 发送的内容',
  `send_status` tinyint unsigned NOT NULL COMMENT 'The send status, 0 unknown 1 success 2 failed | 发送的状态, 0 未知， 1 成功， 2 失败',
  `provider` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The sms service provider | 短信服务提供商',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='SMS Log Table | 短信日志表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mcms_sms_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `mcms_sms_logs` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `mcms_sms_providers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The SMS provider name | 短信服务的提供商',
  `secret_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The secret ID | 密钥 ID',
  `secret_key` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The secret key | 密钥 Key',
  `region` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The service region | 服务器所在地区',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is it the default provider | 是否为默认提供商',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='SMS Provider Configuration Table | 短信服务商配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mcms_sms_providers` DISABLE KEYS */;
/*!40000 ALTER TABLE `mcms_sms_providers` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `mms_members` (
  `id` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `username` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Member''s login name | 登录名',
  `password` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Password | 密码',
  `nickname` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Nickname | 昵称',
  `mobile` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Mobile number | 手机号',
  `email` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Email | 邮箱号',
  `avatar` varchar(512) COLLATE utf8mb4_bin DEFAULT '' COMMENT 'Avatar | 头像路径',
  `wechat_open_id` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Wechat Open ID | 微信 Open ID',
  `expired_at` datetime DEFAULT NULL COMMENT 'Member expired time | 会员到期时间',
  `rank_id` bigint unsigned DEFAULT '2' COMMENT 'Member Rank ID | 会员等级ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `nickname` (`nickname`),
  UNIQUE KEY `member_username_email` (`username`,`email`),
  UNIQUE KEY `member_wechat_open_id` (`wechat_open_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mms_members` DISABLE KEYS */;
/*!40000 ALTER TABLE `mms_members` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `mms_oauth_providers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The provider''s name | 提供商名称',
  `client_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The client id | 客户端 id',
  `client_secret` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The client secret | 客户端密钥',
  `redirect_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The redirect url | 跳转地址',
  `scopes` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The scopes | 权限范围',
  `auth_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The auth url of the provider | 认证地址',
  `token_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The token url of the provider | 获取 token地址',
  `auth_style` bigint unsigned NOT NULL COMMENT 'The auth style, 0: auto detect 1: third party log in 2: log in with username and password',
  `info_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The URL to request user information by token | 用户信息请求地址',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mms_oauth_providers` DISABLE KEYS */;
/*!40000 ALTER TABLE `mms_oauth_providers` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `mms_ranks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Rank name | 等级名称',
  `code` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Rank code | 等级码',
  `description` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Rank description | 等级描述',
  `remark` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Remark | 备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `memberrank_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mms_ranks` DISABLE KEYS */;
/*!40000 ALTER TABLE `mms_ranks` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `mms_tokens` (
  `id` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `uuid` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT ' User''s UUID | 用户的UUID',
  `token` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Token string | Token 字符串',
  `username` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT 'unknown' COMMENT 'Username | 用户名',
  `source` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Log in source such as GitHub | Token 来源 （本地为core, 第三方如github等）',
  `expired_at` timestamp NOT NULL COMMENT ' Expire time | 过期时间',
  PRIMARY KEY (`id`),
  KEY `mms_token_uuid` (`uuid`),
  KEY `mms_token_expired_at` (`expired_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mms_tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `mms_tokens` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `role_menus` (
  `role_id` bigint unsigned NOT NULL,
  `menu_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `role_menus` DISABLE KEYS */;
INSERT IGNORE INTO `role_menus` VALUES (1,1),(1,2),(1,3),(1,4),(1,5),(1,6),(1,7),(1,8),(1,9),(1,10),(1,11),(1,12),(1,13),(1,14),(1,16),(1,17),(2,1),(2,2),(2,3),(2,4),(2,5),(2,6),(2,7),(2,8),(2,9),(2,10),(2,11),(2,12),(2,13),(2,14),(2,16),(2,17),(1,223),(2,223);
/*!40000 ALTER TABLE `role_menus` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `sys_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `path` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'API path | API 路径',
  `description` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'API description | API 描述',
  `api_group` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'API group | API 分组',
  `service_name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT 'Other' COMMENT 'Service name | 服务名称',
  `method` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT 'POST' COMMENT 'HTTP method | HTTP 请求类型',
  `is_required` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Whether is required | 是否必选',
  PRIMARY KEY (`id`),
  UNIQUE KEY `api_path_method` (`path`,`method`)
) ENGINE=InnoDB AUTO_INCREMENT=166 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='API Table | API接口表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sys_apis` DISABLE KEYS */;
INSERT IGNORE INTO `sys_apis` VALUES (1,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/login','apiDesc.userLogin','user','Core','POST',1),(2,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/register','apiDesc.userRegister','user','Core','POST',1),(3,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/create','apiDesc.createUser','user','Core','POST',0),(4,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/update','apiDesc.updateUser','user','Core','POST',0),(5,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/change_password','apiDesc.userChangePassword','user','Core','POST',1),(6,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/info','apiDesc.userInfo','user','Core','GET',1),(7,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/list','apiDesc.userList','user','Core','POST',0),(8,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/delete','apiDesc.deleteUser','user','Core','POST',0),(9,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/perm','apiDesc.userPermissions','user','Core','GET',1),(10,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/profile','apiDesc.userProfile','user','Core','GET',1),(11,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/profile','apiDesc.updateProfile','user','Core','POST',1),(12,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/logout','apiDesc.logout','user','Core','GET',1),(13,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user','apiDesc.getUserById','user','Core','POST',0),(14,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/refresh_token','apiDesc.refreshToken','user','Core','GET',0),(15,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/access_token','apiDesc.accessToken','user','Core','GET',0),(16,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/role/create','apiDesc.createRole','role','Core','POST',0),(17,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/role/update','apiDesc.updateRole','role','Core','POST',0),(18,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/role/delete','apiDesc.deleteRole','role','Core','POST',0),(19,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/role/list','apiDesc.roleList','role','Core','POST',0),(20,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/role','apiDesc.getRoleById','role','Core','POST',0),(21,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/menu/create','apiDesc.createMenu','menu','Core','POST',0),(22,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/menu/update','apiDesc.updateMenu','menu','Core','POST',0),(23,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/menu/delete','apiDesc.deleteMenu','menu','Core','POST',0),(24,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/menu/list','apiDesc.menuList','menu','Core','GET',0),(25,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/menu/role/list','apiDesc.menuRoleList','authority','Core','GET',1),(31,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/menu','apiDesc.getMenuById','menu','Core','POST',0),(32,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/captcha','apiDesc.captcha','captcha','Core','GET',1),(33,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/authority/api/create_or_update','apiDesc.createOrUpdateApiAuthority','authority','Core','POST',0),(34,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/authority/api/role','apiDesc.APIAuthorityOfRole','authority','Core','POST',0),(35,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/authority/menu/create_or_update','apiDesc.createOrUpdateMenuAuthority','authority','Core','POST',0),(36,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/authority/menu/role','apiDesc.menuAuthorityOfRole','authority','Core','POST',0),(37,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/api/create','apiDesc.createApi','api','Core','POST',0),(38,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/api/update','apiDesc.updateApi','api','Core','POST',0),(39,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/api/delete','apiDesc.deleteAPI','api','Core','POST',0),(40,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/api/list','apiDesc.APIList','api','Core','POST',0),(41,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/api','apiDesc.getApiById','api','Core','POST',0),(42,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/dictionary','apiDesc.getDictionaryById','dictionary','Core','POST',0),(43,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/dictionary/create','apiDesc.createDictionary','dictionary','Core','POST',0),(44,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/dictionary/update','apiDesc.updateDictionary','dictionary','Core','POST',0),(45,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/dictionary/delete','apiDesc.deleteDictionary','dictionary','Core','POST',0),(46,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/dictionary_detail/delete','apiDesc.deleteDictionaryDetail','dictionary','Core','POST',0),(47,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/dictionary_detail','apiDesc.getDictionaryDetailById','dictionary','Core','POST',0),(48,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/dictionary_detail/create','apiDesc.createDictionaryDetail','dictionary','Core','POST',0),(49,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/dictionary_detail/update','apiDesc.updateDictionaryDetail','dictionary','Core','POST',0),(50,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/dictionary_detail/list','apiDesc.getDictionaryListDetail','dictionary','Core','POST',0),(51,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/dictionary/list','apiDesc.getDictionaryList','dictionary','Core','POST',0),(52,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/dict/:name','apiDesc.getDictionaryDetailByDictionaryName','dictionary','Core','GET',0),(53,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/oauth_provider/create','apiDesc.createProvider','oauth','Core','POST',0),(54,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/oauth_provider/update','apiDesc.updateProvider','oauth','Core','POST',0),(55,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/oauth_provider/delete','apiDesc.deleteProvider','oauth','Core','POST',0),(56,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/oauth_provider/list','apiDesc.getProviderList','oauth','Core','POST',0),(57,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/oauth/login','apiDesc.oauthLogin','oauth','Core','POST',1),(58,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/oauth_provider','apiDesc.getProviderById','oauth','Core','POST',0),(59,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/token/create','apiDesc.createToken','token','Core','POST',0),(60,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/token/update','apiDesc.updateToken','token','Core','POST',0),(61,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/token/delete','apiDesc.deleteToken','token','Core','POST',0),(62,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/token/list','apiDesc.getTokenList','token','Core','POST',0),(63,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/token/logout','apiDesc.forceLoggingOut','token','Core','POST',0),(64,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/token','apiDesc.getTokenById','token','Core','POST',0),(65,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/department/create','apiDesc.createDepartment','department','Core','POST',0),(66,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/department/update','apiDesc.updateDepartment','department','Core','POST',0),(67,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/department/delete','apiDesc.deleteDepartment','department','Core','POST',0),(68,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/department/list','apiDesc.getDepartmentList','department','Core','POST',0),(69,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/department','apiDesc.getDepartmentById','department','Core','POST',0),(70,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/position/create','apiDesc.createPosition','position','Core','POST',0),(71,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/position/update','apiDesc.updatePosition','position','Core','POST',0),(72,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/position/delete','apiDesc.deletePosition','position','Core','POST',0),(73,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/position/list','apiDesc.getPositionList','position','Core','POST',0),(74,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/position','apiDesc.getPositionById','position','Core','POST',0),(85,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/configuration/create','apiDesc.createConfiguration','configuration','Core','POST',0),(86,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/configuration/update','apiDesc.updateConfiguration','configuration','Core','POST',0),(87,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/configuration/delete','apiDesc.deleteConfiguration','configuration','Core','POST',0),(88,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/configuration/list','apiDesc.getConfigurationList','configuration','Core','POST',0),(89,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/configuration','apiDesc.getConfigurationById','configuration','Core','POST',0),(151,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/core/ping','apiDesc.ping','core','Core','GET',0),(152,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/core/init/database','apiDesc.initDatabase','core','Core','GET',1),(153,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/core/init/job_database','apiDesc.initJobDatabase','core','Core','GET',1),(154,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/core/init/mcms_database','apiDesc.initMcmsDatabase','core','Core','GET',1),(155,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/captcha/email','apiDesc.captchaEmail','captcha','Core','GET',1),(156,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/captcha/sms','apiDesc.captchaSms','captcha','Core','GET',1),(157,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/login_by_email','apiDesc.userLoginByEmail','user','Core','POST',1),(158,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/login_by_sms','apiDesc.userLoginBySms','user','Core','POST',1),(159,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/register_by_email','apiDesc.userRegisterByEmail','user','Core','POST',1),(160,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/register_by_sms','apiDesc.userRegisterBySms','user','Core','POST',1),(161,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/reset_password_by_email','apiDesc.userResetPasswordByEmail','user','Core','POST',1),(162,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/user/reset_password_by_sms','apiDesc.userResetPasswordBySms','user','Core','POST',1),(163,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/oauth/login/callback','apiDesc.oauthCallback','oauth','Core','GET',1),(164,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/configuration/system/list','apiDesc.systemConfiguration','configuration','Core','POST',0),(165,'2026-10-15 06:49:53','2026-10-15 06:49:53','/sys-api/dict/public/:name','apiDesc.publicDict','dictionary','Core','GET',1);
/*!40000 ALTER TABLE `sys_apis` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `sys_configuration` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `sort` int unsigned NOT NULL DEFAULT '1' COMMENT 'Sort Number | 排序编号',
  `state` tinyint(1) DEFAULT '1' COMMENT 'State true: normal false: ban | 状态 true 正常 false 禁用',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Configurarion name | 配置名称',
  `key` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Configuration key | 配置的键名',
  `value` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Configuraion value | 配置的值',
  `category` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Configuration category | 配置的分类',
  `remark` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Remark | 备注',
  PRIMARY KEY (`id`),
  KEY `configuration_key` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Dynamic Configuration Table | 动态配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sys_configuration` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_configuration` ENABLE KEYS */;

-- 种子：系统默认语言（category=system 即公开配置，前端登录页据此设置默认语言）
INSERT INTO `sys_configuration` (`created_at`, `updated_at`, `sort`, `state`, `name`, `key`, `value`, `category`, `remark`)
SELECT NOW(), NOW(), 1, 1, '默认语言', 'sys.ui.defaultLocale', 'zh_CN', 'system', '系统默认语言：zh_CN 中文 / en 英文'
WHERE NOT EXISTS (SELECT 1 FROM `sys_configuration` WHERE `key` = 'sys.ui.defaultLocale');

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `sys_departments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `sort` int unsigned NOT NULL DEFAULT '1' COMMENT 'Sort Number | 排序编号',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Department name | 部门名称',
  `ancestors` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Parents'' IDs | 父级列表',
  `leader` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Department leader | 部门负责人',
  `phone` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Leader''s phone number | 负责人电话',
  `email` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Leader''s email | 部门负责人电子邮箱',
  `remark` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Remark | 备注',
  `parent_id` bigint unsigned DEFAULT '0' COMMENT 'Parent department ID | 父级部门ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Department Table | 部门表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sys_departments` DISABLE KEYS */;
INSERT IGNORE INTO `sys_departments` VALUES (1,'2026-10-15 06:49:53','2026-10-15 06:49:53',1,1,'department.managementDepartment','','admin','18888888888','simpleadmin@gmail.com','Super Administrator',1000000);
/*!40000 ALTER TABLE `sys_departments` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `sys_dictionaries` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `title` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The title shown in the ui | 展示名称 （建议配合i18n）',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The name of dictionary for search | 字典搜索名称',
  `desc` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The description of dictionary | 字典的描述',
  `is_public` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Whether to be public for everyone | 是否公开词典，无需登录即可访问',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Dictionary Table | 字典信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sys_dictionaries` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_dictionaries` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `sys_dictionary_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `sort` int unsigned NOT NULL DEFAULT '1' COMMENT 'Sort Number | 排序编号',
  `title` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The title shown in the ui | 展示名称 （建议配合i18n）',
  `key` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'key | 键',
  `value` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'value | 值',
  `dictionary_id` bigint unsigned DEFAULT NULL COMMENT 'Dictionary ID | 字典ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Dictionary Key/Value Table | 字典键值表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sys_dictionary_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_dictionary_details` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `sys_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `sort` int unsigned NOT NULL DEFAULT '1' COMMENT 'Sort Number | 排序编号',
  `menu_level` int unsigned NOT NULL COMMENT 'Menu level | 菜单层级',
  `menu_type` int unsigned NOT NULL COMMENT 'Menu type | 菜单类型 （菜单或目录）0 目录 1 菜单',
  `path` varchar(255) COLLATE utf8mb4_bin DEFAULT '' COMMENT 'Index path | 菜单路由路径',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Index name | 菜单名称',
  `redirect` varchar(255) COLLATE utf8mb4_bin DEFAULT '' COMMENT 'Redirect path | 跳转路径 （外链）',
  `component` varchar(255) COLLATE utf8mb4_bin DEFAULT '' COMMENT 'The path of vue file | 组件路径',
  `disabled` tinyint(1) DEFAULT '0' COMMENT 'Disable status | 是否停用',
  `service_name` varchar(255) COLLATE utf8mb4_bin DEFAULT 'Other' COMMENT 'Service Name | 服务名称',
  `permission` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Permission symbol | 权限标识',
  `title` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Menu name | 菜单显示标题',
  `icon` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Menu icon | 菜单图标',
  `hide_menu` tinyint(1) DEFAULT '0' COMMENT 'Hide menu | 是否隐藏菜单',
  `hide_breadcrumb` tinyint(1) DEFAULT '0' COMMENT 'Hide the breadcrumb | 隐藏面包屑',
  `ignore_keep_alive` tinyint(1) DEFAULT '0' COMMENT 'Do not keep alive the tab | 取消页面缓存',
  `hide_tab` tinyint(1) DEFAULT '0' COMMENT 'Hide the tab header | 隐藏页头',
  `frame_src` varchar(255) COLLATE utf8mb4_bin DEFAULT '' COMMENT 'Show iframe | 内嵌 iframe',
  `carry_param` tinyint(1) DEFAULT '0' COMMENT 'The route carries parameters or not | 携带参数',
  `hide_children_in_menu` tinyint(1) DEFAULT '0' COMMENT 'Hide children menu or not | 隐藏所有子菜单',
  `affix` tinyint(1) DEFAULT '0' COMMENT 'Affix tab | Tab 固定',
  `dynamic_level` int unsigned DEFAULT '20' COMMENT 'The maximum number of pages the router can open | 能打开的子TAB数',
  `real_path` varchar(255) COLLATE utf8mb4_bin DEFAULT '' COMMENT 'The real path of the route without dynamic part | 菜单路由不包含参数部分',
  `parent_id` bigint unsigned DEFAULT '100000' COMMENT 'Parent menu ID | 父菜单ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `menu_name` (`name`),
  UNIQUE KEY `menu_path` (`path`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Menu Table | 菜单表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sys_menus` DISABLE KEYS */;
INSERT IGNORE INTO `sys_menus` VALUES (1,'2026-10-15 06:49:53','2026-10-15 06:49:53',1,1,0,'/dashboard_dir','DashBoardDir','','LAYOUT',0,'Core',NULL,'route.managementCenter','ic:baseline-view-in-ar',0,0,0,0,'',0,0,0,20,'',1000000),(2,'2026-10-15 06:49:53','2026-10-15 06:49:53',0,2,1,'/dashboard','Dashboard','','/dashboard/workbench/index',0,'Core',NULL,'route.dashboard','ant-design:home-outlined',0,0,0,0,'',0,0,0,20,'',1),(3,'2026-10-15 06:49:53','2026-10-15 06:49:53',999,1,0,'/system','SystemManagement','','LAYOUT',0,'Core',NULL,'route.systemManagementTitle','ant-design:tool-outlined',0,0,0,0,'',0,0,0,20,'',1000000),(4,'2026-10-15 06:49:53','2026-10-15 06:49:53',1,2,1,'/menu','MenuManagement','','/sys/menu/index',0,'Core',NULL,'route.menuManagementTitle','ant-design:bars-outlined',0,0,0,0,'',0,0,0,20,'',3),(5,'2026-10-15 06:49:53','2026-10-15 06:49:53',2,2,1,'/role','RoleManagement','','/sys/role/index',0,'Core',NULL,'route.roleManagementTitle','ant-design:user-outlined',0,0,0,0,'',0,0,0,20,'',3),(6,'2026-10-15 06:49:53','2026-10-15 06:49:53',3,2,1,'/user','UserManagement','','/sys/user/index',0,'Core',NULL,'route.userManagementTitle','ant-design:user-outlined',0,0,0,0,'',0,0,0,20,'',3),(7,'2026-10-15 06:49:53','2026-10-15 06:49:53',4,2,1,'/department','DepartmentManagement','','/sys/department/index',0,'Core',NULL,'route.departmentManagement','ic:outline-people-alt',0,0,0,0,'',0,0,0,20,'',3),(8,'2026-10-15 06:49:53','2026-10-15 06:49:53',5,2,1,'/api','APIManagement','','/sys/api/index',0,'Core',NULL,'route.apiManagementTitle','ant-design:api-outlined',0,0,0,0,'',0,0,0,20,'',3),(9,'2026-10-15 06:49:53','2026-10-15 06:49:53',6,2,1,'/dictionary','DictionaryManagement','','/sys/dictionary/index',0,'Core',NULL,'route.dictionaryManagementTitle','ant-design:book-outlined',0,0,0,0,'',0,1,0,20,'',3),(10,'2026-10-15 06:49:53','2026-10-15 06:49:53',1000,1,0,'/other','OtherPages','','LAYOUT',0,'Core',NULL,'route.otherPages','ant-design:question-circle-outlined',1,0,0,0,'',0,0,0,20,'',1000000),(11,'2026-10-15 06:49:53','2026-10-15 06:49:53',3,1,1,'/profile','Profile','','/sys/profile/index',0,'Core',NULL,'route.userProfileTitle','ant-design:profile-outlined',1,0,0,0,'',0,0,0,20,'',10),(12,'2026-10-15 06:49:53','2026-10-15 06:49:53',6,2,1,'/oauth','OauthManagement','','/sys/oauth/index',0,'Core',NULL,'route.oauthManagement','ant-design:unlock-filled',0,0,0,0,'',0,0,0,20,'',3),(13,'2026-10-15 06:49:53','2026-10-15 06:49:53',7,2,1,'/token','TokenManagement','','/sys/token/index',0,'Core',NULL,'route.tokenManagement','ant-design:lock-outlined',0,0,0,0,'',0,0,0,20,'',3),(14,'2026-10-15 06:49:53','2026-10-15 06:49:53',8,2,1,'/position','PositionManagement','','/sys/position/index',0,'Core',NULL,'route.positionManagement','ic:twotone-work-outline',0,0,0,0,'',0,0,0,20,'',3),(16,'2026-10-15 06:49:53','2026-10-15 06:49:53',9,2,1,'/configuration','ConfigurationManagement','','/sys/configuration/index',0,'Core',NULL,'route.configurationManagement','carbon:data-2',0,0,0,0,'',0,0,0,20,'',3),(17,'2026-10-15 06:49:53','2026-10-15 06:49:53',99,1,1,'/module_store','ModuleStore','','IFrame',0,'Core',NULL,'route.moduleStore','ion:bag-handle-outline',0,0,0,0,'https://simple-admin.simple-hub.top/#/store/index',0,0,0,20,'',1000000),(223,'2026-10-15 06:49:53','2026-10-15 06:49:53',1,2,1,'/dictionary/detail/:dictionaryId','DictionaryDetail','','/sys/dictionaryDetail/index',0,'Core',NULL,'route.dictionaryDetailManagementTitle','',1,0,0,0,'',0,0,0,20,'',9);
/*!40000 ALTER TABLE `sys_menus` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `sys_oauth_providers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The provider''s name | 提供商名称',
  `client_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The client id | 客户端 id',
  `client_secret` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The client secret | 客户端密钥',
  `redirect_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The redirect url | 跳转地址',
  `scopes` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The scopes | 权限范围',
  `auth_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The auth url of the provider | 认证地址',
  `token_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The token url of the provider | 获取 token地址',
  `auth_style` bigint unsigned NOT NULL COMMENT 'The auth style, 0: auto detect 1: third party log in 2: log in with username and password | 鉴权方式 0 自动 1 第三方登录 2 使用用户名密码',
  `info_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The URL to request user information by token | 用户信息请求地址',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Oauth Provider Configuration Table | 三方登录配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sys_oauth_providers` DISABLE KEYS */;
INSERT IGNORE INTO `sys_oauth_providers` VALUES (1,'2026-10-15 06:49:53','2026-10-15 06:49:53','google','your client id','your client secret','http://localhost:3100/oauth/login/callback','email openid','https://accounts.google.com/o/oauth2/auth','https://oauth2.googleapis.com/token',1,'https://www.googleapis.com/oauth2/v2/userinfo?access_token=TOKEN'),(2,'2026-10-15 06:49:53','2026-10-15 06:49:53','github','your client id','your client secret','http://localhost:3100/oauth/login/callback','email openid','https://github.com/login/oauth/authorize','https://github.com/login/oauth/access_token',2,'https://api.github.com/user');
/*!40000 ALTER TABLE `sys_oauth_providers` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `sys_positions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `sort` int unsigned NOT NULL DEFAULT '1' COMMENT 'Sort Number | 排序编号',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Position Name | 职位名称',
  `code` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The code of position | 职位编码',
  `remark` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Remark | 备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `position_code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Position Table | 职位信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sys_positions` DISABLE KEYS */;
INSERT IGNORE INTO `sys_positions` VALUES (1,'2026-10-15 06:49:53','2026-10-15 06:49:53',1,1,'position.ceo','001','CEO');
/*!40000 ALTER TABLE `sys_positions` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `sys_roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Role name | 角色名',
  `code` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Role code for permission control in front end | 角色码，用于前端权限控制',
  `remark` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'Remark | 备注',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT 'Order number | 排序编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Role Table | 角色信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sys_roles` DISABLE KEYS */;
INSERT IGNORE INTO `sys_roles` VALUES (1,'2026-10-15 06:49:53','2026-10-15 18:36:14',1,'role.admin','001','超级管理员',1),(2,'2026-10-15 06:49:53','2026-10-15 06:49:53',1,'role.stuff','002','普通员工',2);
/*!40000 ALTER TABLE `sys_roles` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `sys_task_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `started_at` timestamp NOT NULL COMMENT 'Task Started Time | 任务启动时间',
  `finished_at` timestamp NOT NULL COMMENT 'Task Finished Time | 任务完成时间',
  `result` tinyint unsigned NOT NULL COMMENT 'The Task Process Result | 任务执行结果',
  `task_task_logs` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Task Log Table | 任务日志表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sys_task_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_task_logs` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `sys_tasks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Task Name | 任务名称',
  `task_group` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Task Group | 任务分组',
  `cron_expression` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Cron expression | 定时任务表达式',
  `pattern` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Cron Pattern | 任务的模式 （用于区分和确定要执行的任务）',
  `payload` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'The data used in cron (JSON string) | 任务需要的数据(JSON 字符串)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `task_pattern` (`pattern`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Task Configuration Table | 任务配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sys_tasks` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_tasks` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `sys_tokens` (
  `id` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `uuid` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT ' User''s UUID | 用户的UUID',
  `username` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT 'unknown' COMMENT 'Username | 用户名',
  `token` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Token string | Token 字符串',
  `source` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Log in source such as GitHub | Token 来源 （本地为core, 第三方如github等）',
  `expired_at` timestamp NOT NULL COMMENT ' Expire time | 过期时间',
  PRIMARY KEY (`id`),
  KEY `token_uuid` (`uuid`),
  KEY `token_expired_at` (`expired_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Token Log Table | 令牌信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sys_tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_tokens` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `sys_users` (
  `id` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT 'UUID',
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `status` tinyint unsigned DEFAULT '1' COMMENT 'Status 1: normal 2: ban | 状态 1 正常 2 禁用',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Delete Time | 删除日期',
  `username` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'User''s login name | 登录名',
  `password` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Password | 密码',
  `nickname` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Nickname | 昵称',
  `description` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'The description of user | 用户的描述信息',
  `home_path` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '/dashboard' COMMENT 'The home page that the user enters after logging in | 用户登陆后进入的首页',
  `mobile` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Mobile number | 手机号',
  `email` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Email | 邮箱号',
  `avatar` varchar(512) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Avatar | 头像路径',
  `expired_at` datetime DEFAULT NULL COMMENT 'The expired time | 到期时间',
  `department_id` bigint unsigned DEFAULT '1' COMMENT 'Department ID | 部门ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `nickname` (`nickname`),
  UNIQUE KEY `user_username_email` (`username`,`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='User Table | 用户信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sys_users` DISABLE KEYS */;
INSERT IGNORE INTO `sys_users` VALUES ('01a1400a-8316-7632-933f-d451f527c175','2026-10-15 06:49:53','2026-10-15 06:49:53',1,NULL,'admin','$2a$10$eUJiAWprIWJE/T3dS.PU..6jFC8QDcBb9/2yb1XyyNf/wqcpfUqZC','admin',NULL,'/dashboard',NULL,'simple_admin@gmail.com',NULL,NULL,1);
/*!40000 ALTER TABLE `sys_users` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_positions` (
  `user_id` char(36) COLLATE utf8mb4_bin NOT NULL,
  `position_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`position_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `user_positions` DISABLE KEYS */;
INSERT IGNORE INTO `user_positions` VALUES ('01a1400a-8316-7632-933f-d451f527c175',1);
/*!40000 ALTER TABLE `user_positions` ENABLE KEYS */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_roles` (
  `user_id` char(36) COLLATE utf8mb4_bin NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
INSERT IGNORE INTO `user_roles` VALUES ('01a1400a-8316-7632-933f-d451f527c175',1);
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

