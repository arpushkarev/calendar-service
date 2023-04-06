-- +goose Up
create table calendar
(
    id            bigserial primary key,
    title         text      not null,
    date_event     timestamp not null,
    duration      timestamp not null,
    description   text,
    author_id     bigint    not null,
    reminder_time timestamp not null,
    created_at    timestamp not null default now(),
    updated_at    timestamp not null
);

-- +goose Down
drop table calendar;
