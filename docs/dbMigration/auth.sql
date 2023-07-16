create table auths(
    id serial,
    email text not null,
    password text not null,
    username text not null,
    is_active boolean not null default false,
    created_at timestamptz,
    activated_at timestamptz
)