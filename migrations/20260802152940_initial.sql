-- Create "leaderboard_messages" table
CREATE TABLE `leaderboard_messages` (
  `guild_id` varchar(191) NOT NULL,
  `leaderboard_type` varchar(191) NOT NULL,
  `channel_id` longtext NULL,
  `message_id` longtext NULL,
  PRIMARY KEY (`guild_id`, `leaderboard_type`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "server_preferences" table
CREATE TABLE `server_preferences` (
  `guild_id` varchar(191) NOT NULL,
  `season` bigint NULL,
  `is_season_active` bool NULL,
  `leaderboard_channel_id` longtext NULL,
  PRIMARY KEY (`guild_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "speedruns" table
CREATE TABLE `speedruns` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `user_id` varchar(255) NULL,
  `time_in_seconds` bigint NULL,
  `weapon_type` varchar(255) NULL,
  `proof_link` varchar(255) NULL,
  `season` bigint NULL,
  `is_verified` bool NULL,
  `guild_id` longtext NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_speedruns_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
