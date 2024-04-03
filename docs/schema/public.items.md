# public.items

## Description

## Columns

| Name        | Type                     | Default           | Nullable | Children                                                                          | Parents                                   | Comment      |
| ----------- | ------------------------ | ----------------- | -------- | --------------------------------------------------------------------------------- | ----------------------------------------- | ------------ |
| id          | character(26)            |                   | false    | [public.item_barcodes](public.item_barcodes.md) [public.stocks](public.stocks.md) |                                           | 商品マスタID      |
| name        | varchar(255)             |                   | false    |                                                                                   |                                           | 商品名          |
| price       | integer                  |                   | false    |                                                                                   |                                           | 単価           |
| category_id | character(26)            |                   | false    |                                                                                   | [public.categories](public.categories.md) | カテゴリID       |
| maker_id    | character(26)            |                   | false    |                                                                                   | [public.makers](public.makers.md)         | メーカーID       |
| created_at  | timestamp with time zone | CURRENT_TIMESTAMP | false    |                                                                                   |                                           |              |
| updated_at  | timestamp with time zone | CURRENT_TIMESTAMP | false    |                                                                                   |                                           |              |
| deleted_at  | timestamp with time zone |                   | true     |                                                                                   |                                           |              |

## Constraints

| Name                    | Type        | Definition                                                            |
| ----------------------- | ----------- | --------------------------------------------------------------------- |
| items_maker_id_fkey     | FOREIGN KEY | FOREIGN KEY (maker_id) REFERENCES makers(id) ON DELETE CASCADE        |
| items_category_id_fkey  | FOREIGN KEY | FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE |
| items_pkey              | PRIMARY KEY | PRIMARY KEY (id)                                                      |
| items_name_key          | UNIQUE      | UNIQUE (name)                                                         |
| items_name_maker_id_key | UNIQUE      | UNIQUE (name, maker_id)                                               |

## Indexes

| Name                    | Definition                                                                               |
| ----------------------- | ---------------------------------------------------------------------------------------- |
| items_pkey              | CREATE UNIQUE INDEX items_pkey ON public.items USING btree (id)                          |
| items_name_key          | CREATE UNIQUE INDEX items_name_key ON public.items USING btree (name)                    |
| items_name_maker_id_key | CREATE UNIQUE INDEX items_name_maker_id_key ON public.items USING btree (name, maker_id) |
| idx_items_created_at    | CREATE INDEX idx_items_created_at ON public.items USING btree (created_at)               |

## Relations

```mermaid
erDiagram

"public.item_barcodes" }o--|| "public.items" : "FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE"
"public.sales" }o--|| "public.item_barcodes" : "FOREIGN KEY (barcode_id) REFERENCES item_barcodes(id) ON DELETE CASCADE"
"public.stocks" }o--|| "public.items" : "FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE"
"public.stocks" }o--|| "public.shops" : "FOREIGN KEY (shop_id) REFERENCES shops(id) ON DELETE CASCADE"
"public.items" }o--|| "public.categories" : "FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE"
"public.items" }o--|| "public.makers" : "FOREIGN KEY (maker_id) REFERENCES makers(id) ON DELETE CASCADE"

"public.items" {
  character_26_ id "商品マスタID"
  varchar_255_ name "商品名"
  integer price "単価"
  character_26_ category_id FK "カテゴリID"
  character_26_ maker_id FK "メーカーID"
  timestamp_with_time_zone created_at ""
  timestamp_with_time_zone updated_at ""
  timestamp_with_time_zone deleted_at ""
}
"public.item_barcodes" {
  character_26_ id ""
  character_26_ item_id FK "商品マスタID"
  timestamp_with_time_zone registrated_at "登録日"
  timestamp_with_time_zone created_at ""
  timestamp_with_time_zone updated_at ""
  timestamp_with_time_zone deleted_at ""
}
"public.sales" {
  character_26_ id "売上ID"
  character_26_ barcode_id FK "バーコードID"
  character_26_ finance_id FK "会計ID"
  timestamp_with_time_zone purchased_at "購入日"
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
"public.shops" {
  character_26_ id "店舗ID"
  varchar_255_ name "店舗名"
  timestamp_with_time_zone created_at ""
  timestamp_with_time_zone updated_at ""
  timestamp_with_time_zone deleted_at ""
}
"public.categories" {
  character_26_ id "カテゴリID"
  varchar_255_ name "カテゴリ名"
  timestamp_with_time_zone created_at ""
  timestamp_with_time_zone updated_at ""
  timestamp_with_time_zone deleted_at ""
}
"public.makers" {
  character_26_ id "メーカーID"
  varchar_255_ name "メーカー名"
  timestamp_with_time_zone created_at ""
  timestamp_with_time_zone updated_at ""
  timestamp_with_time_zone deleted_at ""
}
```

---

> Generated by [tbls](https://github.com/k1LoW/tbls)
