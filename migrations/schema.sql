--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.3

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
-- Name: applications; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.applications (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    namespace character varying(255) NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.applications OWNER TO launchbox;

--
-- Name: clusters; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.clusters (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    token character varying(255) NOT NULL,
    last_check_in timestamp without time zone,
    status character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.clusters OWNER TO launchbox;

--
-- Name: projects; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.projects (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    status character varying(255),
    application_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.projects OWNER TO launchbox;

--
-- Name: revisions; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.revisions (
    id uuid NOT NULL,
    status character varying(255),
    commit_sha character varying(255),
    project_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.revisions OWNER TO launchbox;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO launchbox;

--
-- Name: users; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    email character varying(255) NOT NULL,
    email_verified boolean DEFAULT false NOT NULL,
    password_hash character varying(255) NOT NULL,
    avatar_url character varying(255),
    name character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO launchbox;

--
-- Name: applications applications_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT applications_pkey PRIMARY KEY (id);


--
-- Name: clusters clusters_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.clusters
    ADD CONSTRAINT clusters_pkey PRIMARY KEY (id);


--
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- Name: revisions revisions_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.revisions
    ADD CONSTRAINT revisions_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: applications_name_user_id_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX applications_name_user_id_idx ON public.applications USING btree (name, user_id);


--
-- Name: applications_namespace_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX applications_namespace_idx ON public.applications USING btree (namespace);


--
-- Name: clusters_name_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX clusters_name_idx ON public.clusters USING btree (name);


--
-- Name: clusters_token_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX clusters_token_idx ON public.clusters USING btree (token);


--
-- Name: projects_application_id_name_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX projects_application_id_name_idx ON public.projects USING btree (application_id, name);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: applications applications_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT applications_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: projects projects_application_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_application_id_fkey FOREIGN KEY (application_id) REFERENCES public.applications(id);


--
-- Name: revisions revisions_project_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.revisions
    ADD CONSTRAINT revisions_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(id);


--
-- PostgreSQL database dump complete
--

