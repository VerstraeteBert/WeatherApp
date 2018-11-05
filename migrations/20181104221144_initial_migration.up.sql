CREATE TABLE IF NOT EXISTS `readings` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	`timestamp` datetime NOT NULL,
	`degreescelsius` decimal(10,2) NOT NULL,
	PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8;