# Table: galaxies_people_table

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|name|`utf8`|
|email_id|`utf8`|
|role|`utf8`|
|teams|`list<item: utf8, nullable>`|
|streams|`list<item: utf8, nullable>`|
|git_hub_handle|`utf8`|