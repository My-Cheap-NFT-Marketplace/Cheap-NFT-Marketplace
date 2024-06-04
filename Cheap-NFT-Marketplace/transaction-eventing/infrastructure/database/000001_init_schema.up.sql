CREATE TYPE order_status AS ENUM ('created', 'pending', 'on_sale', 'not_for_sale', 'rejected');

CREATE TABLE ntf_orders (
                            "trx" varchar(255) PRIMARY KEY,
                            "token_id" varchar(255) NOT NULL,
                            "owner" varchar(255) NOT NULL,
                            "contract_address" varchar(255) NOT NULL,
                            "creator" varchar(255)  NOT NULL,
                            "token_standard" varchar(255)  NOT NULL,
                            "status" sale_status  NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now()),
                            "updated_at" timestamptz NOT NULL DEFAULT (now())
                          );

CREATE INDEX ON ntf_orders ("owner");
CREATE INDEX ON ntf_orders ("creator");
CREATE INDEX ON ntf_orders ("contract_address");