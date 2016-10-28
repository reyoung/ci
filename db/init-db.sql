CREATE TYPE BuildStatus AS ENUM (
    'queued',
    'running',
    'success',
    'error',
    'failed'
);

CREATE TYPE OutputChannel AS ENUM (
    'stdout',
    'stderr'
);


CREATE TABLE Versions (
    sha VARCHAR(64) PRIMARY KEY NOT NULL,
    createTime TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() at time zone 'utc')
);

CREATE TABLE Builds (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    version_sha VARCHAR(64) NOT NULL REFERENCES Versions(sha),
    outputs TEXT[],
    outputChannels OutputChannel[]
);
