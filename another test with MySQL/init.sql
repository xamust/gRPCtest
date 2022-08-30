--
-- База данных: `gRPCTest`
--
CREATE DATABASE gRPCTest;
use gRPCTest;
SET charset utf8;
--
-- Структура таблицы для писателей...
--
create table TestWritersTable
(
    id     int         not null,
    writer varchar(255) not null,
    constraint TestWritersTable_pk
        primary key (id)
);
create unique index TestWritersTable_Writer_uindex
    on TestWritersTable (writer);
create unique index TestWritersTable_id_uindex
    on TestWritersTable (id);

--
-- Структура таблицы для книг...
--
create table TestBooksTable
(
    id int auto_increment,
    book varchar(255) not null,
    constraint TestBooksTable_pk
        primary key (id)
);
create unique index TestBooksTable_Writer_uindex
    on TestBooksTable (book);
create unique index TestBooksTable_id_uindex
    on TestBooksTable (id);

--
-- Дамп данных таблицы для писателей...
--
INSERT INTO TestWritersTable (writer) VALUES
('Лукъяненко С.'),
('Бредбери Р.'),
('Лем С.'),
('Оруэлл Д.'),
('Стругацкие А. Б.'),
('Толстой А.');

--
-- Дамп данных таблицы для книг...
--
INSERT INTO TestBooksTable (book) VALUES
('Звездая тень'),
('Осенние визиты'),
('Пикник на обочине'),
('1984'),
('Солярис'),
('Звёздные дневники Йона Тихого'),
('Война миров'),
('Марсианские хроники'),
('Извне'),
('Гиперболоид инженера Гарина'),
('Аэлита');
