CREATE TABLE rabbit(
  id SERIAL PRIMARY KEY,
  color character varying(50) NULL
);

INSERT INTO rabbit (id, color) VALUES 
(1, 'white'),
(2, 'black'),
(3, 'gray'),
(4, 'carrot'),
(5, 'pink');