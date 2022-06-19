-- DROP ROLE tech_rota;

CREATE ROLE tech_rota WITH 
	NOSUPERUSER
	CREATEDB
	CREATEROLE
	NOINHERIT
	LOGIN
	NOREPLICATION
	NOBYPASSRLS
	CONNECTION LIMIT -1;

-- public.events definition

-- Drop table

-- DROP TABLE public.events;

CREATE TABLE public.events (
	id int4 NOT NULL,
	event_type varchar NOT NULL,
	date_start timestamp NOT NULL,
	date_end timestamp NULL,
	all_day bool NOT NULL DEFAULT false,
	recur_type varchar NULL,
	recur_interval varchar NULL,
	CONSTRAINT events_pk PRIMARY KEY (id),
	CONSTRAINT events_fk FOREIGN KEY (id) REFERENCES public.events_meta(event_id) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO public.events
(id, event_type, date_start, date_end, all_day, recur_type, recur_interval) VALUES
(0, 'duty_tech_1', '2022-06-19T00:00:00+0000', '2022-06-19T23:59:59+0000', true, '', ''),
(1, 'duty_tech_2', '2022-06-19T00:00:00+0000', '2022-06-19T23:59:59+0000', true, '', '');


-- public.events_meta definition

-- Drop table

-- DROP TABLE public.events_meta;

CREATE TABLE public.events_meta (
	id int4 NOT NULL,
	recur_start int4 NULL,
	event_id int4 NOT NULL,
	CONSTRAINT events_meta_pk PRIMARY KEY (id),
	CONSTRAINT events_meta_un UNIQUE (event_id)
);