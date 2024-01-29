CREATE TABLE "AccessLevel" (
  ID SERIAL PRIMARY KEY,
  LEVEL VARCHAR(10)
);

CREATE TABLE "User" (
  ID VARCHAR(255) PRIMARY KEY,
  Login VARCHAR(255),
  Password VARCHAR(255),
  AccessLevelID INT REFERENCES "AccessLevel"(ID), 
  CreatedAt TIMESTAMP WITH TIME ZONE
);

CREATE TABLE "Customer" (
  ID VARCHAR(255) PRIMARY KEY,
  UserID VARCHAR(255) REFERENCES "User"(ID),  
  Document VARCHAR(255),
  Name VARCHAR(255),
  Email VARCHAR(255)
);

CREATE TABLE "Access" (
  ID VARCHAR(255) PRIMARY KEY,
  AccessLevelID INT REFERENCES "AccessLevel"(ID), 
  USERID VARCHAR(255),
  AccessedAt TIMESTAMP WITH TIME ZONE
);

ALTER TABLE "Customer"
ADD CONSTRAINT customer_user_fk FOREIGN KEY (UserID) REFERENCES "User" (ID);

ALTER TABLE "Access"
ADD CONSTRAINT access_accesslevel_fk FOREIGN KEY (AccessLevelID) REFERENCES "AccessLevel" (ID);

ALTER TABLE "User"
ADD CONSTRAINT user_accesslevel_fk FOREIGN KEY (AccessLevelID) REFERENCES "AccessLevel" (ID);

DO $$ 
BEGIN 
  PERFORM pg_sleep(10); 
END $$;

INSERT INTO "AccessLevel" (LEVEL) VALUES ('Admin');
INSERT INTO "AccessLevel" (LEVEL) VALUES ('Manager');
INSERT INTO "AccessLevel" (LEVEL) VALUES ('Supervisor');
INSERT INTO "AccessLevel" (LEVEL) VALUES ('Operator');
INSERT INTO "AccessLevel" (LEVEL) VALUES ('Guest');