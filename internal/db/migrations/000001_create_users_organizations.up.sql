BEGIN;

CREATE TABLE IF NOT EXISTS organizations (
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS jobs (
    id              BIGSERIAL PRIMARY KEY,
    name            TEXT NOT NULL,
    organization_id BIGINT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users (
    id              BIGSERIAL PRIMARY KEY,
    name            TEXT NOT NULL,
    job_id          BIGINT     REFERENCES jobs(id)           ON DELETE CASCADE,
    organization_id BIGINT     REFERENCES organizations(id)  ON DELETE SET NULL
);

-- helpful indexes for FKs
CREATE INDEX IF NOT EXISTS idx_jobs_org_id  ON jobs(organization_id);
CREATE INDEX IF NOT EXISTS idx_users_job_id ON users(job_id);
CREATE INDEX IF NOT EXISTS idx_users_org_id ON users(organization_id);

COMMIT;