-- File: 2023-07-25-1015_Initial_Schema.sql

-- Up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE
    IF NOT EXISTS public.topic_mappings (
        id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
        node_address TEXT,
        ipfs_id TEXT,
        topic_name TEXT,
        data_path TEXT,
        evaluation_job TEXT,
        evaluation_job_data TEXT

);

-- Down

DROP TABLE IF EXISTS public.topic_mappings;