CREATE TABLE public.events
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    starttime timestamp with time zone,
    endtime timestamp with time zone,
    duration integer,
    typeduration integer,
    title text COLLATE pg_catalog."default",
    note text COLLATE pg_catalog."default",
    CONSTRAINT events_pkey PRIMARY KEY (id)
)