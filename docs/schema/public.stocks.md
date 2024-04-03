# public.stocks

## Description

## Columns

| Name         | Type                     | Default           | Nullable | Parents                         | Comment      |
| ------------ | ------------------------ | ----------------- | -------- | ------------------------------- | ------------ |
| id           | character(26)            |                   | false    |                                 | 在庫ID         |
| item_id      | character(26)            |                   | false    | [public.items](public.items.md) | 商品マスタID      |
| shop_id      | character(26)            |                   | false    | [public.shops](public.shops.md) | 店舗ID         |
| quantity     | integer                  |                   | false    |                                 | 残個数          |
| expirated_at | timestamp with time zone |                   | true     |                                 | 消費期限         |
| created_at   | timestamp with time zone | CURRENT_TIMESTAMP | false    |                                 |              |
| updated_at   | timestamp with time zone | CURRENT_TIMESTAMP | false    |                                 |              |
| deleted_at   | timestamp with time zone |                   | true     |                                 |              |

## Constraints

| Name                | Type        | Definition                                                   |
| ------------------- | ----------- | ------------------------------------------------------------ |
| stocks_shop_id_fkey | FOREIGN KEY | FOREIGN KEY (shop_id) REFERENCES shops(id) ON DELETE CASCADE |
| stocks_item_id_fkey | FOREIGN KEY | FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE |
| stocks_pkey         | PRIMARY KEY | PRIMARY KEY (id)                                             |

## Indexes

| Name                    | Definition                                                                       |
| ----------------------- | -------------------------------------------------------------------------------- |
| stocks_pkey             | CREATE UNIQUE INDEX stocks_pkey ON public.stocks USING btree (id)                |
| idx_stocks_created_at   | CREATE INDEX idx_stocks_created_at ON public.stocks USING btree (created_at)     |
| idx_stocks_expirated_at | CREATE INDEX idx_stocks_expirated_at ON public.stocks USING btree (expirated_at) |

## Relations

```mermaid
erDiagram

"public.stocks" }o--|| "public.items" : "FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE"
"public.item_barcodes" }o--|| "public.items" : "FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE"
"public.items" }o--|| "public.categories" : "FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE"
"public.items" }o--|| "public.makers" : "FOREIGN KEY (maker_id) REFERENCES makers(id) ON DELETE CASCADE"
"public.stocks" }o--|| "public.shops" : "FOREIGN KEY (shop_id) REFERENCES shops(id) ON DELETE CASCADE"
"public.users" }o--|| "public.shops" : "FOREIGN KEY (shop_id) REFERENCES shops(id) ON DELETE CASCADE"
"public.finances" }o--|| "public.shops" : "FOREIGN KEY (shop_id) REFERENCES shops(id) ON DELETE CASCADE"

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
"public.shops" {
  character_26_ id "店舗ID"
  varchar_255_ name "店舗名"
  timestamp_with_time_zone created_at ""
  timestamp_with_time_zone updated_at ""
  timestamp_with_time_zone deleted_at ""
}
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
