# Reddit Archivar

This is my attempt at rescueing as much from my favorite subreddits as possible.

The Web Archive has also a running archiving attempt over at /r/DataHoarder, but
the Archive Warrior just scrapes HTML which is pretty much useless for my OSINT
related work.

This tool here is built on the basis of the old `v1` API on reddit, which downloads
and stores all the `JSON` files directly, so that they can be processed later.


## Limitations

- Each listing (hot/top/new) is limited to 10 pages of 100 results each (1000 results),
  which means that the discovery of older threads is only possible via keyword search.

- Keyword search is also limited to 1000 results, which means the more specific the
  keywords, the better the discovery.


## Usage

The keywords are set inside the [keywords.json](./keywords.json) file, and the
subreddit is searched for the given set of keywords.

The script I was/am using to archive the cybersecurity related subreddits is the
[scrape.sh](./scrape.sh) script. It builds the binary and then calls the binary
with each time with the subreddit as an argument.

Please make sure to use the correct upper/lowercase writing of the subreddit's
name, otherwise the redirects might break the scraping mechanism.

```bash
go build -o ./build/reddit-archivar ./cmds/reddit-archivar/main.go;
cp keywords.json ./build/keywords.json;

cd ./build && reddit-archivar /r/MalwareResearch;
```

## TODO

These subreddits went private too early, so I couldn't archive them :(

- /r/security
- /r/websec


# License

AGPL-3.0

