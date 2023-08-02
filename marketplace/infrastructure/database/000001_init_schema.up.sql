CREATE TYPE sale_status AS ENUM ('on_sale', 'sold', 'not_for_sale');

CREATE TABLE nfts_to_sell (
                            "token_id" varchar(255) PRIMARY KEY,
                            "owner" varchar(255) NOT NULL,
                            "contract_address" varchar(255) NOT NULL,
                            "creator" varchar(255)  NOT NULL,
                            "token_standard" varchar(255)  NOT NULL,
                            "status" sale_status  NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now()),
                            "updated_at" timestamptz NOT NULL DEFAULT (now())
                          );

CREATE INDEX ON nfts_to_sell ("creator");
CREATE INDEX ON nfts_to_sell ("contract_address");