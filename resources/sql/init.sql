/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE IF NOT EXISTS `website` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `website`;

CREATE TABLE IF NOT EXISTS `auth_cookies` (
                                              `userid` varchar(50) DEFAULT NULL,
                                              `value` varchar(50) DEFAULT NULL,
                                              `type` int(11) DEFAULT NULL,
                                              `expire` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `paste` (
                                       `id` varchar(20) NOT NULL DEFAULT '',
                                       `userid` varchar(50) NOT NULL DEFAULT '',
                                       `raw` longtext NOT NULL,
                                       `language` varchar(50) NOT NULL DEFAULT '',
                                       `expire` int(11) NOT NULL DEFAULT 0,
                                       `password` varchar(50) DEFAULT '',
                                       PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `users` (
                                       `id` varchar(50) NOT NULL,
                                       `username` varchar(50) NOT NULL,
                                       `email` varchar(50) NOT NULL,
                                       `password` varchar(50) NOT NULL,
                                       `email_verified` tinyint(4) NOT NULL DEFAULT 0,
                                       `avatar` varchar(20) DEFAULT NULL,
                                       PRIMARY KEY (`id`),
                                       UNIQUE KEY `id` (`id`),
                                       UNIQUE KEY `username` (`username`),
                                       UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `verify_email` (
                                              `userid` varchar(50) DEFAULT NULL,
                                              `value` varchar(50) DEFAULT NULL,
                                              `expire` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
