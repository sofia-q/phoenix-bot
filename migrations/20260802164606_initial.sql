-- Create "leaderboard_messages" table
CREATE TABLE `leaderboard_messages` (
  `guild_id` varchar(20) NOT NULL,
  `leaderboard_type` varchar(64) NOT NULL,
  `channel_id` varchar(20) NOT NULL,
  `message_id` varchar(20) NOT NULL,
  PRIMARY KEY (`guild_id`, `leaderboard_type`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "seasons" table
CREATE TABLE `seasons` (
  `id` char(36) NOT NULL,
  `guild_id` varchar(20) NOT NULL,
  `number` bigint NOT NULL,
  `name` varchar(255) NULL,
  `started_at` datetime(3) NOT NULL,
  `ended_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_seasons_guild_number` (`guild_id`, `number`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "server_preferences" table
CREATE TABLE `server_preferences` (
  `guild_id` varchar(20) NOT NULL,
  `leaderboard_channel_id` varchar(20) NULL,
  PRIMARY KEY (`guild_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "speedruns" table
CREATE TABLE `speedruns` (
  `id` char(36) NOT NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `user_id` varchar(20) NOT NULL,
  `time_in_seconds` bigint NOT NULL,
  `weapon_type` varchar(255) NOT NULL,
  `proof_link` varchar(255) NOT NULL,
  `is_verified` bool NOT NULL DEFAULT 0,
  `season_id` char(36) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_speedruns_deleted_at` (`deleted_at`),
  INDEX `idx_speedruns_season_id` (`season_id`),
  CONSTRAINT `fk_seasons_speedruns` FOREIGN KEY (`season_id`) REFERENCES `seasons` (`id`) ON UPDATE RESTRICT ON DELETE RESTRICT
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
