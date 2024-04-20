CREATE TABLE "order"
(
    OrderUid          VARCHAR(255),
    TrackNumber       VARCHAR(255),
    Entry             VARCHAR(255),
    Locale            VARCHAR(255),
    InternalSignature VARCHAR(255),
    CustomerId        VARCHAR(255),
    DeliveryService   VARCHAR(255),
    ShardKey          VARCHAR(255),
    SmId              INT,
    DateCreated       TIMESTAMP,
    OofShard          VARCHAR(255),
    PRIMARY KEY (OrderUid)
);

CREATE TABLE delivery
(
    OrderUid VARCHAR(255),
    Name     VARCHAR(255),
    Phone    VARCHAR(255),
    Zip      VARCHAR(255),
    City     VARCHAR(255),
    Address  VARCHAR(255),
    Region   VARCHAR(255),
    Email    VARCHAR(255),
    FOREIGN KEY (OrderUid) REFERENCES "order" (OrderUid)
);

CREATE TABLE payment
(
    OrderUid     VARCHAR(255),
    Transaction  VARCHAR(255),
    RequestId    VARCHAR(255),
    Currency     VARCHAR(255),
    Provider     VARCHAR(255),
    Amount       INT,
    PaymentDt    BIGINT,
    Bank         VARCHAR(255),
    DeliveryCost INT,
    GoodsTotal   INT,
    CustomFee    INT,
    FOREIGN KEY (OrderUid) REFERENCES "order" (OrderUid)
);

CREATE TABLE item
(
    OrderUid    VARCHAR(255),
    ChrtId      INT,
    TrackNumber VARCHAR(255),
    Price       INT,
    Rid         VARCHAR(255),
    Name        VARCHAR(255),
    Sale        INT,
    Size        VARCHAR(255),
    TotalPrice  INT,
    NmId        INT,
    Brand       VARCHAR(255),
    Status      INT,
    FOREIGN KEY (OrderUid) REFERENCES "order" (OrderUid)
);