create table episodes
(
    id int auto_increment,
    name varchar(80) null,
    director_id int null,
    serie_id int null,
    constraint episodes_pk
        primary key (id),
    constraint episodes_directors_id_fk
        foreign key (director_id) references directors (id),
    constraint episodes_series_id_fk
        foreign key (serie_id) references series (id)
);