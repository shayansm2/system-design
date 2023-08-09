### requirements
#### functional requirements
- paste a data and get a unique url
- view the paste data with url
- only text, no image and file
- expire urls after some time (can be determined by user)
- custom alias for the url
#### non functional requirements
- reliability (no data loss)
- availability (low server down time)
- not predictable shortened links
- minimum latency
#### extended requirements
- analytics
- rest api
### system interface
```
POST addPaste(userKey, data, customUrl = None, expireDate = Nonde) -> pasteKey
POST delPaste(userkey, pasteKey)
GET getPaste(pasteKey) -> data
```
### back of the envelope estimation
- **assumptions**
	- read heavy $${read \over write} = 5$$
	- 1M new post per day
	- user can store max 10MB data
	- average storage per post is 10KB
	- store data for 10 years
	- 70 capacity model for storage (30% margin)
	- 20-80 rule for cachine
- Traffic estimates 
  $$writeRPS = 1M({request \over day}) \times {1\over3600\times24}({day \over sec}) \approx12({paste \over sec}) $$
  $$readRPS = 12*5=60({reads\over sec})$$
- Storage estimates
$$storage = 10(years)\times365({days\over year}) \times1M({object \over day})\times 10K({B \over object}) \approx36TB$$
using 70% capacity model: 
	$$storageNeed = 36TB \times {100/70} = 51.4 TB$$
- length and storage of keys
	$$numberOfRequiredKeys = 1M({paste \over day}) \times 10(years) \approx 3.6 Billion Paste $$
	if we use base64 encoding: 6 letter keys are enough
	$$64 ^ 5 \approx 1 Billion $$
	$$64 ^ 6 \approx 68.7 Billion $$ $$keyStorage = 6 (chars) \times 1({byte \over char}) \times 3.6 Billion Paste \approx 22 GB$$
- Bandwidth estimates
$$incomingData = 12({object \over sec}) \times 10({KB \over object}) = 120 ({KB \over sec})$$
$$outgoingData = 60({object \over sec}) \times 10({KB \over object}) = 0.6 ({MB \over sec})$$
- Memory estimates
$$$$
### data model and database
### high level design
### detailed design
### identify and resolve bottlenecks