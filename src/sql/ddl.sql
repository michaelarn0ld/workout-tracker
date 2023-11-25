CREATE DATABASE IF NOT EXISTS workout_tracker;
USE workout_tracker;
CREATE TABLE IF NOT EXISTS exercise_progress (
  exercise ENUM('front-squat', 'deadlift', 'weighted-dip', 'barbell-curl', 'barbell-row', 'weighted-pullup', 'incline-dumbbell-press',  'bench-press', 'seated-db-press', 'overhead-press'),
  date_recorded DATE,
  weight_in_lbs SMALLINT
);
