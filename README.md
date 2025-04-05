# gator the blog-aggregator
A blog aggregator CLI as a part of the boot dev course written in GO

## Requirements
- Go installed on your local
- Postgres installed on your local ( psql 15+ )

## Installation
- go install github.com/TheBarnakhil/gator

## Config file
You will need a .gatorConfig file ( as shown below ) to run the program. This file needs to be created in your home directory `~`

```json
{"db_url":"<connection_string>","current_user_name":"<logged_in_user>"}
```

## Available Commands

- **login**  : 
Log in to the gator application via the user in db, user first needs to be registered to be able to login
usage: `gator login <user_name>`
- **register** : 
Register command to add a user into the db
usage: `gator register <user_name>`
- **reset** : 
Resets all the data: users, posts, feeds, follows
usage : `gator reset`
- **users** : 
Lists all the users available along with the current logged in user
usage : `gator users`
- **addfeed**
Adds a feed to be subscribed to by the user saved by name and url, automatically follows the feed
usage: `gator addfeed name url`
- **feeds** : 
Lists all the feeds available to with the current logged in user
usage : `gator feeds`
- **follow**
Follow a feed to be subscribed to by the user by url
usage: `gator follow url`
- **unfollow**
Unfollow a feed already subscribed to by the user by url
usage: `gator unfollow url`
- **following** : 
Lists all the feeds followed by the current logged in user
usage : `gator following`
- **agg** : 
Aggregate all the feeds into posts followed by the current logged in user for every specified duration like 1s, 20s, 1m etc.
usage : `gator agg <time_duration>`
- **browse** : 
Browse all the aggregated feeds as posts followed by the current logged, you can optionally pass in an integer limit for the number of posts to see in a go
usage : `gator browse <limit>` or `gator browse`