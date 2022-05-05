--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.2 (Debian 14.2-1.pgdg110+1)

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
-- Name: accounts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.accounts (
    account_id bigint NOT NULL,
    user_id bigint,
    location text,
    website_url text,
    profile_image text,
    display_name text,
    description text,
    name text,
    "is_employee default:false" boolean,
    "reputation default:0" bigint
);


ALTER TABLE public.accounts OWNER TO postgres;

--
-- Name: accounts_account_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.accounts_account_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.accounts_account_id_seq OWNER TO postgres;

--
-- Name: accounts_account_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.accounts_account_id_seq OWNED BY public.accounts.account_id;


--
-- Name: answers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.answers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text,
    body text,
    question_id bigint,
    "view_count default:0" bigint,
    "down_vote_count default:0" bigint,
    up_vote_count bigint,
    "score default:0" bigint
);


ALTER TABLE public.answers OWNER TO postgres;

--
-- Name: answers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.answers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.answers_id_seq OWNER TO postgres;

--
-- Name: answers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.answers_id_seq OWNED BY public.answers.id;


--
-- Name: collectives; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.collectives (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    slug text,
    logo_url text,
    description text,
    website text,
    github text,
    twitter text,
    facebook text,
    tags character varying(64)[],
    created_by bigint,
    "slug,unique" text
);


ALTER TABLE public.collectives OWNER TO postgres;

--
-- Name: collectives_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.collectives_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.collectives_id_seq OWNER TO postgres;

--
-- Name: collectives_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.collectives_id_seq OWNED BY public.collectives.id;


--
-- Name: question_id; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.question_id (
    question_id bigint NOT NULL,
    up_voted_by_id bigint NOT NULL
);


ALTER TABLE public.question_id OWNER TO postgres;

--
-- Name: question_related; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.question_related (
    question_id bigint NOT NULL,
    top_word_id bigint NOT NULL
);


ALTER TABLE public.question_related OWNER TO postgres;

--
-- Name: question_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.question_tags (
    question_id bigint NOT NULL,
    tag_id bigint NOT NULL
);


ALTER TABLE public.question_tags OWNER TO postgres;

--
-- Name: questions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.questions (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text,
    body text,
    "is_answered default:false" boolean,
    "view_count default:0" bigint,
    "down_vote_count default:0" bigint,
    up_vote_count bigint,
    "answer_count default:0" bigint,
    "score default:0" bigint,
    created_by bigint,
    slug text,
    related_works character varying(64)[],
    related_topics character varying(64)[]
);


ALTER TABLE public.questions OWNER TO postgres;

--
-- Name: questions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.questions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.questions_id_seq OWNER TO postgres;

--
-- Name: questions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.questions_id_seq OWNED BY public.questions.id;


--
-- Name: tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tags (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text
);


ALTER TABLE public.tags OWNER TO postgres;

--
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tags_id_seq OWNER TO postgres;

--
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;


--
-- Name: top_words; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.top_words (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    word text,
    count bigint
);


ALTER TABLE public.top_words OWNER TO postgres;

--
-- Name: top_words_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.top_words_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.top_words_id_seq OWNER TO postgres;

--
-- Name: top_words_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.top_words_id_seq OWNED BY public.top_words.id;


--
-- Name: up_voted_bies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.up_voted_bies (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint
);


ALTER TABLE public.up_voted_bies OWNER TO postgres;

--
-- Name: up_voted_bies_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.up_voted_bies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.up_voted_bies_id_seq OWNER TO postgres;

--
-- Name: up_voted_bies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.up_voted_bies_id_seq OWNED BY public.up_voted_bies.id;


--
-- Name: user_collective; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_collective (
    collective_id bigint NOT NULL,
    user_id bigint NOT NULL
);


ALTER TABLE public.user_collective OWNER TO postgres;

--
-- Name: user_collective_collective_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_collective_collective_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_collective_collective_id_seq OWNER TO postgres;

