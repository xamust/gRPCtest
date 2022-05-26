--
-- База данных: `KVADOTest`
--
CREATE DATABASE KVADOTest;
use KVADOTest;
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
   writerId int not null,
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
INSERT INTO KVADOTest.TestWritersTable (id, writer) VALUES
(1, 'Лукъяненко С.'),
(2, 'Бредбери Р.'),
(3, 'Лем С.'),
(4, 'Оруэлл Д.'),
(5, 'Стругацкие А. Б.'),
(6, 'Толстой А.');

--
-- Дамп данных таблицы для книг...
--
INSERT INTO KVADOTest.TestBooksTable (book, writerId) VALUES
('Звездая тень', 1),
('Осенние визиты', 1),
('Пикник на обочине', 5),
('1984', 4),
('Солярис', 3),
('Звёздные дневники Йона Тихого', 3),
('Война миров', 2),
('Марсианские хроники', 2),
('Извне', 5),
('Гиперболоид инженера Гарина', 6),
('Аэлита', 6);
