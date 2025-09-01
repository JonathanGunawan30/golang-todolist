CREATE TYPE status AS ENUM ('NEW', 'ON PROGRESS', 'EXPIRED');
CREATE TYPE category_type AS ENUM ('TASK', 'EVENT');

CREATE TABLE activities (
                            id SERIAL PRIMARY KEY,
                            title VARCHAR(250) NOT NULL,
                            category category_type NOT NULL,
                            description TEXT NOT NULL,
                            activity_date TIMESTAMPTZ NOT NULL,
                            status status DEFAULT 'NEW'
);