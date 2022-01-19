CREATE TABLE IF NOT EXISTS cluster_metro (
    id uuid NOT NULL default uuid_generate_v4(),
    name character varying(256) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    modified_at timestamp with time zone NOT NULL,
    trash boolean NOT NULL default false,
    latitude character varying(16) NOT NULL,
    longitude character varying(16) NOT NULL,
    city character varying(64),
    state character varying(64),
    country character varying(64),
    cc character varying(2),
    st character varying(3),
    organization_id uuid,
    partner_id uuid NOT NULL
);

ALTER TABLE cluster_metro OWNER TO clusterdbuser;

ALTER TABLE ONLY cluster_metro ADD CONSTRAINT cluster_metro_pkey PRIMARY KEY (id);