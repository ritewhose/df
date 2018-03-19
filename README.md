# df

Tiny discord framework.

## Usage

If you're using the DB package for anything, make sure you actually have a database. There's a `.sql` file with the schema included, so all you have to do is:

```
$ sqlite3 <db-name>.db
sqlite> .read create.sql
```

And you're done.