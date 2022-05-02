-- Measurements table
create table measurement
(
    uuid           uuid not null
        primary key default gen_random_uuid(),
    created_date bigint,
    heart_rate   integer,
    high         integer,
    low          integer,
    username     varchar(255)
);

-- Events table
create table event
(
    uuid           uuid not null
        primary key,
    created_date bigint,
    description  varchar(255),
    username     varchar(255)
);

-- Snapshots table
create table snapshot
(
    uuid         uuid    not null
        primary key,
    created_date bigint,
    description  varchar(255),
    end_date     bigint,
    is_public    boolean not null,
    start_date   bigint,
    username     varchar(255)
);

-- Users table
create table users
(
    email       varchar(255) not null
        primary key,
    enabled     boolean      not null,
    full_name   varchar(255),
    provider    varchar(255),
    provider_id varchar(255)
        constraint ukt5f0hi34g8ylyq8b61gxsmk6i
            unique
);

