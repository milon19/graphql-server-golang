CREATE TABLE IF NOT EXISTS books (
     id serial primary key,
     title VARCHAR(255) NOT NULL,
     price float NOT NULL,
     isbn_no int NOT NULL unique
);

CREATE TABLE IF NOT EXISTS authors (
   id serial primary key,
   name VARCHAR(255) NOT NULL,
   biography VARCHAR(10000) NOT NULL
);

CREATE TABLE books_authors (
   book_id int NOT NULL,
   author_id int NOT NULL,
   PRIMARY KEY (book_id, author_id),
   FOREIGN KEY (book_id) REFERENCES books(id) ON UPDATE CASCADE,
   FOREIGN KEY (author_id) REFERENCES authors(id) ON UPDATE CASCADE
);

INSERT INTO books
    (id, title, price, isbn_no)
VALUES (1, 'Unlocking Android', 520, 1933988673),
       (2, 'Android in Action, Second Edition', 100.12, 1935182722),
       (3, 'Specification by Example', 278, 1617290084),
       (4, 'Flex 3 in Action', 1000, 1933988746),
       (5, 'Flex 4 in Action', 1023.21, 1935182420);

INSERT INTO authors
(id, name, biography)
VALUES (1, 'W. Frank Ableson', 'W. Frank Ableson is and author of Unlocking Android & Android in Action, Second Edition'),
       (2, 'Charlie Collins', 'Charlie Collins is and author of Unlocking Android'),
       (3, 'Robi Sen', 'Robi Sen is and author of Unlocking Android & Android in Action, Second Edition'),
       (4, 'Gojko Adzic', 'Gojko Adžić is a software delivery consultant and author of several books on Serverless computing, Impact Mapping, Specification by example, Behavior Driven Development, Test Driven Development and Agile Testing.'),
       (5, 'Tariq Ahmed', 'Tariq Mahmood Ahmad, Baron Ahmad of Wimbledon (born 3 April 1968), is a British-Pakistani businessman and a Conservative life peer.'),
       (6, 'Faisal Abid', 'Faisal Abid is the CTO and Co-Founder at Eirene . Additionally, Faisal Abid has had 6 past jobs including Chief Technology Officer at CalendarHero . Eirene CTO and Co-Founder Nov 2019.'),
       (7, 'Dan Orlando', 'Dan Orlando is CEO at Creative RIA and a recognized leader in the Flash Platform community. As a long time consultant and Adobe Community Professional'),
       (8, 'John C. Bland II', 'I am, mostly, a self-taught developer who thoroughly enjoys exploring new tech. Coding for hours on end is a joy but only with the purpose of launching.'),
       (9, 'Joel Hooks', 'My name is Joel Hooks. I''m a skilled virtual assistant, software developer, and a collaborator at egghead.io. This is my personal site where I drop notes');

INSERT INTO books_authors
(book_id, author_id)
VALUES (1, 1),
       (1, 2),
       (1, 3),
       (2, 1),
       (2, 3),
       (3, 4),
       (4, 5),
       (4, 6),
       (5, 5),
       (5, 7),
       (5, 8),
       (5, 9);
