### requirements
#### functional requirements
customers:
- generating short urls
- can pick a custom short url optionally
- expire links after a expiration date given by customer
users:
- redirect to the main url after clicking to the short url
#### non functional requirements
- highly available
- minimal latency for redirection
- not predictable shortened links
#### extended requirements
- analytics: number of redirection of each link
- REST API
### system interface
- `GET redirect`
- `POST createUrl(userKey, originalUrl, customAlias = None, expireDate = None)`
- `POST deleteUrl(userKey, url)`
### back of the envelope estimation
- **assumptions**
	- read heavy $${read \over write} = 100$$
	- 500M new url shortenings per month
	- store data for 5 years
	- each stored object is 500 bytes
- Traffic estimates
	$$writeRPS = 500M({request \over month}) \times30({month\over day}) \times 24({day \over hour}) \times 3600({hour \over second}) \approx 200rps $$
	$$readRPS = 200rps \times 100 ({read\over write}) = 20K/s$$
- Storage estimates
$$storage = 500M({object \over month}) \times 5 (year) \times 500 ({B/object}) \times 12 ({year \over month}) \approx 15TB$$
- Bandwidth estimates
### data model
### high level design
### detailed design
### identify and resolve bottlenecks