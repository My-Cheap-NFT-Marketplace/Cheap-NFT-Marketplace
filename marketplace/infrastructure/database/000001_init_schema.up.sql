CREATE TABLE nfts_to_sell (
                            "tokenId" varchar(255) NOT NULL,
                            "owner" varchar(255) NOT NULL,
                            "contractAddress" varchar(255) NOT NULL,
                            "creator" varchar(255)  NOT NULL,
                            "tokenStandard" varchar(255)  NOT NULL,
                            "status" varchar(255)  NOT NULL,
                            "createdAt" timestamptz NOT NULL DEFAULT (now()),
                            "updatedAt" timestamptz NOT NULL DEFAULT (now())
                          );

CREATE INDEX ON nfts_to_sell ("creator");
CREATE INDEX ON nfts_to_sell ("contractAddress");