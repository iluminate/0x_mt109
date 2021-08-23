create table series_actors
(
    id int auto_increment,
    serie_id int,
    actor_id int,
    constraint series_actors_pk
        primary key (id),
    constraint episodes_serie_id_fk
        foreign key (serie_id) references series (id),
    constraint episodes_actor_id_fk
        foreign key (actor_id) references actors (id)
);