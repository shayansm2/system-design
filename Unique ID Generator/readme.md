#### Functional Requirements (FRs)

- The service should generate a unique ID on every call.
- IDs should contain only digits (0-9).
- IDs should be ordered by date, i.e., if ID1 is generated before ID2 then ID1 should be numerically smaller than ID2.

#### Non-Functional Requirements (NFRs)

- An ID should not take more than 8 bytes.
- Total IDs created per day: 1B.

---
- [[solution 1]]
- [[solution 2]]