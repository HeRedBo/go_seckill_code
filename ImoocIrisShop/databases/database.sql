
## 创建数据库
CREATE DATABASE `imooc_shop` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- 产品基础信息表 
CREATE TABLE `product` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `product_name` varchar(100) DEFAULT '' COMMENT '产品名称',
    `product_num` int(10) unsigned DEFAULT '0' COMMENT '产品数量',
    `product_image` varchar(255) DEFAULT '' COMMENT '产品图片地址',
    `product_url` varchar(255) DEFAULT '' COMMENT '产品地址',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品信息表';

-- 商城订单表 简易版
CREATE TABLE `orders` (
   `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
   `user_id`  int(10) DEFAULT '0' COMMENT '用户ID',
   `product_id` int(10) unsigned DEFAULT '0' COMMENT '产品ID',
   `order_status` int(1) DEFAULT '0' COMMENT '订单状态',
   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单信息表';


