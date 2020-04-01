CREATE TABLE public.events1
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    starttime timestamp with time zone,
    endtime timestamp with time zone,
    duration integer,
    typeduration integer,
    title character varying(255) COLLATE pg_catalog."default",
    note character varying(1024) COLLATE pg_catalog."default",
    issending integer NOT NULL DEFAULT 0,
    CONSTRAINT events1_pkey PRIMARY KEY (id)
)