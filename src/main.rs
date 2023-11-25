use mysql::*;
use std::env;
use mysql::prelude::*;
use chrono::*;

struct Exercise {
    exercise: String,
    date_recorded: String,
    weight_in_lbs: i16,
}

fn insert(exercise: Exercise) -> std::result::Result<(), Box<dyn std::error::Error>> {
    // Establish mysql connection
    let url = "mysql://root@localhost:3306/workout_tracker";
    let mut conn = Pool::new(url)?.get_conn()?;


    let query = "INSERT INTO exercise_progress (exercise, date_recorded, weight_in_lbs) VALUES (:exercise, :date_recorded, :weight_in_lbs)";
    let stmt = conn.prep(query)?;
    conn.exec::<String, _, _>(&stmt, params! {"exercise" => exercise.exercise})?;
    Ok(())
}

fn track(date: String, args: &Vec<String>) {
    let workout = &args[2];
    let queries = match workout.as_str() {
        "back" => ["weighted-pullup", "barbell-row"],
        "legs" => ["front-squat", "deadlift"],
        "chest" => ["bench-press", "incline-dumbbell-press"],
        "shoulders" => ["seated-db-press", "overhead-press"],
        "arms" => ["barbell-curl", "weighted-dip"],
        _ => ["", ""]
    };

    for query in &queries {
        if !query.is_empty() {
            let mut line  = String::new();
            // read input line string and store it into line
            println!("Enter the weight for the following exercise: {}", query);
            std::io::stdin().read_line(&mut line).unwrap();
            // convert line to integer
            let weight: i16 = line.trim().parse().unwrap();
            let exercise = Exercise { exercise: (*query).to_string(), date_recorded: date.clone(), weight_in_lbs: weight };
            insert(exercise);
        }
    }
    
}

fn display(args: &Vec<String>) {
    println!("This functionality is not supported yet!");
}

fn main() -> std::result::Result<(), Box<dyn std::error::Error>> {

    // Today is the day that we will use for mysql inserts
    let date = Utc::now().naive_utc().to_string();
    let args: Vec<String> = env::args().collect();

    let command = &args[1];
    match command.as_str() {
        "track" => track(date, &args),
        "display" => display(&args),
        _ =>  return Err("incorrect command, please try a supported operation!".into()),
    }

    Ok(())
}
