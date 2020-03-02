create table if not exists itop_ticket (
    id serial not null primary key,
    ref varchar(32) not null unique,
    request_type varchar(128) not null default '',
    servicesubcategory_name varchar(128) not null default '',
    urgency varchar(1) not null default '0',
    origin varchar(10) not null default '',
    caller_id_friendlyname varchar(32) not null default '',
    impact varchar(1) not null default '',
    title varchar(255) not null default '',
    description varchar(512) not null default ''
);

create table if not exists ding_approve (
    id serial not null primary key,
    process_id varchar(64) not null default ''
);