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
	- cache 20% of requests with 80% of the load
- Traffic estimates
	$$writeRPS = 500M({request \over month}) \times30({month\over day}) \times 24({day \over hour}) \times 3600({hour \over second}) \approx 200rps $$
	$$readRPS = 200rps \times 100 ({read\over write}) = 20K/s$$
- Storage estimates
$$storage = 500M({object \over month}) \times 5 (year) \times 500 ({B \over object}) \times 12 ({year \over month}) \approx 15TB$$
- Bandwidth estimates
$$Incoming Data = wrtieRPS \times ObjectStorage = 200 ({object \over second}) \times 500 ({B \over object}) = 100 KB/s$$
$$OutgoingData = readRPS \times ObjectStorage = 20K ({object \over second}) \times 500 ({B \over object}) = 10 MB/s$$
- Memory estimation
$$cacheMemory = 20K ({object \over second}) \times 20\% \times 500 ({B \over object}) \times 3600 ({second\over hour}) \times 24({hour\over day}) \approx 170GB$$
### data model and database
database: NoSQL ->  DynamoDB or Cassandra

| table | fields                                                                   |
|-------|--------------------------------------------------------------------------|
| users | id<br>name<br>email<br>createdAt<br>lastLogin                            |
| urls  | id<br>hashUrl<br>originalUrl<br>userId<br>creationDate<br>expirationDate |

### high level design
#### 1. Encoding url
Encoding algorithms: [cryptographic hashes in go](./cryptographic_hash.go)
#### 2. generating keys beforehand
![[high level design.canvas]]
### detailed design
### identify and resolve bottlenecks
#### load balancer

#### db cleanup
- fanout on user read: remove the expired links when users try to view them
- cron job: periodically get and remove expired links