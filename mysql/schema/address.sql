CREATE TABLE `address` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
 )

INSERT INTO `address` (`id`, `email`, `created_at`)
  VALUES
	(1, 'test@test.co.jp', '2019-05-04 03:23:39'),
	(2, 'aaa@aaa.co.jp', '2019-05-04 03:23:46');