--
-- Name: user_collective_collective_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_collective_collective_id_seq OWNED BY public.user_collective.collective_id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    username text,
    email text,
    password text,
    "user_type default:0" bigint
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: accounts account_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts ALTER COLUMN account_id SET DEFAULT nextval('public.accounts_account_id_seq'::regclass);


--
-- Name: answers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers ALTER COLUMN id SET DEFAULT nextval('public.answers_id_seq'::regclass);


--
-- Name: collectives id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.collectives ALTER COLUMN id SET DEFAULT nextval('public.collectives_id_seq'::regclass);


--
-- Name: questions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.questions ALTER COLUMN id SET DEFAULT nextval('public.questions_id_seq'::regclass);


--
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);


--
-- Name: top_words id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.top_words ALTER COLUMN id SET DEFAULT nextval('public.top_words_id_seq'::regclass);


--
-- Name: up_voted_bies id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.up_voted_bies ALTER COLUMN id SET DEFAULT nextval('public.up_voted_bies_id_seq'::regclass);


--
-- Name: user_collective collective_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_collective ALTER COLUMN collective_id SET DEFAULT nextval('public.user_collective_collective_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.accounts (account_id, user_id, location, website_url, profile_image, display_name, description, name, "is_employee default:false", "reputation default:0") FROM stdin;
1	1	colombo	http://localhost	https://www.gravatar.com/avatar/4fdabdea961ad7448191d8bdc5a5de4c?s=200	john bow	my description	john doe 	f	0
\.


--
-- Data for Name: answers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.answers (id, created_at, updated_at, deleted_at, title, body, question_id, "view_count default:0", "down_vote_count default:0", up_vote_count, "score default:0") FROM stdin;
1	2022-04-28 17:52:48.731422+00	2022-04-28 17:52:48.731422+00	\N	my answer	you can solve this by doing nothing	1	0	0	0	0
2	2022-04-28 17:54:40.289751+00	2022-04-28 17:54:40.289751+00	\N	second answer	you can solve this by doing nothing	1	0	0	0	0
\.


--
-- Data for Name: collectives; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.collectives (id, created_at, updated_at, deleted_at, name, slug, logo_url, description, website, github, twitter, facebook, tags, created_by, "slug,unique") FROM stdin;
\.


--
-- Data for Name: question_id; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.question_id (question_id, up_voted_by_id) FROM stdin;
1	1
1	2
1	3
1	4
1	5
1	6
1	7
1	8
1	9
1	10
1	11
1	12
\.


--
-- Data for Name: question_related; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.question_related (question_id, top_word_id) FROM stdin;
14	25
14	26
14	27
14	28
14	29
14	30
14	31
14	32
14	33
14	34
14	35
14	36
\.


--
-- Data for Name: question_tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.question_tags (question_id, tag_id) FROM stdin;
1	1
1	2
2	3
2	4
7	5
7	6
8	7
8	8
9	9
9	10
10	11
10	12
11	13
11	14
14	17
\.


--
-- Data for Name: questions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.questions (id, created_at, updated_at, deleted_at, title, body, "is_answered default:false", "view_count default:0", "down_vote_count default:0", up_vote_count, "answer_count default:0", "score default:0", created_by, slug, related_works, related_topics) FROM stdin;
2	2022-04-28 11:03:48.805579+00	2022-04-28 11:03:48.805579+00	\N	New Test 2	hehe	f	0	0	0	0	0	1	\N	\N	\N
3	2022-05-01 17:14:01.218556+00	2022-05-01 17:14:01.218556+00	\N	Just another post	post body	f	0	0	0	0	0	1	\N	\N	\N
4	2022-05-01 17:14:14.744854+00	2022-05-01 17:14:14.744854+00	\N	title of new post	post content will be here	f	0	0	0	0	0	1	\N	\N	\N
5	2022-05-01 17:14:55.185233+00	2022-05-01 17:14:55.185233+00	\N	TS2531 Object is possibly 'null' on basic component example and no Validators fired	Why does the hell can I've got this error on *ngIf=title statements ? error TS2531: Object is possibly 'null'.Pushed in StackBlitz, I don't have this error :-/ But I don't get the validator working either :(	f	0	0	0	0	0	1	\N	\N	\N
6	2022-05-01 17:15:23.413342+00	2022-05-01 17:15:23.413342+00	\N	Android studio - chromeless dialog in the editor with 4 code lines	Android Studio 2020.03 Patch 4, shows 4 lines of code in a chromeless dialog, when I am scrolling through the editor.	f	0	0	0	0	0	1	\N	\N	\N
7	2022-05-01 17:16:29.324825+00	2022-05-01 17:16:29.324825+00	\N	React: use Children.map and edit props and use functions as Children in the same time	I want to use Children.map and edit props and use functions as Children in the same time	f	0	0	0	0	0	1	\N	\N	\N
8	2022-05-01 17:17:08.312439+00	2022-05-01 17:17:08.312439+00	\N	how to fix issue on tampermonkey script	i using a tampermonkey script but i have this error i dont know how to fix this issue .	f	0	0	0	0	0	1	\N	\N	\N
9	2022-05-01 17:17:40.733627+00	2022-05-01 17:17:40.733627+00	\N	submit doesn't trigger required error on untouched custom form control element	 have a simple login form with email and password, i wanted to create my own form control component for the email, to add my own validators and mat-error for error messages and also to allow parent components to detect errors and fetch value and to be able to use it as a form control.	f	0	0	0	0	0	1	\N	\N	\N
10	2022-05-01 17:18:20.588301+00	2022-05-01 17:18:20.588301+00	\N	The method 'map' was called on null. Receiver: null Tried calling: map<Container>(Closure: (ImageModel) => Container)	NoSuchMethodError was thrown building FutureBuilder<List>(dirty, state: _FutureBuilderState<List>#8cb05): The method 'map' was called on null. Receiver: null Tried calling: map(Closure: (ImageModel) => Container)	f	0	0	0	0	0	1	\N	\N	\N
11	2022-05-04 07:11:20.199669+00	2022-05-04 07:11:20.199669+00	\N	Regex expression in grep not respecting conditional clause?	Please add sample input (no descriptions, no images, no links) and your desired output for that sample input to your question (no comment)	f	0	0	0	0	0	1	\N	\N	\N
1	2022-04-28 11:03:10.88658+00	2022-05-04 10:07:06.41717+00	\N	test	hehe	t	0	0	12	2	0	1		\N	\N
14	2022-05-04 15:05:34.707358+00	2022-05-04 15:05:34.707358+00	\N	google-map-react not loading- Uncaught TypeError: Cannot read properties of undefined (reading 'getChildren')	have been trying to solve it for a while. I am trying to run simple example of google-map-react. But, This does not load maps. Instead gives following errors and the page is blank.	f	0	0	0	0	0	1	google-map-react-not-loading--uncaught-typeerror:-cannot-read-properties-of-undefined-(reading-'getchildren')	\N	\N
\.


--
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tags (id, created_at, updated_at, deleted_at, name) FROM stdin;
1	2022-04-28 11:03:10.887171+00	2022-04-28 11:03:10.887171+00	\N	tag
2	2022-04-28 11:03:10.887171+00	2022-04-28 11:03:10.887171+00	\N	programming
3	2022-04-28 11:03:48.806844+00	2022-04-28 11:03:48.806844+00	\N	tag
4	2022-04-28 11:03:48.806844+00	2022-04-28 11:03:48.806844+00	\N	programming
5	2022-05-01 17:16:29.325289+00	2022-05-01 17:16:29.325289+00	\N	reactjs
6	2022-05-01 17:16:29.325289+00	2022-05-01 17:16:29.325289+00	\N	programming
7	2022-05-01 17:17:08.313438+00	2022-05-01 17:17:08.313438+00	\N	javascript
8	2022-05-01 17:17:08.313438+00	2022-05-01 17:17:08.313438+00	\N	programming
9	2022-05-01 17:17:40.734007+00	2022-05-01 17:17:40.734007+00	\N	Angular
10	2022-05-01 17:17:40.734007+00	2022-05-01 17:17:40.734007+00	\N	programming
11	2022-05-01 17:18:20.588676+00	2022-05-01 17:18:20.588676+00	\N	flutter
12	2022-05-01 17:18:20.588676+00	2022-05-01 17:18:20.588676+00	\N	programming
13	2022-05-04 07:11:20.211137+00	2022-05-04 07:11:20.211137+00	\N	regex
14	2022-05-04 07:11:20.211137+00	2022-05-04 07:11:20.211137+00	\N	grep
17	2022-05-04 15:05:34.70778+00	2022-05-04 15:05:34.70778+00	\N	google-map-react
\.


--
-- Data for Name: top_words; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.top_words (id, created_at, updated_at, deleted_at, word, count) FROM stdin;
25	2022-05-04 15:05:34.709688+00	2022-05-04 15:05:34.709688+00	\N	properties	1
26	2022-05-04 15:05:34.709688+00	2022-05-04 15:05:34.709688+00	\N	google-map-react	1
27	2022-05-04 15:05:34.709688+00	2022-05-04 15:05:34.709688+00	\N	not	1
28	2022-05-04 15:05:34.709688+00	2022-05-04 15:05:34.709688+00	\N	loading-	1
29	2022-05-04 15:05:34.709688+00	2022-05-04 15:05:34.709688+00	\N	uncaught	1
30	2022-05-04 15:05:34.709688+00	2022-05-04 15:05:34.709688+00	\N	typeerror	1
31	2022-05-04 15:05:34.709688+00	2022-05-04 15:05:34.709688+00	\N	cannot	1
32	2022-05-04 15:05:34.709688+00	2022-05-04 15:05:34.709688+00	\N	read	1
33	2022-05-04 15:05:34.709688+00	2022-05-04 15:05:34.709688+00	\N	undefined	1
34	2022-05-04 15:05:34.709688+00	2022-05-04 15:05:34.709688+00	\N	reading	1
35	2022-05-04 15:05:34.709688+00	2022-05-04 15:05:34.709688+00	\N	'getchildren'	1
36	2022-05-04 15:05:34.709688+00	2022-05-04 15:05:34.709688+00	\N	of	1
\.


--
-- Data for Name: up_voted_bies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.up_voted_bies (id, created_at, updated_at, deleted_at, user_id) FROM stdin;
1	2022-05-04 09:52:16.511986+00	2022-05-04 09:52:16.511986+00	\N	1
2	2022-05-04 09:52:18.707737+00	2022-05-04 09:52:18.707737+00	\N	1
3	2022-05-04 09:52:19.541386+00	2022-05-04 09:52:19.541386+00	\N	1
4	2022-05-04 09:52:58.700331+00	2022-05-04 09:52:58.700331+00	\N	1
5	2022-05-04 09:56:29.325969+00	2022-05-04 09:56:29.325969+00	\N	1
6	2022-05-04 10:03:15.603948+00	2022-05-04 10:03:15.603948+00	\N	1
7	2022-05-04 10:04:03.944295+00	2022-05-04 10:04:03.944295+00	\N	1
8	2022-05-04 10:04:46.481101+00	2022-05-04 10:04:46.481101+00	\N	1
9	2022-05-04 10:05:40.973513+00	2022-05-04 10:05:40.973513+00	\N	1
10	2022-05-04 10:06:01.166165+00	2022-05-04 10:06:01.166165+00	\N	1
11	2022-05-04 10:06:31.34829+00	2022-05-04 10:06:31.34829+00	\N	1
12	2022-05-04 10:07:06.418578+00	2022-05-04 10:07:06.418578+00	\N	1
\.


--
-- Data for Name: user_collective; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_collective (collective_id, user_id) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, email, password, "user_type default:0") FROM stdin;
1	dasith	dasithvidanage000@gmail.com	$2a$14$3TA5e33QXAPWJR1lceFYtuEmv92xQwQbd8USYvJtJ5NKTghOQkdUe	0
\.


--
-- Name: accounts_account_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.accounts_account_id_seq', 1, true);


--
-- Name: answers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.answers_id_seq', 2, true);


--
-- Name: collectives_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.collectives_id_seq', 1, false);


--
-- Name: questions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.questions_id_seq', 14, true);


--
-- Name: tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tags_id_seq', 17, true);


--
-- Name: top_words_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.top_words_id_seq', 36, true);


--
-- Name: up_voted_bies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.up_voted_bies_id_seq', 12, true);


--
-- Name: user_collective_collective_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_collective_collective_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (account_id);


--
-- Name: answers answers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers
    ADD CONSTRAINT answers_pkey PRIMARY KEY (id);


--
-- Name: collectives collectives_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.collectives
    ADD CONSTRAINT collectives_pkey PRIMARY KEY (id);


--
-- Name: question_id question_id_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.question_id
    ADD CONSTRAINT question_id_pkey PRIMARY KEY (question_id, up_voted_by_id);


--
-- Name: question_related question_related_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.question_related
    ADD CONSTRAINT question_related_pkey PRIMARY KEY (question_id, top_word_id);


--
-- Name: question_tags question_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.question_tags
    ADD CONSTRAINT question_tags_pkey PRIMARY KEY (question_id, tag_id);


--
-- Name: questions questions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.questions
    ADD CONSTRAINT questions_pkey PRIMARY KEY (id);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- Name: top_words top_words_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.top_words
    ADD CONSTRAINT top_words_pkey PRIMARY KEY (id);


--
-- Name: up_voted_bies up_voted_bies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.up_voted_bies
    ADD CONSTRAINT up_voted_bies_pkey PRIMARY KEY (id);


--
-- Name: user_collective user_collective_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_collective
    ADD CONSTRAINT user_collective_pkey PRIMARY KEY (collective_id, user_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: question_id fk_question_id_question; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.question_id
    ADD CONSTRAINT fk_question_id_question FOREIGN KEY (question_id) REFERENCES public.questions(id);


--
-- Name: question_id fk_question_id_up_voted_by; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.question_id
    ADD CONSTRAINT fk_question_id_up_voted_by FOREIGN KEY (up_voted_by_id) REFERENCES public.up_voted_bies(id);


--
-- Name: question_related fk_question_related_question; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.question_related
    ADD CONSTRAINT fk_question_related_question FOREIGN KEY (question_id) REFERENCES public.questions(id);


--
-- Name: question_related fk_question_related_top_word; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.question_related
    ADD CONSTRAINT fk_question_related_top_word FOREIGN KEY (top_word_id) REFERENCES public.top_words(id);


--
-- Name: question_tags fk_question_tags_question; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.question_tags
    ADD CONSTRAINT fk_question_tags_question FOREIGN KEY (question_id) REFERENCES public.questions(id);


--
-- Name: question_tags fk_question_tags_tag; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.question_tags
    ADD CONSTRAINT fk_question_tags_tag FOREIGN KEY (tag_id) REFERENCES public.tags(id);


--
-- Name: answers fk_questions_answers; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answers
    ADD CONSTRAINT fk_questions_answers FOREIGN KEY (question_id) REFERENCES public.questions(id);


--
-- Name: user_collective fk_user_collective_collective; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_collective
    ADD CONSTRAINT fk_user_collective_collective FOREIGN KEY (collective_id) REFERENCES public.collectives(id);


--
-- Name: user_collective fk_user_collective_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_collective
    ADD CONSTRAINT fk_user_collective_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: accounts fk_users_user_acc; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT fk_users_user_acc FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

