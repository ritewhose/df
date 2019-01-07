CREATE TABLE quotes (
    id integer primary key,
    creator varchar(255),
    count integer not null default 0,
    quote text,
    addtime datetime default current_timestamp,
    lasttime datetime default current_timestamp
);
