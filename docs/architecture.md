# NovaDB Architecture

NovaDB is an LSM-Tree-inspired key-value storage engine built in Go.

## Data Flow

Client Request
      |
      v
 Write-Ahead Log (WAL)
      |
      v
   MemTable
      |
      v
 SSTable Flush
      |
      v
   SSTables
      |
      v
  Compaction
      |
      v
 Bloom Filter
      |
      v
 MVCC Snapshot Reads

## Components

### WAL
Stores operations before they are applied to memory. Used for crash recovery.

### MemTable
In-memory storage for recent writes.

### SSTables
Immutable on-disk storage files.

### Compaction
Merges multiple SSTables into fewer files.

### Bloom Filter
Reduces unnecessary SSTable lookups.

### MVCC
Stores multiple versions of data and supports snapshot reads.