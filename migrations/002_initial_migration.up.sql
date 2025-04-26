-- Enable pgcrypto extension (for generating UUIDs)
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Step 1: Add new UUID column (only if not exists)
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_name='users' AND column_name='user_uuid'
    ) THEN
        ALTER TABLE users ADD COLUMN user_uuid UUID DEFAULT gen_random_uuid();
    END IF;
END $$;

-- Step 2: Update only NULL UUIDs (safe to rerun)
UPDATE users SET user_uuid = gen_random_uuid() WHERE user_uuid IS NULL;

-- Step 3: Make UUID column NOT NULL (only if needed)
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_name='users' AND column_name='user_uuid' AND is_nullable='YES'
    ) THEN
        ALTER TABLE users ALTER COLUMN user_uuid SET NOT NULL;
    END IF;
END $$;

-- Step 4: Add UNIQUE constraint (only if not exists)
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 
        FROM pg_constraint 
        WHERE conname = 'unique_user_uuid'
    ) THEN
        ALTER TABLE users ADD CONSTRAINT unique_user_uuid UNIQUE (user_uuid);
    END IF;
END $$;