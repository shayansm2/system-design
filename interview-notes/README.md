## interview steps
1. Requirements clarifications
2. System interface definition
3. Back-of-the-envelope estimation
4. Defining data model
5. High-level design
6. Detailed design
7. Identifying and resolving bottlenecks

### step 1: Requirements
requirements can be clarified as:

1. Functional Requirements
2. Non-Functional Requirements
3. Extended Requirements

### Step 2: Back-of-the-envelope estimation or Capacity Estimation and Constraints

inputs or your own estimations:
- daily active users (DAU)
- request per second (RPS)
- read vs write ratio (read heavy or wright heavy or both)

outputs:
- Traffic estimates
  - request per second (RPS)
  - query per second (QPS)
- Storage estimates (B)
- Bandwidth estimates (in and out B/s)
- Memory estimates
  - 80-20 rule


### Step 3: System interface definition or System APIs
something that you are good at :)

### step 4: Defining data model or DB design
- choose the db
  - SQL and transactional: MySQL, SQL server
  - NoSQL: DynamoDB, Cassandra, Riak
- db schema

### step 5: High-level design
- in the high level design do not mention sharding, replication, redundancy, load balanced and etc.
- also see each component of your system as a service (like a micro-service architecture, it's not a web server doing everything)

### step 6: Detailed design

### step 7: Identifying and resolving bottlenecks
not it's time to:
- add *load balancers* for solving **single point of failures** problem
- solve **reliabaility** with *redundancy (replication)*
- add *partitioning (sharding)* for
- add *cache* for better **availability**
- talk about DB cleanup
- talk about *Telemetry* for **monitoring** and **analytics**