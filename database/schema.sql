--table leave_status
CREATE TABLE leave_status(
    id SERIAL PRIMARY KEY,
    status_name VARCHAR(50) UNIQUE NOT NULL
);

INSERT INTO leave_status(status_name) VALUES
    ('Pending'),
    ('Approved'),
    ('Rejected'),
    ('Cancelled');

-- employees table
CREATE TABLE employees (
    id SERIAL PRIMARY KEY ,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    role_id INT REFERENCES roles(id),
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp
);

-- roles table
CREATE TABLE roles(
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) UNIQUE NOT NULL
);

INSERT INTO roles(role_name) VALUES
    ('Intern'),
    ('Employee'),
    ('Team Lead'),
    ('Manager'),
    ('HR'),
    ('Admin');

-- leave_status table
CREATE TABLE leave_status(
    id SERIAL PRIMARY KEY,
    status_name VARCHAR(50) UNIQUE NOT NULL
);

INSERT INTO leave_status(status_name) VALUES
    ('Pending'),
    ('Approved'),
    ('Rejected'),
    ('Cancelled');

-- leave_request table
CREATE TABLE leave_request(
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    reason TEXT,
    status_id INT REFERENCES leave_status(id),
    leave_type_id INT REFERENCES leave_types(id),
    created_at TIMESTAMP DEFAULT current_timestamp
);

-- leave_types table
CREATE TABLE leave_types(
    id SERIAL PRIMARY KEY,
    leave_type_name VARCHAR(50) UNIQUE NOT NULL
);

INSERT INTO leave_types (leave_type_name) VALUES
    ('Annual Leave'),
    ('Sick Leave'),
    ('Casual Leave'),
    ('Unpaid Leave'),
    ('Maternity Leave'),
    ('Paternity Leave');
