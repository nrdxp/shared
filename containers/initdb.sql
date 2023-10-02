CREATE DATABASE homechart;
CREATE ROLE "homechart" WITH CREATEDB LOGIN PASSWORD 'homechart';
GRANT ALL PRIVILEGES ON DATABASE homechart TO homechart;
\c homechart;
GRANT ALL ON SCHEMA public TO homechart;

CREATE DATABASE homechart_self_hosted;
GRANT ALL PRIVILEGES ON DATABASE homechart_self_hosted TO homechart;
\c homechart_self_hosted;
GRANT ALL ON SCHEMA public TO homechart;

CREATE DATABASE testdb;
CREATE ROLE "testuser" WITH CREATEDB LOGIN PASSWORD 'testuser';
GRANT ALL PRIVILEGES ON DATABASE testdb TO testuser;
\c testdb;
GRANT ALL ON SCHEMA public TO testuser;
