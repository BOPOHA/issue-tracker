
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
create table roles
(
	id uuid primary key default uuid_generate_v4(),
	name text not null
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

drop table roles;