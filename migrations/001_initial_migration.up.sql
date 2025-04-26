-- Create the database (optional)
-- CREATE DATABASE blog_db; -- Usually created manually in PostgreSQL

-- Switch to the database
-- \c blog_db; -- Use this inside psql

-- ---------------------------------------------
-- 1. Users Table
-- ---------------------------------------------
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role TEXT DEFAULT 'reader',
    avatar_url VARCHAR(255),
    bio TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- ---------------------------------------------
-- 2. Posts Table
-- ---------------------------------------------
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    excerpt TEXT,
    content TEXT NOT NULL,
    featured_image VARCHAR(255),
    status TEXT DEFAULT 'draft',
    published_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_posts_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- ---------------------------------------------
-- 3. Categories Table
-- ---------------------------------------------
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- ---------------------------------------------
-- 4. PostCategories Table
-- ---------------------------------------------
CREATE TABLE IF NOT EXISTS post_categories (
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    CONSTRAINT fk_postcategories_post FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_postcategories_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
    UNIQUE(post_id, category_id)
);

-- ---------------------------------------------
-- 5. Tags Table
-- ---------------------------------------------
CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- ---------------------------------------------
-- 6. PostTags Table
-- ---------------------------------------------
CREATE TABLE IF NOT EXISTS post_tags (
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL,
    CONSTRAINT fk_posttags_post FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_posttags_tag FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
    UNIQUE(post_id, tag_id)
);

-- ---------------------------------------------
-- 7. Comments Table
-- ---------------------------------------------
CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL,
    user_id UUID,
    name VARCHAR(255),
    email VARCHAR(255),
    content TEXT NOT NULL,
    parent_id INTEGER,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_comments_post FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_comments_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT fk_comments_parent FOREIGN KEY (parent_id) REFERENCES comments(id) ON DELETE CASCADE
);

-- ---------------------------------------------
-- 8. Likes Table
-- ---------------------------------------------
CREATE TABLE IF NOT EXISTS likes (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    post_id INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_likes_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_likes_post FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    UNIQUE(user_id, post_id)
);

-- ---------------------------------------------
-- 9. Media Table
-- ---------------------------------------------
CREATE TABLE IF NOT EXISTS media (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_url VARCHAR(255) NOT NULL,
    mime_type VARCHAR(100),
    size BIGINT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_media_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- ---------------------------------------------
-- 10. PostViews Table
-- ---------------------------------------------
CREATE TABLE IF NOT EXISTS post_views (
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL,
    user_id UUID,
    ip_address VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_postviews_post FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_postviews_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);

-- ---------------------------------------------
-- Indexes
-- ---------------------------------------------
CREATE INDEX IF NOT EXISTS idx_posts_user ON posts(user_id);
CREATE INDEX IF NOT EXISTS idx_comments_post ON comments(post_id);
CREATE INDEX IF NOT EXISTS idx_likes_post ON likes(post_id);
CREATE INDEX IF NOT EXISTS idx_media_user ON media(user_id);
CREATE INDEX IF NOT EXISTS idx_postviews_post ON post_views(post_id);
