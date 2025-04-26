-- +migrate Down

DROP INDEX IF EXISTS idx_postviews_post;
DROP INDEX IF EXISTS idx_media_user;
DROP INDEX IF EXISTS idx_like_post;
DROP INDEX IF EXISTS idx_comment_post;
DROP INDEX IF EXISTS idx_post_user;

DROP TABLE IF EXISTS post_views;
DROP TABLE IF EXISTS media;
DROP TABLE IF EXISTS likes;
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS post_tags;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS post_categories;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS users;