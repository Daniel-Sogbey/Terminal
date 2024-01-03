--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0
-- Dumped by pg_dump version 16.0

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
-- Name: originals; Type: TABLE; Schema: public; Owner: dansogbey
--

CREATE TABLE public.originals (
    id integer NOT NULL,
    original_url character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.originals OWNER TO dansogbey;

--
-- Name: originals_id_seq; Type: SEQUENCE; Schema: public; Owner: dansogbey
--

CREATE SEQUENCE public.originals_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.originals_id_seq OWNER TO dansogbey;

--
-- Name: originals_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dansogbey
--

ALTER SEQUENCE public.originals_id_seq OWNED BY public.originals.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: dansogbey
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO dansogbey;

--
-- Name: shorts; Type: TABLE; Schema: public; Owner: dansogbey
--

CREATE TABLE public.shorts (
    id integer NOT NULL,
    short_url character varying(255) NOT NULL,
    original_url_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.shorts OWNER TO dansogbey;

--
-- Name: shorts_id_seq; Type: SEQUENCE; Schema: public; Owner: dansogbey
--

CREATE SEQUENCE public.shorts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.shorts_id_seq OWNER TO dansogbey;

--
-- Name: shorts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dansogbey
--

ALTER SEQUENCE public.shorts_id_seq OWNED BY public.shorts.id;


--
-- Name: originals id; Type: DEFAULT; Schema: public; Owner: dansogbey
--

ALTER TABLE ONLY public.originals ALTER COLUMN id SET DEFAULT nextval('public.originals_id_seq'::regclass);


--
-- Name: shorts id; Type: DEFAULT; Schema: public; Owner: dansogbey
--

ALTER TABLE ONLY public.shorts ALTER COLUMN id SET DEFAULT nextval('public.shorts_id_seq'::regclass);


--
-- Name: originals originals_pkey; Type: CONSTRAINT; Schema: public; Owner: dansogbey
--

ALTER TABLE ONLY public.originals
    ADD CONSTRAINT originals_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: dansogbey
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: shorts shorts_pkey; Type: CONSTRAINT; Schema: public; Owner: dansogbey
--

ALTER TABLE ONLY public.shorts
    ADD CONSTRAINT shorts_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: dansogbey
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: shorts shorts_originals_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: dansogbey
--

ALTER TABLE ONLY public.shorts
    ADD CONSTRAINT shorts_originals_id_fk FOREIGN KEY (original_url_id) REFERENCES public.originals(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

