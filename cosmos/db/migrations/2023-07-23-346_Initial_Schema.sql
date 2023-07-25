-- File: 2023-07-23-346_Initial_Schema.sql

-- Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS public.topic_mappings (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    node_id TEXT,
    node_type TEXT,
    topic_name TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Down
DROP TABLE IF EXISTS public.topic_mappings;
