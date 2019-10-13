-- store the keywords
CREATE TABLE keyword (
    id bigserial primary key not null,
    word varchar(128) not null,
    created_at timestamptz not null
);

-- store the domain
CREATE TABLE domain (
    id bigserial primary key not null,
    link text not null,
    created_at timestamptz not null
);

-- store the links to places
CREATE TABLE reference (
    id bigserial primary key not null,
    domain_id bigserial not null,

    -- we store the entire href to avoid joins and extra lookups.
    -- note that I don't think we would need to store parameters/queries?
    href text not null,
    created_at timestamptz not null
);