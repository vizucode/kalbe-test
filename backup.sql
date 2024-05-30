--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.2 (Debian 15.2-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: attandances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.attandances (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    employee_id bigint,
    location_id bigint,
    absent_in timestamp with time zone,
    absent_out timestamp with time zone,
    updated_by text,
    created_by text
);


ALTER TABLE public.attandances OWNER TO postgres;

--
-- Name: attandances_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.attandances_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.attandances_id_seq OWNER TO postgres;

--
-- Name: attandances_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.attandances_id_seq OWNED BY public.attandances.id;


--
-- Name: departements; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.departements (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    departement_name text,
    created_by text,
    updated_by text
);


ALTER TABLE public.departements OWNER TO postgres;

--
-- Name: departements_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.departements_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.departements_id_seq OWNER TO postgres;

--
-- Name: departements_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.departements_id_seq OWNED BY public.departements.id;


--
-- Name: employees; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.employees (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    employee_code text,
    password text,
    departement_id bigint,
    position_id bigint,
    superior boolean,
    created_by text,
    updated_by text,
    employee_name text
);


ALTER TABLE public.employees OWNER TO postgres;

--
-- Name: employees_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.employees_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.employees_id_seq OWNER TO postgres;

--
-- Name: employees_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.employees_id_seq OWNED BY public.employees.id;


--
-- Name: locations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.locations (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    location_name text,
    created_by text,
    updated_by text
);


ALTER TABLE public.locations OWNER TO postgres;

--
-- Name: locations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.locations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.locations_id_seq OWNER TO postgres;

--
-- Name: locations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.locations_id_seq OWNED BY public.locations.id;


--
-- Name: positions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.positions (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    departement_id bigint,
    position_name text,
    created_by text,
    updated_by text
);


ALTER TABLE public.positions OWNER TO postgres;

--
-- Name: positions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.positions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.positions_id_seq OWNER TO postgres;

--
-- Name: positions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.positions_id_seq OWNED BY public.positions.id;


--
-- Name: attandances id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.attandances ALTER COLUMN id SET DEFAULT nextval('public.attandances_id_seq'::regclass);


--
-- Name: departements id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.departements ALTER COLUMN id SET DEFAULT nextval('public.departements_id_seq'::regclass);


--
-- Name: employees id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees ALTER COLUMN id SET DEFAULT nextval('public.employees_id_seq'::regclass);


--
-- Name: locations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.locations ALTER COLUMN id SET DEFAULT nextval('public.locations_id_seq'::regclass);


--
-- Name: positions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.positions ALTER COLUMN id SET DEFAULT nextval('public.positions_id_seq'::regclass);


--
-- Data for Name: attandances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.attandances (id, created_at, updated_at, deleted_at, employee_id, location_id, absent_in, absent_out, updated_by, created_by) FROM stdin;
1	2024-05-30 08:15:21.21573+00	2024-05-30 08:15:21.21573+00	\N	4	1	2024-05-30 08:15:21.209944+00	0001-01-01 00:00:00+00		
\.


--
-- Data for Name: departements; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.departements (id, created_at, updated_at, deleted_at, departement_name, created_by, updated_by) FROM stdin;
1	2024-05-30 03:21:31.244869+00	2024-05-30 03:32:29.630113+00	\N	Information Technology		
2	2024-05-30 03:32:15.510295+00	2024-05-30 03:32:15.510295+00	2024-05-30 03:32:42.912278+00	Pharmacy		
\.


--
-- Data for Name: employees; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.employees (id, created_at, updated_at, deleted_at, employee_code, password, departement_id, position_id, superior, created_by, updated_by, employee_name) FROM stdin;
1	2024-05-30 07:12:10.801856+00	2024-05-30 07:12:10.801856+00	2024-05-30 07:17:59.919901+00	2405100000	$2a$10$w3jdc5qv90xt0pNxTpRtweHrDwdv2vUttQKBIaYiYjVcjQqHGcuGK	1	2	t			Andi Budi
2	2024-05-30 07:15:23.055592+00	2024-05-30 07:15:23.055592+00	2024-05-30 07:17:59.919901+00	2405100000	$2a$10$/ChR81cw/vdNBqqjgCHNc.l.Cv0jmWy2SDb6xNKrJt4lVa4X3zBU6	1	2	t			Budi Komar
3	2024-05-30 07:18:08.956589+00	2024-05-30 07:18:08.956589+00	\N	24050000	$2a$10$pAQudv1O3yScGkjq/1AzbuuGjsFjIkQSSbWJ1AyrVMvrCLBDe6yki	1	2	t			Budi Komar
4	2024-05-30 07:18:22.51636+00	2024-05-30 07:20:54.311752+00	\N	24050001	$2a$10$LNoya4GAg0JOBY6pUWmNHudzFEX8Gf1qlVrpdrlEhMavD/1uHmjdy	1	2	t			Budi Dudi
\.


--
-- Data for Name: locations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.locations (id, created_at, updated_at, deleted_at, location_name, created_by, updated_by) FROM stdin;
1	2024-05-30 04:49:37.11494+00	2024-05-30 04:51:22.583936+00	\N	Samarinda		
2	2024-05-30 04:51:31.578057+00	2024-05-30 04:51:31.578057+00	2024-05-30 04:51:44.864218+00	Pintu Gerbang		
\.


--
-- Data for Name: positions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.positions (id, created_at, updated_at, deleted_at, departement_id, position_name, created_by, updated_by) FROM stdin;
2	2024-05-30 04:06:49.488122+00	2024-05-30 04:06:49.488122+00	2024-05-30 04:07:18.574764+00	2	Tech Lead		
1	2024-05-30 04:06:42.977664+00	2024-05-30 04:15:19.805821+00	\N	1	Tech Lead		
\.


--
-- Name: attandances_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.attandances_id_seq', 1, true);


--
-- Name: departements_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.departements_id_seq', 2, true);


--
-- Name: employees_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.employees_id_seq', 4, true);


--
-- Name: locations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.locations_id_seq', 2, true);


--
-- Name: positions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.positions_id_seq', 3, true);


--
-- Name: attandances attandances_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.attandances
    ADD CONSTRAINT attandances_pkey PRIMARY KEY (id);


--
-- Name: departements departements_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.departements
    ADD CONSTRAINT departements_pkey PRIMARY KEY (id);


--
-- Name: employees employees_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_pkey PRIMARY KEY (id);


--
-- Name: locations locations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT locations_pkey PRIMARY KEY (id);


--
-- Name: positions positions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.positions
    ADD CONSTRAINT positions_pkey PRIMARY KEY (id);


--
-- Name: idx_attandances_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_attandances_deleted_at ON public.attandances USING btree (deleted_at);


--
-- Name: idx_departements_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_departements_deleted_at ON public.departements USING btree (deleted_at);


--
-- Name: idx_employees_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_employees_deleted_at ON public.employees USING btree (deleted_at);


--
-- Name: idx_locations_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_locations_deleted_at ON public.locations USING btree (deleted_at);


--
-- Name: idx_positions_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_positions_deleted_at ON public.positions USING btree (deleted_at);


--
-- Name: attandances fk_attandances_employee; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.attandances
    ADD CONSTRAINT fk_attandances_employee FOREIGN KEY (employee_id) REFERENCES public.employees(id);


--
-- Name: attandances fk_attandances_location; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.attandances
    ADD CONSTRAINT fk_attandances_location FOREIGN KEY (location_id) REFERENCES public.locations(id);


--
-- Name: positions fk_departements_positions; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.positions
    ADD CONSTRAINT fk_departements_positions FOREIGN KEY (departement_id) REFERENCES public.departements(id);


--
-- Name: employees fk_employees_departement; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT fk_employees_departement FOREIGN KEY (departement_id) REFERENCES public.departements(id);


--
-- Name: employees fk_employees_position; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT fk_employees_position FOREIGN KEY (position_id) REFERENCES public.positions(id);


--
-- PostgreSQL database dump complete
--

