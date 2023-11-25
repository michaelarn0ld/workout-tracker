CREATE DATABASE IF NOT EXISTS workout_tracker;
USE workout_tracker;
DROP TABLE IF EXISTS exercise_progress;
CREATE TABLE IF NOT EXISTS exercise_progress (
  exercise ENUM('front-squat', 'deadlift', 'weighted-dip', 'barbell-curl', 'barbell-row', 'weighted-pullup', 'incline-dumbbell-press',  'bench-press', 'arnold-press', 'overhead-press'),
  date_recorded DATETIME,
  weight_in_lbs SMALLINT
);
