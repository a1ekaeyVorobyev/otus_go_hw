PGDMP         %                x            calendar    12.2    12.2 	               0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16393    calendar    DATABASE     �   CREATE DATABASE calendar WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'Russian_Russia.1251' LC_CTYPE = 'Russian_Russia.1251';
    DROP DATABASE calendar;
                postgres    false            �            1259    16506    events    TABLE     �   CREATE TABLE public.events (
    id integer NOT NULL,
    starttime timestamp with time zone,
    endtime timestamp with time zone,
    duration integer,
    typeduration integer,
    title text,
    note text
);
    DROP TABLE public.events;
       public         heap    postgres    false            �            1259    16504    events_id_seq    SEQUENCE     �   ALTER TABLE public.events ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.events_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    203                      0    16506    events 
   TABLE DATA           ]   COPY public.events (id, starttime, endtime, duration, typeduration, title, note) FROM stdin;
    public          postgres    false    203   �                  0    0    events_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.events_id_seq', 60, true);
          public          postgres    false    202            �
           2606    16513    events events_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.events DROP CONSTRAINT events_pkey;
       public            postgres    false    203                  x������ � �     