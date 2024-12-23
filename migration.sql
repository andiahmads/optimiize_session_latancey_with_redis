CREATE TABLE `users` (
  `id` varchar(100) NOT NULL DEFAULT '',
  `username` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `products` (
  `id` varchar(100) NOT NULL DEFAULT '',
  `name` varchar(100) DEFAULT NULL,
  `stock` int(10) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



INSERT INTO `products` (`id`, `name`, `stock`, `created_at`, `updated_at`) VALUES
('1', 'Product 1', 10, current_timestamp(), current_timestamp()),
('2', 'Product 2', 20, current_timestamp(), current_timestamp()),
('3', 'Product 3', 30, current_timestamp(), current_timestamp()),
('4', 'Product 4', 40, current_timestamp(), current_timestamp()),
('5', 'Product 5', 50, current_timestamp(), current_timestamp()),
('6', 'Product 6', 60, current_timestamp(), current_timestamp()),
('7', 'Product 7', 70, current_timestamp(), current_timestamp()),
('8', 'Product 8', 80, current_timestamp(), current_timestamp()),
('9', 'Product 9', 90, current_timestamp(), current_timestamp()),
('10', 'Product 10', 100, current_timestamp(), current_timestamp()),
('11', 'Product 11', 110, current_timestamp(), current_timestamp()),
('12', 'Product 12', 120, current_timestamp(), current_timestamp()),
('13', 'Product 13', 130, current_timestamp(), current_timestamp()),
('14', 'Product 14', 140, current_timestamp(), current_timestamp()),
('15', 'Product 15', 150, current_timestamp(), current_timestamp()),
('16', 'Product 16', 160, current_timestamp(), current_timestamp()),
('17', 'Product 17', 170, current_timestamp(), current_timestamp()),
('18', 'Product 18', 180, current_timestamp(), current_timestamp()),
('19', 'Product 19', 190, current_timestamp(), current_timestamp()),
('20', 'Product 20', 200, current_timestamp(), current_timestamp()),
('21', 'Product 21', 210, current_timestamp(), current_timestamp()),
('22', 'Product 22', 220, current_timestamp(), current_timestamp()),
('23', 'Product 23', 230, current_timestamp(), current_timestamp()),
('24', 'Product 24', 240, current_timestamp(), current_timestamp()),
('25', 'Product 25', 250, current_timestamp(), current_timestamp()),
('26', 'Product 26', 260, current_timestamp(), current_timestamp()),
('27', 'Product 27', 270, current_timestamp(), current_timestamp()),
('28', 'Product 28', 280, current_timestamp(), current_timestamp()),
('29', 'Product 29', 290, current_timestamp(), current_timestamp()),
('30', 'Product 30', 300, current_timestamp(), current_timestamp()),
('31', 'Product 31', 310, current_timestamp(), current_timestamp()),
('32', 'Product 32', 320, current_timestamp(), current_timestamp()),
('33', 'Product 33', 330, current_timestamp(), current_timestamp()),
('34', 'Product 34', 340, current_timestamp(), current_timestamp()),
('35', 'Product 35', 350, current_timestamp(), current_timestamp()),
('36', 'Product 36', 360, current_timestamp(), current_timestamp()),
('37', 'Product 37', 370, current_timestamp(), current_timestamp()),
('38', 'Product 38', 380, current_timestamp(), current_timestamp()),
('39', 'Product 39', 390, current_timestamp(), current_timestamp()),
('40', 'Product 40', 400, current_timestamp(), current_timestamp()),
('41', 'Product 41', 410, current_timestamp(), current_timestamp()),
('42', 'Product 42', 420, current_timestamp(), current_timestamp()),
('43', 'Product 43', 430, current_timestamp(), current_timestamp()),
('44', 'Product 44', 440, current_timestamp(), current_timestamp()),
('45', 'Product 45', 450, current_timestamp(), current_timestamp()),
('46', 'Product 46', 460, current_timestamp(), current_timestamp()),
('47', 'Product 47', 470, current_timestamp(), current_timestamp()),
('48', 'Product 48', 480, current_timestamp(), current_timestamp()),
('49', 'Product 49', 490, current_timestamp(), current_timestamp()),
('50', 'Product 50', 500, current_timestamp(), current_timestamp()),
('51', 'Product 51', 510, current_timestamp(), current_timestamp()),
('52', 'Product 52', 520, current_timestamp(), current_timestamp()),
('53', 'Product 53', 530, current_timestamp(), current_timestamp()),
('54', 'Product 54', 540, current_timestamp(), current_timestamp()),
('55', 'Product 55', 550, current_timestamp(), current_timestamp()),
('56', 'Product 56', 560, current_timestamp(), current_timestamp()),
('57', 'Product 57', 570, current_timestamp(), current_timestamp()),
('58', 'Product 58', 580, current_timestamp(), current_timestamp()),
('59', 'Product 59', 590, current_timestamp(), current_timestamp()),
('60', 'Product 60', 600, current_timestamp(), current_timestamp()),
('61', 'Product 61', 610, current_timestamp(), current_timestamp()),
('62', 'Product 62', 620, current_timestamp(), current_timestamp()),
('63', 'Product 63', 630, current_timestamp(), current_timestamp()),
('64', 'Product 64', 640, current_timestamp(), current_timestamp()),
('65', 'Product 65', 650, current_timestamp(), current_timestamp()),
('66', 'Product 66', 660, current_timestamp(), current_timestamp()),
('67', 'Product 67', 670, current_timestamp(), current_timestamp()),
('68', 'Product 68', 680, current_timestamp(), current_timestamp()),
('69', 'Product 69', 690, current_timestamp(), current_timestamp()),
('70', 'Product 70', 700, current_timestamp(), current_timestamp()),
('71', 'Product 71', 710, current_timestamp(), current_timestamp()),
('72', 'Product 72', 720, current_timestamp(), current_timestamp()),
('73', 'Product 73', 730, current_timestamp(), current_timestamp()),
('74', 'Product 74', 740, current_timestamp(), current_timestamp()),
('75', 'Product 75', 750, current_timestamp(), current_timestamp()),
('76', 'Product 76', 760, current_timestamp(), current_timestamp()),
('77', 'Product 77', 770, current_timestamp(), current_timestamp()),
('78', 'Product 78', 780, current_timestamp(), current_timestamp()),
('79', 'Product 79', 790, current_timestamp(), current_timestamp()),
('80', 'Product 80', 800, current_timestamp(), current_timestamp()),
('81', 'Product 81', 810, current_timestamp(), current_timestamp()),
('82', 'Product 82', 820, current_timestamp(), current_timestamp()),
('83', 'Product 83', 830, current_timestamp(), current_timestamp()),
('84', 'Product 84', 840, current_timestamp(), current_timestamp()),
('85', 'Product 85', 850, current_timestamp(), current_timestamp()),
('86', 'Product 86', 860, current_timestamp(), current_timestamp()),
('87', 'Product 87', 870, current_timestamp(), current_timestamp()),
('88', 'Product 88', 880, current_timestamp(), current_timestamp()),
('89', 'Product 89', 890, current_timestamp(), current_timestamp()),
('90', 'Product 90', 900, current_timestamp(), current_timestamp()),
('91', 'Product 91', 910, current_timestamp(), current_timestamp()),
('92', 'Product 92', 920, current_timestamp(), current_timestamp()),
('93', 'Product 93', 930, current_timestamp(), current_timestamp()),
('94', 'Product 94', 940, current_timestamp(), current_timestamp()),
('95', 'Product 95', 950, current_timestamp(), current_timestamp()),
('96', 'Product 96', 960, current_timestamp(), current_timestamp()),
('97', 'Product 97', 970, current_timestamp(), current_timestamp()),
('98', 'Product 98', 980, current_timestamp(), current_timestamp()),
('99', 'Product 99', 990, current_timestamp(), current_timestamp()),
('100', 'Product 100', 1000, current_timestamp(), current_timestamp());
