CREATE TABLE
    IF NOT EXISTS "auth"."users" (
        "id" VARCHAR PRIMARY KEY NOT NULL,
        "name" VARCHAR(50) NOT NULL,
        "username" VARCHAR(50) NOT NULL,
        "email" VARCHAR(100) NOT NULL,
        "password" VARCHAR NOT NULL,
        "created_at" TIMESTAMP DEFAULT TIMEZONE('utc'::TEXT, CURRENT_TIMESTAMP) NOT NULL,
        "updated_at" TIMESTAMP DEFAULT TIMEZONE('utc'::TEXT, CURRENT_TIMESTAMP) NOT NULL,
        CONSTRAINT "users_username_unique" UNIQUE ("username"),
        CONSTRAINT "users_email_unique" UNIQUE ("email")
    );