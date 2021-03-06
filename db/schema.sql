-- Cleanup
DROP TABLE IF EXISTS userprovider;
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS provider;
DROP TABLE IF EXISTS apiurl;
-- Create tables
CREATE TABLE user (ID INTEGER PRIMARY KEY, Token TEXT, user TEXT, password TEXT, name TEXT, surname TEXT, email TEXT, phone TEXT, level INTEGER);
CREATE TABLE provider (ID INTEGER PRIMARY KEY, name TEXT, contact TEXT);
CREATE TABLE userprovider (
    user_id INTEGER NOT NULL,
    provider_id INTEGER NOT NULL,
    is_active INTEGER NOT NULL,
    PRIMARY KEY (user_id, provider_id),
    FOREIGN KEY (user_id) REFERENCES user (ID),
    FOREIGN KEY (provider_id) REFERENCES provider (ID)
);
CREATE TABLE apiurl (Key TEXT PRIMARY KEY, url TEXT);
-- Inserts
INSERT INTO user (user, token, password, name, surname, email, phone, level) VALUES ('John', 'acs67t23rbhjf987tykgfv', 'Salchichon', 'John', 'Salchichon', 'asdf@omg.god', '666333987', 1);
INSERT INTO provider (name, contact) VALUES ('Balunn', 'balunncorp@evil.death');
INSERT INTO provider (name, contact) VALUES ('PanzerMeals', 'panzercorp@evil.death');
INSERT INTO provider (name, contact) VALUES ('SimplyDelight', 'simplydelightcorp@evil.death');
INSERT INTO provider (name, contact) VALUES ('Gulp', 'gulpcorp@evil.death');
INSERT INTO provider (name, contact) VALUES ('Fooscott', 'Bartiger');
INSERT INTO apiurl VALUES ('LoginUrl', '/api/users/Login');
INSERT INTO apiurl VALUES ('GetApiUrlsUrl', '/api/config/GetApiUrls');
INSERT INTO apiurl VALUES ('GetProvidersUrl', '/api/providers/Get');
INSERT INTO apiurl VALUES ('ConnectToProviderUrl', '/api/providers/Connect');
INSERT INTO apiurl VALUES ('GetOrdersUrl', '/api/orders/Get');
INSERT INTO apiurl VALUES ('AssignOrderUrl', '/api/orders/Assign');
