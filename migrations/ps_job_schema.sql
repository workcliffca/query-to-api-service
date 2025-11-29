-- PeopleSoft PS_JOB Table Schema
-- This is one of the core tables in PeopleSoft HCM

DROP TABLE IF EXISTS ps_job CASCADE;

-- Create PS_JOB table (simplified version with key fields)
CREATE TABLE ps_job (
    emplid VARCHAR(11) NOT NULL,
    empl_rcd NUMERIC(3,0) NOT NULL,
    effdt DATE NOT NULL,
    effseq NUMERIC(3,0) NOT NULL,
    empl_status VARCHAR(1),
    status_dt DATE,
    action VARCHAR(3),
    action_dt DATE,
    action_reason VARCHAR(3),
    position_nbr VARCHAR(8),
    jobcode VARCHAR(6),
    deptid VARCHAR(10),
    location VARCHAR(10),
    business_unit VARCHAR(5),
    company VARCHAR(3),
    paygroup VARCHAR(3),
    empl_class VARCHAR(3),
    full_part_time VARCHAR(1),
    reg_temp VARCHAR(1),
    officer_cd VARCHAR(1),
    supervisor_id VARCHAR(11),
    reports_to VARCHAR(11),
    grade VARCHAR(3),
    step NUMERIC(3,0),
    std_hours NUMERIC(6,2),
    comprate NUMERIC(18,6),
    currency_cd VARCHAR(3),
    comp_frequency VARCHAR(5),
    PRIMARY KEY (emplid, empl_rcd, effdt, effseq)
);

-- Create indexes for performance
CREATE INDEX ps_job_idx1 ON ps_job(emplid);
CREATE INDEX ps_job_idx2 ON ps_job(deptid);
CREATE INDEX ps_job_idx3 ON ps_job(effdt);
CREATE INDEX ps_job_idx4 ON ps_job(empl_status);

-- Insert realistic PeopleSoft job data
INSERT INTO ps_job VALUES
-- Employee 1: John Smith - IT Manager
('1001', 0, '2023-01-15', 0, 'A', '2023-01-15', 'HIR', '2023-01-15', 'NEW', 'P0001234', 'MGR001', 'IT-100', 'LOC-NYC', 'US001', 'ABC', 'PG1', 'FTE', 'F', 'R', 'N', '1000', '1000', 'G12', 5, 40.00, 125000.000000, 'USD', 'A'),

-- Employee 2: Sarah Johnson - Senior Developer
('1002', 0, '2022-06-01', 0, 'A', '2022-06-01', 'HIR', '2022-06-01', 'NEW', 'P0001235', 'DEV002', 'IT-100', 'LOC-NYC', 'US001', 'ABC', 'PG1', 'FTE', 'F', 'R', 'N', '1001', '1001', 'G10', 3, 40.00, 95000.000000, 'USD', 'A'),

-- Employee 3: Michael Chen - HR Director
('1003', 0, '2021-03-10', 0, 'A', '2021-03-10', 'HIR', '2021-03-10', 'NEW', 'P0001236', 'DIR001', 'HR-200', 'LOC-SF', 'US001', 'ABC', 'PG1', 'FTE', 'F', 'R', 'N', '1000', '1000', 'G14', 8, 40.00, 145000.000000, 'USD', 'A'),

-- Employee 4: Emily Davis - HR Specialist
('1004', 0, '2022-09-15', 0, 'A', '2022-09-15', 'HIR', '2022-09-15', 'NEW', 'P0001237', 'HRS001', 'HR-200', 'LOC-SF', 'US001', 'ABC', 'PG1', 'FTE', 'F', 'R', 'N', '1003', '1003', 'G08', 2, 40.00, 68000.000000, 'USD', 'A'),

-- Employee 5: Robert Martinez - Finance Manager
('1005', 0, '2020-11-20', 0, 'A', '2020-11-20', 'HIR', '2020-11-20', 'NEW', 'P0001238', 'FIN001', 'FIN-300', 'LOC-CHI', 'US001', 'ABC', 'PG1', 'FTE', 'F', 'R', 'N', '1000', '1000', 'G12', 6, 40.00, 115000.000000, 'USD', 'A'),

-- Employee 6: Lisa Wong - Senior Accountant
('1006', 0, '2021-07-01', 0, 'A', '2021-07-01', 'HIR', '2021-07-01', 'NEW', 'P0001239', 'ACC002', 'FIN-300', 'LOC-CHI', 'US001', 'ABC', 'PG1', 'FTE', 'F', 'R', 'N', '1005', '1005', 'G09', 4, 40.00, 78000.000000, 'USD', 'A'),

-- Employee 7: David Brown - Part-time Contractor
('1007', 0, '2023-02-01', 0, 'A', '2023-02-01', 'HIR', '2023-02-01', 'NEW', 'P0001240', 'CON001', 'IT-100', 'LOC-NYC', 'US001', 'ABC', 'PG2', 'CON', 'P', 'T', 'N', '1001', '1001', 'G07', 1, 20.00, 45000.000000, 'USD', 'A'),

-- Employee 8: Jennifer Lee - Sales Director
('1008', 0, '2019-05-15', 0, 'A', '2019-05-15', 'HIR', '2019-05-15', 'NEW', 'P0001241', 'DIR002', 'SAL-400', 'LOC-LA', 'US001', 'ABC', 'PG1', 'FTE', 'F', 'R', 'Y', '1000', '1000', 'G14', 10, 40.00, 165000.000000, 'USD', 'A'),

-- Employee 9: James Wilson - Sales Rep
('1009', 0, '2022-01-10', 0, 'A', '2022-01-10', 'HIR', '2022-01-10', 'NEW', 'P0001242', 'REP001', 'SAL-400', 'LOC-LA', 'US001', 'ABC', 'PG1', 'FTE', 'F', 'R', 'N', '1008', '1008', 'G08', 1, 40.00, 65000.000000, 'USD', 'A'),

-- Employee 10: Maria Garcia - Terminated Employee
('1010', 0, '2021-04-01', 0, 'T', '2023-08-31', 'TER', '2023-08-31', 'RES', 'P0001243', 'ADM001', 'HR-200', 'LOC-SF', 'US001', 'ABC', 'PG1', 'FTE', 'F', 'R', 'N', '1003', '1003', 'G07', 2, 40.00, 55000.000000, 'USD', 'A'),

-- Employee 11: Kevin Anderson - Promotion Example (multiple records)
('1011', 0, '2020-01-01', 0, 'A', '2020-01-01', 'HIR', '2020-01-01', 'NEW', 'P0001244', 'DEV001', 'IT-100', 'LOC-NYC', 'US001', 'ABC', 'PG1', 'FTE', 'F', 'R', 'N', '1001', '1001', 'G08', 1, 40.00, 75000.000000, 'USD', 'A'),
('1011', 0, '2022-01-01', 0, 'A', '2022-01-01', 'PRO', '2022-01-01', 'MER', 'P0001245', 'DEV002', 'IT-100', 'LOC-NYC', 'US001', 'ABC', 'PG1', 'FTE', 'F', 'R', 'N', '1001', '1001', 'G09', 1, 40.00, 88000.000000, 'USD', 'A'),

-- Employee 12: Amanda Taylor - Leave of Absence
('1012', 0, '2021-08-15', 0, 'L', '2023-10-01', 'LOA', '2023-10-01', 'PER', 'P0001246', 'HRS002', 'HR-200', 'LOC-SF', 'US001', 'ABC', 'PG1', 'FTE', 'F', 'R', 'N', '1003', '1003', 'G08', 3, 40.00, 70000.000000, 'USD', 'A');
