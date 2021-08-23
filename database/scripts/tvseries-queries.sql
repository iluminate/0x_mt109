/* a. Which actors play in the series Big Sister? */
select actors.name as actors
from actors, series_actors, series
where actors.id = series_actors.actor_id
  and series_actors.serie_id = series.id
  and series.name = 'Big Sister';
/* b. Which director has directed the greatest number of episodes? */
select directors.name from (
   select count(1) as counter, episodes.director_id as director
     from episodes
 group by (episodes.director_id)
) directeds, directors
   where directors.id = directeds.director
order by directeds.counter desc
   limit 1