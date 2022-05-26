DROP DATABASE KVADOTest;

CREATE DATABASE KVADOTest;

use KVADOTest;

create table TestWritersTable
(
    id     int         not null,
    Writer varchar(32) not null,
    constraint TestWritersTable_pk
        primary key (id)
);
create unique index TestWritersTable_Writer_uindex
    on TestWritersTable (Writer);
create unique index TestWritersTable_id_uindex
    on TestWritersTable (id);

create table TestBooksTable
(
    id int auto_increment,
    Book varchar(32) not null,
    WriterId int not null,
    constraint TestBooksTable_pk
        primary key (id)
);
create unique index TestBooksTable_Writer_uindex
    on TestBooksTable (Book);
create unique index TestBooksTable_id_uindex
    on TestBooksTable (id);

INSERT INTO KVADOTest.TestWritersTable (id, Writer) VALUES (1, 'Лукъяненко С.');
INSERT INTO KVADOTest.TestWritersTable (id, Writer) VALUES (2, 'Бредбери Р.');
INSERT INTO KVADOTest.TestWritersTable (id, Writer) VALUES (3, 'Лем С.');
INSERT INTO KVADOTest.TestWritersTable (id, Writer) VALUES (4, 'Оруэлл Д.');
INSERT INTO KVADOTest.TestWritersTable (id, Writer) VALUES (5, 'Стругацкие А. Б.');
INSERT INTO KVADOTest.TestWritersTable (id, Writer) VALUES (6, 'Толстой А.');

INSERT INTO KVADOTest.TestBooksTable (Book, WriterId) VALUES ('Звездая тень', 1);
INSERT INTO KVADOTest.TestBooksTable (Book, WriterId) VALUES ('Осенние визиты', 1);
INSERT INTO KVADOTest.TestBooksTable (Book, WriterId) VALUES ('Пикник на обочине', 5);
INSERT INTO KVADOTest.TestBooksTable (Book, WriterId) VALUES ('1984', 4);
INSERT INTO KVADOTest.TestBooksTable (Book, WriterId) VALUES ('Солярис', 3);
INSERT INTO KVADOTest.TestBooksTable (Book, WriterId) VALUES ('Звёздные дневники Йона Тихого', 3);
INSERT INTO KVADOTest.TestBooksTable (Book, WriterId) VALUES ('Война миров', 2);
INSERT INTO KVADOTest.TestBooksTable (Book, WriterId) VALUES ('Марсианские хроники', 2);
INSERT INTO KVADOTest.TestBooksTable (Book, WriterId) VALUES ('Извне', 5);
INSERT INTO KVADOTest.TestBooksTable (Book, WriterId) VALUES ('Гиперболоид инженера Гарина', 6);
INSERT INTO KVADOTest.TestBooksTable (Book, WriterId) VALUES ('Аэлита', 6);
