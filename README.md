 # Workout Tracker
 The purpose of this project is for me to track the latest progress on my new
 GVT workout program. I will also be able to visualize the data over time with
 some various statistics to see how I am progressing/regressing with the new 
 changes.

 Separately, I will also be tracking weekly photos to have a more subjective
 understanding of how my body is changing.


 Check out the GVT program [here](https://github.com/michaelarn0ld/zettelkasten-public/tree/main/202311250238).

## Setup
Please make sure that mysql server is setup on the machine where you are going
to be tracking the GVT data.

To start the mysql server, please run:
```
mysql.server start
```

The project is configured so that the root user for mysql should be accessible
with no password. Login as the root user with:
```
sudo mysql
```

Once you are in, please execute the DDL with:
```
source /path/to/repository/ddl.sql
```

Confirm that the database and table are created AND that the table is now empty:
```
use workout_tracker; show tables;
select * from exercise_progress;
```

Build the executable from the root of the repository with:
```
./build.sh
```


## Usage
Track one of the following splits:
* shoulders
* back
* chest
* legs
* arms
```
./bin/workout-tracker -split <SPLIT>
```
