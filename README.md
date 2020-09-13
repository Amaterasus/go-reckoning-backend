While learning Go I am also rebuilding the backend of my application reckoning in go which is this project.

To get this running on your local machine you will need to create a .env file structured as so.

export PORT="XXXX" (any port that you know is free on your local machine)

export SECRET="" (use anything you want here it's just required so we can load in the secret)

export DATABASE_URL="host=localhost port=5432 user=(your local machines username) dbname=(the postgres database you will be connecting too) password=(your local machines password) sslmode=disable"