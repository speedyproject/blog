CREATE TABLE `t_admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(15) COLLATE utf8_bin NOT NULL,
  `passwd` varchar(64) COLLATE utf8_bin NOT NULL,
  `email` varchar(45) COLLATE utf8_bin DEFAULT NULL,
  `sign` varchar(64) COLLATE utf8_bin NOT NULL,
  `lock` int(11) DEFAULT '0',
  `last_ip` varchar(20) COLLATE utf8_bin DEFAULT '0.0.0.0',
  `last_login` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
