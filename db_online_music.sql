/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50642
 Source Host           : localhost:3306
 Source Schema         : db_online_music

 Target Server Type    : MySQL
 Target Server Version : 50642
 File Encoding         : 65001

 Date: 04/04/2019 00:01:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_channel
-- ----------------------------
DROP TABLE IF EXISTS `tb_channel`;
CREATE TABLE `tb_channel`  (
  `channel_id` int(10) UNSIGNED NOT NULL COMMENT '歌曲来源渠道id',
  `channel_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '渠道名称',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `create_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '创建人',
  `create_user_id` int(10) UNSIGNED NOT NULL COMMENT '创建人ID',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  `update_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '更新人',
  `update_user_id` int(10) UNSIGNED NOT NULL COMMENT '更新人ID',
  `del_status` tinyint(1) NOT NULL DEFAULT 2 COMMENT '删除状态 1删除，2不删除',
  PRIMARY KEY (`channel_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '歌曲或者歌单来源渠道表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for tb_login
-- ----------------------------
DROP TABLE IF EXISTS `tb_login`;
CREATE TABLE `tb_login`  (
  `login_id` int(10) UNSIGNED NOT NULL COMMENT '登录id',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户id',
  `login_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登录名',
  `password` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `create_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '创建人',
  `create_user_id` int(10) UNSIGNED NOT NULL COMMENT '创建人ID',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  `update_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '更新人',
  `update_user_id` int(10) UNSIGNED NOT NULL COMMENT '更新人ID',
  PRIMARY KEY (`login_id`, `user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户登录表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for tb_song
-- ----------------------------
DROP TABLE IF EXISTS `tb_song`;
CREATE TABLE `tb_song`  (
  `song_id` int(10) UNSIGNED NOT NULL COMMENT '歌曲id',
  `song_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '歌曲名称',
  `song_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '歌曲url',
  `song_lyric` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '歌词url',
  `channel_id` int(11) NOT NULL COMMENT '歌曲来源渠道id',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `create_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '创建人',
  `create_user_id` int(10) UNSIGNED NOT NULL COMMENT '创建人ID',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  `update_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '更新人',
  `update_user_id` int(10) UNSIGNED NOT NULL COMMENT '更新人ID',
  `del_status` tinyint(1) NOT NULL DEFAULT 2 COMMENT '删除状态 1删除，2不删除',
  PRIMARY KEY (`song_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '歌曲表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for tb_song_cover
-- ----------------------------
DROP TABLE IF EXISTS `tb_song_cover`;
CREATE TABLE `tb_song_cover`  (
  `song_cover_id` int(10) UNSIGNED NOT NULL COMMENT '歌单id',
  `type` tinyint(4) NOT NULL COMMENT '歌单类型',
  `channel_id` int(10) UNSIGNED NOT NULL COMMENT '歌曲来源渠道id',
  `song_cover_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '歌单名称',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `create_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '创建人',
  `create_user_id` int(10) UNSIGNED NOT NULL COMMENT '创建人ID',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  `update_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '更新人',
  `update_user_id` int(10) UNSIGNED NOT NULL COMMENT '更新人ID',
  `del_status` tinyint(1) NOT NULL DEFAULT 2 COMMENT '删除状态 1删除，2不删除',
  PRIMARY KEY (`song_cover_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '歌单表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for tb_song_cover_song
-- ----------------------------
DROP TABLE IF EXISTS `tb_song_cover_song`;
CREATE TABLE `tb_song_cover_song`  (
  `song_cover_song_id` int(10) UNSIGNED NOT NULL COMMENT '歌单歌曲id',
  `song_id` int(10) UNSIGNED NOT NULL COMMENT '歌曲id',
  `song_cover_id` int(10) UNSIGNED NOT NULL COMMENT '歌单id',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `create_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '创建人',
  `create_user_id` int(10) UNSIGNED NOT NULL COMMENT '创建人ID',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  `update_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '更新人',
  `update_user_id` int(10) UNSIGNED NOT NULL COMMENT '更新人ID',
  `del_status` tinyint(1) NOT NULL DEFAULT 2 COMMENT '删除状态 1删除，2不删除',
  PRIMARY KEY (`song_cover_song_id`, `song_id`, `song_cover_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '歌单歌曲表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user`  (
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户id',
  `user_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '姓名',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `create_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '创建人',
  `create_user_id` int(10) UNSIGNED NOT NULL COMMENT '创建人ID',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  `update_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '更新人',
  `update_user_id` int(10) UNSIGNED NOT NULL COMMENT '更新人ID',
  `del_state` tinyint(1) NOT NULL DEFAULT 2 COMMENT '删除状态 1删除，2不删除',
  `gender` char(3) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '性别',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for tb_user_song_cover
-- ----------------------------
DROP TABLE IF EXISTS `tb_user_song_cover`;
CREATE TABLE `tb_user_song_cover`  (
  `user_song_cover_id` int(10) UNSIGNED NOT NULL COMMENT '用户歌单id',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户id',
  `song_cover_id` int(10) UNSIGNED NOT NULL COMMENT '歌单id',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `create_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '创建人',
  `create_user_id` int(10) UNSIGNED NOT NULL COMMENT '创建人ID',
  `update_time` datetime(0) NOT NULL COMMENT '更新时间',
  `update_user` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '更新人',
  `update_user_id` int(10) UNSIGNED NOT NULL COMMENT '更新人ID',
  `del_state` tinyint(1) NOT NULL DEFAULT 2 COMMENT '删除状态 1删除，2不删除',
  PRIMARY KEY (`user_song_cover_id`, `user_id`, `song_cover_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户歌单表' ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
