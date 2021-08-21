CREATE TABLE IF NOT EXISTS email (
	email_id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	email varchar(255) NOT NULL,
	CONSTRAINT email_pk PRIMARY KEY (email_id)
);

CREATE TABLE IF NOT EXISTS friend_relationship (
	relation_id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	your_id int8 NOT NULL,
	friend_id int8 NOT NULL,
	status varchar(10) NOT NULL,
	CONSTRAINT friend_relationship_pk PRIMARY KEY (relation_id)
);


-- public.friend_relationship foreign keys

ALTER TABLE public.friend_relationship ADD CONSTRAINT friend_email FOREIGN KEY (friend_id) REFERENCES email(email_id);
ALTER TABLE public.friend_relationship ADD CONSTRAINT relation_email FOREIGN KEY (your_id) REFERENCES email(email_id);

insert into email(email)
values ('quan12yt@gmail.com'),
('letoan@gmail.com'),
('tonhut@gmail.com'),
('quang@gmail.com'),
('len@gmail.com');

insert into friend_relationship (your_id, friend_id, status)
values (5, 2, 'FRIEND'),
(1, 3, 'BLOCK'),
(2, 4, 'SUBSCRIBE'),
(3, 1, 'FRIEND'),
(4, 1, 'FRIEND');