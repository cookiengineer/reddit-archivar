# Reddit Archivar

This is my attempt at rescueing as much from my favorite subreddits as possible.

The Web Archive has also a running archiving attempt over at /r/DataHoarder, but
the Archive Warrior just scrapes HTML which is pretty much useless for my OSINT
related work.

This tool here is built on the basis of the old `v1` API on reddit, which downloads
and stores all the `JSON` files directly, so that they can be processed later.

The keywords are kind of statically set inside the `cmds/reddit-archivar/main.go`
file for now, and they focus on the important discussion topics related to my
cyber security work.


## TODO

Currently, only the links are gathered together, and each listing is stored in
the `archive/<subreddit>/listings/<thread-id>.json` files. In order to have all
the comments, too, it's necessary to parse out all those JSON files and download
the thread's comments data as well.


# License

AGPL-3.0
