create table verify_auths(
    id serial,
    auth_id int not null,
    token text not null,
    created_at timestamptz,
    activated_at timestamptz,
    expired_at timestamptz not null
);
create table auths(
    id serial,
    email text not null,
    password text not null,
    username text not null,
    first_name text not null,
    last_name text not null,
    is_active boolean not null default false,
    created_at timestamptz,
    activated_at timestamptz
);