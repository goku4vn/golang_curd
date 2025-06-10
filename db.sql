-- Tạo custom enum type
CREATE TYPE customer_status AS ENUM ('active', 'inactive', 'deleted');

-- Tạo bảng customers
CREATE TABLE "customers" (
     "id" varchar(36) NOT NULL PRIMARY KEY,
     "name" varchar(100) NOT NULL,
     "email" varchar(100) UNIQUE NOT NULL,
     "status" customer_status  DEFAULT 'active',
     "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
     "updated_at" timestamp  DEFAULT CURRENT_TIMESTAMP
); 