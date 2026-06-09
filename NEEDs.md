# File Indexer — Feature Summary

## Core Requirements

- **Incremental updates** — only re-index changed files, not the whole tree
- **Filesystem event watching** — inotify (Linux) / FSEvents (macOS) for a live index
- **Symlink/hardlink safety** — no infinite loops on circular references

## Metadata

- Name, path, size, timestamps (created/modified/accessed)
- MIME type detection by content, not just extension
- Permissions and ownership
- Checksums/hashes for deduplication and integrity

## Query Capabilities

- Full-text search (if content indexing is in scope)
- Filtering by type, size, date range, extension, owner
- Glob and regex path matching
- Ranked results by relevance, recency, or access frequency

## Robustness

- Graceful handling of permission errors, broken symlinks, unreadable files
- No memory blowup on very large directories
- Crash recovery — index stays consistent if interrupted mid-scan

## Performance

- Bounded parallel/concurrent scanning
- Low CPU/IO impact during background indexing
- Persistent index storage — no full re-scan on restart

## Content Indexing

- Pluggable extractors per MIME type (PDFs, Office docs, source code, etc.)
- Configurable depth — metadata-only vs. full text per file type
- Respects `.gitignore` and configurable exclude patterns

## Operability

- CLI and/or query API
- Index stats (file count, size, last scan time, errors)
- Configurable watch paths and exclusions
- Detailed logging for debugging missed files

---

> The two things that separate a mediocre indexer from a good one are **incremental updates** and **robust error handling** — real filesystems are messy.
