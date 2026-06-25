NovaDB
Overview
NovaDB is a custom LSM-Tree-inspired key-value storage engine built in Go. The project was developed to explore database internals including persistence, crash recovery, storage optimization, and concurrency concepts used by modern database systems.
Features
Write-Ahead Logging (WAL)
Crash Recovery
MemTable for in-memory writes
SSTable-based persistent storage
Multi-SSTable search
SSTable Compaction
Bloom Filter optimization
Multi-Version Concurrency Control (MVCC)
Transaction Snapshot Reads
Architecture
Client Request
→ WAL
→ MemTable
→ SSTable Flush
→ SSTables
→ Compaction
→ Bloom Filter Assisted Lookup
→ MVCC Snapshot Reads
Project Structure
novadb1/
├── cmd/server
├── internal/wal
├── internal/memtable
├── internal/sstable
├── internal/bloom
├── internal/mvcc
├── docs
└── benchmarks
Benchmark Results
100000 records inserted
Execution Time: 139.4453 ms
Technologies Used
Go (Golang)
Git
GitHub
Future Improvements
Distributed Replication
Raft Consensus Algorithm
B+ Tree Indexing
Query Language Support
Concurrent Transaction Manager
Learning Outcomes
This project demonstrates practical understanding of:
Storage Engine Design
Database Internals
Log Structured Merge Trees (LSM Trees)
Crash Recovery Mechanisms
Probabilistic Data Structures
Versioned Data Management