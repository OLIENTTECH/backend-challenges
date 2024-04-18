# public.users

## Description

## Columns

| Name            | Type                     | Default           | Nullable | Parents                         |
| --------------- | ------------------------ | ----------------- | -------- | ------------------------------- |
| id              | character(26)            |                   | false    |                                 |
| shop_id         | character(26)            |                   | false    | [public.shops](public.shops.md) |
| name            | varchar(255)             |                   | false    |                                 |
| email           | varchar(255)             |                   | false    |                                 |
| password        | varchar(255)             |                   | false    |                                 |
| is_shop_manager | boolean                  |                   | false    |                                 |
| last_logined_at | timestamp with time zone |                   | true     |                                 |
| created_at      | timestamp with time zone | CURRENT_TIMESTAMP | false    |                                 |
| updated_at      | timestamp with time zone | CURRENT_TIMESTAMP | false    |                                 |
| deleted_at      | timestamp with time zone |                   | true     |                                 |

## Constraints

| Name               | Type        | Definition                                                   |
| ------------------ | ----------- | ------------------------------------------------------------ |
| users_shop_id_fkey | FOREIGN KEY | FOREIGN KEY (shop_id) REFERENCES shops(id) ON DELETE CASCADE |
| users_pkey         | PRIMARY KEY | PRIMARY KEY (id)                                             |

## Indexes

| Name                 | Definition                                                                 |
| -------------------- | -------------------------------------------------------------------------- |
| users_pkey           | CREATE UNIQUE INDEX users_pkey ON public.users USING btree (id)            |
| idx_users_created_at | CREATE INDEX idx_users_created_at ON public.users USING btree (created_at) |

## Relations

```mermaid
erDiagram

"public.users" }o--|| "public.shops" : "FOREIGN KEY (shop_id) REFERENCES shops(id) ON DELETE CASCADE"
"public.stocks" }o--|| "public.shops" : "FOREIGN KEY (shop_id) REFERENCES shops(id) ON DELETE CASCADE"
"public.finances" }o--|| "public.shops" : "FOREIGN KEY (shop_id) REFERENCES shops(id) ON DELETE CASCADE"

"public.users" {
  character_26_ id ""
  character_26_ shop_id FK ""
  varchar_255_ name ""
  varchar_255_ email ""
  varchar_255_ password ""
  boolean is_shop_manager ""
  timestamp_with_time_zone last_logined_at ""
  timestamp_with_time_zone created_at ""
  timestamp_with_time_zone updated_at ""
  timestamp_with_time_zone deleted_at ""
}
"public.shops" {
  character_26_ id "店舗ID"
  varchar_255_ name "店舗名"
  timestamp_with_time_zone created_at ""
  timestamp_with_time_zone updated_at ""
  timestamp_with_time_zone deleted_at ""
}
"public.stocks" {
  character_26_ id "在庫ID"
  character_26_ item_id FK "商品マスタID"
  character_26_ shop_id FK "店舗ID"
  integer quantity "残個数"
  timestamp_with_time_zone expirated_at "消費期限"
  timestamp_with_time_zone created_at ""
  timestamp_with_time_zone updated_at ""
  timestamp_with_time_zone deleted_at ""
}
"public.finances" {
  character_26_ id "会計ID"
  character_26_ shop_id FK "店舗ID"
  integer total_amount "合計金額"
  timestamp_with_time_zone purchased_at "購入日"
  timestamp_with_time_zone created_at ""
  timestamp_with_time_zone updated_at ""
  timestamp_with_time_zone deleted_at ""
}
```

---

> Generated by [tbls](https://github.com/k1LoW/tbls)