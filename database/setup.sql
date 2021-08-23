CREATE USER 'tvseriesuser' IDENTIFIED BY '36a0dd17';
GRANT ALL PRIVILEGES ON * . * TO 'tvseriesuser';
CREATE DATABASE tvseriesdb;
USE tvseriesdb;
source tables/actors.sql
source tables/directors.sql
source tables/series.sql
source tables/episodes.sql
source tables/series_actors.sql