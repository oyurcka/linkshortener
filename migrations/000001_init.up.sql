CREATE TABLE links
(
    id integer NOT NULL PRIMARY KEY,
    url character varying(200) NOT NULL
);

CREATE TABLE shortlinks
(
    id integer NOT NULL PRIMARY KEY,
    shorturl character varying(10) NOT NULL,
    links_id integer NOT NULL
);