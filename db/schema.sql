-- Cleanup
DROP TABLE IF EXISTS userprovider;
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS provider;
DROP TABLE IF EXISTS apiurls;
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
CREATE TABLE apiurls (Key TEXT PRIMARY KEY, url TEXT);
-- Inserts
INSERT INTO user (user, token, password, name, surname, email, phone, level) VALUES ('John', 'acs67t23rbhjf987tykgfv', 'Salchichon', 'John', 'Salchichon', 'asdf@omg.god', '666333987', 1);
INSERT INTO provider (name, contact) VALUES ('Balloon', 'ballooncorp@evil.death');
INSERT INTO provider (name, contact) VALUES ('PanzerChomps', 'panzercorp@evil.death');
INSERT INTO provider (name, contact) VALUES ('SimplyDelight', 'simplydelightcorp@evil.death');
INSERT INTO provider (name, contact) VALUES ('Nom', 'nomcorp@evil.death');
INSERT INTO apiurls VALUES ('LoginUrl', '/api/users/Login');
INSERT INTO apiurls VALUES ('GetApiUrlsUrl', '/api/config/GetApiUrls');
INSERT INTO apiurls VALUES ('GetProvidersUrl', '/api/providers/Get');
INSERT INTO apiurls VALUES ('ConnectToProviderUrl', '/api/providers/Connect');
INSERT INTO apiurls VALUES ('GetOrdersUrl', '/api/orders/Get');
INSERT INTO apiurls VALUES ('AssignOrderUrl', '/api/AssignOrder');
