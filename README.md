# otz [![Build Status](https://travis-ci.org/otz1/otz.svg?branch=master)](https://travis-ci.org/otz1/otz)
how it works: scraper service will run periodically

persistent database that stores keywords (search terms) and
references.

redis cache containing keyword mappings to references.
each redis entry has an expiration date. the more something is
searched the expiration length increases slightly - maybe this should be
based off traffic?
in effect, search terms that aren't looked for much will be forgotten.

## unknowns:
- how often do we scrape?
- what do we scrape for: common keywords? everything?
- what should the expiry times for search 'links' be
- horizontally scalable persistent db. perhaps NoSQL

# setup
## requirements

* postgres
* redis

## redis
set the redis cache to be in LRU mode

```bash
$ heroku redis:maxmemory redis-concave-92155 --policy volatile-lru --app otz1
```