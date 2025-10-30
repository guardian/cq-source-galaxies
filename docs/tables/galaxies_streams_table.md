# Table: galaxies_streams_table

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|stream_id|`utf8`|
|stream_name|`utf8`|
|stream_description|`utf8`|
|stream_members|`list<item: utf8, nullable>`|